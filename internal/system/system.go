package system

import (
	"os/exec"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
	"github.com/wailsapp/wails/v2/internal/system/operatingsystem"
	"github.com/wailsapp/wails/v2/internal/system/packagemanager"
)

var IsAppleSilicon bool

// Info 存储当前操作系统、包管理器以及所需依赖项的相关信息
type Info struct {
	OS           *operatingsystem.OS
	PM           packagemanager.PackageManager
	Dependencies packagemanager.DependencyList
}

// GetInfo 扫描系统以获取操作系统详情、
// 系统包管理器以及必需依赖项的状态。

// ff:
func GetInfo() (*Info, error) {
	var result Info
	err := result.discover()
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func checkNodejs() *packagemanager.Dependency {
	// Check for Nodejs
	output, err := exec.Command("node", "-v").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		if len(output) > 0 {
			version = strings.TrimSpace(strings.Split(string(output), "\n")[0])[1:]
		}
	}
	return &packagemanager.Dependency{
		Name:           "Nodejs",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://nodejs.org/en/download/",
		Version:        version,
		Optional:       false,
		External:       false,
	}
}

func checkNPM() *packagemanager.Dependency {
	// Check for npm
	output, err := exec.Command("npm", "-version").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		version = strings.TrimSpace(strings.Split(string(output), "\n")[0])
	}
	return &packagemanager.Dependency{
		Name:           "npm ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://nodejs.org/en/download/",
		Version:        version,
		Optional:       false,
		External:       false,
	}
}

func checkUPX() *packagemanager.Dependency {
	// Check for npm
	output, err := exec.Command("upx", "-V").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		version = strings.TrimSpace(strings.Split(string(output), "\n")[0])
	}
	return &packagemanager.Dependency{
		Name:           "upx ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://upx.github.io/",
		Version:        version,
		Optional:       true,
		External:       false,
	}
}

func checkNSIS() *packagemanager.Dependency {
	// 检查是否存在nsis安装程序
	output, err := exec.Command("makensis", "-VERSION").Output()
	installed := true
	version := ""
	if err != nil {
		installed = false
	} else {
		version = strings.TrimSpace(strings.Split(string(output), "\n")[0])
	}
	return &packagemanager.Dependency{
		Name:           "nsis ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "More info at https://wails.io/docs/guides/windows-installer/",
		Version:        version,
		Optional:       true,
		External:       false,
	}
}

func checkLibrary(name string) func() *packagemanager.Dependency {
	return func() *packagemanager.Dependency {
		output, _, _ := shell.RunCommand(".", "pkg-config", "--cflags", name)
		installed := len(strings.TrimSpace(output)) > 0

		return &packagemanager.Dependency{
			Name:           "lib" + name + " ",
			PackageName:    "N/A",
			Installed:      installed,
			InstallCommand: "Install via your package manager",
			Version:        "N/A",
			Optional:       false,
			External:       false,
		}
	}
}

func checkDocker() *packagemanager.Dependency {
	// Check for npm
	output, err := exec.Command("docker", "version").Output()
	installed := true
	version := ""

	// 如果Docker未运行，则会引发错误，因此需要检查这一点
	if len(output) == 0 && err != nil {
		installed = false
	} else {
		// Version 在一行中，其格式如：" 版本:           20.10.5"
		versionOutput := strings.Split(string(output), "\n")
		for _, line := range versionOutput[1:] {
			splitLine := strings.Split(line, ":")
			if len(splitLine) > 1 {
				key := strings.TrimSpace(splitLine[0])
				if key == "Version" {
					version = strings.TrimSpace(splitLine[1])
					break
				}
			}
		}
	}
	return &packagemanager.Dependency{
		Name:           "docker ",
		PackageName:    "N/A",
		Installed:      installed,
		InstallCommand: "Available at https://www.docker.com/products/docker-desktop",
		Version:        version,
		Optional:       true,
		External:       false,
	}
}
