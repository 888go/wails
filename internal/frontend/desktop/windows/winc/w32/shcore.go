//go:build windows

package w32

import (
	"syscall"
	"unsafe"
)

var (
	modshcore = syscall.NewLazyDLL("shcore.dll")

	procGetDpiForMonitor = modshcore.NewProc("GetDpiForMonitor")
)


// ff:
func HasGetDPIForMonitorFunc() bool {
	err := procGetDpiForMonitor.Find()
	return err == nil
}


// ff:
// dpiY:
// dpiX:
// dpiType:
// hmonitor:
func GetDPIForMonitor(hmonitor HMONITOR, dpiType MONITOR_DPI_TYPE, dpiX *UINT, dpiY *UINT) uintptr {
	ret, _, _ := procGetDpiForMonitor.Call(
		hmonitor,
		uintptr(dpiType),
		uintptr(unsafe.Pointer(dpiX)),
		uintptr(unsafe.Pointer(dpiY)))

	return ret
}
