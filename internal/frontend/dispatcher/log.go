package dispatcher

import (
	"github.com/pkg/errors"
	"github.com/888go/wails/internal/logger"
	pkgLogger "github.com/888go/wails/pkg/logger"
)

var logLevelMap = map[byte]logger.LogLevel{
	'1': pkgLogger.X常量_日志级别_追踪,
	'2': pkgLogger.X常量_日志级别_调试,
	'3': pkgLogger.X常量_日志级别_信息,
	'4': pkgLogger.X常量_日志级别_警告,
	'5': pkgLogger.X常量_日志级别_错误,
}

func (d *Dispatcher) processLogMessage(message string) (string, error) {
	if len(message) < 3 {
		return "", errors.New("Invalid Log Message: " + message)
	}

	messageText := message[2:]

	switch message[1] {
	case 'T':
		d.log.Trace(messageText)
	case 'P':
		d.log.Print(messageText)
	case 'D':
		d.log.Debug(messageText)
	case 'I':
		d.log.Info(messageText)
	case 'W':
		d.log.Warning(messageText)
	case 'E':
		d.log.Error(messageText)
	case 'F':
		d.log.Fatal(messageText)
	case 'S':
		loglevel, exists := logLevelMap[message[2]]
		if !exists {
			return "", errors.New("Invalid Set Log Level Message: " + message)
		}
		d.log.SetLogLevel(loglevel)
	default:
		return "", errors.New("Invalid Log Message: " + message)
	}
	return "", nil
}
