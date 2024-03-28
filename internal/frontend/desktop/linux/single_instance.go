//go:build linux
// +build linux

package linux

import (
	"encoding/json"
	"github.com/godbus/dbus/v5"
	"github.com/888go/wails/pkg/options"
	"log"
	"os"
	"strings"
)

type dbusHandler func(string)


// ff:
// message:
func (f dbusHandler) SendMessage(message string) *dbus.Error {
	f(message)
	return nil
}


// ff:
// uniqueID:
func SetupSingleInstance(uniqueID string) {
	id := "wails_app_" + strings.ReplaceAll(strings.ReplaceAll(uniqueID, "-", "_"), ".", "_")

	dbusName := "org." + id + ".SingleInstance"
	dbusPath := "/org/" + id + "/SingleInstance"

	conn, err := dbus.ConnectSessionBus()
// 如果我们在建立连接或发送消息过程中遇到任何错误，我们将直接继续。
// 实际上不应该发生这样的情况，但以防万一。
	if err != nil {
		return
	}

	f := dbusHandler(func(message string) {
		var secondInstanceData options.SecondInstanceData

		err := json.Unmarshal([]byte(message), &secondInstanceData)
		if err == nil {
			secondInstanceBuffer <- secondInstanceData
		}
	})

	err = conn.Export(f, dbus.ObjectPath(dbusPath), dbusName)
	if err != nil {
		return
	}

	reply, err := conn.RequestName(dbusName, dbus.NameFlagDoNotQueue)
	if err != nil {
		return
	}

	// 如果名称已被占用，则尝试将参数发送给现有实例，如果没有成功，则启动新实例
// 这段Go语言代码的注释翻译成中文为：
// ```go
// 若名称已存在，则尝试将参数传递给已存在的实例；若无法成功传递，则直接启动新的实例
	if reply == dbus.RequestNameReplyExists {
		data := options.SecondInstanceData{
			Args: os.Args[1:],
		}
		data.WorkingDirectory, err = os.Getwd()
		if err != nil {
			log.Printf("Failed to get working directory: %v", err)
			return
		}

		serialized, err := json.Marshal(data)
		if err != nil {
			log.Printf("Failed to marshal data: %v", err)
			return
		}

		err = conn.Object(dbusName, dbus.ObjectPath(dbusPath)).Call(dbusName+".SendMessage", 0, string(serialized)).Store()
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
