package main

import (
	"github.com/zenfire-cn/commkit/logger"
	"go.uber.org/zap"
)

func main() {
	logger.Init(&logger.Option{
		"logs/app.log",
		0, 0, 0,
		"debug", false, true})
	// logger
	zap.L().Info("success")
	// Sugar
	zap.S().Infof("%s success", "sugar")
}
