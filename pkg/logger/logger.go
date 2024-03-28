package logger

import (
	"fmt"
	"strings"
)

// LogLevel 是一个无符号8位整型变量
type LogLevel uint8

const (
	// TRACE level
	TRACE LogLevel = 1 //hs:常量_日志级别_追踪     

	// DEBUG level logging
	DEBUG LogLevel = 2 //hs:常量_日志级别_调试     

	// INFO level logging
	INFO LogLevel = 3 //hs:常量_日志级别_信息     

	// WARNING level logging
	WARNING LogLevel = 4 //hs:常量_日志级别_警告     

	// ERROR level logging
	ERROR LogLevel = 5 //hs:常量_日志级别_错误     
)

var logLevelMap = map[string]LogLevel{
	"trace":   TRACE,
	"debug":   DEBUG,
	"info":    INFO,
	"warning": WARNING,
	"error":   ERROR,
}


// ff:字符串到日志级别
// LogLevel:
// input:日志级别
func StringToLogLevel(input string) (LogLevel, error) {
	result, ok := logLevelMap[strings.ToLower(input)]
	if !ok {
		return ERROR, fmt.Errorf("invalid log level: %s", input)
	}
	return result, nil
}

// Logger 指定了需要附加到 Wails 应用程序的日志器所需的方法
type Logger interface {
	Print(message string) //hs:日志     
	Trace(message string) //hs:日志追踪     
	Debug(message string) //hs:日志调试     
	Info(message string) //hs:日志信息     
	Warning(message string) //hs:日志警告     
	Error(message string) //hs:日志错误     
	Fatal(message string) //hs:日志致命     
}
