package runtime

import (
	"context"
	"fmt"

	"github.com/888go/wails/pkg/logger"
)

// LogPrint 打印一个级别为 Print 的消息
func LogPrint(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Print(message)
}

// LogTrace 打印一条 Trace 级别的消息
func LogTrace(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Trace(message)
}

// LogDebug 打印一条 Debug 级别的消息
func LogDebug(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Debug(message)
}

// LogInfo 打印一条 Info 级别的消息
func LogInfo(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Info(message)
}

// LogWarning 打印一条 Warning 级别的消息
func LogWarning(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Warning(message)
}

// LogError 打印一条 Error 级别的消息
func LogError(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Error(message)
}

// LogFatal 打印一条 Fatal 级别的消息
func LogFatal(ctx context.Context, message string) {
	myLogger := getLogger(ctx)
	myLogger.Fatal(message)
}

// LogPrintf 打印一个级别为 Print 的消息
func LogPrintf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Print(msg)
}

// LogTracef 打印一条 Trace 级别的消息
func LogTracef(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Trace(msg)
}

// LogDebugf 打印一条 Debug 级别的消息
func LogDebugf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Debug(msg)
}

// LogInfof 打印一条 Info 级别的消息
func LogInfof(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Info(msg)
}

// LogWarningf 打印一条 Warning 级别的消息
func LogWarningf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Warning(msg)
}

// LogErrorf 打印一条 Error 级别的消息
func LogErrorf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Error(msg)
}

// LogFatalf 打印一条Fatal级别消息
func LogFatalf(ctx context.Context, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	myLogger := getLogger(ctx)
	myLogger.Fatal(msg)
}

// LogSetLogLevel 设置日志级别
func LogSetLogLevel(ctx context.Context, level logger.LogLevel) {
	myLogger := getLogger(ctx)
	myLogger.SetLogLevel(level)
}
