package buildassets

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	iofs "io/fs"
	"os"
	"path/filepath"
	"text/template"

	"github.com/leaanthony/gosod"
	"github.com/samber/lo"
	"github.com/888go/wails/internal/fs"
	"github.com/888go/wails/internal/project"
)

//go:embed build
var assets embed.FS

// 与 assets 相同，但以 /build/ 为根目录进行操作
var buildAssets iofs.FS

func init() {
	buildAssets = lo.Must(iofs.Sub(assets, "build"))
}

// Install 将安装所有默认项目资源
func Install(targetDir string) error {
	templateDir := gosod.New(assets)
	err := templateDir.Extract(targetDir, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetLocalPath 返回请求构建资源文件的本地路径
func GetLocalPath(projectData *project.Project, file string) string {
	return filepath.Clean(filepath.Join(projectData.GetBuildDir(), filepath.FromSlash(file)))
}

// ReadFile 从项目构建文件夹中读取文件。
// 如果文件不存在，则回退到嵌入的文件，并将文件写入磁盘以便进行自定义。
func X读文件(projectData *project.Project, file string) ([]byte, error) {
	localFilePath := GetLocalPath(projectData, file)

	content, err := os.ReadFile(localFilePath)
	if errors.Is(err, iofs.ErrNotExist) {
		// 文件不存在，让我们从资源文件系统读取该文件并将其写入磁盘
		content, err := iofs.ReadFile(buildAssets, file)
		if err != nil {
			return nil, err
		}

		if err := writeFileSystemFile(projectData, file, content); err != nil {
			return nil, fmt.Errorf("Unable to create file in build folder: %s", err)
		}
		return content, nil
	}

	return content, err
}

// ReadFileWithProjectData 从项目构建文件夹读取文件，并在必要时替换 ProjectInfo。
// 如果文件不存在，则回退到嵌入的文件，该文件将被写入磁盘以便进行自定义。
// 写入的文件是原始未解析的文件。
func ReadFileWithProjectData(projectData *project.Project, file string) ([]byte, error) {
	content, err := X读文件(projectData, file)
	if err != nil {
		return nil, err
	}

	content, err = resolveProjectData(content, projectData)
	if err != nil {
		return nil, fmt.Errorf("Unable to resolve data in %s: %w", file, err)
	}
	return content, nil
}

// ReadOriginalFileWithProjectDataAndSave 从嵌入的资源中读取文件，并在必要时替换项目信息（ProjectInfo）。
// 同时，它还会将解析后的最终文件写回到项目的构建目录中。
func ReadOriginalFileWithProjectDataAndSave(projectData *project.Project, file string) ([]byte, error) {
	content, err := iofs.ReadFile(buildAssets, file)
	if err != nil {
		return nil, fmt.Errorf("Unable to read file %s: %w", file, err)
	}

	content, err = resolveProjectData(content, projectData)
	if err != nil {
		return nil, fmt.Errorf("Unable to resolve data in %s: %w", file, err)
	}

	if err := writeFileSystemFile(projectData, file, content); err != nil {
		return nil, fmt.Errorf("Unable to create file in build folder: %w", err)
	}
	return content, nil
}

type assetData struct {
	Name string
	Info project.Info
}

func resolveProjectData(content []byte, projectData *project.Project) ([]byte, error) {
	tmpl, err := template.New("").Parse(string(content))
	if err != nil {
		return nil, err
	}

	data := &assetData{
		Name: projectData.Name,
		Info: projectData.Info,
	}

	var out bytes.Buffer
	if err := tmpl.Execute(&out, data); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func writeFileSystemFile(projectData *project.Project, file string, content []byte) error {
	targetPath := GetLocalPath(projectData, file)

	if dir := filepath.Dir(targetPath); !fs.DirExists(dir) {
		if err := fs.MkDirs(dir, 0o755); err != nil {
			return fmt.Errorf("Unable to create directory: %w", err)
		}
	}

	if err := os.WriteFile(targetPath, content, 0o644); err != nil {
		return err
	}
	return nil
}
