
<原文开始>
// The display name of the filter (That is shown to the user)
<原文结束>

# <翻译开始>
// 这个过滤器的显示名称（展示给用户看的）
# <翻译结束>


<原文开始>
// The filter pattern. Eg. "*.txt;*.png" to select all txt and png files, "*.*" to select any files, etc.
<原文结束>

# <翻译开始>
// 过滤器模式。例如 "*.txt;*.png" 用于选择所有 txt 和 png 文件，"*.*" 用于选择任何文件等。
# <翻译结束>


<原文开始>
// The title of the dialog
<原文结束>

# <翻译开始>
// 对话框的标题
# <翻译结束>


<原文开始>
	// The role of the dialog. This is used to derive the dialog's GUID, which the
	// OS will use to differentiate it from dialogs that are intended for other purposes.
	// This means that, for example, a dialog with role "Import" will have a different
	// previous location that it will open to than a dialog with role "Open". Can be any string.
<原文结束>

# <翻译开始>
// 对话框的角色。这个属性用于推导对话框的GUID，操作系统会使用这个GUID来区分其与用于其他目的的对话框。
// 这意味着，例如，具有“Import”角色的对话框在打开时将有不同的初始位置，与具有“Open”角色的对话框不同。此值可以是任何字符串。
# <翻译结束>


<原文开始>
	// The default folder - the folder that is used the first time the user opens it
	// (after the first time their last used location is used).
<原文结束>

# <翻译开始>
// 默认文件夹 - 当用户首次打开时使用的文件夹（在首次打开之后，将使用他们上次使用的文件位置）。
# <翻译结束>


<原文开始>
	// The initial folder - the folder that the dialog always opens to if not empty.
	// If this is not empty, it will override the "default folder" behaviour and
	// the dialog will always open to this folder.
<原文结束>

# <翻译开始>
// 初始文件夹 - 如果非空，则对话框始终会打开到此文件夹。
// 如果此值不为空，它将覆盖“默认文件夹”行为，
// 并且对话框将始终打开到这个文件夹。
# <翻译结束>


<原文开始>
	// The file filters that restrict which types of files the dialog is able to choose.
	// Ignored by Select Folder Dialog.
<原文结束>

# <翻译开始>
// 该文件过滤器用于限制对话框可以选择的文件类型。
// 被“选择文件夹对话框”忽略。
# <翻译结束>


<原文开始>
	// Sets the initially selected file filter. This is an index of FileFilters.
	// Ignored by Select Folder Dialog.
<原文结束>

# <翻译开始>
// 设置初始选择的文件过滤器。这是 FileFilters 的索引。
// 被“选择文件夹对话框”忽略。
# <翻译结束>


<原文开始>
	// The initial name of the file (I.E. the text in the file name text box) when the user opens the dialog.
	// For the Select Folder Dialog, this sets the initial folder name.
<原文结束>

# <翻译开始>
// 当用户打开对话框时，文件的初始名称（即文件名文本框中的文本）。
// 对于选择文件夹对话框，这将设置初始文件夹名称。
# <翻译结束>


<原文开始>
	// The default extension applied when a user does not provide one as part of the file name.
	// If the user selects a different file filter, the default extension will be automatically updated to match the new file filter.
	// For Open / Open Multiple File Dialog, this only has an effect when the user specifies a file name with no extension and a file with the default extension exists.
	// For Save File Dialog, this extension will be used whenever a user does not specify an extension.
	// Ignored by Select Folder Dialog.
<原文结束>

# <翻译开始>
// 当用户在文件名中未提供扩展名时，应用的默认扩展名。
// 如果用户选择了不同的文件过滤器，该默认扩展名将自动更新以匹配新的文件过滤器。
// 对于“打开”/“打开多个文件”对话框，只有当用户指定了无扩展名的文件名，并且存在具有默认扩展名的文件时，此设置才有效。
// 对于“保存文件”对话框，每当用户未指定扩展名时，都将使用此扩展名。
// “选择文件夹”对话框忽略此设置。
# <翻译结束>


<原文开始>
	// ParentWindowHandle is the handle (HWND) to the parent window of the dialog.
	// If left as 0 / nil, the dialog will have no parent window.
<原文结束>

# <翻译开始>
// ParentWindowHandle 是对话框的父窗口句柄（HWND）。
// 如果设置为 0 / nil，则对话框将没有父窗口。
# <翻译结束>

