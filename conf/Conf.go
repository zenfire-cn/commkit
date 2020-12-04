package conf

import (
	"github.com/spf13/viper"
	"github.com/zenfire-cn/commkit/utility"
)

// 传入配置文件名
// 从可执行程序所在文件夹或者当前工作路径进行加载，优先从可执行程序所在文件夹加载
func Init(fileName string) error {
	viper.SetConfigFile(utility.FindConfigFile(fileName))
	return viper.ReadInConfig()
}

