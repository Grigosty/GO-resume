package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("task.data")
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	var counter int
	for{
		s,err:=bufio.NewReader(file).ReadString(';')
	if err != nil && err!=io.EOF{
		panic(err)
	} else{
		s,_=strings.CutSuffix(s,";")
		counter++
		if s=="0"{
			fmt.Println(counter)
		}
		if err==io.EOF{
			break
		}
	}
	
	
	}
	
	// An artificial input source.


	//план действий
	/*
	Вариант обхода невозможности задания разделителя в bufio.ScanWords
	1) запихиваем текст из файла в стринг переменную
	2) создаем из string переменной массив []byte по разделителю ";"
	3) запихиваем созданный массив в переменную-функцию split
	4) радуемся жизни

	*/
	/*const input = "1234;5678;0;1234567901234567890;0;222222222222222222222222222222"
	scanner := bufio.NewScanner(strings.NewReader(input).ReadString(';'))
	// Create a custom split function by wrapping the existing ScanWords function.
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_ = string(token)
			//fmt.Print(advance)
		}
		return
	}
	// Set the split function for the scanning operation.
	scanner.Split(split)
	// Validate the input
	for scanner.Scan() {
		fmt.Printf("%s", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		
	}
	*/
}