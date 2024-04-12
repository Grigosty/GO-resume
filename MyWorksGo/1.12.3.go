package main

import "fmt"

//На вход: сперва число N>=4
//Затем все N чисел подряд

func main() {
	
	array := [5]int{}
	var a int
	counter := 0
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}
	counter = array[0]
	for _,elem:=range array{
		//fmt.Println("элемент", elem)
		if(elem>counter){
			counter = elem
		}
	}
fmt.Println(counter)
}

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/