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
	InfoDialog     = frontend.InfoDialog
	WarningDialog  = frontend.WarningDialog
	ErrorDialog    = frontend.ErrorDialog
	QuestionDialog = frontend.QuestionDialog
)

// MessageDialogOptions 包含了用于消息对话框（如Info、Warning等运行时方法）的选项。
type MessageDialogOptions = frontend.MessageDialogOptions

// OpenDirectoryDialog 提示用户选择一个目录
func OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	appFrontend := getFrontend(ctx)
	if dialogOptions.DefaultDirectory != "" {
		if !fs.DirExists(dialogOptions.DefaultDirectory) {
			return "", fmt.Errorf("default directory '%s' does not exist", dialogOptions.DefaultDirectory)
		}
	}
	return appFrontend.OpenDirectoryDialog(dialogOptions)
}

// OpenFileDialog 提示用户选择一个文件
func OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	appFrontend := getFrontend(ctx)
	if dialogOptions.DefaultDirectory != "" {
		if !fs.DirExists(dialogOptions.DefaultDirectory) {
			return "", fmt.Errorf("default directory '%s' does not exist", dialogOptions.DefaultDirectory)
		}
	}
	return appFrontend.OpenFileDialog(dialogOptions)
}

// OpenMultipleFilesDialog 提示用户选择一个或多个文件
func OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error) {
	appFrontend := getFrontend(ctx)
	if dialogOptions.DefaultDirectory != "" {
		if !fs.DirExists(dialogOptions.DefaultDirectory) {
			return nil, fmt.Errorf("default directory '%s' does not exist", dialogOptions.DefaultDirectory)
		}
	}
	return appFrontend.OpenMultipleFilesDialog(dialogOptions)
}

// SaveFileDialog 弹出文件选择对话框，提示用户选择一个文件
func SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error) {
	appFrontend := getFrontend(ctx)
	if dialogOptions.DefaultDirectory != "" {
		if !fs.DirExists(dialogOptions.DefaultDirectory) {
			return "", fmt.Errorf("default directory '%s' does not exist", dialogOptions.DefaultDirectory)
		}
	}
	return appFrontend.SaveFileDialog(dialogOptions)
}

// MessageDialog 向用户展示一条消息对话框
func MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error) {
	appFrontend := getFrontend(ctx)
	return appFrontend.MessageDialog(dialogOptions)
}
