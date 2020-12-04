package main

import (
	"github.com/zenfire-cn/commkit/database"
	"fmt"
	"log"
)

func main() {
	dsn := "sqlserver://sa:Helloworld555@127.0.0.1:1433?database=opermon&connection+timeout=30"
	err := database.Init("mssql", dsn, database.NewOption())

	if err != nil {
		log.Fatal(err)
	}

	type Result struct {
		ID   int
		Name string
	}

	db := database.GetDB()
	var result Result
	db.Raw("SELECT id, name FROM users WHERE id = ?", 6).Scan(&result)
	fmt.Printf("%+v\n", result)


	fmt.Printf("%+v\n", db)

	var results []map[string]interface{}
	db.Table("users").Find(&results)

	for _, r := range results {
		for k, v := range r {
			fmt.Printf("%s => %v, ", k, v)
		}
		fmt.Println()
	}

}
