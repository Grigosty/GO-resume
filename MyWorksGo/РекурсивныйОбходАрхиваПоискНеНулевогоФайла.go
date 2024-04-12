package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	const root = "./testzzz/"
	if err := filepath.Walk(root, MyWalk); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
}

func MyWalk(path string, info os.FileInfo, err error) error{
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	if info.IsDir() {
		return nil // Проигнорируем директории
	}
	file,err:=os.Open(path)
	if err!=nil{
		panic(err)
	}
	defer file.Close()
	
	if data, _ := csv.NewReader(file).ReadAll();len(data) == 10 {
		fmt.Println(data[4][2]) // печатаем значения
	}
	return nil
}

/*
if strings.Contains(info.Name(),".csv"){
		fmt.Println(info.Name())

	}
file,err:=os.Open(path)
		if err != nil{
			panic(err)
		}
		defer file.Close()
		r:=csv.NewReader(file)
		data,err:=r.ReadAll()
		if err!=nil{
			panic(err)
		}
		for _,row := range data{
			fmt.Println(row)
		}
*/