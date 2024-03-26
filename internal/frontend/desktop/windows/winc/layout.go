//go:build windows

/*
 * Copyright (C) 2019 The Winc Authors. All Rights Reserved.
 */

package winc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"unsafe"

	"github.com/888go/wails/internal/frontend/desktop/windows/winc/w32"
)

// 可停靠组件必须满足接口要求才能被停靠。
type Dockable interface {
	Handle() w32.HWND

	Pos() (x, y int)
	Width() int
	Height() int
	Visible() bool

	SetPos(x, y int)
	SetSize(width, height int)

	OnMouseMove() *EventManager
	OnLBUp() *EventManager
}

// DockAllow 是窗口、面板或其他满足接口的组件。
type DockAllow interface {
	Handle() w32.HWND
	ClientWidth() int
	ClientHeight() int
	SetLayout(mng LayoutManager)
}

// 各种布局管理器
type Direction int

const (
	Top Direction = iota
	Bottom
	Left
	Right
	Fill
)

type LayoutControl struct {
	child Dockable
	dir   Direction
}

type LayoutControls []*LayoutControl

type SimpleDock struct {
	parent      DockAllow
	layoutCtl   LayoutControls
	loadedState bool
}

// DockState 会从 json 中保存和加载
type CtlState struct {
	X, Y, Width, Height int
}

type LayoutState struct {
	WindowState string
	Controls    []*CtlState
}

func (lc LayoutControls) Len() int           { return len(lc) }
func (lc LayoutControls) Swap(i, j int)      { lc[i], lc[j] = lc[j], lc[i] }
func (lc LayoutControls) Less(i, j int) bool { return lc[i].dir < lc[j].dir }

func NewSimpleDock(parent DockAllow) *SimpleDock {
	d := &SimpleDock{parent: parent}
	parent.SetLayout(d)
	return d
}

// 子控件的布局管理。
func (sd *SimpleDock) Dock(child Dockable, dir Direction) {
	sd.layoutCtl = append(sd.layoutCtl, &LayoutControl{child, dir})
}

// 保存布局的状态。仅适用于父级设置为主窗口的停靠栏。
func (sd *SimpleDock) SaveState(w io.Writer) error {
	var ls LayoutState

	var wp w32.WINDOWPLACEMENT
	wp.Length = uint32(unsafe.Sizeof(wp))
	if !w32.GetWindowPlacement(sd.parent.Handle(), &wp) {
		return fmt.Errorf("GetWindowPlacement failed")
	}

	ls.WindowState = fmt.Sprint(
		wp.Flags, wp.ShowCmd,
		wp.PtMinPosition.X, wp.PtMinPosition.Y,
		wp.PtMaxPosition.X, wp.PtMaxPosition.Y,
		wp.RcNormalPosition.Left, wp.RcNormalPosition.Top,
		wp.RcNormalPosition.Right, wp.RcNormalPosition.Bottom)

	for _, c := range sd.layoutCtl {
		x, y := c.child.Pos()
		w, h := c.child.Width(), c.child.Height()

		ctl := &CtlState{X: x, Y: y, Width: w, Height: h}
		ls.Controls = append(ls.Controls, ctl)
	}

	if err := json.NewEncoder(w).Encode(ls); err != nil {
		return err
	}

	return nil
}

// 加载布局状态。仅适用于父级设置为主窗体的停靠栏。
func (sd *SimpleDock) LoadState(r io.Reader) error {
	var ls LayoutState

	if err := json.NewDecoder(r).Decode(&ls); err != nil {
		return err
	}

	var wp w32.WINDOWPLACEMENT
	if _, err := fmt.Sscan(ls.WindowState,
		&wp.Flags, &wp.ShowCmd,
		&wp.PtMinPosition.X, &wp.PtMinPosition.Y,
		&wp.PtMaxPosition.X, &wp.PtMaxPosition.Y,
		&wp.RcNormalPosition.Left, &wp.RcNormalPosition.Top,
		&wp.RcNormalPosition.Right, &wp.RcNormalPosition.Bottom); err != nil {
		return err
	}
	wp.Length = uint32(unsafe.Sizeof(wp))

	if !w32.SetWindowPlacement(sd.parent.Handle(), &wp) {
		return fmt.Errorf("SetWindowPlacement failed")
	}

// 如果保存的布局中的控件数量与当前屏幕上的控件数量不匹配
// 意味着某些内容发生了变化，此时我们不重新加载
// 来自json的其余控件大小信息
	if len(sd.layoutCtl) != len(ls.Controls) {
		return nil
	}

	for i, c := range sd.layoutCtl {
		c.child.SetPos(ls.Controls[i].X, ls.Controls[i].Y)
		c.child.SetSize(ls.Controls[i].Width, ls.Controls[i].Height)
	}
	return nil
}

// SaveStateFile 便捷函数。
func (sd *SimpleDock) SaveStateFile(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	return sd.SaveState(f)
}

// LoadStateFile 加载状态文件，如果文件未找到，则忽略错误。
func (sd *SimpleDock) LoadStateFile(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return nil // 如果文件未找到或无法访问，则忽略它
	}
	return sd.LoadState(f)
}

// Update 在需要根据布局方向调整子项大小时被调用。
func (sd *SimpleDock) Update() {
	sort.Stable(sd.layoutCtl)

	x, y := 0, 0
	w, h := sd.parent.ClientWidth(), sd.parent.ClientHeight()
	winw, winh := w, h

	for _, c := range sd.layoutCtl {
		// 非可见控件不保留空间。
		if !c.child.Visible() {
			continue
		}

		switch c.dir {
		case Top:
			c.child.SetPos(x, y)
			c.child.SetSize(w, c.child.Height())
			h -= c.child.Height()
			y += c.child.Height()
		case Bottom:
			c.child.SetPos(x, winh-c.child.Height())
			c.child.SetSize(w, c.child.Height())
			h -= c.child.Height()
			winh -= c.child.Height()
		case Left:
			c.child.SetPos(x, y)
			c.child.SetSize(c.child.Width(), h)
			w -= c.child.Width()
			x += c.child.Width()
		case Right:
			c.child.SetPos(winw-c.child.Width(), y)
			c.child.SetSize(c.child.Width(), h)
			w -= c.child.Width()
			winw -= c.child.Width()
		case Fill:
			// fill available space
			c.child.SetPos(x, y)
			c.child.SetSize(w, h)
		}
		// c.child.Invalidate(true) // 清除c.child的缓存，参数true表示递归清除子节点的缓存
	}
}
