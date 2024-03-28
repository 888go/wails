package dev

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/internal/fs"

	"github.com/fsnotify/fsnotify"
	gitignore "github.com/sabhiram/go-gitignore"
	"github.com/samber/lo"
)

type Watcher interface {
	Add(name string) error
}

// initialiseWatcher 创建项目目录监视器，当检测到变化时将触发重新编译
func initialiseWatcher(cwd string) (*fsnotify.Watcher, error) {
	// 默认情况下忽略 dot 文件、node_modules 和 build 目录
	ignoreDirs := getIgnoreDirs(cwd)

	// Get all subdirectories
	dirs, err := fs.GetSubdirectories(cwd)
	if err != nil {
		return nil, err
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	for _, dir := range processDirectories(dirs.AsSlice(), ignoreDirs) {
		err := watcher.Add(dir)
		if err != nil {
			return nil, err
		}
	}
	return watcher, nil
}

func getIgnoreDirs(cwd string) []string {
	ignoreDirs := []string{filepath.Join(cwd, "build/*"), ".*", "node_modules"}
	baseDir := filepath.Base(cwd)
	// 将.gitignore读取到ignoreDirs中
	f, err := os.Open(filepath.Join(cwd, ".gitignore"))
	if err == nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := scanner.Text()
			if line != baseDir {
				ignoreDirs = append(ignoreDirs, line)
			}
		}
	}

	return lo.Uniq(ignoreDirs)
}

func processDirectories(dirs []string, ignoreDirs []string) []string {
	ignorer := gitignore.CompileIgnoreLines(ignoreDirs...)
	return lo.Filter(dirs, func(dir string, _ int) bool {
		return !ignorer.MatchesPath(dir)
	})
}
