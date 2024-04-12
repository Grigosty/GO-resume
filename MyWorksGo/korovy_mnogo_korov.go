package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if 10<=n%100&&n%100<=20{
		fmt.Print(n," korov")
	} else{
		if n%10==1{
			fmt.Print(n," korova")
		} else if n%10==2||n%10==3||n%10==4{
			fmt.Print(n," korovy")
		} else{
			fmt.Print(n," korov")
		}
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