package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	counterStudents:=0.0
	counterRating:=0.0
	var mySt MyStruct
	var average AverStruct
	data,err:=bufio.NewReader(os.Stdin).ReadString('\n')
	if err!=nil&&err!=io.EOF{
		panic(err)
	}
	
	if err := json.Unmarshal([]byte(data), &mySt); err != nil {
		fmt.Println(err)
		return
	}
	for i:=range mySt.Students{
		counterStudents++
		for range mySt.Students[i].Rating{

			counterRating++
		}
	}
	average.Average=counterRating/counterStudents
	averageJSON,err:=json.MarshalIndent(average,"","    ")
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s",averageJSON)
	fmt.Println()
	fmt.Println(mySt.MyStrId)

}

type MyStruct struct{
	MyStrId int `json:"ID"`
	Number string
	Yeat int
	Students[]struct{
		Rating[]float64 `json:"Rating"`
	} `json: "Students"`
}
type AverStruct struct{
	Average float64
}
