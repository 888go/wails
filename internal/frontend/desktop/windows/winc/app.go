//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */
package winc

import (
	"runtime"
	"unsafe"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

var (
// 资源编译工具将app.ico的ID设置为3
// rsrc命令使用manifest文件（app.manifest）和ico图标文件（app.ico）生成资源文件rsrc.syso
	AppIconID = 3
)

func init() {
	runtime.LockOSThread()

	gAppInstance = w32.GetModuleHandle("")
	if gAppInstance == 0 {
		panic("Error occurred in App.Init")
	}

	// 初始化公共控件
	var initCtrls w32.INITCOMMONCONTROLSEX
	initCtrls.DwSize = uint32(unsafe.Sizeof(initCtrls))
	initCtrls.DwICC =
		w32.ICC_LISTVIEW_CLASSES | w32.ICC_PROGRESS_CLASS | w32.ICC_TAB_CLASSES |
			w32.ICC_TREEVIEW_CLASSES | w32.ICC_BAR_CLASSES

	w32.InitCommonControlsEx(&initCtrls)
}

// SetAppIconID 为应用程序窗口设置资源图标ID。

// ff:
// appIconID:
func SetAppIcon(appIconID int) {
	AppIconID = appIconID
}


// ff:
func GetAppInstance() w32.HINSTANCE {
	return gAppInstance
}


// ff:
// msg:
func PreTranslateMessage(msg *w32.MSG) bool {
// 此函数由MessageLoop调用。它处理键盘加速键，并对键盘和鼠标事件调用Controller.PreTranslateMessage。

	processed := false

	if (msg.Message >= w32.WM_KEYFIRST && msg.Message <= w32.WM_KEYLAST) ||
		(msg.Message >= w32.WM_MOUSEFIRST && msg.Message <= w32.WM_MOUSELAST) {

		if msg.Hwnd != 0 {
			if controller := GetMsgHandler(msg.Hwnd); controller != nil {
				// 在父级链中搜索预翻译的消息。
				for p := controller; p != nil; p = p.Parent() {

					if processed = p.PreTranslateMessage(msg); processed {
						break
					}
				}
			}
		}
	}

	return processed
}

// RunMainLoop 在主应用程序循环中处理消息。

// ff:
func RunMainLoop() int {
	m := (*w32.MSG)(unsafe.Pointer(w32.GlobalAlloc(0, uint32(unsafe.Sizeof(w32.MSG{})))))
	defer w32.GlobalFree(w32.HGLOBAL(unsafe.Pointer(m)))

	for w32.GetMessage(m, 0, 0, 0) != 0 {

		if !PreTranslateMessage(m) {
			w32.TranslateMessage(m)
			w32.DispatchMessage(m)
		}
	}

	w32.GdiplusShutdown()
	return int(m.WParam)
}

// PostMessages 处理最近的消息。有时有助于即时窗口刷新。

// ff:
func PostMessages() {
	m := (*w32.MSG)(unsafe.Pointer(w32.GlobalAlloc(0, uint32(unsafe.Sizeof(w32.MSG{})))))
	defer w32.GlobalFree(w32.HGLOBAL(unsafe.Pointer(m)))

	for i := 0; i < 10; i++ {
		if w32.GetMessage(m, 0, 0, 0) != 0 {
			if !PreTranslateMessage(m) {
				w32.TranslateMessage(m)
				w32.DispatchMessage(m)
			}
		}
	}
}


// ff:
func Exit() {
	w32.PostQuitMessage(0)
}
