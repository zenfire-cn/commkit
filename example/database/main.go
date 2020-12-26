package main

import (
	"fmt"
	"github.com/zenfire-cn/commkit/database"
)

func main() {
	var (
		dsn = "root:joker8133xx@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local" // mysql
		// dsn    = "sqlserver://sa:123456@127.0.0.1:1433?database=test&connection+timeout=30"        // mssql
		option = database.NewOption()
	)

	// option.SlowQueryTime = 300 * time.Millisecond  // 设置慢查询时间
	// logFile, _ := os.OpenFile("SlowQuery.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	// option.SlowQueryLog = logFile  // 设置慢查询日志

	if err := database.Init("mysql", dsn, option); err != nil {
		fmt.Println("err", err)
	}

	db := database.GetDB()
	var result []map[string]interface{}
	db.Raw("SELECT * FROM users").Find(&result)

	fmt.Println(result)
}
