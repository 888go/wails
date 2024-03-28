package logger

import (
	"os"
)

// DefaultLogger 是一个实用工具，用于将消息记录到多个目标地
type DefaultLogger struct{}

// NewDefaultLogger 创建一个新的 Logger。

// ff:创建并按默认
func NewDefaultLogger() Logger {
	return &DefaultLogger{}
}

// Print 函数的工作方式类似于 Sprintf。

// ff:日志
// message:消息
func (l *DefaultLogger) Print(message string) {
	println(message)
}

// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。

// ff:日志追踪
// message:消息
func (l *DefaultLogger) Trace(message string) {
	println("TRA | " + message)
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。

// ff:日志调试
// message:消息
func (l *DefaultLogger) Debug(message string) {
	println("DEB | " + message)
}

// 信息级别日志记录。功能类似于 Sprintf。

// ff:日志信息
// message:消息
func (l *DefaultLogger) Info(message string) {
	println("INF | " + message)
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。

// ff:日志警告
// message:消息
func (l *DefaultLogger) Warning(message string) {
	println("WAR | " + message)
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。

// ff:日志错误
// message:消息
func (l *DefaultLogger) Error(message string) {
	println("ERR | " + message)
}

// Fatal级别日志记录。其工作方式类似于Sprintf。

// ff:日志致命
// message:消息
func (l *DefaultLogger) Fatal(message string) {
	println("FAT | " + message)
	os.Exit(1)
}
