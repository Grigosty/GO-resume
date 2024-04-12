package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	max:=a
	for i:=a;i<=b;i++{
		if(i%7==0){
			max=i
		}
	}
	if(max==a&&a%7!=0){
		fmt.Println("NO")
	}else{
		fmt.Println(max)
	}
}

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/