package clilogger

import (
	"fmt"
	"io"
	"os"

	"github.com/888go/wails/internal/colour"
)

// CLILogger 是被 cli 使用的
type CLILogger struct {
	Writer io.Writer
	mute   bool
}

// New cli logger

// writer:
func X创建(writer io.Writer) *CLILogger {
	return &CLILogger{
		Writer: writer,
	}
}

// Mute 设置是否应该禁用日志器

// ff:禁用日志
// value:禁用
func (c *CLILogger) Mute(value bool) {
	c.mute = value
}

// Print 函数类似于 Printf 函数

// args:
func (c *CLILogger) X日志输出(消息 string, args ...interface{}) {
	if c.mute {
		return
	}

	_, err := fmt.Fprintf(c.Writer, 消息, args...)
	if err != nil {
		c.X日志输出并停止("FATAL: " + err.Error())
	}
}

// Println 工作方式类似于 Printf，但在末尾添加换行符

// args:
func (c *CLILogger) X日志输出并换行(消息 string, args ...interface{}) {
	if c.mute {
		return
	}
	temp := fmt.Sprintf(消息, args...)
	_, err := fmt.Fprintln(c.Writer, temp)
	if err != nil {
		c.X日志输出并停止("FATAL: " + err.Error())
	}
}

// Fatal 打印给定的消息，然后中止程序

// args:
func (c *CLILogger) X日志输出并停止(消息 string, args ...interface{}) {
	temp := fmt.Sprintf(消息, args...)
	_, err := fmt.Fprintln(c.Writer, colour.Red("FATAL: "+temp))
	if err != nil {
		println(colour.Red("FATAL: " + err.Error()))
	}
	os.Exit(1)
}
