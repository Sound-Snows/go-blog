package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	//LogSavePath 保存路径
	LogSavePath = "runtime/logs/"
	//LogSaveName 名称
	LogSaveName = "log"
	//LogFileExt 名称
	LogFileExt = "log"
	//TimeFormat 时间
	TimeFormat = "20060102"
)

//getLogFolePath 获取日志路径
func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

//getLogFileFullPath 获取完整路径
func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

//openLogFile is
func openLogFile(filePath string) *os.File {
	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("fail to OpenFile:%v", err)
	}
	return handle
}

//mkDir is
func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
