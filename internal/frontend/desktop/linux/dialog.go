//go:build linux
// +build linux

package linux

import (
	"github.com/wailsapp/wails/v2/internal/frontend"
	"unsafe"
)

/*
#include <stdlib.h>
#include "gtk/gtk.h"
*/
import "C"

const (
	GTK_FILE_CHOOSER_ACTION_OPEN          C.GtkFileChooserAction = C.GTK_FILE_CHOOSER_ACTION_OPEN
	GTK_FILE_CHOOSER_ACTION_SAVE          C.GtkFileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SAVE
	GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER C.GtkFileChooserAction = C.GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER
)

var openFileResults = make(chan []string)
var messageDialogResult = make(chan string)


// ff:
// err:
// result:
// dialogOptions:
func (f *Frontend) OpenFileDialog(dialogOptions frontend.OpenDialogOptions) (result string, err error) {
	f.mainWindow.OpenFileDialog(dialogOptions, 0, GTK_FILE_CHOOSER_ACTION_OPEN)
	results := <-openFileResults
	if len(results) == 1 {
		return results[0], nil
	}
	return "", nil
}


// ff:
// dialogOptions:
func (f *Frontend) OpenMultipleFilesDialog(dialogOptions frontend.OpenDialogOptions) ([]string, error) {
	f.mainWindow.OpenFileDialog(dialogOptions, 1, GTK_FILE_CHOOSER_ACTION_OPEN)
	result := <-openFileResults
	return result, nil
}


// ff:
// dialogOptions:
func (f *Frontend) OpenDirectoryDialog(dialogOptions frontend.OpenDialogOptions) (string, error) {
	f.mainWindow.OpenFileDialog(dialogOptions, 0, GTK_FILE_CHOOSER_ACTION_SELECT_FOLDER)
	result := <-openFileResults
	if len(result) == 1 {
		return result[0], nil
	}
	return "", nil
}


// ff:
// dialogOptions:
func (f *Frontend) SaveFileDialog(dialogOptions frontend.SaveDialogOptions) (string, error) {
	options := frontend.OpenDialogOptions{
		DefaultDirectory:     dialogOptions.DefaultDirectory,
		DefaultFilename:      dialogOptions.DefaultFilename,
		Title:                dialogOptions.Title,
		Filters:              dialogOptions.Filters,
		ShowHiddenFiles:      dialogOptions.ShowHiddenFiles,
		CanCreateDirectories: dialogOptions.CanCreateDirectories,
	}
	f.mainWindow.OpenFileDialog(options, 0, GTK_FILE_CHOOSER_ACTION_SAVE)
	results := <-openFileResults
	if len(results) == 1 {
		return results[0], nil
	}
	return "", nil
}


// ff:
// dialogOptions:
func (f *Frontend) MessageDialog(dialogOptions frontend.MessageDialogOptions) (string, error) {
	f.mainWindow.MessageDialog(dialogOptions)
	return <-messageDialogResult, nil
}

//export processOpenFileResult
// 导出processOpenFileResult函数（供C语言或其他外部语言调用）
func processOpenFileResult(carray **C.char) {
	// 从C语言数组创建一个Go语言切片
	var result []string
	goArray := (*[1024]*C.char)(unsafe.Pointer(carray))[:1024:1024]
	for _, s := range goArray {
		if s == nil {
			break
		}
		result = append(result, C.GoString(s))
	}
	openFileResults <- result
}

//export processMessageDialogResult
// 导出processMessageDialogResult函数，供C语言或其他外部环境调用
func processMessageDialogResult(result *C.char) {
	messageDialogResult <- C.GoString(result)
}
