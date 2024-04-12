package main

import "fmt"

func main() {
	var N int
	var min int
	count:=1
	fmt.Scan(&N)
	fmt.Scan(&min)
	//fmt.Println(min)
	for i:=0;i<N-1;i++{
		var j int
		fmt.Scan(&j)
		if j<min{
			min=j
			count=1
		}else if j==min{
			count++
		}
		fmt.Println("j:",j,"min:",min,"count:",count)
	}
	fmt.Print(count)
}

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/