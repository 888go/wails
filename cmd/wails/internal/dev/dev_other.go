//go:build darwin || linux
// +build darwin linux

package dev

import (
	"os/exec"
	"syscall"

	"github.com/888go/wails/cmd/wails/internal/logutils"
	"golang.org/x/sys/unix"
)

func setParentGID(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true,
	}
}

func killProc(cmd *exec.Cmd, devCommand string) {
	if cmd == nil || cmd.Process == nil {
		return
	}

// 在 macOS BigSur 系统上遇到同样的问题
// 我正在使用 Vite，但我想这可能是 Node（npm）本身普遍存在的问题
// 此外，在经过多次编辑/重建循环后，任何异常关闭（如崩溃或按 CTRL-C）都可能使 Node 保持运行状态
// 来源：https://stackoverflow.com/a/29552044/14764450 （与上述 Windows 解决方案在同一页面）
// 未在 *nix 系统上进行测试
	pgid, err := syscall.Getpgid(cmd.Process.Pid)
	if err == nil {
		err := syscall.Kill(-pgid, unix.SIGTERM) // note the minus sign
		if err != nil {
			logutils.LogRed("Error from '%s' when attempting to kill the process: %s", devCommand, err.Error())
		}
	}
}
