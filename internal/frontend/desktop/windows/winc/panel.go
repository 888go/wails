//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 */

package winc

import (
	"fmt"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc/w32"
)

type Panel struct {
	ControlBase
	layoutMng LayoutManager
}

func NewPanel(parent Controller) *Panel {
	pa := new(Panel)

	RegClassOnlyOnce("winc_Panel")
	pa.hwnd = CreateWindow("winc_Panel", parent, w32.WS_EX_CONTROLPARENT, w32.WS_CHILD|w32.WS_VISIBLE)
	pa.parent = parent
	RegMsgHandler(pa)

	pa.SetFont(DefaultFont)
	pa.SetText("")
	pa.SetSize(200, 65)
	return pa
}

// SetLayout 面板实现了 DockAllow 接口。
func (pa *Panel) SetLayout(mng LayoutManager) {
	pa.layoutMng = mng
}

func (pa *Panel) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_SIZE, w32.WM_PAINT:
		if pa.layoutMng != nil {
			pa.layoutMng.Update()
		}
	}
	return w32.DefWindowProc(pa.hwnd, msg, wparam, lparam)
}

var errorPanelPen = NewPen(w32.PS_GEOMETRIC, 2, NewSolidColorBrush(RGB(255, 128, 128)))
var errorPanelOkPen = NewPen(w32.PS_GEOMETRIC, 2, NewSolidColorBrush(RGB(220, 220, 220)))

// ErrorPanel 用于显示错误或重要消息。
// 它旨在从屏幕上的其他控件中突出显示。
type ErrorPanel struct {
	ControlBase
	pen    *Pen
	margin int
}

// NewErrorPanel.
func NewErrorPanel(parent Controller) *ErrorPanel {
	f := new(ErrorPanel)
	f.init(parent)

	f.SetFont(DefaultFont)
	f.SetText("No errors")
	f.SetSize(200, 65)
	f.margin = 5
	f.pen = errorPanelOkPen
	return f
}

func (epa *ErrorPanel) init(parent Controller) {
	RegClassOnlyOnce("winc_ErrorPanel")

	epa.hwnd = CreateWindow("winc_ErrorPanel", parent, w32.WS_EX_CONTROLPARENT, w32.WS_CHILD|w32.WS_VISIBLE)
	epa.parent = parent

	RegMsgHandler(epa)
}

func (epa *ErrorPanel) SetMargin(margin int) {
	epa.margin = margin
}

func (epa *ErrorPanel) Printf(format string, v ...interface{}) {
	epa.SetText(fmt.Sprintf(format, v...))
	epa.ShowAsError(false)
}

func (epa *ErrorPanel) Errorf(format string, v ...interface{}) {
	epa.SetText(fmt.Sprintf(format, v...))
	epa.ShowAsError(true)
}

func (epa *ErrorPanel) ShowAsError(show bool) {
	if show {
		epa.pen = errorPanelPen
	} else {
		epa.pen = errorPanelOkPen
	}
	epa.Invalidate(true)
}

func (epa *ErrorPanel) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_ERASEBKGND:
		canvas := NewCanvasFromHDC(w32.HDC(wparam))
		r := epa.Bounds()
		r.rect.Left += int32(epa.margin)
		r.rect.Right -= int32(epa.margin)
		r.rect.Top += int32(epa.margin)
		r.rect.Bottom -= int32(epa.margin)
		// 旧代码中使用了 NewSystemColorBrush(w32.COLOR_BTNFACE)
		canvas.DrawFillRect(r, epa.pen, NewSystemColorBrush(w32.COLOR_WINDOW))

		r.rect.Left += 5
		canvas.DrawText(epa.Text(), r, 0, epa.Font(), RGB(0, 0, 0))
		canvas.Dispose()
		return 1
	}
	return w32.DefWindowProc(epa.hwnd, msg, wparam, lparam)
}

// MultiPanel 包含多个面板，并且只显示其中的一个。
type MultiPanel struct {
	ControlBase
	current int
	panels  []*Panel
}

func NewMultiPanel(parent Controller) *MultiPanel {
	mpa := new(MultiPanel)

	RegClassOnlyOnce("winc_MultiPanel")
	mpa.hwnd = CreateWindow("winc_MultiPanel", parent, w32.WS_EX_CONTROLPARENT, w32.WS_CHILD|w32.WS_VISIBLE)
	mpa.parent = parent
	RegMsgHandler(mpa)

	mpa.SetFont(DefaultFont)
	mpa.SetText("")
	mpa.SetSize(300, 200)
	mpa.current = -1
	return mpa
}

func (mpa *MultiPanel) Count() int { return len(mpa.panels) }

// AddPanel 将面板添加到内部列表中，其中第一个面板可见，其他所有面板均隐藏。
func (mpa *MultiPanel) AddPanel(panel *Panel) {
	if len(mpa.panels) > 0 {
		panel.Hide()
	}
	mpa.current = 0
	mpa.panels = append(mpa.panels, panel)
}

// ReplacePanel 用于替换面板，适用于屏幕上的控件刷新。
func (mpa *MultiPanel) ReplacePanel(index int, panel *Panel) {
	mpa.panels[index] = panel
}

// DeletePanel 删除面板
func (mpa *MultiPanel) DeletePanel(index int) {
	mpa.panels = append(mpa.panels[:index], mpa.panels[index+1:]...)
}

func (mpa *MultiPanel) Current() int {
	return mpa.current
}

func (mpa *MultiPanel) SetCurrent(index int) {
	if index >= len(mpa.panels) {
		panic("index greater than number of panels")
	}
	if mpa.current == -1 {
		panic("no current panel, add panels first")
	}
	for i := range mpa.panels {
		if i != index {
			mpa.panels[i].Hide()
			mpa.panels[i].Invalidate(true)
		}
	}
	mpa.panels[index].Show()
	mpa.panels[index].Invalidate(true)
	mpa.current = index
}

func (mpa *MultiPanel) WndProc(msg uint32, wparam, lparam uintptr) uintptr {
	switch msg {
	case w32.WM_SIZE:
		// 调整容器中面板的大小
		for _, p := range mpa.panels {
			p.SetPos(0, 0)
			p.SetSize(mpa.Size())
		}
	}
	return w32.DefWindowProc(mpa.hwnd, msg, wparam, lparam)
}
