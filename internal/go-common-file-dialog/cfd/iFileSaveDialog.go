//go:build windows
// +build windows

package cfd

import (
	"github.com/go-ole/go-ole"
	"github.com/888go/wails/internal/go-common-file-dialog/util"
	"unsafe"
)

var (
	saveFileDialogCLSID = ole.NewGUID("{C0B4E2F3-BA21-4773-8DBA-335EC946EB8B}")
	saveFileDialogIID   = ole.NewGUID("{84bccd23-5fde-4cdb-aea4-af64b83d78ab}")
)

type iFileSaveDialog struct {
	vtbl               *iFileSaveDialogVtbl
	parentWindowHandle uintptr
}

type iFileSaveDialogVtbl struct {
	iFileDialogVtbl

	SetSaveAsItem          uintptr
	SetProperties          uintptr
	SetCollectedProperties uintptr
	GetProperties          uintptr
	ApplyProperties        uintptr
}

func newIFileSaveDialog() (*iFileSaveDialog, error) {
	if unknown, err := ole.CreateInstance(saveFileDialogCLSID, saveFileDialogIID); err == nil {
		return (*iFileSaveDialog)(unsafe.Pointer(unknown)), nil
	} else {
		return nil, err
	}
}


// ff:
func (fileSaveDialog *iFileSaveDialog) Show() error {
	return fileSaveDialog.vtbl.show(unsafe.Pointer(fileSaveDialog), fileSaveDialog.parentWindowHandle)
}


// ff:
// hwnd:
func (fileSaveDialog *iFileSaveDialog) SetParentWindowHandle(hwnd uintptr) {
	fileSaveDialog.parentWindowHandle = hwnd
}


// ff:
func (fileSaveDialog *iFileSaveDialog) ShowAndGetResult() (string, error) {
	if err := fileSaveDialog.Show(); err != nil {
		return "", err
	}
	return fileSaveDialog.GetResult()
}


// ff:
// title:
func (fileSaveDialog *iFileSaveDialog) SetTitle(title string) error {
	return fileSaveDialog.vtbl.setTitle(unsafe.Pointer(fileSaveDialog), title)
}


// ff:
func (fileSaveDialog *iFileSaveDialog) GetResult() (string, error) {
	return fileSaveDialog.vtbl.getResultString(unsafe.Pointer(fileSaveDialog))
}


// ff:
func (fileSaveDialog *iFileSaveDialog) Release() error {
	return fileSaveDialog.vtbl.release(unsafe.Pointer(fileSaveDialog))
}


// ff:
// defaultFolderPath:
func (fileSaveDialog *iFileSaveDialog) SetDefaultFolder(defaultFolderPath string) error {
	return fileSaveDialog.vtbl.setDefaultFolder(unsafe.Pointer(fileSaveDialog), defaultFolderPath)
}


// ff:
// defaultFolderPath:
func (fileSaveDialog *iFileSaveDialog) SetFolder(defaultFolderPath string) error {
	return fileSaveDialog.vtbl.setFolder(unsafe.Pointer(fileSaveDialog), defaultFolderPath)
}


// ff:
// filter:
func (fileSaveDialog *iFileSaveDialog) SetFileFilters(filter []FileFilter) error {
	return fileSaveDialog.vtbl.setFileTypes(unsafe.Pointer(fileSaveDialog), filter)
}


// ff:
// role:
func (fileSaveDialog *iFileSaveDialog) SetRole(role string) error {
	return fileSaveDialog.vtbl.setClientGuid(unsafe.Pointer(fileSaveDialog), util.StringToUUID(role))
}


// ff:
// defaultExtension:
func (fileSaveDialog *iFileSaveDialog) SetDefaultExtension(defaultExtension string) error {
	return fileSaveDialog.vtbl.setDefaultExtension(unsafe.Pointer(fileSaveDialog), defaultExtension)
}


// ff:
// initialFileName:
func (fileSaveDialog *iFileSaveDialog) SetFileName(initialFileName string) error {
	return fileSaveDialog.vtbl.setFileName(unsafe.Pointer(fileSaveDialog), initialFileName)
}


// ff:
// index:
func (fileSaveDialog *iFileSaveDialog) SetSelectedFileFilterIndex(index uint) error {
	return fileSaveDialog.vtbl.setSelectedFileFilterIndex(unsafe.Pointer(fileSaveDialog), index)
}
