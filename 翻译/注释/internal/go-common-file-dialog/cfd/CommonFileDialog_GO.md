
<原文开始>
	// Show the dialog to the user.
	// Blocks until the user has closed the dialog.
<原文结束>

# <翻译开始>
// 向用户展示对话框。
// 将阻塞直到用户关闭对话框。
# <翻译结束>


<原文开始>
// Sets the dialog's parent window. Use 0 to set the dialog to have no parent window.
<原文结束>

# <翻译开始>
// 设置对话框的父窗口。使用0将对话框设置为无父窗口。
# <翻译结束>


<原文开始>
	// Show the dialog to the user.
	// Blocks until the user has closed the dialog and returns their selection.
	// Returns an error if the user cancelled the dialog.
	// Do not use for the Open Multiple Files dialog. Use ShowAndGetResults instead.
<原文结束>

# <翻译开始>
// 向用户展示对话框。
// 将阻塞直到用户关闭对话框并返回他们的选择。
// 如果用户取消对话框，则返回错误。
// 不要用于“打开多个文件”对话框。请改用ShowAndGetResults方法。
# <翻译结束>


<原文开始>
// Sets the title of the dialog window.
<原文结束>

# <翻译开始>
// 设置对话框窗口的标题。
# <翻译结束>


<原文开始>
	// Sets the "role" of the dialog. This is used to derive the dialog's GUID, which the
	// OS will use to differentiate it from dialogs that are intended for other purposes.
	// This means that, for example, a dialog with role "Import" will have a different
	// previous location that it will open to than a dialog with role "Open". Can be any string.
<原文结束>

# <翻译开始>
// 设置对话框的"角色"。这将用于推导对话框的GUID，操作系统会使用这个GUID来区分其与用于其他目的的对话框。
// 这意味着，例如，具有“Import”角色的对话框将打开到一个与其他“Open”角色对话框不同的先前位置。此值可以为任何字符串。
# <翻译结束>


<原文开始>
// Sets the folder used as a default if there is not a recently used folder value available
<原文结束>

# <翻译开始>
// 如果没有可用的最近使用文件夹值，设置默认使用的文件夹
# <翻译结束>


<原文开始>
	// Sets the folder that the dialog always opens to.
	// If this is set, it will override the "default folder" behaviour and the dialog will always open to this folder.
<原文结束>

# <翻译开始>
// 设置对话框始终打开的文件夹。
// 如果设置了这个值，它将覆盖“默认文件夹”行为，对话框将始终打开到此文件夹。
# <翻译结束>


<原文开始>
	// Gets the selected file or folder path, as an absolute path eg. "C:\Folder\file.txt"
	// Do not use for the Open Multiple Files dialog. Use GetResults instead.
<原文结束>

# <翻译开始>
// 获取所选文件或文件夹的路径，以绝对路径形式返回，例如 "C:\Folder\file.txt"
// 不适用于打开多个文件对话框。这种情况下，请使用 GetResults 方法替代。
# <翻译结束>


<原文开始>
	// Sets the file name, I.E. the contents of the file name text box.
	// For Select Folder Dialog, sets folder name.
<原文结束>

# <翻译开始>
// 设置文件名，即文件名文本框的内容。
// 对于选择文件夹对话框，设置文件夹名称。
# <翻译结束>


<原文开始>
	// Release the resources allocated to this Dialog.
	// Should be called when the dialog is finished with.
<原文结束>

# <翻译开始>
// 释放分配给此对话框的资源。
// 当对话框完成其功能时，应调用此方法。
# <翻译结束>


<原文开始>
// Set the list of file filters that the user can select.
<原文结束>

# <翻译开始>
// 设置用户可以选择的文件过滤器列表。
# <翻译结束>


<原文开始>
// Set the selected item from the list of file filters (set using SetFileFilters) by its index. Defaults to 0 (the first item in the list) if not called.
<原文结束>

# <翻译开始>
// 通过索引设置文件过滤器列表（通过调用SetFileFilters设置）中选定的项。如果不调用此方法，默认为0（列表中的第一个项目）。
# <翻译结束>


<原文开始>
	// Sets the default extension applied when a user does not provide one as part of the file name.
	// If the user selects a different file filter, the default extension will be automatically updated to match the new file filter.
	// For Open / Open Multiple File Dialog, this only has an effect when the user specifies a file name with no extension and a file with the default extension exists.
	// For Save File Dialog, this extension will be used whenever a user does not specify an extension.
<原文结束>

# <翻译开始>
// 设置默认扩展名，当用户在文件名中未提供扩展名时使用。
// 如果用户选择了不同的文件过滤器，那么默认扩展名将自动更新以匹配新的文件过滤器。
// 对于“打开”/“打开多个文件”对话框，只有当用户指定了无扩展名的文件名，并且存在相应默认扩展名的文件时，此设置才会生效。
// 对于“保存文件”对话框，每当用户未指定扩展名时，都将使用此扩展名。
# <翻译结束>


<原文开始>
	// Show the dialog to the user.
	// Blocks until the user has closed the dialog and returns the selected files.
<原文结束>

# <翻译开始>
// 向用户展示对话框。
// 将阻塞直到用户关闭对话框，并返回所选文件。
# <翻译结束>


<原文开始>
// Gets the selected file paths, as absolute paths eg. "C:\Folder\file.txt"
<原文结束>

# <翻译开始>
// 获取所选文件路径，以绝对路径形式，例如："C:\Folder\file.txt"
# <翻译结束>

