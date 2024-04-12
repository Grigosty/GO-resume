package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	textOrig := []rune(text)
	if len(textOrig)>=5{
		for idx,elem :=range textOrig{
			if unicode.IsDigit(elem)||unicode.Is(unicode.Latin,elem){
				if idx == len(textOrig)-1{
					fmt.Println("Ok")
				} else{
					continue
				}
			} else{
				fmt.Println("Wrong password")
				break
			}
		}
	} else {
		fmt.Println("Wrong password")
	}
}

/*
Что бы связать два файла нужно подвязать их под один пакет
при этом если в файле main2.go будет использоваться и импортироваться пакет math - при вызове функции второго файла из первого
мы можем не импортировать в первый файл пакеты, использованные во втором файле.
*/

/*
Синтаксис функции:
func имя_функции (Список параметров) (список возвращаемых значений){
	Выполняемые параметры
}
*/

/*
Пример функции:
func hello(x int, y int){
	fmt.Println("Hello")
}
*/

/*
Что бы изучить какую-то функцию нужно вконсоли прописать go doc fmt.Println
*/

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/

/*
Советы по выгрузке:

1)
Если вы хотите сделать все правильно, то алгоритм ваших действий  примерно таков:
Создать проект в папке /src/ваш_любой_ник/имя_проекта.
Если у вас уже есть github аккаунт то можете создать проект так: /src/github.com/username/имя_проекта.
2)

*/