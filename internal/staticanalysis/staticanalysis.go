package staticanalysis

import (
	"go/ast"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
)

type EmbedDetails struct {
	BaseDir   string
	EmbedPath string
	All       bool
}


// ff:
func (e *EmbedDetails) GetFullPath() string {
	return filepath.Join(e.BaseDir, e.EmbedPath)
}


// ff:
// sourcePath:
func GetEmbedDetails(sourcePath string) ([]*EmbedDetails, error) {
// 读取项目文件，并确定用于嵌入的目录
// 返回一个目录列表

	absPath, err := filepath.Abs(sourcePath)
	if err != nil {
		return nil, err
	}
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedCompiledGoFiles,
		Dir:  absPath,
	}, "./...")
	if err != nil {
		return nil, err
	}
	var result []*EmbedDetails
	for _, pkg := range pkgs {
		for index, file := range pkg.Syntax {
			baseDir := filepath.Dir(pkg.CompiledGoFiles[index])
			embedPaths := GetEmbedDetailsForFile(file, baseDir)
			if len(embedPaths) > 0 {
				result = append(result, embedPaths...)
			}
		}
	}
	return result, nil
}


// ff:
// baseDir:
// file:
func GetEmbedDetailsForFile(file *ast.File, baseDir string) []*EmbedDetails {
	var result []*EmbedDetails
	for _, comment := range file.Comments {
		for _, c := range comment.List {
			if strings.HasPrefix(c.Text, "//go:embed") {
				sl := strings.Split(c.Text, " ")
				path := ""
				all := false
				if len(sl) == 1 {
					continue
				}
				embedPath := strings.TrimSpace(sl[1])
				switch true {
				case strings.HasPrefix(embedPath, "all:"):
					path = strings.TrimPrefix(embedPath, "all:")
					all = true
				default:
					path = embedPath
				}
				result = append(result, &EmbedDetails{
					EmbedPath: path,
					All:       all,
					BaseDir:   baseDir,
				})
			}
		}
	}
	return result
}
