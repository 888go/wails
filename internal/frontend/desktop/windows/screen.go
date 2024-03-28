//go:build windows
// +build windows

package windows

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/pkg/errors"
	"github.com/888go/wails/internal/frontend/desktop/windows/winc"
	"github.com/888go/wails/internal/frontend/desktop/windows/winc/w32"
)


// ff:
// second:
// first:
func MonitorsEqual(first w32.MONITORINFO, second w32.MONITORINFO) bool {
// 检查确保所有字段都相同。
// 一个更简洁的方法是检查设备的身份。但我没找到使用win32 API实现该功能的方法。
	return first.DwFlags == second.DwFlags &&
		first.RcMonitor.Top == second.RcMonitor.Top &&
		first.RcMonitor.Bottom == second.RcMonitor.Bottom &&
		first.RcMonitor.Right == second.RcMonitor.Right &&
		first.RcMonitor.Left == second.RcMonitor.Left &&
		first.RcWork.Top == second.RcWork.Top &&
		first.RcWork.Bottom == second.RcWork.Bottom &&
		first.RcWork.Right == second.RcWork.Right &&
		first.RcWork.Left == second.RcWork.Left
}


// ff:
// hMonitor:
func GetMonitorInfo(hMonitor w32.HMONITOR) (*w32.MONITORINFO, error) {
// 该段代码改编自 winc.utils.getMonitorInfo，待办：将其添加至 win32 库中
// 参考文档：
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-getmonitorinfoa
// （注：此网址为微软官方文档，关于 Windows API 函数 GetMonitorInfoA 的说明）

	var info w32.MONITORINFO
	info.CbSize = uint32(unsafe.Sizeof(info))
	succeeded := w32.GetMonitorInfo(hMonitor, &info)
	if !succeeded {
		return &info, errors.New("Windows call to getMonitorInfo failed")
	}
	return &info, nil
}


// ff:
// screenContainer:
// lprcMonitor:
// hdcMonitor:
// hMonitor:
func EnumProc(hMonitor w32.HMONITOR, hdcMonitor w32.HDC, lprcMonitor *w32.RECT, screenContainer *ScreenContainer) uintptr {
	// 该代码段改编自 StackOverflow 网站上的回答：https://stackoverflow.com/a/23492886/4188138

// 为了更好地理解此函数，请查看以下页面的文档
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-enumdisplaymonitors // （中文版链接，可能需要手动添加或替换）
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nc-winuser-monitorenumproc
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/ns-winuser-monitorinfo
// https://docs.microsoft.com/zh-cn/windows/win32/api/winuser/nf-winuser-monitorfromwindow
// 上述链接分别对应：
// - `EnumDisplayMonitors` 函数的说明
// - `MONITORENUMPROC` 函数指针类型的说明
// - `MONITORINFO` 结构体的定义
// - `MonitorFromWindow` 函数的说明

	ourMonitorData := Screen{}
	currentMonHndl := w32.MonitorFromWindow(screenContainer.mainWinHandle, w32.MONITOR_DEFAULTTONEAREST)
	currentMonInfo, currErr := GetMonitorInfo(currentMonHndl)

	if currErr != nil {
		screenContainer.errors = append(screenContainer.errors, currErr)
		screenContainer.monitors = append(screenContainer.monitors, Screen{})
		// 不确定返回false会带来什么后果，所以我们先返回true并自行处理这个问题
		return w32.TRUE
	}

	monInfo, err := GetMonitorInfo(hMonitor)
	if err != nil {
		screenContainer.errors = append(screenContainer.errors, err)
		screenContainer.monitors = append(screenContainer.monitors, Screen{})
		return w32.TRUE
	}

	width := lprcMonitor.Right - lprcMonitor.Left
	height := lprcMonitor.Bottom - lprcMonitor.Top
	ourMonitorData.IsPrimary = monInfo.DwFlags&w32.MONITORINFOF_PRIMARY == 1
	ourMonitorData.Height = int(height)
	ourMonitorData.Width = int(width)
	ourMonitorData.IsCurrent = MonitorsEqual(*currentMonInfo, *monInfo)

	ourMonitorData.PhysicalSize.Width = int(width)
	ourMonitorData.PhysicalSize.Height = int(height)

	var dpiX, dpiY uint
	w32.GetDPIForMonitor(hMonitor, w32.MDT_EFFECTIVE_DPI, &dpiX, &dpiY)
	if dpiX == 0 || dpiY == 0 {
		screenContainer.errors = append(screenContainer.errors, fmt.Errorf("unable to get DPI for screen"))
		screenContainer.monitors = append(screenContainer.monitors, Screen{})
		return w32.TRUE
	}
	ourMonitorData.Size.Width = winc.ScaleToDefaultDPI(ourMonitorData.PhysicalSize.Width, dpiX)
	ourMonitorData.Size.Height = winc.ScaleToDefaultDPI(ourMonitorData.PhysicalSize.Height, dpiY)

// 我们需要一个容器的原因是我们并不知道这个函数会被调用多少次
// 这个 "append" 调用可能潜在地进行内存分配并重写 monitors 指针。所以我们把指针保存在 screenContainer.monitors 中
// 并在所有 EnumProc 调用结束后获取其值
// 如果 EnumProc 是多线程的，这可能会存在问题。尽管我认为并不是这样。
	screenContainer.monitors = append(screenContainer.monitors, ourMonitorData)
	// 让我们保持 screenContainer.errors 与 screenContainer.monitors 的大小相同，以便后续如有必要时可以将它们对应匹配起来
	screenContainer.errors = append(screenContainer.errors, nil)
	return w32.TRUE
}

type ScreenContainer struct {
	monitors      []Screen
	errors        []error
	mainWinHandle w32.HWND
}


// ff:
// mainWinHandle:
func GetAllScreens(mainWinHandle w32.HWND) ([]Screen, error) {
	// TODO：通过在Windows与运行时之间建立合适的数据共享机制，修复容器共享的临时解决方案
	monitorContainer := ScreenContainer{mainWinHandle: mainWinHandle}
	returnErr := error(nil)
	errorStrings := []string{}

	dc := w32.GetDC(0)
	defer w32.ReleaseDC(0, dc)
	succeeded := w32.EnumDisplayMonitors(dc, nil, syscall.NewCallback(EnumProc), unsafe.Pointer(&monitorContainer))
	if !succeeded {
		return monitorContainer.monitors, errors.New("Windows call to EnumDisplayMonitors failed")
	}
	for idx, err := range monitorContainer.errors {
		if err != nil {
			errorStrings = append(errorStrings, fmt.Sprintf("Error from monitor #%v, %v", idx+1, err))
		}
	}

	if len(errorStrings) > 0 {
		returnErr = fmt.Errorf("%v errors encountered: %v", len(errorStrings), errorStrings)
	}
	return monitorContainer.monitors, returnErr
}
