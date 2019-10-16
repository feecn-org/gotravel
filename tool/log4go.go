package tool

import (
	"github.com/wonderivan/logger"
)

func InitLogger(systemCode string) {

	var logConfigFile = ""
	if systemCode == "windows" {
		logConfigFile = LOG_FILE_PATH_WINDOWS
	} else if systemCode == "linux" {
		logConfigFile = LOG_FILE_PATH_LINUS
	}

	error := logger.SetLogger(logConfigFile)
	if error != nil {
		logger.Error("日志组件初始化失败%s", error)
	}
	logger.Debug("日志组件初始化")
}
