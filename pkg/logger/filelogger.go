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
func NewFileLogger(filename string) Logger {
	return &FileLogger{
		filename: filename,
	}
}

// Print 函数的工作方式类似于 Sprintf。

// ff:日志
// message:消息
func (l *FileLogger) Print(message string) {
	f, err := os.OpenFile(l.filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.WriteString(message); err != nil {
		f.Close()
		log.Fatal(err)
	}
	f.Close()
}


// ff:日志并换行
// message:消息
func (l *FileLogger) Println(message string) {
	l.Print(message + "\n")
}

// 以下是将该段Go语言代码注释翻译成中文：
// 跟踪级别日志记录。其工作方式类似于Sprintf（格式化字符串函数）。

// ff:日志追踪
// message:消息
func (l *FileLogger) Trace(message string) {
	l.Println("TRACE | " + message)
}

// 调试级别日志记录。其工作方式类似于 Sprintf（格式化字符串并写入）。

// ff:日志调试
// message:消息
func (l *FileLogger) Debug(message string) {
	l.Println("DEBUG | " + message)
}

// 信息级别日志记录。功能类似于 Sprintf。

// ff:日志信息
// message:消息
func (l *FileLogger) Info(message string) {
	l.Println("INFO  | " + message)
}

// 警告级别日志记录。其工作方式类似于 Sprintf（格式化字符串输出）。

// ff:日志警告
// message:消息
func (l *FileLogger) Warning(message string) {
	l.Println("WARN  | " + message)
}

// 错误级别日志记录。其工作方式类似于 Sprintf（格式化字符串并输出）。

// ff:日志错误
// message:消息
func (l *FileLogger) Error(message string) {
	l.Println("ERROR | " + message)
}

// Fatal级别日志记录。其工作方式类似于Sprintf。

// ff:日志致命
// message:消息
func (l *FileLogger) Fatal(message string) {
	l.Println("FATAL | " + message)
	os.Exit(1)
}
