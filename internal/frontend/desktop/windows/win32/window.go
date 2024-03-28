//go:build windows

package win32

import (
	"fmt"
	"log"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/wailsapp/wails/v2/internal/frontend/desktop/windows/winc"
)

const (
	WS_MAXIMIZE = 0x01000000
	WS_MINIMIZE = 0x20000000

	GWL_STYLE = -16

	MONITOR_DEFAULTTOPRIMARY = 0x00000001
)

const (
	SW_HIDE            = 0
	SW_NORMAL          = 1
	SW_SHOWNORMAL      = 1
	SW_SHOWMINIMIZED   = 2
	SW_MAXIMIZE        = 3
	SW_SHOWMAXIMIZED   = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
)

const (
	GCLP_HBRBACKGROUND int32 = -10
)

// Power
const (
	// WM_POWERBROADCAST - 当发生电源管理事件时，用于通知应用程序。
	WM_POWERBROADCAST = 536

	// PBT_APMPOWERSTATUSCHANGE - 电源状态已改变。
	PBT_APMPOWERSTATUSCHANGE = 10

	// PBT_APMRESUMEAUTOMATIC - 操作正在从低功耗状态自动恢复。每当系统恢复时，都会发送此消息。
	PBT_APMRESUMEAUTOMATIC = 18

	// PBT_APMRESUMESUSPEND - 操作从低功耗状态恢复。如果恢复是由用户输入触发的，例如按下某个键，则在接收到 PBT_APMRESUMEAUTOMATIC 之后发送此消息。
	PBT_APMRESUMESUSPEND = 7

	// PBT_APMSUSPEND - 系统正在进行挂起操作。
	PBT_APMSUSPEND = 4

	// PBT_POWERSETTINGCHANGE - 收到了电源设置更改事件。
	PBT_POWERSETTINGCHANGE = 32787
)

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/bb773244.aspx
// （该注释仅提供了一个参考链接，没有具体针对某段代码进行解释，故不涉及翻译具体注释内容。原始链接为微软开发者网络（MSDN）上关于某个Windows桌面开发相关主题的技术文档。）
type MARGINS struct {
	CxLeftWidth, CxRightWidth, CyTopHeight, CyBottomHeight int32
}

// 参考链接：http://msdn.microsoft.com/en-us/library/windows/desktop/dd162897.aspx （MSDN官网关于某个Windows桌面开发API的详细说明）
type RECT struct {
	Left, Top, Right, Bottom int32
}

// 参考网址：http://msdn.microsoft.com/en-us/library/windows/desktop/dd145065.aspx
// （由于您提供的代码片段中没有具体的代码内容，此处仅翻译了注释部分，该注释表明这是一个参考链接，指向的是MSDN（微软开发者网络）上关于某个Windows桌面开发相关主题的技术文档。）
type MONITORINFO struct {
	CbSize    uint32
	RcMonitor RECT
	RcWork    RECT
	DwFlags   uint32
}

func ExtendFrameIntoClientArea(hwnd uintptr, extend bool) {
// -1: 添加默认窗口边框样式（如Windows 11的Aero阴影和圆角等）
//     如果窗口是透明或半透明，也会显示标题栏按钮，但这些按钮无法正常工作。
//  0: 添加默认窗口边框样式，但不包括Aero阴影，并且不显示标题栏按钮。
//  1: 添加默认窗口边框样式（如Windows 11的Aero阴影和圆角等），但如果窗口是透明或半透明，则不显示标题栏按钮。
	var margins MARGINS
	if extend {
		margins = MARGINS{1, 1, 1, 1} // 仅扩展1像素以具有默认边框样式，但无标题栏按钮
	}
	if err := dwmExtendFrameIntoClientArea(hwnd, &margins); err != nil {
		log.Fatal(fmt.Errorf("DwmExtendFrameIntoClientArea failed: %s", err))
	}
}

func IsVisible(hwnd uintptr) bool {
	ret, _, _ := procIsWindowVisible.Call(hwnd)
	return ret != 0
}

func IsWindowFullScreen(hwnd uintptr) bool {
	wRect := GetWindowRect(hwnd)
	m := MonitorFromWindow(hwnd, MONITOR_DEFAULTTOPRIMARY)
	var mi MONITORINFO
	mi.CbSize = uint32(unsafe.Sizeof(mi))
	if !GetMonitorInfo(m, &mi) {
		return false
	}
	return wRect.Left == mi.RcMonitor.Left &&
		wRect.Top == mi.RcMonitor.Top &&
		wRect.Right == mi.RcMonitor.Right &&
		wRect.Bottom == mi.RcMonitor.Bottom
}

func IsWindowMaximised(hwnd uintptr) bool {
	style := uint32(getWindowLong(hwnd, GWL_STYLE))
	return style&WS_MAXIMIZE != 0
}
func IsWindowMinimised(hwnd uintptr) bool {
	style := uint32(getWindowLong(hwnd, GWL_STYLE))
	return style&WS_MINIMIZE != 0
}

func RestoreWindow(hwnd uintptr) {
	showWindow(hwnd, SW_RESTORE)
}

func ShowWindow(hwnd uintptr) {
	showWindow(hwnd, SW_SHOW)
}

func ShowWindowMaximised(hwnd uintptr) {
	showWindow(hwnd, SW_MAXIMIZE)
}
func ShowWindowMinimised(hwnd uintptr) {
	showWindow(hwnd, SW_MINIMIZE)
}

func SetBackgroundColour(hwnd uintptr, r, g, b uint8) {
	col := winc.RGB(r, g, b)
	hbrush, _, _ := procCreateSolidBrush.Call(uintptr(col))
	setClassLongPtr(hwnd, GCLP_HBRBACKGROUND, hbrush)
}

func IsWindowNormal(hwnd uintptr) bool {
	return !IsWindowMaximised(hwnd) && !IsWindowMinimised(hwnd) && !IsWindowFullScreen(hwnd)
}

func dwmExtendFrameIntoClientArea(hwnd uintptr, margins *MARGINS) error {
	ret, _, _ := procDwmExtendFrameIntoClientArea.Call(
		hwnd,
		uintptr(unsafe.Pointer(margins)))

	if ret != 0 {
		return syscall.GetLastError()
	}

	return nil
}

func setClassLongPtr(hwnd uintptr, param int32, val uintptr) bool {
	proc := procSetClassLongPtr
	if strconv.IntSize == 32 {
		/*
			https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setclasslongptrw
			Note: 	To write code that is compatible with both 32-bit and 64-bit Windows, use SetClassLongPtr.
					When compiling for 32-bit Windows, SetClassLongPtr is defined as a call to the SetClassLong function

			=> We have to do this dynamically when directly calling the DLL procedures
		*/
		proc = procSetClassLong
	}

	ret, _, _ := proc.Call(
		hwnd,
		uintptr(param),
		val,
	)
	return ret != 0
}

func getWindowLong(hwnd uintptr, index int) int32 {
	ret, _, _ := procGetWindowLong.Call(
		hwnd,
		uintptr(index))

	return int32(ret)
}

func showWindow(hwnd uintptr, cmdshow int) bool {
	ret, _, _ := procShowWindow.Call(
		hwnd,
		uintptr(cmdshow))
	return ret != 0
}

func GetWindowRect(hwnd uintptr) *RECT {
	var rect RECT
	procGetWindowRect.Call(
		hwnd,
		uintptr(unsafe.Pointer(&rect)))

	return &rect
}

func MonitorFromWindow(hwnd uintptr, dwFlags uint32) HMONITOR {
	ret, _, _ := procMonitorFromWindow.Call(
		hwnd,
		uintptr(dwFlags),
	)
	return HMONITOR(ret)
}

func GetMonitorInfo(hMonitor HMONITOR, lmpi *MONITORINFO) bool {
	ret, _, _ := procGetMonitorInfo.Call(
		uintptr(hMonitor),
		uintptr(unsafe.Pointer(lmpi)),
	)
	return ret != 0
}
