//go:build linux
// +build linux

package packagemanager

import (
	"os/exec"
	"regexp"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// Pacman 表示 Pacman 包管理系统
type Pacman struct {
	name string
	osid string
}

// NewPacman 创建一个全新的 Pacman 实例

// ff:
// osid:
func NewPacman(osid string) *Pacman {
	return &Pacman{
		name: "pacman",
		osid: osid,
	}
}

// Packages 返回 Wails 编译所需的库
// 在不同的发行版或版本中，这些库可能会有所不同

// ff:
func (p *Pacman) Packages() packagemap {
	return packagemap{
		"libgtk-3": []*Package{
			{Name: "gtk3", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: "webkit2gtk", SystemPackage: true, Library: true},
		},
		"gcc": []*Package{
			{Name: "gcc", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: "pkgconf", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: "npm", SystemPackage: true},
		},
		"docker": []*Package{
			{Name: "docker", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称

// ff:
func (p *Pacman) Name() string {
	return p.name
}

// PackageInstalled 测试给定的包名是否已安装

// ff:
// pkg:
func (p *Pacman) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	stdout, _, err := shell.RunCommand(".", "pacman", "-Q", pkg.Name)
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	splitoutput := strings.Split(stdout, "\n")
	for _, line := range splitoutput {
		if strings.HasPrefix(line, pkg.Name) {
			splitline := strings.Split(line, " ")
			pkg.Version = strings.TrimSpace(splitline[1])
		}
	}

	return true, err
}

// PackageAvailable 测试给定的包是否可供安装

// ff:
// pkg:
func (p *Pacman) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}
	output, _, err := shell.RunCommand(".", "pacman", "-Si", pkg.Name)
	// 我们添加一个空格以确保获得完整匹配，而非部分匹配
	if err != nil {
		_, ok := err.(*exec.ExitError)
		if ok {
			return false, nil
		}
		return false, err
	}

	reg := regexp.MustCompile(`.*Version.*?:\s+(.*)`)
	matches := reg.FindStringSubmatch(output)
	pkg.Version = ""
	noOfMatches := len(matches)
	if noOfMatches > 1 {
		pkg.Version = strings.TrimSpace(matches[1])
	}

	return true, nil
}

// InstallCommand 返回特定包管理器用于安装包的命令

// ff:
// pkg:
func (p *Pacman) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[p.osid]
	}
	return "sudo pacman -S " + pkg.Name
}
