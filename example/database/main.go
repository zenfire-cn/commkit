package main

import (
	"fmt"
	"github.com/zenfire-cn/commkit/database"
)

func main() {
	// 初始化
	// dsn := "sqlserver://sa:123456@127.0.0.1:1433?database=test&connection+timeout=30"
	dsn := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local" // mysql
	if err := database.Init("mysql", dsn, database.NewOption()); err != nil {
		fmt.Println("err", err)
	}

	db := database.GetDB()
	var result []map[string]interface{}
	db.Raw("SELECT * FROM users").Find(&result)

	fmt.Println(result)
}
