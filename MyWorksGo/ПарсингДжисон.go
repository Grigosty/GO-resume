package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var itemSt []Items
	var counter int = 0
	dataFromFile, err := os.ReadFile("data-20190514T0100.json")
	if err != nil {
	panic(err)
	}
	if err:=json.Unmarshal([]byte(dataFromFile),&itemSt);err!=nil{
		fmt.Println(err)
	}
	for _,elem:=range itemSt{
		counter+=elem.Global_id
	}
	fmt.Println(counter)
}

type Items struct {
	Global_id int `json : "global_id"`
}
