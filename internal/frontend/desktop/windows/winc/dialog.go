//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 * Copyright (C) 2010-2013 Allen Dang. All Rights Reserved.
 */

package winc

import "github.com/888go/wails/internal/frontend/desktop/windows/winc/w32"

// 此对话框将以最高级窗口的方式显示，直到被关闭。
// 同时，它会禁用父窗口，使其无法被点击。
type Dialog struct {
	Form
	isModal bool

	btnOk     *PushButton
	btnCancel *PushButton

	onLoad   EventManager
	onOk     EventManager
	onCancel EventManager
}

func NewDialog(parent Controller) *Dialog {
	dlg := new(Dialog)

	dlg.isForm = true
	dlg.isModal = true
	RegClassOnlyOnce("winc_Dialog")

	dlg.hwnd = CreateWindow("winc_Dialog", parent, w32.WS_EX_CONTROLPARENT, /* IMPORTANT */
		w32.WS_SYSMENU|w32.WS_CAPTION|w32.WS_THICKFRAME /*|w32.WS_BORDER|w32.WS_POPUP*/)
	dlg.parent = parent

	// 如果图标资源未嵌入到二进制文件中，dlg可能会失败
	if ico, err := NewIconFromResource(GetAppInstance(), uint16(AppIconID)); err == nil {
		dlg.SetIcon(0, ico)
	}

	// Dlg 在用户开始输入时强制显示焦点矩形框。
	w32.SendMessage(dlg.hwnd, w32.WM_CHANGEUISTATE, w32.UIS_INITIALIZE, 0)
	RegMsgHandler(dlg)

	dlg.SetFont(DefaultFont)
	dlg.SetText("Form")
	dlg.SetSize(200, 100)
	return dlg
}

func (dlg *Dialog) SetModal(modal bool) {
	dlg.isModal = modal
}

// SetButtons 将对话框事件与按钮连接起来。 btnCancel 可以为 nil。
// （该函数用于设置或绑定按钮到相应的对话框事件，其中 btnCancel 参数表示取消按钮，如果不需要可以传入 nil 值。）
func (dlg *Dialog) SetButtons(btnOk *PushButton, btnCancel *PushButton) {
	dlg.btnOk = btnOk
	dlg.btnOk.SetDefault()
	dlg.btnCancel = btnCancel
}

// Events
func (dlg *Dialog) OnLoad() *EventManager {
	return &dlg.onLoad
}

func (dlg *Dialog) OnOk() *EventManager {
	return &dlg.onOk
}

func (dlg *Dialog) OnCancel() *EventManager {
	return &dlg.onCancel
}

// PreTranslateMessage 处理与对话框相关的特定消息。非常重要。
func (dlg *Dialog) PreTranslateMessage(msg *w32.MSG) bool {
	if msg.Message >= w32.WM_KEYFIRST && msg.Message <= w32.WM_KEYLAST {
		if w32.IsDialogMessage(dlg.hwnd, msg) {
			return true
		}
	}
	return false
}

// ShowDialog 为对话框窗口执行特殊设置。
func (dlg *Dialog) Show() {
	if dlg.isModal {
		dlg.Parent().SetEnabled(false)
	}
	dlg.onLoad.Fire(NewEvent(dlg, nil))
	dlg.Form.Show()
}

// 当你完成对话框操作后，关闭它。
func (dlg *Dialog) Close() {
	if dlg.isModal {
		dlg.Parent().SetEnabled(true)
	}
	dlg.ControlBase.Close()
}

func (dlg *Dialog) cancel() {
	if dlg.btnCancel != nil {
		dlg.btnCancel.onClick.Fire(NewEvent(dlg.btnCancel, nil))
	}
	dlg.onCancel.Fire(NewEvent(dlg, nil))
}

func (dlg *Dialog) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_COMMAND:
		switch w32.LOWORD(uint32(wparam)) {
		case w32.IDOK:
			if dlg.btnOk != nil {
				dlg.btnOk.onClick.Fire(NewEvent(dlg.btnOk, nil))
			}
			dlg.onOk.Fire(NewEvent(dlg, nil))
			return w32.TRUE

		case w32.IDCANCEL:
			dlg.cancel()
			return w32.TRUE
		}

	case w32.WM_CLOSE:
		dlg.cancel() // 使用onCancel或dlg.btnCancel.OnClick来关闭
		return 0

	case w32.WM_DESTROY:
		if dlg.isModal {
			dlg.Parent().SetEnabled(true)
		}
	}
	return w32.DefWindowProc(dlg.hwnd, msg, wparam, lparam)
}
