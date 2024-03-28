package logger

import (
	"log"
	"os"
)

// FileLogger 是一个工具，用于将消息记录到多个目标中
type FileLogger struct {
	filename string
}

// NewFileLogger 创建一个新的 Logger。

// ff:创建文件日志
// filename:文件路径
func X创建文件日志(文件路径 string) Logger {
	return &FileLogger{
		filename: 文件路径,
	}
}

// Print 函数的工作方式类似于 Sprintf。

// ff:日志
// message:消息
func (l *FileLogger) X日志(消息 string) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(消息); err != nil {
		f.Close()
		log.Fatal(err)
	}
	f.Close()
}


// ff:日志并换行
// message:消息
func (l *FileLogger) X日志并换行(消息 string) {
	l.X日志(消息 + "\n")
}

// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。

// ff:日志追踪
// message:消息
func (l *FileLogger) X日志追踪(消息 string) {
	l.X日志并换行("TRACE | " + 消息)
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。

// ff:日志调试
// message:消息
func (l *FileLogger) X日志调试(消息 string) {
	l.X日志并换行("DEBUG | " + 消息)
}

// 信息级别日志记录。功能类似于 Sprintf。

// ff:日志信息
// message:消息
func (l *FileLogger) X日志信息(消息 string) {
	l.X日志并换行("INFO  | " + 消息)
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。

// ff:日志警告
// message:消息
func (l *FileLogger) X日志警告(消息 string) {
	l.X日志并换行("WARN  | " + 消息)
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。

// ff:日志错误
// message:消息
func (l *FileLogger) X日志错误(消息 string) {
	l.X日志并换行("ERROR | " + 消息)
}

// Fatal级别日志记录。其工作方式类似于Sprintf。

// ff:日志致命
// message:消息
func (l *FileLogger) X日志致命(消息 string) {
	l.X日志并换行("FATAL | " + 消息)
	os.Exit(1)
}
