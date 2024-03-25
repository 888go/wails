package fs

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"github.com/leaanthony/slicer"
)

// RelativeToCwd 返回一个基于当前工作目录（cwd）和给定的相对路径的绝对路径

// ff:
// relativePath:
func RelativeToCwd(relativePath string) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(cwd, relativePath), nil
}

// Mkdir 将创建给定的目录

// ff:
// dirname:
func Mkdir(dirname string) error {
	return os.Mkdir(dirname, 0o755)
}

// MkDirs 创建给定的嵌套目录。
// 若创建失败，返回错误

// ff:
// mode:
// fullPath:
func MkDirs(fullPath string, mode ...os.FileMode) error {
	var perms os.FileMode
	perms = 0o755
	if len(mode) == 1 {
		perms = mode[0]
	}
	return os.MkdirAll(fullPath, perms)
}

// MoveFile尝试将源文件移动到目标位置
// 目标是一个指向文件名的完整路径，而不是一个目录

// ff:
// target:
// source:
func MoveFile(source string, target string) error {
	return os.Rename(source, target)
}

// DeleteFile 将会删除给定的文件

// ff:
// filename:
func DeleteFile(filename string) error {
	return os.Remove(filename)
}

// CopyFile 从源文件复制到目标文件

// ff:
// target:
// source:
func CopyFile(source string, target string) error {
	s, err := os.Open(source)
	if err != nil {
		return err
	}
	defer s.Close()
	d, err := os.Create(target)
	if err != nil {
		return err
	}
	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		return err
	}
	return d.Close()
}

// DirExists - 如果给定的路径在文件系统中解析为一个目录，则返回true

// ff:
// path:
func DirExists(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode().IsDir()
}

// FileExists 返回一个布尔值，表示给定的文件是否存在

// ff:
// path:
func FileExists(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode().IsRegular()
}

// RelativePath 函数返回一个由调用文件所在目录与给定的相对路径组合而成的完整路径。
//
// 示例：在 *本* 文件中调用 RelativePath("..") 将会得到 '/path/to/wails2/v2/internal'

// ff:
// optionalpaths:
// relativepath:
func RelativePath(relativepath string, optionalpaths ...string) string {
	_, thisFile, _, _ := runtime.Caller(1)
	localDir := filepath.Dir(thisFile)

	// 如果我们有可选路径，将其与相对路径连接起来
	if len(optionalpaths) > 0 {
		paths := []string{relativepath}
		paths = append(paths, optionalpaths...)
		relativepath = filepath.Join(paths...)
	}
	result, err := filepath.Abs(filepath.Join(localDir, relativepath))
	if err != nil {
// 我仅出于一个原因允许这样做：如果提供的路径不正确，那将是致命的，因为它只在Wails内部使用。如果我们获取的路径错误，我们应立即得知。另一个原因是，这可以大量减少不必要的错误处理。
		panic(err)
	}
	return result
}

// MustLoadString尝试加载一个字符串，如果出现任何错误，将会输出一条致命消息并终止程序

// ff:
// filename:
func MustLoadString(filename string) string {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("FATAL: Unable to load file '%s': %s\n", filename, err.Error())
		os.Exit(1)
	}
	return *(*string)(unsafe.Pointer(&data))
}

// MD5File 返回给定文件的 md5 哈希值

// ff:
// filename:
func MD5File(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}

// MustMD5File将会调用MD5File函数，并在出现错误时终止程序运行

// ff:
// filename:
func MustMD5File(filename string) string {
	result, err := MD5File(filename)
	if err != nil {
		println("FATAL: Unable to MD5Sum file:", err.Error())
		os.Exit(1)
	}
	return result
}

// MustWriteString 将尝试将给定的数据写入给定的文件名
// 如果发生失败，它将中止程序

// ff:
// data:
// filename:
func MustWriteString(filename string, data string) {
	err := os.WriteFile(filename, []byte(data), 0o755)
	if err != nil {
		fatal("Unable to write file", filename, ":", err.Error())
		os.Exit(1)
	}
}

// fatal会打印可选的消息并终止程序
func fatal(message ...string) {
	if len(message) > 0 {
		print("FATAL:")
		for text := range message {
			print(text)
		}
	}
	os.Exit(1)
}

// GetSubdirectories 返回给定根目录下的子目录列表

// ff:
// rootDir:
func GetSubdirectories(rootDir string) (*slicer.StringSlicer, error) {
	var result slicer.StringSlicer

	// Iterate root dir
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果我们有一个目录，保存它
		if info.IsDir() {
			result.Add(path)
		}
		return nil
	})
	return &result, err
}


// ff:
// dir:
func DirIsEmpty(dir string) (bool, error) {
	// CREDIT: 代码来源：https://stackoverflow.com/a/30708914/8325411
	f, err := os.Open(dir)
	if err != nil {
		return false, err
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, err // 不为空或者存在错误，适用于这两种情况
}

// CopyDir 递归地复制一个目录树，尝试保持原始权限设置。
// 源目录必须存在，目标目录必须不存在。
// 符号链接会被忽略并跳过。
// 来源：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

// ff:
// err:
// dst:
// src:
func CopyDir(src string, dst string) (err error) {
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = MkDirs(dst)
	if err != nil {
		return
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Type()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}

// SetPermissions 递归地设置目录及其下文件的权限

// ff:
// perm:
// dir:
func SetPermissions(dir string, perm os.FileMode) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return os.Chmod(path, perm)
	})
}

// CopyDirExtended递归地复制一个目录树，尝试保持文件权限。
// 源目录必须存在，目标目录必须不存在。它会忽略通过ignore参数给出的所有文件或目录。
// 符号链接会被忽略并跳过。
// 来源：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

// ff:
// err:
// ignore:
// dst:
// src:
func CopyDirExtended(src string, dst string, ignore []string) (err error) {
	ignoreList := slicer.String(ignore)
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	if err != nil {
		return err
	}
	if !si.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		return
	}
	if err == nil {
		return fmt.Errorf("destination already exists")
	}

	err = MkDirs(dst)
	if err != nil {
		return
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		if ignoreList.Contains(entry.Name()) {
			continue
		}
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = CopyDir(srcPath, dstPath)
			if err != nil {
				return
			}
		} else {
			// Skip symlinks.
			if entry.Type()&os.ModeSymlink != 0 {
				continue
			}

			err = CopyFile(srcPath, dstPath)
			if err != nil {
				return
			}
		}
	}

	return
}


// ff:
// file:
// fsys:
func FindPathToFile(fsys fs.FS, file string) (string, error) {
	stat, _ := fs.Stat(fsys, file)
	if stat != nil {
		return ".", nil
	}
	var indexFiles slicer.StringSlicer
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(path, file) {
			indexFiles.Add(path)
		}
		return nil
	})
	if err != nil {
		return "", err
	}

	if indexFiles.Length() > 1 {
		selected := indexFiles.AsSlice()[0]
		for _, f := range indexFiles.AsSlice() {
			if len(f) < len(selected) {
				selected = f
			}
		}
		path, _ := filepath.Split(selected)
		return path, nil
	}
	if indexFiles.Length() > 0 {
		path, _ := filepath.Split(indexFiles.AsSlice()[0])
		return path, nil
	}
	return "", fmt.Errorf("%s: %w", file, os.ErrNotExist)
}

// FindFileInParents 在当前目录及其所有父目录中搜索指定文件。
// 如果找到该文件，则返回该文件的绝对路径，否则返回一个空字符串

// ff:
// filename:
// path:
func FindFileInParents(path string, filename string) string {
	// Check for bad paths
	if _, err := os.Stat(path); err != nil {
		return ""
	}

	var pathToFile string
	for {
		pathToFile = filepath.Join(path, filename)
		if _, err := os.Stat(pathToFile); err == nil {
			break
		}
		parent := filepath.Dir(path)
		if parent == path {
			return ""
		}
		path = parent
	}
	return pathToFile
}
