//go:build windows
// +build windows

package cfd

import (
	"fmt"
	"github.com/go-ole/go-ole"
	"strings"
	"syscall"
	"unsafe"
)

func hresultToError(hr uintptr) error {
	if hr < 0 {
		return ole.NewError(hr)
	}
	return nil
}

func (vtbl *iUnknownVtbl) release(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Release,
		0,
		uintptr(objPtr),
		0,
		0)
	return hresultToError(ret)
}

func (vtbl *iModalWindowVtbl) show(objPtr unsafe.Pointer, hwnd uintptr) error {
	ret, _, _ := syscall.Syscall(vtbl.Show,
		1,
		uintptr(objPtr),
		hwnd,
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setFileTypes(objPtr unsafe.Pointer, filters []FileFilter) error {
	cFileTypes := len(filters)
	if cFileTypes < 0 {
		return fmt.Errorf("must specify at least one filter")
	}
	comDlgFilterSpecs := make([]comDlgFilterSpec, cFileTypes)
	for i := 0; i < cFileTypes; i++ {
		filter := &filters[i]
		comDlgFilterSpecs[i] = comDlgFilterSpec{
			pszName: ole.SysAllocString(filter.DisplayName),
			pszSpec: ole.SysAllocString(filter.Pattern),
		}
	}
	ret, _, _ := syscall.Syscall(vtbl.SetFileTypes,
		2,
		uintptr(objPtr),
		uintptr(cFileTypes),
		uintptr(unsafe.Pointer(&comDlgFilterSpecs[0])))
	return hresultToError(ret)
}

// 可选参数如下：
// FOS_OVERWRITEPROMPT = 0x2，      // 提示用户是否覆盖已存在的文件
// FOS_STRICTFILETYPES = 0x4，     // 仅允许在过滤器中指定的文件类型
// FOS_NOCHANGEDIR = 0x8，        // 打开或保存文件时不改变当前目录
// FOS_PICKFOLDERS = 0x20，       // 允许选择文件夹而非文件
// FOS_FORCEFILESYSTEM = 0x40，    // 必须选择文件系统中的项目
// FOS_ALLNONSTORAGEITEMS = 0x80， // 显示非存储设备项（如网络位置）
// FOS_NOVALIDATE = 0x100，       // 不验证所选文件路径的有效性
// FOS_ALLOWMULTISELECT = 0x200， // 允许多选文件或文件夹
// FOS_PATHMUSTEXIST = 0x800，    // 路径必须存在
// FOS_FILEMUSTEXIST = 0x1000，   // 文件必须存在
// FOS_CREATEPROMPT = 0x2000，    // 如果文件不存在则提示创建
// FOS_SHAREAWARE = 0x4000，      // 在打开文件时检查共享冲突
// FOS_NOREADONLYRETURN = 0x8000, // 返回的文件不是只读属性
// FOS_NOTESTFILECREATE = 0x10000,// 不预先测试文件创建操作
// FOS_HIDEMRUPLACES = 0x20000,   // 隐藏最近使用的文件夹列表
// FOS_HIDEPINNEDPLACES = 0x40000,// 隐藏固定到当前位置的文件夹列表
// FOS_NODEREFERENCELINKS = 0x100000,// 不解析符号链接
// FOS_OKBUTTONNEEDSINTERACTION = 0x200000,// 确定按钮需要用户交互后才能启用
// FOS_DONTADDTORECENT = 0x2000000,// 不将选择的文件添加到最近使用的列表
// FOS_FORCESHOWHIDDEN = 0x10000000,// 强制显示隐藏文件和文件夹
// FOS_DEFAULTNOMINIMODE = 0x20000000,// 默认不使用最小化模式
// FOS_FORCEPREVIEWPANEON = 0x40000000,// 强制开启预览窗格
// FOS_SUPPORTSTREAMABLEITEMS = 0x80000000,// 支持流式传输的项目
func (vtbl *iFileDialogVtbl) setOptions(objPtr unsafe.Pointer, options uint32) error {
	ret, _, _ := syscall.Syscall(vtbl.SetOptions,
		1,
		uintptr(objPtr),
		uintptr(options),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) getOptions(objPtr unsafe.Pointer) (uint32, error) {
	var options uint32
	ret, _, _ := syscall.Syscall(vtbl.GetOptions,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&options)),
		0)
	return options, hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) addOption(objPtr unsafe.Pointer, option uint32) error {
	if options, err := vtbl.getOptions(objPtr); err == nil {
		return vtbl.setOptions(objPtr, options|option)
	} else {
		return err
	}
}

func (vtbl *iFileDialogVtbl) removeOption(objPtr unsafe.Pointer, option uint32) error {
	if options, err := vtbl.getOptions(objPtr); err == nil {
		return vtbl.setOptions(objPtr, options&^option)
	} else {
		return err
	}
}

func (vtbl *iFileDialogVtbl) setDefaultFolder(objPtr unsafe.Pointer, path string) error {
	shellItem, err := newIShellItem(path)
	if err != nil {
		return err
	}
	defer shellItem.vtbl.release(unsafe.Pointer(shellItem))
	ret, _, _ := syscall.Syscall(vtbl.SetDefaultFolder,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(shellItem)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setFolder(objPtr unsafe.Pointer, path string) error {
	shellItem, err := newIShellItem(path)
	if err != nil {
		return err
	}
	defer shellItem.vtbl.release(unsafe.Pointer(shellItem))
	ret, _, _ := syscall.Syscall(vtbl.SetFolder,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(shellItem)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setTitle(objPtr unsafe.Pointer, title string) error {
	titlePtr := ole.SysAllocString(title)
	ret, _, _ := syscall.Syscall(vtbl.SetTitle,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(titlePtr)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) close(objPtr unsafe.Pointer) error {
	ret, _, _ := syscall.Syscall(vtbl.Close,
		1,
		uintptr(objPtr),
		0,
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) getResult(objPtr unsafe.Pointer) (*iShellItem, error) {
	var shellItem *iShellItem
	ret, _, _ := syscall.Syscall(vtbl.GetResult,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&shellItem)),
		0)
	return shellItem, hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) getResultString(objPtr unsafe.Pointer) (string, error) {
	shellItem, err := vtbl.getResult(objPtr)
	if err != nil {
		return "", err
	}
	if shellItem == nil {
		return "", ErrCancelled
	}
	defer shellItem.vtbl.release(unsafe.Pointer(shellItem))
	return shellItem.vtbl.getDisplayName(unsafe.Pointer(shellItem))
}

func (vtbl *iFileDialogVtbl) setClientGuid(objPtr unsafe.Pointer, guid *ole.GUID) error {
	ret, _, _ := syscall.Syscall(vtbl.SetClientGuid,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(guid)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setDefaultExtension(objPtr unsafe.Pointer, defaultExtension string) error {
	if defaultExtension[0] == '.' {
		defaultExtension = strings.TrimPrefix(defaultExtension, ".")
	}
	defaultExtensionPtr := ole.SysAllocString(defaultExtension)
	ret, _, _ := syscall.Syscall(vtbl.SetDefaultExtension,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(defaultExtensionPtr)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setFileName(objPtr unsafe.Pointer, fileName string) error {
	fileNamePtr := ole.SysAllocString(fileName)
	ret, _, _ := syscall.Syscall(vtbl.SetFileName,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(fileNamePtr)),
		0)
	return hresultToError(ret)
}

func (vtbl *iFileDialogVtbl) setSelectedFileFilterIndex(objPtr unsafe.Pointer, index uint) error {
	ret, _, _ := syscall.Syscall(vtbl.SetFileTypeIndex,
		1,
		uintptr(objPtr),
		uintptr(index+1), // SetFileTypeIndex 从1开始计数
		0)
	return hresultToError(ret)
}
