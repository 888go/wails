// Cross-platform.

// Common File Dialogs
package cfd

type Dialog interface {
// 向用户展示对话框。
// 将阻塞直到用户关闭对话框。
	Show() error
	// 设置对话框的父窗口。使用0将对话框设置为无父窗口。
	SetParentWindowHandle(hwnd uintptr)
// 向用户展示对话框。
// 将阻塞直到用户关闭对话框并返回他们的选择。
// 如果用户取消对话框，则返回错误。
// 不要用于“打开多个文件”对话框。请改用ShowAndGetResults方法。
	ShowAndGetResult() (string, error)
	// 设置对话框窗口的标题。
	SetTitle(title string) error
// 设置对话框的"角色"。这将用于推导对话框的GUID，操作系统会使用这个GUID来区分其与用于其他目的的对话框。
// 这意味着，例如，具有“Import”角色的对话框将打开到一个与其他“Open”角色对话框不同的先前位置。此值可以为任何字符串。
	SetRole(role string) error
	// 如果没有可用的最近使用文件夹值，设置默认使用的文件夹
	SetDefaultFolder(defaultFolder string) error
// 设置对话框始终打开的文件夹。
// 如果设置了这个值，它将覆盖“默认文件夹”行为，对话框将始终打开到此文件夹。
	SetFolder(folder string) error
// 获取所选文件或文件夹的路径，以绝对路径形式返回，例如 "C:\Folder\file.txt"
// 不适用于打开多个文件对话框。这种情况下，请使用 GetResults 方法替代。
	GetResult() (string, error)
// 设置文件名，即文件名文本框的内容。
// 对于选择文件夹对话框，设置文件夹名称。
	SetFileName(fileName string) error
// 释放分配给此对话框的资源。
// 当对话框完成其功能时，应调用此方法。
	Release() error
}

type FileDialog interface {
	Dialog
	// 设置用户可以选择的文件过滤器列表。
	SetFileFilters(fileFilter []FileFilter) error
	// 通过索引设置文件过滤器列表（通过调用SetFileFilters设置）中选定的项。如果不调用此方法，默认为0（列表中的第一个项目）。
	SetSelectedFileFilterIndex(index uint) error
// 设置默认扩展名，当用户在文件名中未提供扩展名时使用。
// 如果用户选择了不同的文件过滤器，那么默认扩展名将自动更新以匹配新的文件过滤器。
// 对于“打开”/“打开多个文件”对话框，只有当用户指定了无扩展名的文件名，并且存在相应默认扩展名的文件时，此设置才会生效。
// 对于“保存文件”对话框，每当用户未指定扩展名时，都将使用此扩展名。
	SetDefaultExtension(defaultExtension string) error
}

type OpenFileDialog interface {
	FileDialog
}

type OpenMultipleFilesDialog interface {
	FileDialog
// 向用户展示对话框。
// 将阻塞直到用户关闭对话框，并返回所选文件。
	ShowAndGetResults() ([]string, error)
	// 获取所选文件路径，以绝对路径形式，例如："C:\Folder\file.txt"
	GetResults() ([]string, error)
}

type SelectFolderDialog interface {
	Dialog
}

type SaveFileDialog interface { // TODO Properties
	FileDialog
}
