package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
)

var connStr = "user=postgres password=password dbname=test sslmode=disable"
var mu sync.Mutex

//функция создает для нового скрипта запись в бд и передает id записи воркеру
func createNewId () int{

	var currentId int

	db, err := sql.Open("postgres", connStr)
	if err!=nil{
		fmt.Println("Ошибка при подключении к БД из createNewId():",err)
	}
	defer db.Close()
	
	mu.Lock()//блокируем, что бы значение id скрипта было уникально

	rows, err := db.Query("select max(id) from scripts")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

	for rows.Next(){
		err = rows.Scan(&currentId)
		if err!=nil{
			currentId=0
		} 
	}

	currentId++
	_, err = db.Exec("insert into scripts (id,script,scriptresult) values ($1,'Скрипт в процессе выполнения','Результаты выполнения скрипта будут здесь')", currentId)
    if err != nil{
        panic(err)
    }

	mu.Unlock()

	return currentId
}

//Обработка скрипта в форме текста
func postScriptsFromText(w http.ResponseWriter, r  *http.Request, done chan bool, stopByUser chan int){

	var currentId = createNewId()

	//блок 1 - получение текста скрипта, сохранение текста в файл
	
	type TextContext struct {
		Text string `json:"text"`
	}

	decoder := json.NewDecoder(r.Body)
	var data TextContext
	if err := decoder.Decode(&data); err != nil {
	 return
	}
	scriptTextContent:=data.Text

	file,err:=os.Create("../API_Server/scriptsStdout/scriptFromText"+strconv.Itoa(currentId)+".sh")
	defer os.Remove(file.Name())
	if err!=nil{
		fmt.Println("Ошибка при создании файла: ",err)
	}
	defer file.Close()

	//Пишем полученную из формы команду в файл .sh
	_,err=file.Write([]byte("#!/bin/bash\n"))
	if err!=nil{
		fmt.Println("Ошибка из postScriptsFromText при записи шапки скрипта:\n","Файл:\n",file.Name(),"\nОшибка:\n", err)
	}
	_,err=file.Write([]byte(scriptTextContent))
	if err!=nil{
		fmt.Println("Ошибка при записи скрипта в файл:",file.Name(),"\nОшибка:\n",err)
	}
	fmt.Println("Текст успешно загружен на сервер")
	
	db, err := sql.Open("postgres", connStr)
    if err != nil {
		fmt.Println("Ошибка при попытке подключения к базе данных:")
        panic(err)
    } 
	defer db.Close()

	// Блок 2: Выполнение скрипта, динамическое сохранение вывода скрипта построчно в БД
	cmd := exec.Command("C:/Program Files/Git/bin/bash", "../API_Server/scriptsStdout/scriptFromText"+strconv.Itoa(currentId)+".sh")
	stdout, err := cmd.StdoutPipe()
		if err != nil {
			return
		}
	err = cmd.Start()
		if err != nil {
			return
		}


	var output string //переменная для хранения вывода скрипта
	
	scanner := bufio.NewScanner(stdout)
	fmt.Println("Запущен скрипт №",currentId)
	for scanner.Scan() {
			select{
				case checkId:=<-stopByUser:
					if checkId==currentId{
						fmt.Println("Получен id остановки:",checkId)
						fmt.Println("Выполнение скрипта остановлено пользователем")
						return
					}else{
						stopByUser<-checkId
						continue
					}
				default:
					line := scanner.Text()
					output+=line+"\n"
					fmt.Println(line) //раскомментить для проверки вывода
					_, err = db.Exec("update scripts set script=$1,scriptresult=$2 where id = $3", scriptTextContent, output,currentId)
						if err != nil{
							panic(err)
						}	
						
			}
		}
	err = cmd.Wait()//ожидаем стандартное или пользовательское завершение скрипта
	if err != nil {
		return
	}

	// ждем поступления true
	fmt.Println("Выполнение скрипта завершилось штатно")
	done <- true
}

func postScriptsFromFile(w http.ResponseWriter, r *http.Request, done chan bool){

	var currentId = createNewId()
	
	// Получение файла из формы
	file, _, err := r.FormFile("loadFileInput")
	if err != nil {
		
		return
	}
	defer file.Close()

	// Создание нового файла на сервере
	f, err := os.Create("../API_Server/scriptsStdout/scriptFromFile"+strconv.Itoa(currentId)+".sh")
	if err != nil {
		fmt.Println("Ошибка при создании файла на сервере:", err)
		return
	}

	defer os.Remove(f.Name())//удаляем 
	defer f.Close()

	// Копирование содержимого файла в созданный файл на сервере
	_, err = io.Copy(f, file)
	if err != nil {
		fmt.Println("Ошибка при копировании файла 3:", err)
		return
	}
	
	fmt.Println("Файл успешно загружен на сервер.")

	//Выполняем скрипт и сохраняем результаты работы в файл scriptFromFile.sh
	 // Создаем команду для запуска bash-скрипта
	cmd := exec.Command("C:/Program Files/Git/bin/bash", "../API_Server/scriptsStdout/scriptFromFile"+ strconv.Itoa(currentId)+".sh")
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    } 
    defer db.Close()
	stdout, err := cmd.StdoutPipe()
		if err != nil {
			return
		}
	err = cmd.Start()
		if err != nil {
			return
		}

	var output string
	
	filePath := f.Name()

	 // Чтение файла
	content, err := os.ReadFile(filePath)
	if err != nil {
		// Обработка ошибки, если файл не может быть прочитан
		fmt.Println("Ошибка при чтении файла:", err)
	  return
	}

	scanner := bufio.NewScanner(stdout)
	fmt.Println("Запущен скрипт №",currentId)
	for scanner.Scan() {
			select{
			case checkId:=<-stopByUser:
				if checkId==currentId{
					fmt.Println("recived StopId:",checkId)
					fmt.Println("Выполнение скрипта остановлено пользователем")
					return
					} else{
						stopByUser<-checkId
						continue
					}
			default:
				line := scanner.Text()
				output+=line+"\n" // Сохраняем строку в файл или базу данных
				fmt.Println(line)
				_, err = db.Exec("update scripts set script=$1,scriptresult=$2 where id = $3", content, output,currentId)
					if err != nil{
						panic(err)
					}	
						
			}
		}
	err = cmd.Wait()
	if err != nil {
		return
	}
	fmt.Println("Выполнение скрипта ",currentId," завершилось штатно")
	done <- true	 
}

//получение списка скриптов из базы данных
func getScripts()[]Command{

	db, err := sql.Open("postgres", connStr)
	if err!=nil{
		fmt.Println("Ошибка при подключении к БД",err)
	}
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
	}
	rows, err := db.Query("select * from scripts order by id asc;")
    if err != nil {
        panic(err)
    }
    defer rows.Close()

	commands:=[]Command{}
	for rows.Next(){

        c := Command{}
		var scriptFromDb *string
		var scriptResultFromDb *string
        err := rows.Scan(&c.Id, &scriptFromDb, &scriptResultFromDb)
		c.Script=*scriptFromDb
		c.ScriptResult=*scriptResultFromDb

        if err != nil{
            fmt.Println(err)
			
            continue
        }
        commands = append(commands, c)
	}
	return commands
}




func getOneScript(w http.ResponseWriter,r *http.Request) Command{
	db, err := sql.Open("postgres", connStr)
	if err!=nil{
		fmt.Println("Ошибка при создании соединения с БД: ", err)
	}
	defer db.Close()
	if err!=nil{
		fmt.Println(err)
	}
	id := filepath.Base(r.URL.Path)
	rows, err := db.Query(fmt.Sprintf("select * from scripts where id = %s",id))
    if err != nil {
        w.Write([]byte("Данный id не найден в базе данных"))
    }
    defer rows.Close()
	c := Command{}
	for rows.Next(){
		
		err = rows.Scan(&c.Id, &c.Script, &c.ScriptResult)
		if err != nil{
			fmt.Println(err)
		}
	}
	return c
}