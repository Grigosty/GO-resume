package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	
	var summator1 int
	var summator2 int
	
	for n>0{
		summator1+=n%10
		n/=10
	}
	for summator1>0{
		summator2+=summator1%10
		summator1/=10
	}	
	fmt.Println(summator2)
}

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/