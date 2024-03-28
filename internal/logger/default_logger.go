package logger

import (
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/logger"
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
		output = logger.NewDefaultLogger()
	}
	result := &Logger{
		logLevel:       logger.INFO,
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
	l.output.Print(message)
}

// Write 直接将内容写入输出，不带有日志级别

// ff:
// message:
func (l *Logger) Write(message string) {
	l.output.Print(message)
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
	if l.logLevel <= logger.TRACE {
		l.output.Trace(fmt.Sprintf(format, args...))
	}
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。

// ff:
// args:
// format:
func (l *Logger) Debug(format string, args ...interface{}) {
	if l.logLevel <= logger.DEBUG {
		l.output.Debug(fmt.Sprintf(format, args...))
	}
}

// 信息级别日志记录。功能类似于 Sprintf。

// ff:
// args:
// format:
func (l *Logger) Info(format string, args ...interface{}) {
	if l.logLevel <= logger.INFO {
		l.output.Info(fmt.Sprintf(format, args...))
	}
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。

// ff:
// args:
// format:
func (l *Logger) Warning(format string, args ...interface{}) {
	if l.logLevel <= logger.WARNING {
		l.output.Warning(fmt.Sprintf(format, args...))
	}
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。

// ff:
// args:
// format:
func (l *Logger) Error(format string, args ...interface{}) {
	if l.logLevel <= logger.ERROR {
		l.output.Error(fmt.Sprintf(format, args...))
	}
}

// Fatal级别日志记录。其工作方式类似于Sprintf。

// ff:
// args:
// format:
func (l *Logger) Fatal(format string, args ...interface{}) {
	l.output.Fatal(fmt.Sprintf(format, args...))
	os.Exit(1)
}
