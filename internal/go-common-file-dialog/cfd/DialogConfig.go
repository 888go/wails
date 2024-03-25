// Cross-platform.

package cfd

type FileFilter struct {
	// 这个过滤器的显示名称（展示给用户看的）
	DisplayName string
	// 过滤器模式。例如 "*.txt;*.png" 用于选择所有 txt 和 png 文件，"*.*" 用于选择任何文件等。
	Pattern string
}

type DialogConfig struct {
	// 对话框的标题
	Title string
// 对话框的角色。这个属性用于推导对话框的GUID，操作系统会使用这个GUID来区分其与用于其他目的的对话框。
// 这意味着，例如，具有“Import”角色的对话框在打开时将有不同的初始位置，与具有“Open”角色的对话框不同。此值可以是任何字符串。
	Role string
// 默认文件夹 - 当用户首次打开时使用的文件夹（在首次打开之后，将使用他们上次使用的文件位置）。
	DefaultFolder string
// 初始文件夹 - 如果非空，则对话框始终会打开到此文件夹。
// 如果此值不为空，它将覆盖“默认文件夹”行为，
// 并且对话框将始终打开到这个文件夹。
	Folder string
// 该文件过滤器用于限制对话框可以选择的文件类型。
// 被“选择文件夹对话框”忽略。
	FileFilters []FileFilter
// 设置初始选择的文件过滤器。这是 FileFilters 的索引。
// 被“选择文件夹对话框”忽略。
	SelectedFileFilterIndex uint
// 当用户打开对话框时，文件的初始名称（即文件名文本框中的文本）。
// 对于选择文件夹对话框，这将设置初始文件夹名称。
	FileName string
// 当用户在文件名中未提供扩展名时，应用的默认扩展名。
// 如果用户选择了不同的文件过滤器，该默认扩展名将自动更新以匹配新的文件过滤器。
// 对于“打开”/“打开多个文件”对话框，只有当用户指定了无扩展名的文件名，并且存在具有默认扩展名的文件时，此设置才有效。
// 对于“保存文件”对话框，每当用户未指定扩展名时，都将使用此扩展名。
// “选择文件夹”对话框忽略此设置。
	DefaultExtension string
// ParentWindowHandle 是对话框的父窗口句柄（HWND）。
// 如果设置为 0 / nil，则对话框将没有父窗口。
	ParentWindowHandle uintptr
}

var defaultFilters = []FileFilter{
	{
		DisplayName: "All Files (*.*)",
		Pattern:     "*.*",
	},
}

func (config *DialogConfig) apply(dialog Dialog) (err error) {
	if config.Title != "" {
		err = dialog.SetTitle(config.Title)
		if err != nil {
			return
		}
	}

	if config.Role != "" {
		err = dialog.SetRole(config.Role)
		if err != nil {
			return
		}
	}

	if config.Folder != "" {
		err = dialog.SetFolder(config.Folder)
		if err != nil {
			return
		}
	}

	if config.DefaultFolder != "" {
		err = dialog.SetDefaultFolder(config.DefaultFolder)
		if err != nil {
			return
		}
	}

	if config.FileName != "" {
		err = dialog.SetFileName(config.FileName)
		if err != nil {
			return
		}
	}

	dialog.SetParentWindowHandle(config.ParentWindowHandle)

	if dialog, ok := dialog.(FileDialog); ok {
		var fileFilters []FileFilter
		if config.FileFilters != nil && len(config.FileFilters) > 0 {
			fileFilters = config.FileFilters
		} else {
			fileFilters = defaultFilters
		}
		err = dialog.SetFileFilters(fileFilters)
		if err != nil {
			return
		}

		if config.SelectedFileFilterIndex != 0 {
			err = dialog.SetSelectedFileFilterIndex(config.SelectedFileFilterIndex)
			if err != nil {
				return
			}
		}

		if config.DefaultExtension != "" {
			err = dialog.SetDefaultExtension(config.DefaultExtension)
			if err != nil {
				return
			}
		}
	}

	return
}
