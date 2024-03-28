//go:build linux
// +build linux

package packagemanager

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// Emerge 代表了 Emerge 包管理器
type Emerge struct {
	name string
	osid string
}

// NewEmerge 创建一个新的 Emerge 实例

// ff:
// osid:
func NewEmerge(osid string) *Emerge {
	return &Emerge{
		name: "emerge",
		osid: osid,
	}
}

// Packages 返回 Wails 编译所需的库
// 在不同的发行版或版本中，这些库可能会有所不同

// ff:
func (e *Emerge) Packages() packagemap {
	return packagemap{
		"libgtk-3": []*Package{
			{Name: "x11-libs/gtk+", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: "net-libs/webkit-gtk", SystemPackage: true, Library: true},
		},
		"gcc": []*Package{
			{Name: "sys-devel/gcc", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: "dev-util/pkgconf", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: "net-libs/nodejs", SystemPackage: true},
		},
		"docker": []*Package{
			{Name: "app-emulation/docker", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称

// ff:
func (e *Emerge) Name() string {
	return e.name
}

// PackageInstalled 测试给定的包名是否已安装

// ff:
// pkg:
func (e *Emerge) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "emerge", "-s", pkg.Name+"$")
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	regex := `.*\*\s+` + regexp.QuoteMeta(pkg.Name) + `\n(?:\S|\s)+?Latest version installed: (.*)`
	installedRegex := regexp.MustCompile(regex)
	matches := installedRegex.FindStringSubmatch(stdout)
	pkg.Version = ""
	noOfMatches := len(matches)
	installed := false
	if noOfMatches > 1 && matches[1] != "[ Not Installed ]" {
		installed = true
		pkg.Version = strings.TrimSpace(matches[1])
	}
	return installed, err
}

// PackageAvailable 测试给定的包是否可供安装

// ff:
// pkg:
func (e *Emerge) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "emerge", "-s", pkg.Name+"$")
	// 我们添加一个空格以确保获得完整匹配，而非部分匹配
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	installedRegex := regexp.MustCompile(`.*\*\s+` + regexp.QuoteMeta(pkg.Name) + `\n(?:\S|\s)+?Latest version available: (.*)`)
	matches := installedRegex.FindStringSubmatch(stdout)
	pkg.Version = ""
	noOfMatches := len(matches)
	available := false
	if noOfMatches > 1 {
		available = true
		pkg.Version = strings.TrimSpace(matches[1])
	}
	return available, nil
}

// InstallCommand 返回特定包管理器用于安装包的命令

// ff:
// pkg:
func (e *Emerge) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[e.osid]
	}
	return "sudo emerge " + pkg.Name
}
