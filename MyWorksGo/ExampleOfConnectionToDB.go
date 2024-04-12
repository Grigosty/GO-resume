package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	_ "github.com/miekg/pkcs11"
)
func main() { 
 
    db, err := sql.Open("sqlite3", "testdb.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    result, err := db.Exec("insert into tb1 (good,price) values ($1, $2)", 
        "стол", "5000.00")
    if err != nil{
        panic(err)
    }
    fmt.Println(result.LastInsertId())  // id последнего добавленного объекта
    fmt.Println(result.RowsAffected())  // количество добавленных строк
     
}
