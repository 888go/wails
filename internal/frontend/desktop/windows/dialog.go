//go:build windows
// +build windows

package windows

import (
	"path/filepath"
	"strings"
	"syscall"

	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/internal/frontend/desktop/windows/winc/w32"
	"github.com/888go/wails/internal/go-common-file-dialog/cfd"
	"golang.org/x/sys/windows"
)

func (f *Frontend) getHandleForDialog() w32.HWND {
	if f.mainWindow.IsVisible() {
		return f.mainWindow.Handle()
	}
	return 0
}

func getDefaultFolder(folder string) (string, error) {
	if folder == "" {
		return "", nil
	}
	return filepath.Abs(folder)
}

// OpenDirectoryDialog 提示用户选择一个目录

// ff:对话框选择目录
// dialogOptions:选项
func (f *Frontend) OpenDirectoryDialog(options frontend.OpenDialogOptions) (string, error) {

	defaultFolder, err := getDefaultFolder(options.X默认目录)
	if err != nil {
		return "", err
	}

	config := cfd.DialogConfig{
		Title:  options.X标题,
		Role:   "PickFolder",
		Folder: defaultFolder,
	}

	result, err := f.showCfdDialog(
		func() (cfd.Dialog, error) {
			return cfd.NewSelectFolderDialog(config)
		}, false)

	if err != nil && err != cfd.ErrCancelled {
		return "", err
	}
	return result.(string), nil
}

// OpenFileDialog 提示用户选择一个文件

// ff:对话框选择文件
// options:选项
func (f *Frontend) OpenFileDialog(options frontend.OpenDialogOptions) (string, error) {
	defaultFolder, err := getDefaultFolder(options.X默认目录)
	if err != nil {
		return "", err
	}

	config := cfd.DialogConfig{
		Folder:      defaultFolder,
		FileFilters: convertFilters(options.X过滤器),
		FileName:    options.X默认文件名,
		Title:       options.X标题,
	}

	result, err := f.showCfdDialog(
		func() (cfd.Dialog, error) {
			return cfd.NewOpenFileDialog(config)
		}, false)

	if err != nil && err != cfd.ErrCancelled {
		return "", err
	}
	return result.(string), nil
}

// OpenMultipleFilesDialog 提示用户选择一个或多个文件

// ff:对话框多选文件
// dialogOptions:选项
func (f *Frontend) OpenMultipleFilesDialog(options frontend.OpenDialogOptions) ([]string, error) {

	defaultFolder, err := getDefaultFolder(options.X默认目录)
	if err != nil {
		return nil, err
	}

	config := cfd.DialogConfig{
		Title:       options.X标题,
		Role:        "OpenMultipleFiles",
		FileFilters: convertFilters(options.X过滤器),
		FileName:    options.X默认文件名,
		Folder:      defaultFolder,
	}

	result, err := f.showCfdDialog(
		func() (cfd.Dialog, error) {
			return cfd.NewOpenMultipleFilesDialog(config)
		}, true)

	if err != nil && err != cfd.ErrCancelled {
		return nil, err
	}
	return result.([]string), nil
}

// SaveFileDialog 弹出文件选择对话框，提示用户选择一个文件

// ff:对话框保存文件
// dialogOptions:选项
func (f *Frontend) SaveFileDialog(options frontend.SaveDialogOptions) (string, error) {

	defaultFolder, err := getDefaultFolder(options.X默认目录)
	if err != nil {
		return "", err
	}

	config := cfd.DialogConfig{
		Title:       options.X标题,
		Role:        "SaveFile",
		FileFilters: convertFilters(options.X过滤器),
		FileName:    options.X默认文件名,
		Folder:      defaultFolder,
	}

	if len(options.X过滤器) > 0 {
		config.DefaultExtension = strings.TrimPrefix(strings.Split(options.X过滤器[0].X扩展名列表, ";")[0], "*")
	}

	result, err := f.showCfdDialog(
		func() (cfd.Dialog, error) {
			return cfd.NewSaveFileDialog(config)
		}, false)

	if err != nil && err != cfd.ErrCancelled {
		return "", err
	}
	return result.(string), nil
}

func (f *Frontend) showCfdDialog(newDlg func() (cfd.Dialog, error), isMultiSelect bool) (any, error) {
	return invokeSync(f.mainWindow, func() (any, error) {
		dlg, err := newDlg()
		if err != nil {
			return nil, err
		}
		defer func() {
			err := dlg.Release()
			if err != nil {
				println("ERROR: Unable to release dialog:", err.Error())
			}
		}()

		dlg.SetParentWindowHandle(f.getHandleForDialog())
		if multi, _ := dlg.(cfd.OpenMultipleFilesDialog); multi != nil && isMultiSelect {
			return multi.ShowAndGetResults()
		}
		return dlg.ShowAndGetResult()
	})
}

func calculateMessageDialogFlags(options frontend.MessageDialogOptions) uint32 {
	var flags uint32

	switch options.X对话框类型 {
	case frontend.X常量_对话框_信息:
		flags = windows.MB_OK | windows.MB_ICONINFORMATION
	case frontend.X常量_对话框_错误:
		flags = windows.MB_ICONERROR | windows.MB_OK
	case frontend.X常量_对话框_问题:
		flags = windows.MB_YESNO
		if strings.TrimSpace(strings.ToLower(options.X默认按钮)) == "no" {
			flags |= windows.MB_DEFBUTTON2
		}
	case frontend.X常量_对话框_警告:
		flags = windows.MB_OK | windows.MB_ICONWARNING
	}

	return flags
}

// MessageDialog 向用户展示一条消息对话框

// ff:对话框弹出消息
// options:选项
func (f *Frontend) MessageDialog(options frontend.MessageDialogOptions) (string, error) {

	title, err := syscall.UTF16PtrFromString(options.X标题)
	if err != nil {
		return "", err
	}
	message, err := syscall.UTF16PtrFromString(options.X消息)
	if err != nil {
		return "", err
	}

	flags := calculateMessageDialogFlags(options)

	button, _ := windows.MessageBox(windows.HWND(f.getHandleForDialog()), message, title, flags|windows.MB_SYSTEMMODAL)
	// 这个映射将MessageBox返回值转换为字符串
	responses := []string{"", "Ok", "Cancel", "Abort", "Retry", "Ignore", "Yes", "No", "", "", "Try Again", "Continue"}
	result := "Error"
	if int(button) < len(responses) {
		result = responses[button]
	}
	return result, nil
}

func convertFilters(filters []frontend.FileFilter) []cfd.FileFilter {
	var result []cfd.FileFilter
	for _, filter := range filters {
		result = append(result, cfd.FileFilter(filter))
	}
	return result
}
