//go:build windows
// +build windows

package cfd

import (
	"github.com/go-ole/go-ole"
	"github.com/888go/wails/internal/go-common-file-dialog/util"
	"syscall"
	"unsafe"
)

var (
	fileOpenDialogCLSID = ole.NewGUID("{DC1C5A9C-E88A-4dde-A5A1-60F82A20AEF7}")
	fileOpenDialogIID   = ole.NewGUID("{d57c7288-d4ad-4768-be02-9d969532d960}")
)

type iFileOpenDialog struct {
	vtbl               *iFileOpenDialogVtbl
	parentWindowHandle uintptr
}

type iFileOpenDialogVtbl struct {
	iFileDialogVtbl

	GetResults       uintptr // 函数定义：(ppenum **IShellItemArray) HRESULT
// 
// 参数：
// - ppenum：指向IShellItemArray接口指针的指针，用于接收函数返回的Shell项目数组对象
// 
// 返回值：
// - HRESULT：一个COM（Component Object Model）标准的错误代码，表示函数执行成功或失败的具体情况
	GetSelectedItems uintptr
}

func newIFileOpenDialog() (*iFileOpenDialog, error) {
	if unknown, err := ole.CreateInstance(fileOpenDialogCLSID, fileOpenDialogIID); err == nil {
		return (*iFileOpenDialog)(unsafe.Pointer(unknown)), nil
	} else {
		return nil, err
	}
}

func (fileOpenDialog *iFileOpenDialog) Show() error {
	return fileOpenDialog.vtbl.show(unsafe.Pointer(fileOpenDialog), fileOpenDialog.parentWindowHandle)
}

func (fileOpenDialog *iFileOpenDialog) SetParentWindowHandle(hwnd uintptr) {
	fileOpenDialog.parentWindowHandle = hwnd
}

func (fileOpenDialog *iFileOpenDialog) ShowAndGetResult() (string, error) {
	isMultiselect, err := fileOpenDialog.isMultiselect()
	if err != nil {
		return "", err
	}
	if isMultiselect {
		// 我们应该引发panic，因为这个错误是由使用该库的开发者导致的
		panic("use ShowAndGetResults for open multiple files dialog")
	}
	if err := fileOpenDialog.Show(); err != nil {
		return "", err
	}
	return fileOpenDialog.GetResult()
}

func (fileOpenDialog *iFileOpenDialog) ShowAndGetResults() ([]string, error) {
	isMultiselect, err := fileOpenDialog.isMultiselect()
	if err != nil {
		return nil, err
	}
	if !isMultiselect {
		// 我们应该引发panic，因为这个错误是由使用该库的开发者导致的
		panic("use ShowAndGetResult for open single file dialog")
	}
	if err := fileOpenDialog.Show(); err != nil {
		return nil, err
	}
	return fileOpenDialog.GetResults()
}

func (fileOpenDialog *iFileOpenDialog) SetTitle(title string) error {
	return fileOpenDialog.vtbl.setTitle(unsafe.Pointer(fileOpenDialog), title)
}

func (fileOpenDialog *iFileOpenDialog) GetResult() (string, error) {
	isMultiselect, err := fileOpenDialog.isMultiselect()
	if err != nil {
		return "", err
	}
	if isMultiselect {
		// 我们应该引发panic，因为这个错误是由使用该库的开发者导致的
		panic("use GetResults for open multiple files dialog")
	}
	return fileOpenDialog.vtbl.getResultString(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) Release() error {
	return fileOpenDialog.vtbl.release(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) SetDefaultFolder(defaultFolderPath string) error {
	return fileOpenDialog.vtbl.setDefaultFolder(unsafe.Pointer(fileOpenDialog), defaultFolderPath)
}

func (fileOpenDialog *iFileOpenDialog) SetFolder(defaultFolderPath string) error {
	return fileOpenDialog.vtbl.setFolder(unsafe.Pointer(fileOpenDialog), defaultFolderPath)
}

func (fileOpenDialog *iFileOpenDialog) SetFileFilters(filter []FileFilter) error {
	return fileOpenDialog.vtbl.setFileTypes(unsafe.Pointer(fileOpenDialog), filter)
}

func (fileOpenDialog *iFileOpenDialog) SetRole(role string) error {
	return fileOpenDialog.vtbl.setClientGuid(unsafe.Pointer(fileOpenDialog), util.StringToUUID(role))
}

// 这段代码应当仅在用户请求多选时才可调用，因为
// 否则会提供Dialog接口，该接口并未公开这个函数。
func (fileOpenDialog *iFileOpenDialog) GetResults() ([]string, error) {
	isMultiselect, err := fileOpenDialog.isMultiselect()
	if err != nil {
		return nil, err
	}
	if !isMultiselect {
		// 我们应该引发panic，因为这个错误是由使用该库的开发者导致的
		panic("use GetResult for open single file dialog")
	}
	return fileOpenDialog.vtbl.getResultsStrings(unsafe.Pointer(fileOpenDialog))
}

func (fileOpenDialog *iFileOpenDialog) SetDefaultExtension(defaultExtension string) error {
	return fileOpenDialog.vtbl.setDefaultExtension(unsafe.Pointer(fileOpenDialog), defaultExtension)
}

func (fileOpenDialog *iFileOpenDialog) SetFileName(initialFileName string) error {
	return fileOpenDialog.vtbl.setFileName(unsafe.Pointer(fileOpenDialog), initialFileName)
}

func (fileOpenDialog *iFileOpenDialog) SetSelectedFileFilterIndex(index uint) error {
	return fileOpenDialog.vtbl.setSelectedFileFilterIndex(unsafe.Pointer(fileOpenDialog), index)
}

func (fileOpenDialog *iFileOpenDialog) setPickFolders(pickFolders bool) error {
	const FosPickfolders = 0x20
	if pickFolders {
		return fileOpenDialog.vtbl.addOption(unsafe.Pointer(fileOpenDialog), FosPickfolders)
	} else {
		return fileOpenDialog.vtbl.removeOption(unsafe.Pointer(fileOpenDialog), FosPickfolders)
	}
}

const FosAllowMultiselect = 0x200

func (fileOpenDialog *iFileOpenDialog) isMultiselect() (bool, error) {
	options, err := fileOpenDialog.vtbl.getOptions(unsafe.Pointer(fileOpenDialog))
	if err != nil {
		return false, err
	}
	return options&FosAllowMultiselect != 0, nil
}

func (fileOpenDialog *iFileOpenDialog) setIsMultiselect(isMultiselect bool) error {
	if isMultiselect {
		return fileOpenDialog.vtbl.addOption(unsafe.Pointer(fileOpenDialog), FosAllowMultiselect)
	} else {
		return fileOpenDialog.vtbl.removeOption(unsafe.Pointer(fileOpenDialog), FosAllowMultiselect)
	}
}

func (vtbl *iFileOpenDialogVtbl) getResults(objPtr unsafe.Pointer) (*iShellItemArray, error) {
	var shellItemArray *iShellItemArray
	ret, _, _ := syscall.Syscall(vtbl.GetResults,
		1,
		uintptr(objPtr),
		uintptr(unsafe.Pointer(&shellItemArray)),
		0)
	return shellItemArray, hresultToError(ret)
}

func (vtbl *iFileOpenDialogVtbl) getResultsStrings(objPtr unsafe.Pointer) ([]string, error) {
	shellItemArray, err := vtbl.getResults(objPtr)
	if err != nil {
		return nil, err
	}
	if shellItemArray == nil {
		return nil, ErrCancelled
	}
	defer shellItemArray.vtbl.release(unsafe.Pointer(shellItemArray))
	count, err := shellItemArray.vtbl.getCount(unsafe.Pointer(shellItemArray))
	if err != nil {
		return nil, err
	}
	var results []string
	for i := uintptr(0); i < count; i++ {
		newItem, err := shellItemArray.vtbl.getItemAt(unsafe.Pointer(shellItemArray), i)
		if err != nil {
			return nil, err
		}
		results = append(results, newItem)
	}
	return results, nil
}
