package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)


type Command struct{
	Id int `json:"id"`
	Script string `json:"script"`
	ScriptResult string `json:"scriptResult"`
}

var stopByUser = make(chan int, 100)

func main() {
	
	http.HandleFunc("/Command/All", HandleFuncGet) //Получение списка всех команд
	http.HandleFunc("/Command/Create/", HandleFuncPost) //Создание новой команды
	http.HandleFunc("/Command/",HandleFuncGetId) //Получение одной команды
	http.HandleFunc("/save-text", HandleScriptSaver) //Обработка формы для записи команды
	http.HandleFunc("/upload-file", HandleFileSaver) //Обработка формы для записи команды
	http.HandleFunc("/Command/Stop", HandleStopCommand) //Обработка формы для остановки команды
    log.Fatal(http.ListenAndServe(":80", nil))
	
}

//обработчики верхнего уровня

func HandleFuncGet(w http.ResponseWriter, r * http.Request){
	tmpl,err := template.ParseFiles("./templates/allCommandsForm.tmpl")
	if err!=nil{
		fmt.Println("Ошибка при парсинге формы allCommandsForm: ",err)
	}
	commands:=getScripts()
	if len(commands)<1{
		c := Command{}
		c.Id=0
		c.Script="Здесь будет ваш первый скрипт"
		c.ScriptResult="Здесь будет результат вашего первого скрипта"
		commands = append(commands, c)
		if err := tmpl.Execute(w, commands); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}else{
		if err := tmpl.Execute(w, commands); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func HandleFuncGetId(w http.ResponseWriter, r *http.Request){
	tmpl,err := template.ParseFiles("./templates/oneCommandForm.tmpl")
	if err!=nil{
		fmt.Println("Ошибка при парсинге формы allCommandsForm: ",err)
	}
	
	command:=getOneScript(w,r)
	if err := tmpl.Execute(w, command); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}

func HandleFuncPost(w http.ResponseWriter, r *http.Request){
	tmpl,err := template.ParseFiles("./templates/createScriptForm.tmpl")
	if err!=nil{
		fmt.Println("Ошибка при парсинге шаблона для страницы создания скрипта из файла /templates/createScriptForm.tmpl",err)
	}
   	err = tmpl.Execute(w, r)
   	if err != nil {
	panic(err)
   	}
}

func HandleScriptSaver(w http.ResponseWriter, r  * http.Request) {
	done:=make(chan bool)
	go postScriptsFromText(w,r,done,stopByUser)
	<-done
	w.WriteHeader(http.StatusOK)
}

func HandleFileSaver(w http.ResponseWriter, r  * http.Request){
	done:=make(chan bool)
	go postScriptsFromFile(w,r,done)
	<-done	
	http.Redirect(w, r, "/Command/Create/", http.StatusTemporaryRedirect)
	w.WriteHeader(http.StatusOK)	
}

func HandleStopCommand(w http.ResponseWriter, r  * http.Request){

	//Получение id команды для остановки
	type idToStop struct  {
		Id string `json:id`
	}
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var id idToStop
	if err := json.NewDecoder(r.Body).Decode(&id); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
	return
	}
	stoppId,err:=strconv.Atoi(id.Id)
	if err!=nil{
		fmt.Println("Пользователь остановил уже остановленную программу:",err)
	}
	fmt.Println("Получен id остановки:",stoppId)
	func(){
		stopByUser<-stoppId
	}()
}





