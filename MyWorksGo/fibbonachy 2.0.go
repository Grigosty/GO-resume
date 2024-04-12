package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibonacci(n))
}

func fibonacci(n int) int {
	f1 := 1
	f2 := 1
	var fn int
	if n>2{
		for i := 0; i < n-2; i++ {
			fn = f1 + f2
			f1 = f2
			f2 = fn
		}
		return fn
	}else{
		if n==1{
			return 1
		} else{
			return 1
		}
	}
}

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