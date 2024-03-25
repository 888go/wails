package logger

import (
	"fmt"
)

// CustomLogger 定义了用户可以对日志器执行的操作
type CustomLogger interface {
	// Writeln 直接将内容写入输出，不带有日志级别，并在末尾添加换行符
	Writeln(message string)

	// Write 直接将内容写入输出，不带有日志级别
	Write(message string)

	// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。
	Trace(format string, args ...interface{})

	// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。
	Debug(format string, args ...interface{})

	// 信息级别日志记录。功能类似于 Sprintf。
	Info(format string, args ...interface{})

	// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。
	Warning(format string, args ...interface{})

	// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。
	Error(format string, args ...interface{})

	// Fatal级别日志记录。其工作方式类似于Sprintf。
	Fatal(format string, args ...interface{})
}

// customLogger 是一个工具，用于将消息记录到多个目标中
type customLogger struct {
	logger *Logger
	name   string
}

// New 创建一个新的 customLogger。你可以传入任意数量的 `io.Writer` 对象，
// 这些对象将作为日志的目标输出地
func newcustomLogger(logger *Logger, name string) *customLogger {
	result := &customLogger{
		name:   name,
		logger: logger,
	}
	return result
}

// Writeln 直接将内容写入输出，不带日志级别
// 在消息末尾追加回车符

// ff:
// message:
func (l *customLogger) Writeln(message string) {
	l.logger.Writeln(message)
}

// Write 直接将内容写入输出，不带有日志级别

// ff:
// message:
func (l *customLogger) Write(message string) {
	l.logger.Write(message)
}

// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。

// ff:
// args:
// format:
func (l *customLogger) Trace(format string, args ...interface{}) {
	format = fmt.Sprintf("%s | %s", l.name, format)
	l.logger.Trace(format, args...)
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。

// ff:
// args:
// format:
func (l *customLogger) Debug(format string, args ...interface{}) {
	format = fmt.Sprintf("%s | %s", l.name, format)
	l.logger.Debug(format, args...)
}

// 信息级别日志记录。功能类似于 Sprintf。

// ff:
// args:
// format:
func (l *customLogger) Info(format string, args ...interface{}) {
	format = fmt.Sprintf("%s | %s", l.name, format)
	l.logger.Info(format, args...)
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。

// ff:
// args:
// format:
func (l *customLogger) Warning(format string, args ...interface{}) {
	format = fmt.Sprintf("%s | %s", l.name, format)
	l.logger.Warning(format, args...)
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。

// ff:
// args:
// format:
func (l *customLogger) Error(format string, args ...interface{}) {
	format = fmt.Sprintf("%s | %s", l.name, format)
	l.logger.Error(format, args...)
}

// Fatal级别日志记录。其工作方式类似于Sprintf。

// ff:
// args:
// format:
func (l *customLogger) Fatal(format string, args ...interface{}) {
	format = fmt.Sprintf("%s | %s", l.name, format)
	l.logger.Fatal(format, args...)
}
