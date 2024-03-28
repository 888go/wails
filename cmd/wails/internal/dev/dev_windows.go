//go:build windows
// +build windows

package dev

import (
	"bytes"
	"os/exec"
	"strconv"

	"github.com/wailsapp/wails/v2/cmd/wails/internal/logutils"
)

func setParentGID(_ *exec.Cmd) {}

func killProc(cmd *exec.Cmd, devCommand string) {
// 代码来源：https://stackoverflow.com/a/44551450
// 出于某种原因，在Windows系统上终止npm脚本执行时，仅使用取消命令无法正常退出
	if cmd != nil && cmd.Process != nil {
		kill := exec.Command("TASKKILL", "/T", "/F", "/PID", strconv.Itoa(cmd.Process.Pid))
		var errorBuffer bytes.Buffer
		var stdoutBuffer bytes.Buffer
		kill.Stderr = &errorBuffer
		kill.Stdout = &stdoutBuffer
		err := kill.Run()
		if err != nil {
			if err.Error() != "exit status 1" {
				println(stdoutBuffer.String())
				println(errorBuffer.String())
				logutils.LogRed("Error from '%s': %s", devCommand, err.Error())
			}
		}
	}
}
