package main

import (
	"github.com/zenfire-cn/commkit/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Init(&logger.Option{
		Path:    "logs/app.log",
		Level:   "debug",
		MaxSize: 10,
	})
	// logger
	zap.L().Info("success")
	// Sugar
	zap.S().Infof("%s success", "sugar")
}
