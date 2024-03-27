package s

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bitfield/script"
)

var (
	Output         io.Writer = io.Discard
	IndentSize     int
	originalOutput io.Writer
	currentIndent  int
)

func checkError(err error) {
	if err != nil {
		println("\nERROR:", err.Error())
		os.Exit(1)
	}
}

func mute() {
	originalOutput = Output
	Output = ioutil.Discard
}

func unmute() {
	Output = originalOutput
}

func indent() {
	currentIndent += IndentSize
}

func unindent() {
	currentIndent -= IndentSize
}

func log(message string, args ...interface{}) {
	indent := strings.Repeat(" ", currentIndent)
	_, err := fmt.Fprintf(Output, indent+message+"\n", args...)
	checkError(err)
}

// 重命名文件或目录

// ff:
// target:
// source:
func RENAME(source string, target string) {
	log("RENAME %s -> %s", source, target)
	err := os.Rename(source, target)
	checkError(err)
}

// MUSTDELETE a file.

// ff:
// filename:
func MUSTDELETE(filename string) {
	log("DELETE %s", filename)
	err := os.Remove(filepath.Join(CWD(), filename))
	checkError(err)
}

// DELETE a file.

// ff:
// filename:
func DELETE(filename string) {
	log("DELETE %s", filename)
	_ = os.Remove(filepath.Join(CWD(), filename))
}


// ff:
// dir:
func CD(dir string) {
	err := os.Chdir(dir)
	checkError(err)
	log("CD %s [%s]", dir, CWD())
}


// ff:
// mode:
// path:
func MKDIR(path string, mode ...os.FileMode) {
	var perms os.FileMode
	perms = 0o755
	if len(mode) == 1 {
		perms = mode[0]
	}
	log("MKDIR %s (perms: %v)", path, perms)
	err := os.MkdirAll(path, perms)
	checkError(err)
}

// ENDIR 确保如果路径不存在，则创建该路径

// ff:
// mode:
// path:
func ENDIR(path string, mode ...os.FileMode) {
	var perms os.FileMode
	perms = 0o755
	if len(mode) == 1 {
		perms = mode[0]
	}
	_ = os.MkdirAll(path, perms)
}

// COPYDIR递归地复制一个目录树，尝试保持文件权限。
// 源目录必须存在，目标目录必须不存在。
// 符号链接将被忽略并跳过。
// 致谢：https://gist.github.com/r0l1/92462b38df26839a3ca324697c8cba04

// ff:
// dst:
// src:
func COPYDIR(src string, dst string) {
	log("COPYDIR %s -> %s", src, dst)
	src = filepath.Clean(src)
	dst = filepath.Clean(dst)

	si, err := os.Stat(src)
	checkError(err)
	if !si.IsDir() {
		checkError(fmt.Errorf("source is not a directory"))
	}

	_, err = os.Stat(dst)
	if err != nil && !os.IsNotExist(err) {
		checkError(err)
	}
	if err == nil {
		checkError(fmt.Errorf("destination already exists"))
	}

	indent()
	MKDIR(dst)

	entries, err := os.ReadDir(src)
	checkError(err)

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			COPYDIR(srcPath, dstPath)
		} else {
			// Skip symlinks.
			if entry.Type()&os.ModeSymlink != 0 {
				continue
			}

			COPY(srcPath, dstPath)
		}
	}
	unindent()
}

// 将文件从源复制到目标

// ff:
// target:
// source:
func COPY(source string, target string) {
	log("COPY %s -> %s", source, target)
	src, err := os.Open(source)
	checkError(err)
	defer closefile(src)
	d, err := os.Create(target)
	checkError(err)
	_, err = io.Copy(d, src)
	checkError(err)
}


// ff:
func CWD() string {
	result, err := os.Getwd()
	checkError(err)
	log("CWD [%s]", result)
	return result
}


// ff:
// target:
func RMDIR(target string) {
	log("RMDIR %s", target)
	err := os.RemoveAll(target)
	checkError(err)
}


// ff:
// target:
func RM(target string) {
	log("RM %s", target)
	err := os.Remove(target)
	checkError(err)
}


// ff:
// message:
func ECHO(message string) {
	println(message)
}


// ff:
// filepath:
func TOUCH(filepath string) {
	log("TOUCH %s", filepath)
	f, err := os.Create(filepath)
	checkError(err)
	closefile(f)
}


// ff:
// command:
func EXEC(command string) {
	log("EXEC %s", command)
	gen := script.Exec(command)
	gen.Wait()
	checkError(gen.Error())
}

// EXISTS - 如果给定路径存在则返回真（true）

// ff:
// path:
func EXISTS(path string) bool {
	_, err := os.Lstat(path)
	log("EXISTS %s (%T)", path, err == nil)
	return err == nil
}

// ISDIR 函数返回 true，如果给定的目录存在

// ff:
// path:
func ISDIR(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode().IsDir()
}

// ISDIREMPTY 判断给定的目录是否为空，若为空则返回 true

// ff:
// dir:
func ISDIREMPTY(dir string) bool {
	// CREDIT: 代码来源：https://stackoverflow.com/a/30708914/8325411
	f, err := os.Open(dir)
	checkError(err)
	defer closefile(f)

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	return err == io.EOF
}

// ISFILE在给定的文件存在时返回true

// ff:
// path:
func ISFILE(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}

	return fi.Mode().IsRegular()
}

// SUBDIRS 函数返回给定目录下的子目录列表

// ff:
// rootDir:
func SUBDIRS(rootDir string) []string {
	var result []string

	// Iterate root dir
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		checkError(err)
		// 如果我们有一个目录，保存它
		if info.IsDir() {
			result = append(result, path)
		}
		return nil
	})
	checkError(err)
	return result
}

// SAVESTRING 将使用给定的字符串创建一个文件

// ff:
// data:
// filename:
func SAVESTRING(filename string, data string) {
	log("SAVESTRING %s", filename)
	mute()
	SAVEBYTES(filename, []byte(data))
	unmute()
}

// LOADSTRING 函数返回给定文件名的文件内容作为字符串

// ff:
// filename:
func LOADSTRING(filename string) string {
	log("LOADSTRING %s", filename)
	mute()
	data := LOADBYTES(filename)
	unmute()
	return string(data)
}

// SAVEBYTES 将使用给定的字符串创建一个文件

// ff:
// data:
// filename:
func SAVEBYTES(filename string, data []byte) {
	log("SAVEBYTES %s", filename)
	err := os.WriteFile(filename, data, 0o755)
	checkError(err)
}

// LOADBYTES 函数返回给定文件名的文件内容作为字符串

// ff:
// filename:
func LOADBYTES(filename string) []byte {
	log("LOADBYTES %s", filename)
	data, err := os.ReadFile(filename)
	checkError(err)
	return data
}

func closefile(f *os.File) {
	err := f.Close()
	checkError(err)
}

// MD5FILE 返回给定文件的 md5 校验和

// ff:
// filename:
func MD5FILE(filename string) string {
	f, err := os.Open(filename)
	checkError(err)
	defer closefile(f)

	h := md5.New()
	_, err = io.Copy(h, f)
	checkError(err)

	return hex.EncodeToString(h.Sum(nil))
}

// Sub 是替换类型
type Sub map[string]string

// REPLACEALL 将给定文件中所有替换键替换为关联的值

// ff:
// substitutions:
// filename:
func REPLACEALL(filename string, substitutions Sub) {
	log("REPLACEALL %s (%v)", filename, substitutions)
	data := LOADSTRING(filename)
	for old, newText := range substitutions {
		data = strings.ReplaceAll(data, old, newText)
	}
	SAVESTRING(filename, data)
}
