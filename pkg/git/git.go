package git

import (
	"html/template"
	"runtime"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

func gitcommand() string {
	gitcommand := "git"
	if runtime.GOOS == "windows" {
		gitcommand = "git.exe"
	}

	return gitcommand
}

// IsInstalled在给定平台下，如果已安装git则返回true

// ff:是否已安装
func IsInstalled() bool {
	return shell.CommandExists(gitcommand())
}

// Email 尝试检索

// ff:取邮件地址
func Email() (string, error) {
	stdout, _, err := shell.RunCommand(".", gitcommand(), "config", "user.email")
	return stdout, err
}

// Name 尝试检索

// ff:取名称
func Name() (string, error) {
	stdout, _, err := shell.RunCommand(".", gitcommand(), "config", "user.name")
	name := template.JSEscapeString(strings.TrimSpace(stdout))
	return name, err
}


// ff:
// projectDir:
func InitRepo(projectDir string) error {
	_, _, err := shell.RunCommand(projectDir, gitcommand(), "init")
	return err
}
