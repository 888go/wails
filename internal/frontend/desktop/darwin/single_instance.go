//go:build darwin
// +build darwin

package darwin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa
#import "AppDelegate.h"

*/
import "C"

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"

	"github.com/888go/wails/pkg/options"
)


// ff:
// uniqueID:
func SetupSingleInstance(uniqueID string) {
	lockFilePath := getTempDir()
	lockFileName := uniqueID + ".lock"
	_, err := createLockFile(lockFilePath + "/" + lockFileName)
	// if lockFile exist – send notification to second instance
	if err != nil {
		c := NewCalloc()
		defer c.Free()
		singleInstanceUniqueId := c.String(uniqueID)

		data, err := options.NewSecondInstanceData()
		if err != nil {
			return
		}

		serialized, err := json.Marshal(data)
		if err != nil {
			return
		}

		C.SendDataToFirstInstance(singleInstanceUniqueId, c.String(string(serialized)))

		os.Exit(0)
	}
}

//export HandleSecondInstanceData
// 导出HandleSecondInstanceData函数，以便在C语言或其他外部语言中调用
// （由于上下文不完整，无法提供更详尽的翻译，但大体上，Go语言中的`//export`注释是用来标记一个函数，表示该函数需要被cgo暴露给C代码或者其他使用CGO的环境，使得它们可以调用Go编写的这个函数。）

// ff:
// secondInstanceMessage:
func HandleSecondInstanceData(secondInstanceMessage *C.char) {
	message := C.GoString(secondInstanceMessage)

	var secondInstanceData options.SecondInstanceData

	err := json.Unmarshal([]byte(message), &secondInstanceData)
	if err == nil {
		secondInstanceBuffer <- secondInstanceData
	}
}

// CreateLockFile尝试使用给定的名称创建一个文件并获取其独占锁。
// 如果文件已存在且仍然被锁定，则操作将会失败。
func createLockFile(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		fmt.Printf("Failed to open lockfile %s: %s", filename, err)
		return nil, err
	}

	err = syscall.Flock(int(file.Fd()), syscall.LOCK_EX|syscall.LOCK_NB)
	if err != nil {
		// 如果 flock（文件锁定）由于除其他实例已锁定之外的其他原因失败，则将其打印到日志中，以便于可能的调试。
		if !strings.Contains(err.Error(), "resource temporarily unavailable") {
			fmt.Printf("Failed to lock lockfile %s: %s", filename, err)
		}
		file.Close()
		return nil, err
	}

	return file, nil
}

// 如果应用处于沙箱环境，golang 的 os.TempDir() 函数会返回一个无法访问的路径。因此，这里使用 macOS 原生的临时目录函数。
func getTempDir() string {
	cstring := C.GetMacOsNativeTempDir()
	path := C.GoString(cstring)
	C.free(unsafe.Pointer(cstring))

	return path
}
