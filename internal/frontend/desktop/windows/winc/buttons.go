//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import (
	"fmt"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

type Button struct {
	ControlBase
	onClick EventManager
}

func (bt *Button) OnClick() *EventManager {
	return &bt.onClick
}

func (bt *Button) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_COMMAND:
		bt.onClick.Fire(NewEvent(bt, nil))
		/*case w32.WM_LBUTTONDOWN:
			w32.SetCapture(bt.Handle())
		case w32.WM_LBUTTONUP:
			w32.ReleaseCapture()*/
		/*case win.WM_GETDLGCODE:
		println("GETDLGCODE")*/
	}
	return w32.DefWindowProc(bt.hwnd, msg, wparam, lparam)
	// 返回 bt.W32Control.WndProc 方法处理的消息结果，参数分别为 msg、wparam 和 lparam
}

func (bt *Button) Checked() bool {
	result := w32.SendMessage(bt.hwnd, w32.BM_GETCHECK, 0, 0)
	return result == w32.BST_CHECKED
}

func (bt *Button) SetChecked(checked bool) {
	wparam := w32.BST_CHECKED
	if !checked {
		wparam = w32.BST_UNCHECKED
	}
	w32.SendMessage(bt.hwnd, w32.BM_SETCHECK, uintptr(wparam), 0)
}

// SetIcon 设置按钮上的图标。建议使用的图标大小为32x32像素，颜色深度为32位。
func (bt *Button) SetIcon(ico *Icon) {
	w32.SendMessage(bt.hwnd, w32.BM_SETIMAGE, w32.IMAGE_ICON, uintptr(ico.handle))
}

func (bt *Button) SetResIcon(iconID uint16) {
	if ico, err := NewIconFromResource(GetAppInstance(), iconID); err == nil {
		bt.SetIcon(ico)
		return
	}
	panic(fmt.Sprintf("missing icon with icon ID: %d", iconID))
}

type PushButton struct {
	Button
}

func NewPushButton(parent Controller) *PushButton {
	pb := new(PushButton)

	pb.InitControl("BUTTON", parent, 0, w32.BS_PUSHBUTTON|w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD)
	RegMsgHandler(pb)

	pb.SetFont(DefaultFont)
	pb.SetText("Button")
	pb.SetSize(100, 22)

	return pb
}

// SetDefault 用于对话框设置默认按钮。
func (pb *PushButton) SetDefault() {
	pb.SetAndClearStyleBits(w32.BS_DEFPUSHBUTTON, w32.BS_PUSHBUTTON)
}

// IconButton 不显示文本，需要调用 SetResIcon 方法。
type IconButton struct {
	Button
}

func NewIconButton(parent Controller) *IconButton {
	pb := new(IconButton)

	pb.InitControl("BUTTON", parent, 0, w32.BS_ICON|w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD)
	RegMsgHandler(pb)

	pb.SetFont(DefaultFont)
	// 即便text被设置，它也不会被显示
	pb.SetText("")
	pb.SetSize(100, 22)

	return pb
}

type CheckBox struct {
	Button
}

func NewCheckBox(parent Controller) *CheckBox {
	cb := new(CheckBox)

	cb.InitControl("BUTTON", parent, 0, w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD|w32.BS_AUTOCHECKBOX)
	RegMsgHandler(cb)

	cb.SetFont(DefaultFont)
	cb.SetText("CheckBox")
	cb.SetSize(100, 22)

	return cb
}

type RadioButton struct {
	Button
}

func NewRadioButton(parent Controller) *RadioButton {
	rb := new(RadioButton)

	rb.InitControl("BUTTON", parent, 0, w32.WS_TABSTOP|w32.WS_VISIBLE|w32.WS_CHILD|w32.BS_AUTORADIOBUTTON)
	RegMsgHandler(rb)

	rb.SetFont(DefaultFont)
	rb.SetText("RadioButton")
	rb.SetSize(100, 22)

	return rb
}

type GroupBox struct {
	Button
}

func NewGroupBox(parent Controller) *GroupBox {
	gb := new(GroupBox)

	gb.InitControl("BUTTON", parent, 0, w32.WS_CHILD|w32.WS_VISIBLE|w32.WS_GROUP|w32.BS_GROUPBOX)
	RegMsgHandler(gb)

	gb.SetFont(DefaultFont)
	gb.SetText("GroupBox")
	gb.SetSize(100, 100)

	return gb
}
