package project

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/samber/lo"
)

// Project 结构体持有与Wails项目相关联的数据
type Project struct {
	/*** Application Data ***/
	Name           string `json:"name"`
	AssetDirectory string `json:"assetdir,omitempty"`

	ReloadDirectories string `json:"reloaddirs,omitempty"`

	BuildCommand   string `json:"frontend:build"`
	InstallCommand string `json:"frontend:install"`

	// 以下命令在 `wails dev` 中使用
	DevCommand        string `json:"frontend:dev"`
	DevBuildCommand   string `json:"frontend:dev:build"`
	DevInstallCommand string `json:"frontend:dev:install"`
	DevWatcherCommand string `json:"frontend:dev:watcher"`
	// 外部Wails开发服务器的URL。如果设置了这个值，那么将使用这个服务器作为前端服务。默认为 ""
	FrontendDevServerURL string `json:"frontend:dev:serverUrl"`

	// 用于生成API模块的目录
	WailsJSDir string `json:"wailsjsdir"`

	Version string `json:"version"`

	/*** Internal Data ***/

	// 项目目录的路径
	Path string `json:"projectdir"`

	// Build directory
	BuildDir string `json:"build:dir"`

	// The output filename
	OutputFilename string `json:"outputfilename"`

	// 应用程序的类型。例如：桌面应用、服务器应用等
	OutputType string

	// The platform to target
	Platform string

	// RunNonNativeBuildHooks 将运行构建钩子，即使这些钩子是为与主机操作系统不同的 GOOS 环境定义的。
	RunNonNativeBuildHooks bool `json:"runNonNativeBuildHooks"`

// 为不同目标构建钩子，这些钩子按照以下顺序执行：
// Key: GOOS/GOARCH - 在特定平台和架构的构建级别上，在构建前后执行
// Key: GOOS/*      - 在特定平台的构建级别上，在构建前后执行
// Key: */*         - 在构建级别上，在所有构建前后执行
// 下列键目前还不支持。
// Key: GOOS        - 在特定平台级别上，在该平台的所有构建之前/之后执行
// Key: *           - 在平台级别上，在所有平台的所有构建之前/之后执行
// Key: [空]        - 在全局级别上，在所有平台的所有构建之前/之后执行
	PostBuildHooks map[string]string `json:"postBuildHooks"`
	PreBuildHooks  map[string]string `json:"preBuildHooks"`

	// The application author
	Author Author

	// 应用程序信息
	Info Info

	// 完全限定文件名
	filename string

	// 内置开发服务器热重载的防抖时间，默认为100
// （注：debounce 时间通常是指在连续触发事件后，会等待一段固定的时间再去执行回调函数，用于限制函数在一定时间段内只能被执行一次，从而避免短时间内大量无用计算或网络请求。这里的“热重载”一般指的是当代码发生变化时，自动重新加载并应用更改到运行中的程序。）
	DebounceMS int `json:"debounceMS"`

	// 绑定wails开发服务器的地址。默认为 "localhost:34115"
	DevServer string `json:"devServer"`

	// 在开发模式下传递给应用程序的参数
	AppArgs string `json:"appargs"`

	// NSISType to be build
	NSISType string `json:"nsisType"`

	// Garble
	Obfuscated bool   `json:"obfuscated"`
	GarbleArgs string `json:"garbleargs"`

	// Frontend directory
	FrontendDir string `json:"frontend:dir"`

	Bindings Bindings `json:"bindings"`
}


// ff:
func (p *Project) GetFrontendDir() string {
	if filepath.IsAbs(p.FrontendDir) {
		return p.FrontendDir
	}
	return filepath.Join(p.Path, p.FrontendDir)
}


// ff:
func (p *Project) GetWailsJSDir() string {
	if filepath.IsAbs(p.WailsJSDir) {
		return p.WailsJSDir
	}
	return filepath.Join(p.Path, p.WailsJSDir)
}


// ff:
func (p *Project) GetBuildDir() string {
	if filepath.IsAbs(p.BuildDir) {
		return p.BuildDir
	}
	return filepath.Join(p.Path, p.BuildDir)
}


// ff:
func (p *Project) GetDevBuildCommand() string {
	if p.DevBuildCommand != "" {
		return p.DevBuildCommand
	}
	if p.DevCommand != "" {
		return p.DevCommand
	}
	return p.BuildCommand
}


// ff:
func (p *Project) GetDevInstallerCommand() string {
	if p.DevInstallCommand != "" {
		return p.DevInstallCommand
	}
	return p.InstallCommand
}


// ff:
func (p *Project) IsFrontendDevServerURLAutoDiscovery() bool {
	return p.FrontendDevServerURL == "auto"
}


// ff:
func (p *Project) Save() error {
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p.filename, data, 0o755)
}

func (p *Project) setDefaults() {
	if p.Path == "" {
		p.Path = lo.Must(os.Getwd())
	}
	if p.Version == "" {
		p.Version = "2"
	}
	// 如果未提供名称，则创建默认名称
	if p.Name == "" {
		p.Name = "wailsapp"
	}
	if p.OutputFilename == "" {
		p.OutputFilename = p.Name
	}
	if p.FrontendDir == "" {
		p.FrontendDir = "frontend"
	}
	if p.WailsJSDir == "" {
		p.WailsJSDir = p.FrontendDir
	}
	if p.BuildDir == "" {
		p.BuildDir = "build"
	}
	if p.DebounceMS == 0 {
		p.DebounceMS = 100
	}
	if p.DevServer == "" {
		p.DevServer = "localhost:34115"
	}
	if p.NSISType == "" {
		p.NSISType = "multiple"
	}
	if p.Info.CompanyName == "" {
		p.Info.CompanyName = p.Name
	}
	if p.Info.ProductName == "" {
		p.Info.ProductName = p.Name
	}
	if p.Info.ProductVersion == "" {
		p.Info.ProductVersion = "1.0.0"
	}
	if p.Info.Copyright == nil {
		v := "Copyright........."
		p.Info.Copyright = &v
	}
	if p.Info.Comments == nil {
		v := "Built using Wails (https://wails.io)"
		p.Info.Comments = &v
	}

	// Fix up OutputFilename
	switch runtime.GOOS {
	case "windows":
		if !strings.HasSuffix(p.OutputFilename, ".exe") {
			p.OutputFilename += ".exe"
		}
	case "darwin", "linux":
		p.OutputFilename = strings.TrimSuffix(p.OutputFilename, ".exe")
	}
}

// Author 用于存储应用程序作者的详细信息
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Info struct {
	CompanyName      string            `json:"companyName"`
	ProductName      string            `json:"productName"`
	ProductVersion   string            `json:"productVersion"`
	Copyright        *string           `json:"copyright"`
	Comments         *string           `json:"comments"`
	FileAssociations []FileAssociation `json:"fileAssociations"`
	Protocols        []Protocol        `json:"protocols"`
}

type FileAssociation struct {
	Ext         string `json:"ext"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IconName    string `json:"iconName"`
	Role        string `json:"role"`
}

type Protocol struct {
	Scheme      string `json:"scheme"`
	Description string `json:"description"`
	Role        string `json:"role"`
}

type Bindings struct {
	TsGeneration TsGeneration `json:"ts_generation"`
}

type TsGeneration struct {
	Prefix     string `json:"prefix"`
	Suffix     string `json:"suffix"`
	OutputType string `json:"outputType"`
}

// 将给定的JSON数据解析为一个Project结构体

// ff:
// projectData:
func Parse(projectData []byte) (*Project, error) {
	project := &Project{}
	err := json.Unmarshal(projectData, project)
	if err != nil {
		return nil, err
	}
	project.setDefaults()
	return project, nil
}

// 从当前工作目录加载项目

// ff:
// projectPath:
func Load(projectPath string) (*Project, error) {
	projectFile := filepath.Join(projectPath, "wails.json")
	rawBytes, err := os.ReadFile(projectFile)
	if err != nil {
		return nil, err
	}
	result, err := Parse(rawBytes)
	if err != nil {
		return nil, err
	}
	result.filename = projectFile
	return result, nil
}
