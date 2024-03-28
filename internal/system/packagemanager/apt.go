//go:build linux
// +build linux

package packagemanager

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// Apt代表Apt管理器
type Apt struct {
	name string
	osid string
}

// NewApt 创建一个新的Apt实例

// ff:
// osid:
func NewApt(osid string) *Apt {
	return &Apt{
		name: "apt",
		osid: osid,
	}
}

// Packages 返回 Wails 编译所需的库
// 在不同的发行版或版本中，这些库可能会有所不同

// ff:
func (a *Apt) Packages() packagemap {
	return packagemap{
		"libgtk-3": []*Package{
			{Name: "libgtk-3-dev", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: "libwebkit2gtk-4.0-dev", SystemPackage: true, Library: true},
		},
		"gcc": []*Package{
			{Name: "build-essential", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: "pkg-config", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: "npm", SystemPackage: true},
		},
		"docker": []*Package{
			{Name: "docker.io", SystemPackage: true, Optional: true},
		},
		"nsis": []*Package{
			{Name: "nsis", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称

// ff:
func (a *Apt) Name() string {
	return a.name
}

// PackageInstalled 测试给定的包名是否已安装

// ff:
// pkg:
func (a *Apt) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	cmd := exec.Command("apt", "list", "-qq", pkg.Name)
	var stdo, stde bytes.Buffer
	cmd.Stdout = &stdo
	cmd.Stderr = &stde
	cmd.Env = append(os.Environ(), "LANGUAGE=en")
	err := cmd.Run()
	return strings.Contains(stdo.String(), "[installed]"), err
}

// PackageAvailable 测试给定的包是否可供安装

// ff:
// pkg:
func (a *Apt) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "apt", "list", "-qq", pkg.Name)
	// 我们添加一个空格以确保获得完整匹配，而非部分匹配
	output := a.removeEscapeSequences(stdout)
	installed := strings.HasPrefix(output, pkg.Name)
	a.getPackageVersion(pkg, output)
	return installed, err
}

// InstallCommand 返回特定包管理器用于安装包的命令

// ff:
// pkg:
func (a *Apt) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[a.osid]
	}
	return "sudo apt install " + pkg.Name
}

func (a *Apt) removeEscapeSequences(in string) string {
	escapechars, _ := regexp.Compile(`\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])`)
	return escapechars.ReplaceAllString(in, "")
}

func (a *Apt) getPackageVersion(pkg *Package, output string) {

	splitOutput := strings.Split(output, " ")
	if len(splitOutput) > 1 {
		pkg.Version = splitOutput[1]
	}
}
