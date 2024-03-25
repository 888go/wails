package bindings

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/samber/lo"
	"github.com/888go/wails/internal/colour"
	"github.com/888go/wails/internal/shell"
	"github.com/888go/wails/pkg/commands/buildtags"
)

// 用于生成绑定的选项
type Options struct {
	Filename         string
	Tags             []string
	ProjectDirectory string
	Compiler         string
	GoModTidy        bool
	TsPrefix         string
	TsSuffix         string
	TsOutputType     string
}

// GenerateBindings 为给定 Wails 项目目录生成绑定。如果未指定项目目录，则使用当前工作目录。
func GenerateBindings(options Options) (string, error) {
	filename, _ := lo.Coalesce(options.Filename, "wailsbindings")
	if runtime.GOOS == "windows" {
		filename += ".exe"
	}

	// 使用go build命令编译代码，添加-tags bindings标记，并将输出文件命名为bindings.exe
	tempDir := os.TempDir()
	filename = filepath.Join(tempDir, filename)

	workingDirectory, _ := lo.Coalesce(options.ProjectDirectory, lo.Must(os.Getwd()))

	var stdout, stderr string
	var err error

	tags := append(options.Tags, "bindings")
	genModuleTags := lo.Without(tags, "desktop", "production", "debug", "dev")
	tagString := buildtags.Stringify(genModuleTags)

	if options.GoModTidy {
		stdout, stderr, err = shell.RunCommand(workingDirectory, options.Compiler, "mod", "tidy")
		if err != nil {
			return stdout, fmt.Errorf("%s\n%s\n%s", stdout, stderr, err)
		}
	}

	stdout, stderr, err = shell.RunCommand(workingDirectory, options.Compiler, "build", "-tags", tagString, "-o", filename)
	if err != nil {
		return stdout, fmt.Errorf("%s\n%s\n%s", stdout, stderr, err)
	}

	if runtime.GOOS == "darwin" {
		// 删除隔离属性
		stdout, stderr, err = shell.RunCommand(workingDirectory, "xattr", "-rc", filename)
		if err != nil {
			return stdout, fmt.Errorf("%s\n%s\n%s", stdout, stderr, err)
		}
	}

	defer func() {
		// 尽最大努力移除临时文件
		_ = os.Remove(filename)
	}()

	// 根据实际情况设置环境变量
	env := os.Environ()
	env = shell.SetEnv(env, "tsprefix", options.TsPrefix)
	env = shell.SetEnv(env, "tssuffix", options.TsSuffix)
	env = shell.SetEnv(env, "tsoutputtype", options.TsOutputType)

	stdout, stderr, err = shell.RunCommandWithEnv(env, workingDirectory, filename)
	if err != nil {
		return stdout, fmt.Errorf("%s\n%s\n%s", stdout, stderr, err)
	}

	if stderr != "" {
		log.Println(colour.DarkYellow(stderr))
	}

	return stdout, nil
}
