package runtime

import (
	"context"
	"fmt"

	"github.com/888go/wails/pkg/logger"
)

// LogPrint 打印一个级别为 Print 的消息
func X日志(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Print(消息)
}

// LogTrace 打印一条 Trace 级别的消息
func X日志追踪(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Trace(消息)
}

// LogDebug 打印一条 Debug 级别的消息
func X日志调试(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Debug(消息)
}

// LogInfo 打印一条 Info 级别的消息
func X日志信息(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Info(消息)
}

// LogWarning 打印一条 Warning 级别的消息
func X日志警告(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Warning(消息)
}

// LogError 打印一条 Error 级别的消息
func X日志错误(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Error(消息)
}

// LogFatal 打印一条 Fatal 级别的消息
func X日志致命(上下文 context.Context, 消息 string) {
	myLogger := getLogger(上下文)
	myLogger.Fatal(消息)
}

// LogPrintf 打印一个级别为 Print 的消息
func X日志消息F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Print(msg)
}

// LogTracef 打印一条 Trace 级别的消息
func X日志追踪消息F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Trace(msg)
}

// LogDebugf 打印一条 Debug 级别的消息
func X日志调试消息F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Debug(msg)
}

// LogInfof 打印一条 Info 级别的消息
func X日志信息F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Info(msg)
}

// LogWarningf 打印一条 Warning 级别的消息
func X日志警告F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Warning(msg)
}

// LogErrorf 打印一条 Error 级别的消息
func X日志错误F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Error(msg)
}

// LogFatalf 打印一条Fatal级别消息
func X日志致命F(上下文 context.Context, 格式化 string, 参数 ...interface{}) {
	msg := fmt.Sprintf(格式化, 参数...)
	myLogger := getLogger(上下文)
	myLogger.Fatal(msg)
}

// LogSetLogLevel 设置日志级别
func X设置日志级别(上下文 context.Context, 常量_日志级别 logger.LogLevel) {
	myLogger := getLogger(上下文)
	myLogger.SetLogLevel(常量_日志级别)
}
