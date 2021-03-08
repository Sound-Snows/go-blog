package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//Level 日志级别
type Level int

var (
	// F is
	F *os.File
	//DefaultPrefix is
	DefaultPrefix = ""
	//DefaultCallerDepth is
	DefaultCallerDepth = 2

	//logger is
	logger *log.Logger
	//logPrefix is
	logPrefix = ""
	//levelFlags is
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	//DEBUG 日志类型为DEBUG
	DEBUG Level = iota
	//INFO is 普通
	INFO
	//WARNING is 警告
	WARNING
	//ERROR is 错误
	ERROR
	//FATAL is 紧急
	FATAL
)

func inti() {
	filePath := getLogFileFullPath()
	F = openLogFile(filePath)
	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

//Debug is
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

//Info is
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

//Warn is
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

//Error is
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

//Fatal is
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}
