package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zenfire-cn/commkit/conf"
)

func main() {
	// 初始化
	conf.Init("config.ini")
	// 使用 （源码可知，viper包下的Get方法，本质上是调用了viper.Viper对象上的Get方法）
	fmt.Println(viper.GetString("app.app_name"))
	// 获取 viper.Viper 对象
	v := viper.GetViper()
	fmt.Println(v.GetString("app.app_name"))
}
