package main

import (
	"fmt"
)

func main() {//пример хорошей обработки ошибки в функции
	var a,b int
    fmt.Scan(&a,&b)
    result,err:=divide(a,b)
    if err!=nil{
        fmt.Println("ошибка")
    } else{
        fmt.Println(result)
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