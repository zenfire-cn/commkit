package conf

import (
	"github.com/spf13/viper"
	"github.com/zenfire-cn/commkit/utility"
	"log"
	"strings"
)

func Init(fileName string) {
	path := utility.ExePath(fileName)

	if path == "" {
		log.Fatal("Conf Error: Could Not Found Config File.")
	}
	path = strings.Replace(path, fileName, "", -1)

	split := strings.Split(fileName, ".")
	suffix := split[len(split)-1]
	viper.SetConfigName(strings.Replace(fileName, "."+suffix, "", -1))
	viper.SetConfigType(suffix)
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}

