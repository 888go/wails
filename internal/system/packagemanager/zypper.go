//go:build linux
// +build linux

package packagemanager

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// Zypper代表Zypper包管理器
type Zypper struct {
	name string
	osid string
}

// NewZypper 创建一个新的 Zypper 实例

// ff:
// osid:
func NewZypper(osid string) *Zypper {
	return &Zypper{
		name: "zypper",
		osid: osid,
	}
}

// Packages 返回 Wails 编译所需的库
// 在不同的发行版或版本中，这些库可能会有所不同

// ff:
func (z *Zypper) Packages() packagemap {
	return packagemap{
		"libgtk-3": []*Package{
			{Name: "gtk3-devel", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: "webkit2gtk3-soup2-devel", SystemPackage: true, Library: true},
			{Name: "webkit2gtk3-devel", SystemPackage: true, Library: true},
		},
		"gcc": []*Package{
			{Name: "gcc-c++", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: "pkg-config", SystemPackage: true},
			{Name: "pkgconf-pkg-config", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: "npm10", SystemPackage: true},
			{Name: "npm20", SystemPackage: true},
		},
		"docker": []*Package{
			{Name: "docker", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称

// ff:
func (z *Zypper) Name() string {
	return z.name
}

// PackageInstalled 测试给定的包名是否已安装

// ff:
// pkg:
func (z *Zypper) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	var env []string
	env = shell.SetEnv(env, "LANGUAGE", "en_US.utf-8")
	stdout, _, err := shell.RunCommandWithEnv(env, ".", "zypper", "info", pkg.Name)
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}
	reg := regexp.MustCompile(`.*Installed\s*:\s*(Yes)\s*`)
	matches := reg.FindStringSubmatch(stdout)
	pkg.Version = ""
	noOfMatches := len(matches)
	if noOfMatches > 1 {
		z.getPackageVersion(pkg, stdout)
	}
	return noOfMatches > 1, err
}

// PackageAvailable 测试给定的包是否可供安装

// ff:
// pkg:
func (z *Zypper) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	var env []string
	env = shell.SetEnv(env, "LANGUAGE", "en_US.utf-8")
	stdout, _, err := shell.RunCommandWithEnv(env, ".", "zypper", "info", pkg.Name)
	// 我们添加一个空格以确保获得完整匹配，而非部分匹配
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	available := strings.Contains(stdout, "Information for package")
	if available {
		z.getPackageVersion(pkg, stdout)
	}

	return available, nil
}

// InstallCommand 返回特定包管理器用于安装包的命令

// ff:
// pkg:
func (z *Zypper) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[z.osid]
	}
	return "sudo zypper in " + pkg.Name
}

func (z *Zypper) getPackageVersion(pkg *Package, output string) {

	reg := regexp.MustCompile(`.*Version.*:(.*)`)
	matches := reg.FindStringSubmatch(output)
	pkg.Version = ""
	noOfMatches := len(matches)
	if noOfMatches > 1 {
		pkg.Version = strings.TrimSpace(matches[1])
	}
}
