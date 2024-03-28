//go:build linux
// +build linux

package packagemanager

import (
	"os/exec"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// Dnf代表了Dnf管理器
type Dnf struct {
	name string
	osid string
}

// NewDnf 创建一个新的 Dnf 实例
func NewDnf(osid string) *Dnf {
	return &Dnf{
		name: "dnf",
		osid: osid,
	}
}

// Packages 返回 Wails 编译所需的库
// 在不同的发行版或版本中，这些库可能会有所不同
func (y *Dnf) Packages() packagemap {
	return packagemap{
		"libgtk-3": []*Package{
			{Name: "gtk3-devel", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: "webkit2gtk4.0-devel", SystemPackage: true, Library: true},
			{Name: "webkit2gtk3-devel", SystemPackage: true, Library: true},
			// {名称: "webkitgtk3-devel", 系统包: true, 库: true}, 
// 这段Go语言代码的注释描述了一个结构体或map的键值对，翻译成中文如下：
// 该元素表示一个软件包信息：
// 名称： "webkitgtk3-devel"，表示软件包的名称是webkitgtk3-devel
// 系统包： true，表示这是一个系统级别的软件包
// 库： true，表示此软件包提供了库文件（在安装后可供其他程序链接和调用）
		},
		"gcc": []*Package{
			{Name: "gcc-c++", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: "pkgconf-pkg-config", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: "npm", SystemPackage: true},
			{Name: "nodejs-npm", SystemPackage: true},
		},
		"upx": []*Package{
			{Name: "upx", SystemPackage: true, Optional: true},
		},
		"docker": []*Package{
			{
				SystemPackage: false,
				Optional:      true,
				InstallCommand: map[string]string{
					"centos": "Follow the guide: https://docs.docker.com/engine/install/centos/",
					"fedora": "Follow the guide: https://docs.docker.com/engine/install/fedora/",
				},
			},
			{Name: "moby-engine", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称
func (y *Dnf) Name() string {
	return y.name
}

// PackageInstalled 测试给定的包名是否已安装
func (y *Dnf) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "dnf", "info", "installed", pkg.Name)
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	splitoutput := strings.Split(stdout, "\n")
	for _, line := range splitoutput {
		if strings.HasPrefix(line, "Version") {
			splitline := strings.Split(line, ":")
			pkg.Version = strings.TrimSpace(splitline[1])
		}
	}

	return true, err
}

// PackageAvailable 测试给定的包是否可供安装
func (y *Dnf) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "dnf", "info", pkg.Name)
	// 我们添加一个空格以确保获得完整匹配，而非部分匹配
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}
	splitoutput := strings.Split(stdout, "\n")
	for _, line := range splitoutput {
		if strings.HasPrefix(line, "Version") {
			splitline := strings.Split(line, ":")
			pkg.Version = strings.TrimSpace(splitline[1])
		}
	}
	return true, nil
}

// InstallCommand 返回特定包管理器用于安装包的命令
func (y *Dnf) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[y.osid]
	}
	return "sudo dnf install " + pkg.Name
}

func (y *Dnf) getPackageVersion(pkg *Package, output string) {
	splitOutput := strings.Split(output, " ")
	if len(splitOutput) > 0 {
		pkg.Version = splitOutput[1]
	}
}
