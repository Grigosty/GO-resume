package main

import "fmt"

//На вход: сперва число N>=4
//Затем все N чисел подряд

func main() {
	
	var N int
	var counter int
	fmt.Scan(&N)
	var sliceOfNumb []int
	for idx := 0; idx < N; idx++{
	fmt.Scan(&counter)
	sliceOfNumb = append(sliceOfNumb, counter)
	}
	fmt.Println(sliceOfNumb[3])


}

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/