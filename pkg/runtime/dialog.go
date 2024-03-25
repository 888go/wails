package runtime

import (
	"context"
	"fmt"

	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/internal/fs"
)

// FileFilter 定义了对话框的文件过滤器
type FileFilter = frontend.FileFilter

// OpenDialogOptions 包含了OpenDialogOptions运行时方法的选项参数
type OpenDialogOptions = frontend.OpenDialogOptions

// SaveDialogOptions 包含了SaveDialog运行时方法的选项参数
type SaveDialogOptions = frontend.SaveDialogOptions

type DialogType = frontend.DialogType

const (
	InfoDialog     = frontend.X常量_对话框_信息
	WarningDialog  = frontend.X常量_对话框_警告
	ErrorDialog    = frontend.X常量_对话框_错误
	QuestionDialog = frontend.X常量_对话框_问题
)

// MessageDialogOptions 包含了用于消息对话框（如Info、Warning等运行时方法）的选项。
type MessageDialogOptions = frontend.MessageDialogOptions

// OpenDirectoryDialog 提示用户选择一个目录
func X对话框选择目录(上下文 context.Context, 选项 OpenDialogOptions) (string, error) {
	appFrontend := getFrontend(上下文)
	if 选项.X默认目录 != "" {
		if !fs.DirExists(选项.X默认目录) {
			return "", fmt.Errorf("default directory '%s' does not exist", 选项.X默认目录)
		}
	}
	return appFrontend.X对话框选择目录(选项)
}

// OpenFileDialog 提示用户选择一个文件
func X对话框选择文件(上下文 context.Context, 选项 OpenDialogOptions) (string, error) {
	appFrontend := getFrontend(上下文)
	if 选项.X默认目录 != "" {
		if !fs.DirExists(选项.X默认目录) {
			return "", fmt.Errorf("default directory '%s' does not exist", 选项.X默认目录)
		}
	}
	return appFrontend.X对话框选择文件(选项)
}

// OpenMultipleFilesDialog 提示用户选择一个或多个文件
func X对话框多选文件(上下文 context.Context, 选项 OpenDialogOptions) ([]string, error) {
	appFrontend := getFrontend(上下文)
	if 选项.X默认目录 != "" {
		if !fs.DirExists(选项.X默认目录) {
			return nil, fmt.Errorf("default directory '%s' does not exist", 选项.X默认目录)
		}
	}
	return appFrontend.X对话框多选文件(选项)
}

// SaveFileDialog 弹出文件选择对话框，提示用户选择一个文件
func X对话框保存文件(上下文 context.Context, 选项 SaveDialogOptions) (string, error) {
	appFrontend := getFrontend(上下文)
	if 选项.X默认目录 != "" {
		if !fs.DirExists(选项.X默认目录) {
			return "", fmt.Errorf("default directory '%s' does not exist", 选项.X默认目录)
		}
	}
	return appFrontend.X对话框保存文件(选项)
}

// MessageDialog 向用户展示一条消息对话框
func X对话框弹出消息(上下文 context.Context, 选项 MessageDialogOptions) (string, error) {
	appFrontend := getFrontend(上下文)
	return appFrontend.X对话框弹出消息(选项)
}
