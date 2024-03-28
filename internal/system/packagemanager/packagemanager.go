//go:build linux
// +build linux

package packagemanager

import (
	"sort"
	"strings"

	"github.com/wailsapp/wails/v2/internal/shell"
)

// 一系列包管理器命令
var pmcommands = []string{
	"eopkg",
	"apt",
	"dnf",
	"pacman",
	"emerge",
	"zypper",
	"nix-env",
}

// Find 尝试查找系统包管理器
func Find(osid string) PackageManager {

	// Loop over pmcommands
	for _, pmname := range pmcommands {
		if shell.CommandExists(pmname) {
			return newPackageManager(pmname, osid)
		}
	}
	return nil
}

func newPackageManager(pmname string, osid string) PackageManager {
	switch pmname {
	case "eopkg":
		return NewEopkg(osid)
	case "apt":
		return NewApt(osid)
	case "dnf":
		return NewDnf(osid)
	case "pacman":
		return NewPacman(osid)
	case "emerge":
		return NewEmerge(osid)
	case "zypper":
		return NewZypper(osid)
	case "nix-env":
		return NewNixpkgs(osid)
	}
	return nil
}

// Dependencies 扫描系统以查找所需的依赖项
// 返回一个依赖项列表，包含已搜索的依赖项、它们是否被找到
// 以及它们是否已被安装
func Dependencies(p PackageManager) (DependencyList, error) {

	var dependencies DependencyList

	for name, packages := range p.Packages() {
		dependency := &Dependency{Name: name}
		for _, pkg := range packages {
			dependency.Optional = pkg.Optional
			dependency.External = !pkg.SystemPackage
			dependency.InstallCommand = p.InstallCommand(pkg)
			packageavailable, err := p.PackageAvailable(pkg)
			if err != nil {
				return nil, err
			}
			if packageavailable {
				dependency.Version = pkg.Version
				dependency.PackageName = pkg.Name
				installed, err := p.PackageInstalled(pkg)
				if err != nil {
					return nil, err
				}
				if installed {
					dependency.Installed = true
					dependency.Version = pkg.Version
					if !pkg.SystemPackage {
						dependency.Version = AppVersion(name)
					}
				} else {
					dependency.InstallCommand = p.InstallCommand(pkg)
				}
				break
			}
		}
		dependencies = append(dependencies, dependency)
	}

	// Sort dependencies
	sort.Slice(dependencies, func(i, j int) bool {
		return dependencies[i].Name < dependencies[j].Name
	})

	return dependencies, nil
}

// AppVersion 返回与给定包相关的应用程序的版本
func AppVersion(name string) string {

	if name == "gcc" {
		return gccVersion()
	}

	if name == "pkg-config" {
		return pkgConfigVersion()
	}

	if name == "npm" {
		return npmVersion()
	}

	if name == "docker" {
		return dockerVersion()
	}

	return ""

}

func gccVersion() string {

	var version string
	var err error

	// Try "-dumpfullversion"
	version, _, err = shell.RunCommand(".", "gcc", "-dumpfullversion")
	if err != nil {

// 尝试执行 -dumpversion 操作
// 我们忽略这个函数返回的错误，因为该函数并非用于检测应用程序是否存在，
// 而仅是为了确认我们能够获取到版本号
		dumpversion, _, err := shell.RunCommand(".", "gcc", "-dumpversion")
		if err == nil {
			version = dumpversion
		}
	}
	return strings.TrimSpace(version)
}

func pkgConfigVersion() string {
	version, _, _ := shell.RunCommand(".", "pkg-config", "--version")
	return strings.TrimSpace(version)
}

func npmVersion() string {
	version, _, _ := shell.RunCommand(".", "npm", "--version")
	return strings.TrimSpace(version)
}

func dockerVersion() string {
	version, _, _ := shell.RunCommand(".", "docker", "--version")
	version = strings.TrimPrefix(version, "Docker version ")
	version = strings.ReplaceAll(version, ", build ", " (")
	version = strings.TrimSpace(version) + ")"
	return version
}
