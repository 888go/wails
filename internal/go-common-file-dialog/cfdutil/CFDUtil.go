package cfdutil

import (
	"github.com/wailsapp/wails/v2/internal/go-common-file-dialog/cfd"
)

// TODO doc

// ff:
// config:
func ShowOpenFileDialog(config cfd.DialogConfig) (string, error) {
	dialog, err := cfd.NewOpenFileDialog(config)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = dialog.Release()
	}()
	return dialog.ShowAndGetResult()
}

// TODO doc

// ff:
// config:
func ShowOpenMultipleFilesDialog(config cfd.DialogConfig) ([]string, error) {
	dialog, err := cfd.NewOpenMultipleFilesDialog(config)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = dialog.Release()
	}()
	return dialog.ShowAndGetResults()
}

// TODO doc

// ff:
// config:
func ShowPickFolderDialog(config cfd.DialogConfig) (string, error) {
	dialog, err := cfd.NewSelectFolderDialog(config)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = dialog.Release()
	}()
	return dialog.ShowAndGetResult()
}

// TODO doc

// ff:
// config:
func ShowSaveFileDialog(config cfd.DialogConfig) (string, error) {
	dialog, err := cfd.NewSaveFileDialog(config)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = dialog.Release()
	}()
	return dialog.ShowAndGetResult()
}
