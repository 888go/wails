package logger

import (
	"fmt"
	"os"

	"github.com/888go/wails/pkg/logger"
)

// LogLevel 是 public LogLevel 的别名
type LogLevel = logger.LogLevel

// Logger 是一个实用工具，用于将消息记录到多个目标地点
type Logger struct {
	output         logger.Logger
	logLevel       LogLevel
	showLevelInLog bool
}

// New 创建一个新的 Logger。你可以传入任意数量的 `io.Writer` 实例作为日志的目标输出地

// ff:
// output:
func New(output logger.Logger) *Logger {
	if output == nil {
		output = logger.X创建并按默认()
	}
	result := &Logger{
		logLevel:       logger.X常量_日志级别_信息,
		showLevelInLog: true,
		output:         output,
	}

	return result
}

// CustomLogger 创建一个新的自定义日志器，在消息前打印名称/ID

// ff:
// name:
func (l *Logger) CustomLogger(name string) CustomLogger {
	return newcustomLogger(l, name)
}

// HideLogLevel 从每条日志记录的开头移除日志级别文本

// ff:
func (l *Logger) HideLogLevel() {
	l.showLevelInLog = true
}

// SetLogLevel 设置输出日志的最低级别

// ff:
// level:
func (l *Logger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

// Writeln 直接将内容写入输出，不带日志级别
// 在消息末尾追加回车符

// ff:
// message:
func (l *Logger) Writeln(message string) {
	l.output.X日志(message)
}

// Write 直接将内容写入输出，不带有日志级别

// ff:
// message:
func (l *Logger) Write(message string) {
	l.output.X日志(message)
}

// Print 直接将内容写入输出，不带有日志级别
// 在消息后附加一个回车符

// ff:
// message:
func (l *Logger) Print(message string) {
	l.Write(message)
}

// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。

// ff:
// args:
// format:
func (l *Logger) Trace(format string, args ...interface{}) {
	if l.logLevel <= logger.X常量_日志级别_追踪 {
		l.output.X日志追踪(fmt.Sprintf(format, args...))
	}
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。

// ff:
// args:
// format:
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.logLevel <= logger.X常量_日志级别_调试 {
		l.output.X日志调试(fmt.Sprintf(format, args...))
	}
}

// 信息级别日志记录。功能类似于 Sprintf。

// ff:
// args:
// format:
func (l *Logger) Info(format string, args ...interface{}) {
	if l.logLevel <= logger.X常量_日志级别_信息 {
		l.output.X日志信息(fmt.Sprintf(format, args...))
	}
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。

// ff:
// args:
// format:
func (l *Logger) Warning(format string, args ...interface{}) {
	if l.logLevel <= logger.X常量_日志级别_警告 {
		l.output.X日志警告(fmt.Sprintf(format, args...))
	}
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。

// ff:
// args:
// format:
func (l *Logger) Error(format string, args ...interface{}) {
	if l.logLevel <= logger.X常量_日志级别_错误 {
		l.output.X日志错误(fmt.Sprintf(format, args...))
	}
}

// Fatal级别日志记录。其工作方式类似于Sprintf。

// ff:
// args:
// format:
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.output.X日志致命(fmt.Sprintf(format, args...))
	os.Exit(1)
}
