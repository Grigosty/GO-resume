package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)
	counter:=0
	first:=0
	second:=1
	real:=1
	for real<=N{
		real=first+second
		first=second
		second=real
		counter++
	}
	if N!=real&&N!=first{
		fmt.Println("-1")
	} else{
		fmt.Println(counter)
	}
}

/*
1 korova
2 korovy
3 korovy
4 korovy
5 korov
6 korov
7 korov
8 korov
9 korov
10 korov
11 korov
12 korov
13 korov
14 korov
15 korov
16 korov
17 korov
18 korov
19 korov
20 korov
21 korova
22 korovy
23 korovy
24 korovy
25 korov

*/

/*
Инициализация среза
var a []int
var b []int = []int{1, 2, 3}
c := []int{1, 2, 3}
d := []int{1: 12}
*/