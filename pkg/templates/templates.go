package templates

import (
	"embed"
	"encoding/json"
	"fmt"
	gofs "io/fs"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/pkg/errors"

	"github.com/leaanthony/debme"
	"github.com/leaanthony/gosod"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/pkg/clilogger"
)

//go:embed all:templates
var templates embed.FS

//go:embed all:ides/*
var ides embed.FS

// 模板缓存
// 我们使用这个是因为我们需要对相同数据有不同的视图
var templateCache []Template = nil

// Data 包含我们在模板安装期间希望嵌入的数据
type Data struct {
	ProjectName        string
	BinaryName         string
	WailsVersion       string
	NPMProjectName     string
	AuthorName         string
	AuthorEmail        string
	AuthorNameAndEmail string
	WailsDirectory     string
	GoSDKPath          string
	WindowsFlags       string
	CGOEnabled         string
	OutputFile         string
}

// 用于安装模板的选项
type Options struct {
	ProjectName         string
	TemplateName        string
	BinaryName          string
	TargetDir           string
	Logger              *clilogger.CLILogger
	PathToDesktopBinary string
	PathToServerBinary  string
	InitGit             bool
	AuthorName          string
	AuthorEmail         string
	IDE                 string
	ProjectNameFilename string // 作为合法文件名的项目名称
	WailsVersion        string
	GoSDKPath           string
	WindowsFlags        string
	CGOEnabled          string
	CGOLDFlags          string
	OutputFile          string
}

// Template 用于存储与模板相关的数据，
// 包括保存在template.json中的元数据
type Template struct {
	// Template details
	Name        string `json:"name"`
	ShortName   string `json:"shortname"`
	Author      string `json:"author"`
	Description string `json:"description"`
	HelpURL     string `json:"helpurl"`

	// Other data
	FS gofs.FS `json:"-"`
}

func parseTemplate(template gofs.FS) (Template, error) {
	var result Template
	data, err := gofs.ReadFile(template, "template.json")
	if err != nil {
		return result, errors.Wrap(err, "Error parsing template")
	}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, err
	}
	result.FS = template
	return result, nil
}

// List 返回可用模板的列表
func List() ([]Template, error) {
	// 如果缓存未加载，则加载它
	if templateCache == nil {
		err := loadTemplateCache()
		if err != nil {
			return nil, err
		}
	}

	return templateCache, nil
}

// getTemplateByShortname 根据给定的短名称返回模板
func getTemplateByShortname(shortname string) (Template, error) {
	var result Template

	// 如果缓存未加载，则加载它
	if templateCache == nil {
		err := loadTemplateCache()
		if err != nil {
			return result, err
		}
	}

	for _, template := range templateCache {
		if template.ShortName == shortname {
			return template, nil
		}
	}

	return result, fmt.Errorf("shortname '%s' is not a valid template shortname", shortname)
}

// 加载模板缓存
func loadTemplateCache() error {
	templatesFS, err := debme.FS(templates, "templates")
	if err != nil {
		return err
	}

	// Get directories
	files, err := templatesFS.ReadDir(".")
	if err != nil {
		return err
	}

	// Reset cache
	templateCache = []Template{}

	for _, file := range files {
		if file.IsDir() {
			templateFS, err := templatesFS.FS(file.Name())
			if err != nil {
				return err
			}
			template, err := parseTemplate(templateFS)
			if err != nil {
				// 无法解析此模板，继续
				continue
			}
			templateCache = append(templateCache, template)
		}
	}

	return nil
}

// 安装给定的模板。如果模板为远程，则返回 true。
func Install(options *Options) (bool, *Template, error) {
	// Get cwd
	cwd, err := os.Getwd()
	if err != nil {
		return false, nil, err
	}

	// 用户是否希望在当前目录进行安装？
	if options.TargetDir == "" {
		options.TargetDir = filepath.Join(cwd, options.ProjectName)
		if fs.DirExists(options.TargetDir) {
			return false, nil, fmt.Errorf("cannot create project directory. Dir exists: %s", options.TargetDir)
		}
	} else {
		// 获取给定目录的绝对路径
		targetDir, err := filepath.Abs(options.TargetDir)
		if err != nil {
			return false, nil, err
		}
		options.TargetDir = targetDir
		if !fs.DirExists(options.TargetDir) {
			err := fs.Mkdir(options.TargetDir)
			if err != nil {
				return false, nil, err
			}
		}
	}

	// 标志表示远程模板
	remoteTemplate := false

	// Is this a shortname?
	template, err := getTemplateByShortname(options.TemplateName)
	if err != nil {
		// Is this a filepath?
		templatePath, err := filepath.Abs(options.TemplateName)
		if fs.DirExists(templatePath) {
			templateFS := os.DirFS(templatePath)
			template, err = parseTemplate(templateFS)
			if err != nil {
				return false, nil, errors.Wrap(err, "Error installing template")
			}
		} else {
			// 将git仓库克隆到临时目录中
			tempdir, err := gitclone(options)
			defer func(path string) {
				err := os.RemoveAll(path)
				if err != nil {
					log.Fatal(err)
				}
			}(tempdir)
			if err != nil {
				return false, nil, err
			}
			// 移除.git目录
			err = os.RemoveAll(filepath.Join(tempdir, ".git"))
			if err != nil {
				return false, nil, err
			}

			templateFS := os.DirFS(tempdir)
			template, err = parseTemplate(templateFS)
			if err != nil {
				return false, nil, err
			}
			remoteTemplate = true
		}
	}

	// 使用Gosod安装模板
	installer := gosod.New(template.FS)

	// 忽略 template.json 文件
	installer.IgnoreFile("template.json")

// 设置数据
// 我们使用目录名称作为二进制文件名，如同Go的风格
	BinaryName := filepath.Base(options.TargetDir)
	NPMProjectName := strings.ToLower(strings.ReplaceAll(BinaryName, " ", ""))
	localWailsDirectory := fs.RelativePath("../../../../../..")

	templateData := &Data{
		ProjectName:    options.ProjectName,
		BinaryName:     filepath.Base(options.TargetDir),
		NPMProjectName: NPMProjectName,
		WailsDirectory: localWailsDirectory,
		AuthorEmail:    options.AuthorEmail,
		AuthorName:     options.AuthorName,
		WailsVersion:   options.WailsVersion,
		GoSDKPath:      options.GoSDKPath,
	}

	// 创建一个格式化的名字和邮箱组合。
	if options.AuthorName != "" {
		templateData.AuthorNameAndEmail = options.AuthorName + " "
	}
	if options.AuthorEmail != "" {
		templateData.AuthorNameAndEmail += "<" + options.AuthorEmail + ">"
	}
	templateData.AuthorNameAndEmail = strings.TrimSpace(templateData.AuthorNameAndEmail)

	installer.RenameFiles(map[string]string{
		"gitignore.txt": ".gitignore",
	})

	// Extract the template
	err = installer.Extract(options.TargetDir, templateData)
	if err != nil {
		return false, nil, err
	}

	err = generateIDEFiles(options)
	if err != nil {
		return false, nil, err
	}

	return remoteTemplate, &template, nil
}

// 深度复制给定的uri，并返回临时克隆的目录
func gitclone(options *Options) (string, error) {
	// 创建临时目录
	dirname, err := os.MkdirTemp("", "wails-template-*")
	if err != nil {
		return "", err
	}

	// 解析远程模板URL和版本号
	templateInfo := strings.Split(options.TemplateName, "@")
	cloneOption := &git.CloneOptions{
		URL: templateInfo[0],
	}
	if len(templateInfo) > 1 {
		cloneOption.ReferenceName = plumbing.NewTagReferenceName(templateInfo[1])
	}

	_, err = git.PlainClone(dirname, false, cloneOption)

	return dirname, err
}

func generateIDEFiles(options *Options) error {
	switch options.IDE {
	case "vscode":
		return generateVSCodeFiles(options)
	case "goland":
		return generateGolandFiles(options)
	}

	return nil
}

type ideOptions struct {
	name         string
	targetDir    string
	options      *Options
	renameFiles  map[string]string
	ignoredFiles []string
}

func generateGolandFiles(options *Options) error {
	ideoptions := ideOptions{
		name:      "goland",
		targetDir: filepath.Join(options.TargetDir, ".idea"),
		options:   options,
		renameFiles: map[string]string{
			"projectname.iml": options.ProjectNameFilename + ".iml",
			"gitignore.txt":   ".gitignore",
			"name":            ".name",
		},
	}
	if !options.InitGit {
		ideoptions.ignoredFiles = []string{"vcs.xml"}
	}
	err := installIDEFiles(ideoptions)
	if err != nil {
		return errors.Wrap(err, "generating Goland IDE files")
	}

	return nil
}

func generateVSCodeFiles(options *Options) error {
	ideoptions := ideOptions{
		name:      "vscode",
		targetDir: filepath.Join(options.TargetDir, ".vscode"),
		options:   options,
	}
	return installIDEFiles(ideoptions)
}

func installIDEFiles(o ideOptions) error {
	source, err := debme.FS(ides, "ides/"+o.name)
	if err != nil {
		return err
	}

	// 使用Gosod安装模板
	installer := gosod.New(source)

	if o.renameFiles != nil {
		installer.RenameFiles(o.renameFiles)
	}

	for _, ignoreFile := range o.ignoredFiles {
		installer.IgnoreFile(ignoreFile)
	}

	binaryName := filepath.Base(o.options.TargetDir)
	o.options.WindowsFlags = ""
	o.options.CGOEnabled = "1"

	switch runtime.GOOS {
	case "windows":
		binaryName += ".exe"
		o.options.WindowsFlags = " -H windowsgui"
		o.options.CGOEnabled = "0"
	case "darwin":
		o.options.CGOLDFlags = "-framework UniformTypeIdentifiers"
	}

	o.options.PathToDesktopBinary = filepath.ToSlash(filepath.Join("build", "bin", binaryName))

	err = installer.Extract(o.targetDir, o.options)
	if err != nil {
		return err
	}

	return nil
}
