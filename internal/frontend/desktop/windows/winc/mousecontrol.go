//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 */

package winc

import (
	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

// MouseControl 用于创建自定义控件，这些控件需要鼠标悬停或鼠标离开事件。
type MouseControl struct {
	ControlBase
	isMouseLeft bool
}

func (cc *MouseControl) Init(parent Controller, className string, exStyle, style uint) {
	RegClassOnlyOnce(className)
	cc.hwnd = CreateWindow(className, parent, exStyle, style)
	cc.parent = parent
	RegMsgHandler(cc)

	cc.isMouseLeft = true
	cc.SetFont(DefaultFont)
}

func (cc *MouseControl) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	sender := GetMsgHandler(cc.hwnd)
	switch msg {
	case w32.WM_CREATE:
		internalTrackMouseEvent(cc.hwnd)
		cc.onCreate.Fire(NewEvent(sender, nil))
	case w32.WM_CLOSE:
		cc.onClose.Fire(NewEvent(sender, nil))
	case w32.WM_MOUSEMOVE:
		//if cc.isMouseLeft {

		cc.onMouseHover.Fire(NewEvent(sender, nil))
		// 内部调用TrackMouseEvent函数，参数为cc.hwnd
		cc.isMouseLeft = false

		//}
	case w32.WM_MOUSELEAVE:
		cc.onMouseLeave.Fire(NewEvent(sender, nil))
		cc.isMouseLeft = true
	}
	return w32.DefWindowProc(cc.hwnd, msg, wparam, lparam)
}
