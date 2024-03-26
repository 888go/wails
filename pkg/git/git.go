package git

import (
	"html/template"
	"runtime"
	"strings"

	"github.com/888go/wails/internal/shell"
)

func gitcommand() string {
	gitcommand := "git"
	if runtime.GOOS == "windows" {
		gitcommand = "git.exe"
	}

	return gitcommand
}

// IsInstalled在给定平台下，如果已安装git则返回true
func X是否已安装() bool {
	return shell.CommandExists(gitcommand())
}

// Email 尝试检索
func X取邮件地址() (string, error) {
	stdout, _, err := shell.RunCommand(".", gitcommand(), "config", "user.email")
	return stdout, err
}

// Name 尝试检索
func X取名称() (string, error) {
	stdout, _, err := shell.RunCommand(".", gitcommand(), "config", "user.name")
	name := template.JSEscapeString(strings.TrimSpace(stdout))
	return name, err
}

func InitRepo(projectDir string) error {
	_, _, err := shell.RunCommand(projectDir, gitcommand(), "init")
	return err
}
