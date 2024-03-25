//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import (
	"syscall"

	"github.com/888go/wails/internal/frontend/desktop/windows/winc/w32"
)

// 私有全局变量。
var (
	gAppInstance        w32.HINSTANCE
	gControllerRegistry map[w32.HWND]Controller
	gRegisteredClasses  []string
)

// 公共全局变量。
var (
	GeneralWndprocCallBack = syscall.NewCallback(generalWndProc)
	DefaultFont            *Font
)
