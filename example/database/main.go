package main

import (
	"fmt"
	"github.com/zenfire-cn/commkit/database"
)

func selectAll() {
	db := database.DB()
	rows, err := db.Raw("select * from users").Rows()

	if err != nil {
		fmt.Println("err:", err)
	} else {
		defer rows.Close()

		data := database.RowsToMaps(rows)
		for _, item := range data {
			fmt.Println(item)
		}
	}
}

func main() {
	// 配置
	option := &database.Option{
		database.DBType.MSSQL,
		"127.0.0.1", "1433",
		"SA", "123456",
		"test", "", true,
		0, 0, 0}

	// 初始化
	err := database.Init(option)
	fmt.Println(err)
	// 失败后不断重连的初始化
	// database.InitUntilSuccess(option, 5 * time.Second)

	selectAll()
}
