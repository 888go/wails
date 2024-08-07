package assetserver

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// FindEmbedRootPath 在嵌入式文件系统中查找根路径。这个根路径是包含所有文件的目录。

// ff:
// fsys:
func FindEmbedRootPath(fsys embed.FS) (string, error) {
	stopErr := fmt.Errorf("files or multiple dirs found")

	fPath := ""
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			fPath = path
			if entries, dErr := fs.ReadDir(fsys, path); dErr != nil {
				return dErr
			} else if len(entries) <= 1 {
				return nil
			}
		}

		return stopErr
	})

	if err != nil && err != stopErr {
		return "", err
	}

	return fPath, nil
}


// ff:查找文件路径
// file:
// fsys:
func X查找文件路径(fsys fs.FS, file string) (string, error) {
	stat, _ := fs.Stat(fsys, file)
	if stat != nil {
		return ".", nil
	}
	var indexFiles []string
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, file) {
			indexFiles = append(indexFiles, path)
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	if len(indexFiles) > 1 {
		selected := indexFiles[0]
		for _, f := range indexFiles {
			if len(f) < len(selected) {
				selected = f
			}
		}
		path, _ := filepath.Split(selected)
		return path, nil
	}
	if len(indexFiles) > 0 {
		path, _ := filepath.Split(indexFiles[0])
		return path, nil
	}
	return "", fmt.Errorf("%s: %w", file, os.ErrNotExist)
}
