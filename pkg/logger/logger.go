package logger

import (
	"fmt"
	"strings"
)

// LogLevel 是一个无符号8位整型变量
type LogLevel uint8

const (
	// TRACE level
	TRACE LogLevel = 1

	// DEBUG level logging
	DEBUG LogLevel = 2

	// INFO level logging
	INFO LogLevel = 3

	// WARNING level logging
	WARNING LogLevel = 4

	// ERROR level logging
	ERROR LogLevel = 5
)

var logLevelMap = map[string]LogLevel{
	"trace":   TRACE,
	"debug":   DEBUG,
	"info":    INFO,
	"warning": WARNING,
	"error":   ERROR,
}

func StringToLogLevel(input string) (LogLevel, error) {
	result, ok := logLevelMap[strings.ToLower(input)]
	if !ok {
		return ERROR, fmt.Errorf("invalid log level: %s", input)
	}
	return result, nil
}

// Logger 指定了需要附加到 Wails 应用程序的日志器所需的方法
type Logger interface {
	Print(message string)
	Trace(message string)
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
	Fatal(message string)
}
