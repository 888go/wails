//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2012 The W32 Authors. All Rights Reserved.
 */
package w32

import (
	"syscall"
	"unsafe"
)

var (
	modcomdlg32 = syscall.NewLazyDLL("comdlg32.dll")

	procGetSaveFileName      = modcomdlg32.NewProc("GetSaveFileNameW")
	procGetOpenFileName      = modcomdlg32.NewProc("GetOpenFileNameW")
	procCommDlgExtendedError = modcomdlg32.NewProc("CommDlgExtendedError")
)


// ff:
// ofn:
func GetOpenFileName(ofn *OPENFILENAME) bool {
	ret, _, _ := procGetOpenFileName.Call(
		uintptr(unsafe.Pointer(ofn)))

	return ret != 0
}


// ff:
// ofn:
func GetSaveFileName(ofn *OPENFILENAME) bool {
	ret, _, _ := procGetSaveFileName.Call(
		uintptr(unsafe.Pointer(ofn)))

	return ret != 0
}


// ff:
func CommDlgExtendedError() uint {
	ret, _, _ := procCommDlgExtendedError.Call()

	return uint(ret)
}
