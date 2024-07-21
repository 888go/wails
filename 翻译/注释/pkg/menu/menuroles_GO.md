
<原文开始>
// Package menu provides all the functions and structs related to menus in a Wails application.
// Heavily inspired by Electron (c) 2013-2020 Github Inc.
// Electron License: https://github.com/electron/electron/blob/master/LICENSE
<原文结束>

# <翻译开始>
// package menu 提供了Wails应用程序中与菜单相关的所有函数和结构体。
// 受到Electron（c）2013-2020 Github Inc.的深度启发。
// Electron 许可证：https://github.com/electron/electron/blob/master/LICENSE
# <翻译结束>


<原文开始>
// Role is a type to identify menu roles
<原文结束>

# <翻译开始>
// Role 是一种用于标识菜单角色的类型
# <翻译结束>


<原文开始>
// These constants need to be kept in sync with `v2/internal/frontend/desktop/darwin/Role.h`
<原文结束>

# <翻译开始>
// 这些常量需要与`v2/internal/frontend/desktop/darwin/Role.h`中的内容保持同步
# <翻译结束>


<原文开始>
	// AboutRole              Role = "about"
	// UndoRole               Role = "undo"
	// RedoRole               Role = "redo"
	// CutRole                Role = "cut"
	// CopyRole               Role = "copy"
	// PasteRole              Role = "paste"
	// PasteAndMatchStyleRole Role = "pasteAndMatchStyle"
	// SelectAllRole          Role = "selectAll"
	// DeleteRole             Role = "delete"
	// MinimizeRole           Role = "minimize"
	// QuitRole               Role = "quit"
	// TogglefullscreenRole   Role = "togglefullscreen"
	// FileMenuRole           Role = "fileMenu"
	// ViewMenuRole           Role = "viewMenu"
	// WindowMenuRole         Role = "windowMenu"
	// HideRole               Role = "hide"
	// HideOthersRole         Role = "hideOthers"
	// UnhideRole             Role = "unhide"
	// FrontRole              Role = "front"
	// ZoomRole               Role = "zoom"
	// WindowSubMenuRole      Role = "windowSubMenu"
	// HelpSubMenuRole        Role = "helpSubMenu"
	// SeparatorItemRole      Role = "separatorItem"
<原文结束>

# <翻译开始>
	// 关于Role              Role = "about" 	// 撤销Role               Role = "undo" 	// 重做Role               Role = "redo" 	// 剪切Role                Role = "cut" 	// 复制Role               Role = "copy" 	// 粘贴Role              Role = "paste" 	// 粘贴并匹配样式Role   Role = "pasteAndMatchStyle" 	// 全选Role          Role = "selectAll" 	// 删除Role             Role = "delete" 	// 最小化Role           Role = "minimize" 	// 退出Role               Role = "quit" 	// 切换全屏Role       Role = "togglefullscreen" 	// 文件菜单Role         Role = "fileMenu" 	// 查看菜单Role           Role = "viewMenu" 	// 窗口菜单Role         Role = "windowMenu" 	// 隐藏Role               Role = "hide" 	// 隐藏其他Role         Role = "hideOthers" 	// 显示Role               Role = "unhide" 	// 置顶Role              Role = "front" 	// 放大Role               Role = "zoom" 	// 窗口子菜单Role      Role = "windowSubMenu" 	// 帮助子菜单Role    Role = "helpSubMenu" 	// 分隔符ItemRole      Role = "separatorItem" 
	// 以上Go语言代码中的注释翻译成中文，主要定义了一系列与GUI界面操作相关的角色（Role）名称，如“关于”、“撤销”、“剪切”、“复制”、“粘贴”等，这些角色通常用于菜单栏、上下文菜单以及快捷键绑定等功能。
# <翻译结束>


<原文开始>
// About provides a MenuItem with the About role
<原文结束>

# <翻译开始>
// About 提供了一个具有“关于”角色的MenuItem
# <翻译结束>


<原文开始>
// Undo provides a MenuItem with the Undo role
<原文结束>

# <翻译开始>
// Undo 提供了一个具有撤销角色的 MenuItem
# <翻译结束>


<原文开始>
// Redo provides a MenuItem with the Redo role
<原文结束>

# <翻译开始>
// Redo 提供了一个具有 Redo 角色的 MenuItem
# <翻译结束>


<原文开始>
// Cut provides a MenuItem with the Cut role
<原文结束>

# <翻译开始>
// Cut 提供了一个具有“Cut”角色的 MenuItem
# <翻译结束>


<原文开始>
// Copy provides a MenuItem with the Copy role
<原文结束>

# <翻译开始>
// Copy 提供了一个具有“复制”角色的 MenuItem
# <翻译结束>


<原文开始>
// Paste provides a MenuItem with the Paste role
<原文结束>

# <翻译开始>
// Paste 提供了一个具有“粘贴”角色的 MenuItem
# <翻译结束>


<原文开始>
// PasteAndMatchStyle provides a MenuItem with the PasteAndMatchStyle role
<原文结束>

# <翻译开始>
// PasteAndMatchStyle 提供了一个具有 PasteAndMatchStyle 角色的 MenuItem
# <翻译结束>


<原文开始>
// SelectAll provides a MenuItem with the SelectAll role
<原文结束>

# <翻译开始>
// SelectAll 提供了一个具有 SelectAll 角色的 MenuItem
# <翻译结束>


<原文开始>
// Delete provides a MenuItem with the Delete role
<原文结束>

# <翻译开始>
// Delete 为具有“删除”角色的MenuItem提供功能
# <翻译结束>


<原文开始>
// Minimize provides a MenuItem with the Minimize role
<原文结束>

# <翻译开始>
// Minimize 提供了一个具有最小化角色的MenuItem
# <翻译结束>


<原文开始>
// Quit provides a MenuItem with the Quit role
<原文结束>

# <翻译开始>
// Quit 为MenuItem提供了一个具有Quit角色的功能
# <翻译结束>


<原文开始>
// ToggleFullscreen provides a MenuItem with the ToggleFullscreen role
<原文结束>

# <翻译开始>
// ToggleFullscreen 提供了一个具有 ToggleFullscreen 角色的 MenuItem
# <翻译结束>


<原文开始>
// FileMenu provides a MenuItem with the whole default "File" menu (Close / Quit)
<原文结束>

# <翻译开始>
// FileMenu 提供一个具有默认“文件”菜单（关闭/退出）的 MenuItem
# <翻译结束>


<原文开始>
// EditMenu provides a MenuItem with the whole default "Edit" menu (Undo, Copy, etc.).
<原文结束>

# <翻译开始>
// EditMenu 提供一个具有默认“编辑”菜单（撤销、复制等）的 MenuItem。
# <翻译结束>


<原文开始>
// ViewMenu provides a MenuItem with the whole default "View" menu (Reload, Toggle Developer Tools, etc.)
<原文结束>

# <翻译开始>
// ViewMenu 提供一个具有默认“查看”菜单（如：重新加载、切换开发者工具等）的 MenuItem
# <翻译结束>


<原文开始>
// WindowMenu provides a MenuItem with the whole default "Window" menu (Minimize, Zoom, etc.).
// On MacOS currently all options in there won't work if the window is frameless.
<原文结束>

# <翻译开始>
// WindowMenu 提供一个带有默认“窗口”菜单（最小化、缩放等）的 MenuItem。
// 在 MacOS 中，如果窗口无边框，则当前其中的所有选项将无法正常工作。
# <翻译结束>


<原文开始>
// These roles are Mac only
<原文结束>

# <翻译开始>
// 这些角色仅适用于Mac系统
# <翻译结束>


<原文开始>
// AppMenu provides a MenuItem with the whole default "App" menu (About, Services, etc.)
<原文结束>

# <翻译开始>
// AppMenu 提供一个具有默认“应用”菜单（关于、服务等）的 MenuItem
# <翻译结束>


<原文开始>
// Hide provides a MenuItem that maps to the hide action.
<原文结束>

# <翻译开始>
// Hide 提供了一个MenuItem，该MenuItem映射到隐藏动作。
# <翻译结束>


<原文开始>
// HideOthers provides a MenuItem that maps to the hideOtherApplications action.
<原文结束>

# <翻译开始>
// HideOthers 提供了一个MenuItem，该MenuItem映射到 hideOtherApplications 动作。
# <翻译结束>


<原文开始>
// UnHide provides a MenuItem that maps to the unHideAllApplications action.
<原文结束>

# <翻译开始>
// UnHide 提供了一个MenuItem，该MenuItem映射到unHideAllApplications动作。
# <翻译结束>


<原文开始>
// Front provides a MenuItem that maps to the arrangeInFront action.
<原文结束>

# <翻译开始>
// Front 提供了一个 MenuItem，该 MenuItem 与 arrangeInFront 动作相对应。
# <翻译结束>


<原文开始>
// Zoom provides a MenuItem that maps to the performZoom action.
<原文结束>

# <翻译开始>
// Zoom 提供了一个MenuItem，该MenuItem映射到执行缩放操作的performZoom动作。
# <翻译结束>


<原文开始>
// WindowSubMenu provides a MenuItem with the "Window" submenu.
<原文结束>

# <翻译开始>
// WindowSubMenu 提供一个具有“窗口”子菜单的 MenuItem。
# <翻译结束>


<原文开始>
// HelpSubMenu provides a MenuItem with the "Help" submenu.
<原文结束>

# <翻译开始>
// HelpSubMenu 提供了一个带有“帮助”子菜单的 MenuItem。
# <翻译结束>

