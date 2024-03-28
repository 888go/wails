package logger

import (
	"fmt"
	"strings"
)

// LogLevel 是一个无符号8位整型变量
type LogLevel uint8

const (
	// TRACE level
	X常量_日志级别_追踪 LogLevel = 1 //hs:常量_日志级别_追踪     

	// DEBUG level logging
	X常量_日志级别_调试 LogLevel = 2 //hs:常量_日志级别_调试     

	// INFO level logging
	X常量_日志级别_信息 LogLevel = 3 //hs:常量_日志级别_信息     

	// WARNING level logging
	X常量_日志级别_警告 LogLevel = 4 //hs:常量_日志级别_警告     

	// ERROR level logging
	X常量_日志级别_错误 LogLevel = 5 //hs:常量_日志级别_错误     
)

var logLevelMap = map[string]LogLevel{
	"trace":   X常量_日志级别_追踪,
	"debug":   X常量_日志级别_调试,
	"info":    X常量_日志级别_信息,
	"warning": X常量_日志级别_警告,
	"error":   X常量_日志级别_错误,
}


// ff:字符串到日志级别
// LogLevel:
// input:日志级别
func X字符串到日志级别(日志级别 string) (LogLevel, error) {
	result, ok := logLevelMap[strings.ToLower(日志级别)]
	if !ok {
		return X常量_日志级别_错误, fmt.Errorf("invalid log level: %s", 日志级别)
	}
	return result, nil
}

// Logger 指定了需要附加到 Wails 应用程序的日志器所需的方法
type Logger interface {
	X日志(message string) //hs:日志     
	X日志追踪(message string) //hs:日志追踪     
	X日志调试(message string) //hs:日志调试     
	X日志信息(message string) //hs:日志信息     
	X日志警告(message string) //hs:日志警告     
	X日志错误(message string) //hs:日志错误     
	X日志致命(message string) //hs:日志致命     
}
