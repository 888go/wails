//go:build windows
// +build windows

package cfd

import (
	"github.com/go-ole/go-ole"
	"syscall"
	"unsafe"
)

const (
	iidShellItemArrayGUID = "{b63ea76d-1f85-456f-a19c-48159efa858b}"
)

var (
	iidShellItemArray *ole.GUID
)

func init() {
	iidShellItemArray, _ = ole.IIDFromString(iidShellItemArrayGUID)
}

type iShellItemArray struct {
	vtbl *iShellItemArrayVtbl
}

type iShellItemArrayVtbl struct {
	iUnknownVtbl
	BindToHandler              uintptr
	GetPropertyStore           uintptr
	GetPropertyDescriptionList uintptr
	GetAttributes              uintptr
	GetCount                   uintptr // 函数声明：(pdwNumItems *DWORD) HRESULT
// 
// 参数：
// - pdwNumItems *DWORD: 指向一个 DWORD 类型的指针，用于接收返回的项数量
// 
// 返回值：
// - HRESULT：在 Windows 系统编程中，HRESULT 是一种表示函数调用结果的状态代码。通常用来表示函数执行成功与否及错误信息。
// 整体翻译为：
// 函数（输出参数为指向 DWORD 类型的指针，用于接收项的数量）返回 HRESULT 类型值
// 
// 参数：
// - pdwNumItems：指向 DWORD 类型的指针，用于存储获取到的项的数量
// 
// 返回值：
// - HRESULT：表示函数执行结果的状态码，在Windows系统编程中使用，用于指示函数调用是否成功及其相关错误信息
	GetItemAt                  uintptr // 函数定义：给定一个DWORD类型的dwIndex参数和一个指向IShellItem接口指针的指针(ppsi)，该函数返回一个HRESULT值。
// 
// 参数：
// - dwIndex: 类型为DWORD的索引值
// - ppsi: 指向IShellItem接口指针的指针，用于接收函数返回的Shell Item对象
// 
// 返回值：
// - HRESULT：表示函数调用结果的成功或失败，通常是一个32位的错误代码。
	EnumItems                  uintptr
}

func (vtbl *iShellItemArrayVtbl) getCount(objPtr unsafe.Pointer) (uintptr, error) {
	var count uintptr
	ret, _, _ := syscall.Syscall(vtbl.GetCount,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&count)),
		0)
	if err := hresultToError(ret); err != nil {
		return 0, err
	}
	return count, nil
}

func (vtbl *iShellItemArrayVtbl) getItemAt(objPtr unsafe.Pointer, index uintptr) (string, error) {
	var shellItem *iShellItem
	ret, _, _ := syscall.Syscall(vtbl.GetItemAt,
		2,
		uintptr(objPtr),
		index,
		uintptr(unsafe.Pointer(&shellItem)))
	if err := hresultToError(ret); err != nil {
		return "", err
	}
	if shellItem == nil {
		return "", ErrCancelled
	}
	defer shellItem.vtbl.release(unsafe.Pointer(shellItem))
	return shellItem.vtbl.getDisplayName(unsafe.Pointer(shellItem))
}
