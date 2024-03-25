package operatingsystem

import (
	"strings"

	"github.com/888go/wails/internal/shell"
)

func getSysctlValue(key string) (string, error) {
	stdout, _, err := shell.RunCommand(".", "sysctl", key)
	if err != nil {
		return "", err
	}
	version := strings.TrimPrefix(stdout, key+": ")
	return strings.TrimSpace(version), nil
}

func platformInfo() (*OS, error) {
	// Default value
	var result OS
	result.ID = "Unknown"
	result.Name = "MacOS"
	result.Version = "Unknown"

	version, err := getSysctlValue("kern.osproductversion")
	if err != nil {
		return nil, err
	}
	result.Version = version
	ID, err := getSysctlValue("kern.osversion")
	if err != nil {
		return nil, err
	}
	result.ID = ID

// 创建一个命令，其中directory为工作目录，command为要执行的命令，args...为传递给命令的参数列表
// 创建两个缓冲区stdo和stde，用于存储命令的标准输出和标准错误信息
// 将cmd的标准输出重定向到stdo缓冲区，标准错误重定向到stde缓冲区
// 执行命令，并将可能的错误信息保存在err变量中
// 返回stdo缓冲区的内容（即命令的标准输出）、stde缓冲区的内容（即命令的标准错误）以及可能发生的错误err
// 结束当前函数
// 初始化一个新的shell命令，该命令为"sysctl"
// 下面两行是示例输出，表示系统类型为Darwin，操作系统版本为20.1.0，内核修订版本为199506
// 注意：最后一行`kern.osrevision: 199506` 应该也是作为`sysctl`命令运行后的输出，而非代码部分。

	return &result, nil
}
