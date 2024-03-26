//go:build linux
// +build linux

package packagemanager

import (
	"encoding/json"
	"github.com/888go/wails/internal/shell"
)

// Nixpkgs 表示 Nixpkgs 管理器
type Nixpkgs struct {
	name string
	osid string
}

type NixPackageDetail struct {
	Name    string
	Pname   string
	Version string
}

var available map[string]NixPackageDetail

// NewNixpkgs 创建一个新的 Nixpkgs 实例
func NewNixpkgs(osid string) *Nixpkgs {
	available = map[string]NixPackageDetail{}

	return &Nixpkgs{
		name: "nixpkgs",
		osid: osid,
	}
}

// Packages 返回 Wails 编译所需的库
// 在不同的发行版或版本中，这些库可能会有所不同
func (n *Nixpkgs) Packages() packagemap {
	// 目前，仅支持检查默认通道。
	channel := "nixpkgs"
	if n.osid == "nixos" {
		channel = "nixos"
	}

	return packagemap{
		"libgtk-3": []*Package{
			{Name: channel + ".gtk3", SystemPackage: true, Library: true},
		},
		"libwebkit": []*Package{
			{Name: channel + ".webkitgtk", SystemPackage: true, Library: true},
		},
		"gcc": []*Package{
			{Name: channel + ".gcc", SystemPackage: true},
		},
		"pkg-config": []*Package{
			{Name: channel + ".pkg-config", SystemPackage: true},
		},
		"npm": []*Package{
			{Name: channel + ".nodejs", SystemPackage: true},
		},
		"upx": []*Package{
			{Name: channel + ".upx", SystemPackage: true, Optional: true},
		},
		"docker": []*Package{
			{Name: channel + ".docker", SystemPackage: true, Optional: true},
		},
		"nsis": []*Package{
			{Name: channel + ".nsis", SystemPackage: true, Optional: true},
		},
	}
}

// Name 返回包管理器的名称
func (n *Nixpkgs) Name() string {
	return n.name
}

// PackageInstalled 测试给定的包名是否已安装
func (n *Nixpkgs) PackageInstalled(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}

	stdout, _, err := shell.RunCommand(".", "nix-env", "--json", "-qA", pkg.Name)
	if err != nil {
		return false, nil
	}

	var attributes map[string]NixPackageDetail
	err = json.Unmarshal([]byte(stdout), &attributes)
	if err != nil {
		return false, err
	}

	// Did we get one?
	installed := false
	for attribute, detail := range attributes {
		if attribute == pkg.Name {
			installed = true
			pkg.Version = detail.Version
		}
		break
	}

	// 如果在NixOS系统上，该包可能通过系统配置进行安装，所以需要检查nix存储。
	detail, ok := available[pkg.Name]
	if !installed && n.osid == "nixos" && ok {
		cmd := "nix-store --query --requisites /run/current-system | cut -d- -f2- | sort | uniq | grep '^" + detail.Pname + "'"

		if pkg.Library {
			cmd += " | grep 'dev$'"
		}

		stdout, _, err = shell.RunCommand(".", "sh", "-c", cmd)
		if err != nil {
			return false, nil
		}

		if len(stdout) > 0 {
			installed = true
		}
	}

	return installed, nil
}

// PackageAvailable 测试给定的包是否可供安装
func (n *Nixpkgs) PackageAvailable(pkg *Package) (bool, error) {
	if pkg.SystemPackage == false {
		return false, nil
	}

	stdout, _, err := shell.RunCommand(".", "nix-env", "--json", "-qaA", pkg.Name)
	if err != nil {
		return false, nil
	}

	var attributes map[string]NixPackageDetail
	err = json.Unmarshal([]byte(stdout), &attributes)
	if err != nil {
		return false, err
	}

	// Grab first version.
	for attribute, detail := range attributes {
		pkg.Version = detail.Version
		available[attribute] = detail
		break
	}

	return len(pkg.Version) > 0, nil
}

// InstallCommand 返回特定包管理器用于安装包的命令
func (n *Nixpkgs) InstallCommand(pkg *Package) string {
	if pkg.SystemPackage == false {
		return pkg.InstallCommand[n.osid]
	}
	return "nix-env -iA " + pkg.Name
}
