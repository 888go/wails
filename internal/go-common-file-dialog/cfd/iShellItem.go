//go:build windows
// +build windows

package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

var (
	procSHCreateItemFromParsingName = syscall.NewLazyDLL("Shell32.dll").NewProc("SHCreateItemFromParsingName")
	iidShellItem                    = ole.NewGUID("43826d1e-e718-42ee-bc55-a1e261c37bfe")
)

type iShellItem struct {
	vtbl *iShellItemVtbl
}

type iShellItemVtbl struct {
	iUnknownVtbl
	BindToHandler  uintptr
	GetParent      uintptr
	GetDisplayName uintptr // 函数签名：(sigdnName SIGDN, ppszName *LPWSTR) HRESULT
// 
// 对于给定的SIGDN枚举值sigdnName，此函数接收一个指向LPWSTR类型的指针(ppszName)，并返回一个HRESULT值。
// 
// 参数：
// - sigdnName: 类型为SIGDN的枚举值，用于指定文件或对象的显示名称格式。
// - ppszName: 指向LPWSTR类型的指针，用于接收根据指定格式转换后的文件或对象的名称。
// 返回值：
// - HRESULT：表示函数执行成功或失败的状态。在Windows API中，HRESULT是一个32位值，用于表示系统组件（如COM对象）的方法调用结果。
	GetAttributes  uintptr
	Compare        uintptr
}

func newIShellItem(path string) (*iShellItem, error) {
	var shellItem *iShellItem
	pathPtr := ole.SysAllocString(path)
	ret, _, _ := procSHCreateItemFromParsingName.Call(
		uintptr(unsafe.Pointer(pathPtr)),
		0,
		uintptr(unsafe.Pointer(iidShellItem)),
		uintptr(unsafe.Pointer(&shellItem)))
	return shellItem, hresultToError(ret)
}

func (vtbl *iShellItemVtbl) getDisplayName(objPtr unsafe.Pointer) (string, error) {
	var ptr *uint16
	ret, _, _ := syscall.Syscall(vtbl.GetDisplayName,
		2,
		uintptr(objPtr),
		0x80058000, // SIGDN_FILESYSPATH
		uintptr(unsafe.Pointer(&ptr)))
	if err := hresultToError(ret); err != nil {
		return "", err
	}
	defer ole.CoTaskMemFree(uintptr(unsafe.Pointer(ptr)))
	return ole.LpOleStrToString(ptr), nil
}
