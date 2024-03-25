package logger

import (
	"os"
)

// DefaultLogger 是一个实用工具，用于将消息记录到多个目标地
type DefaultLogger struct{}

// NewDefaultLogger 创建一个新的 Logger。
func X创建并按默认() Logger {
	return &DefaultLogger{}
}

// Print 函数的工作方式类似于 Sprintf。
func (l *DefaultLogger) X日志(消息 string) {
	println(消息)
}

// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。
func (l *DefaultLogger) X日志追踪(消息 string) {
	println("TRA | " + 消息)
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。
func (l *DefaultLogger) X日志调试(消息 string) {
	println("DEB | " + 消息)
}

// 信息级别日志记录。功能类似于 Sprintf。
func (l *DefaultLogger) X日志信息(消息 string) {
	println("INF | " + 消息)
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。
func (l *DefaultLogger) X日志警告(消息 string) {
	println("WAR | " + 消息)
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。
func (l *DefaultLogger) X日志错误(消息 string) {
	println("ERR | " + 消息)
}

// Fatal级别日志记录。其工作方式类似于Sprintf。
func (l *DefaultLogger) X日志致命(消息 string) {
	println("FAT | " + 消息)
	os.Exit(1)
}
