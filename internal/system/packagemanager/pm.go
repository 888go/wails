package packagemanager

// Package 包含有关系统包的信息
type Package struct {
	Name           string
	Version        string
	InstallCommand map[string]string
	SystemPackage  bool
	Library        bool
	Optional       bool
}

type packagemap = map[string][]*Package

// PackageManager 是所有包管理器通用的接口
type PackageManager interface {
	Name() string
	Packages() packagemap
	PackageInstalled(pkg *Package) (bool, error)
	PackageAvailable(pkg *Package) (bool, error)
	InstallCommand(pkg *Package) string
}

// Dependency 表示我们所需的系统包
type Dependency struct {
	Name           string
	PackageName    string
	Installed      bool
	InstallCommand string
	Version        string
	Optional       bool
	External       bool
}

// DependencyList 是一组 Dependency 实例的列表
type DependencyList []*Dependency

// InstallAllRequiredCommand 返回你需要使用的命令，以便安装所有必需的依赖项
func (d DependencyList) InstallAllRequiredCommand() string {
	result := ""
	for _, dependency := range d {
		if !dependency.Installed && !dependency.Optional {
			result += "  - " + dependency.Name + ": " + dependency.InstallCommand + "\n"
		}
	}

	return result
}

// InstallAllOptionalCommand 返回用于安装所有可选依赖项的命令
func (d DependencyList) InstallAllOptionalCommand() string {
	result := ""
	for _, dependency := range d {
		if !dependency.Installed && dependency.Optional {
			result += "  - " + dependency.Name + ": " + dependency.InstallCommand + "\n"
		}
	}

	return result
}
