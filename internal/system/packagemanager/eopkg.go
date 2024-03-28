//go:build linux
// +build linux

package packagemanager

import (
	"regexp"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// Eopkg 表示 Eopkg 包管理系统
type Eopkg struct {
	name string
	osid string
}

// NewEopkg 创建一个新的 Eopkg 实例
func NewEopkg(osid string) *Eopkg {
	result := &Eopkg{
		name: "eopkg",
		osid: osid,
	}
	result.intialiseName()
	return result
}

// Packages 返回 Wails 编译所需的包
// 在不同的发行版或版本中，这些包可能会有所不同
func (e *Eopkg) Packages() packagemap {
	return packagemap{
		"libgtk-3": []*Package{
			{Name: "libgtk-3-devel", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: "libwebkit-gtk-devel", SystemPackage: true, Library: true},
		},
		"gcc": []*Package{
			{Name: "gcc", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: "pkg-config", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: "nodejs", SystemPackage: true},
		},
		"docker": []*Package{
			{Name: "docker", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称
func (e *Eopkg) Name() string {
	return e.name
}

// PackageInstalled 测试给定的包是否已安装
func (e *Eopkg) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "eopkg", "info", pkg.Name)
	return strings.HasPrefix(stdout, "Installed"), err
}

// PackageAvailable 测试给定的包是否可供安装
func (e *Eopkg) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "eopkg", "info", pkg.Name)
	// 我们添加一个空格以确保获得完整匹配，而非部分匹配
	output := e.removeEscapeSequences(stdout)
	installed := strings.Contains(output, "Package found in Solus repository")
	e.getPackageVersion(pkg, output)
	return installed, err
}

// InstallCommand 返回特定包管理器用于安装包的命令
func (e *Eopkg) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[e.osid]
	}
	return "sudo eopkg it " + pkg.Name
}

func (e *Eopkg) removeEscapeSequences(in string) string {
	escapechars, _ := regexp.Compile(`\x1B(?:[@-Z\\-_]|\[[0-?]*[ -/]*[@-~])`)
	return escapechars.ReplaceAllString(in, "")
}

func (e *Eopkg) intialiseName() {
	result := "eopkg"
	stdout, _, err := shell.RunCommand(".", "eopkg", "--version")
	if err == nil {
		result = strings.TrimSpace(stdout)
	}
	e.name = result
}

func (e *Eopkg) getPackageVersion(pkg *Package, output string) {

	versionRegex := regexp.MustCompile(`.*Name.*version:\s+(.*)+, release: (.*)`)
	matches := versionRegex.FindStringSubmatch(output)
	pkg.Version = ""
	noOfMatches := len(matches)
	if noOfMatches > 1 {
		pkg.Version = matches[1]
		if noOfMatches > 2 {
			pkg.Version += " (r" + matches[2] + ")"
		}
	}
}
