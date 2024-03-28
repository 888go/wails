//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

func genOFN(parent Controller, title, filter string, filterIndex uint, initialDir string, buf []uint16) *w32.OPENFILENAME {
	var ofn w32.OPENFILENAME
	ofn.StructSize = uint32(unsafe.Sizeof(ofn))
	ofn.Owner = parent.Handle()

	if filter != "" {
		filterBuf := make([]uint16, len(filter)+1)
		copy(filterBuf, syscall.StringToUTF16(filter))
		// 将 '|' 替换为预期的 '\0'
		for i, c := range filterBuf {
			if byte(c) == '|' {
				filterBuf[i] = uint16(0)
			}
		}
		ofn.Filter = &filterBuf[0]
		ofn.FilterIndex = uint32(filterIndex)
	}

	ofn.File = &buf[0]
	ofn.MaxFile = uint32(len(buf))

	if initialDir != "" {
		ofn.InitialDir = syscall.StringToUTF16Ptr(initialDir)
	}
	if title != "" {
		ofn.Title = syscall.StringToUTF16Ptr(title)
	}

	ofn.Flags = w32.OFN_FILEMUSTEXIST
	return &ofn
}


// ff:
// accepted:
// filePath:
// initialDir:
// filterIndex:
// filter:
// title:
// parent:
func ShowOpenFileDlg(parent Controller, title, filter string, filterIndex uint, initialDir string) (filePath string, accepted bool) {
	buf := make([]uint16, 1024)
	ofn := genOFN(parent, title, filter, filterIndex, initialDir, buf)

	if accepted = w32.GetOpenFileName(ofn); accepted {
		filePath = syscall.UTF16ToString(buf)
	}
	return
}


// ff:
// accepted:
// filePath:
// initialDir:
// filterIndex:
// filter:
// title:
// parent:
func ShowSaveFileDlg(parent Controller, title, filter string, filterIndex uint, initialDir string) (filePath string, accepted bool) {
	buf := make([]uint16, 1024)
	ofn := genOFN(parent, title, filter, filterIndex, initialDir, buf)

	if accepted = w32.GetSaveFileName(ofn); accepted {
		filePath = syscall.UTF16ToString(buf)
	}
	return
}


// ff:
// accepted:
// folder:
// title:
// parent:
func ShowBrowseFolderDlg(parent Controller, title string) (folder string, accepted bool) {
	var bi w32.BROWSEINFO
	bi.Owner = parent.Handle()
	bi.Title = syscall.StringToUTF16Ptr(title)
	bi.Flags = w32.BIF_RETURNONLYFSDIRS | w32.BIF_NEWDIALOGSTYLE

	w32.CoInitialize()
	ret := w32.SHBrowseForFolder(&bi)
	w32.CoUninitialize()

	folder = w32.SHGetPathFromIDList(ret)
	accepted = folder != ""
	return
}

// MsgBoxOkCancel 基础弹出消息框。返回值为1表示点击了“确定”，返回值为2表示点击了“取消”。

// ff:
// caption:
// title:
// parent:
func MsgBoxOkCancel(parent Controller, title, caption string) int {
	return MsgBox(parent, title, caption, w32.MB_ICONEXCLAMATION|w32.MB_OKCANCEL)
}


// ff:
// caption:
// title:
// parent:
func MsgBoxYesNo(parent Controller, title, caption string) int {
	return MsgBox(parent, title, caption, w32.MB_ICONEXCLAMATION|w32.MB_YESNO)
}


// ff:
// caption:
// title:
// parent:
func MsgBoxOk(parent Controller, title, caption string) {
	MsgBox(parent, title, caption, w32.MB_ICONINFORMATION|w32.MB_OK)
}

// Warningf 是一个通用警告信息函数，带有“确定”和“取消”按钮。当用户点击“确定”时返回1。

// ff:
// data:
// format:
// parent:
func Warningf(parent Controller, format string, data ...interface{}) int {
	caption := fmt.Sprintf(format, data...)
	return MsgBox(parent, "Warning", caption, w32.MB_ICONWARNING|w32.MB_OKCANCEL)
}

// Printf 是一个通用的信息消息，带有一个“确定”按钮。

// ff:
// data:
// format:
// parent:
func Printf(parent Controller, format string, data ...interface{}) {
	caption := fmt.Sprintf(format, data...)
	MsgBox(parent, "Information", caption, w32.MB_ICONINFORMATION|w32.MB_OK)
}

// Errorf 是一个通用错误消息，带有“确定”按钮。

// ff:
// data:
// format:
// parent:
func Errorf(parent Controller, format string, data ...interface{}) {
	caption := fmt.Sprintf(format, data...)
	MsgBox(parent, "Error", caption, w32.MB_ICONERROR|w32.MB_OK)
}


// ff:
// flags:
// caption:
// title:
// parent:
func MsgBox(parent Controller, title, caption string, flags uint) int {
	var result int
	if parent != nil {
		result = w32.MessageBox(parent.Handle(), caption, title, flags)
	} else {
		result = w32.MessageBox(0, caption, title, flags)
	}

	return result
}
