// package menu 提供了Wails应用程序中与菜单相关的所有函数和结构体。
// 受到Electron（c）2013-2020 Github Inc.的深度启发。
// Electron 许可证：https://github.com/electron/electron/blob/master/LICENSE
package menu

// Role 是一种用于标识菜单角色的类型
type Role int

// 这些常量需要与`v2/internal/frontend/desktop/darwin/Role.h`中的内容保持同步
const (
	AppMenuRole    Role = 1
	EditMenuRole        = 2
	WindowMenuRole      = 3
// 关于Role              Role = "about" // 撤销Role               Role = "undo" // 重做Role               Role = "redo" // 剪切Role                Role = "cut" // 复制Role               Role = "copy" // 粘贴Role              Role = "paste" // 粘贴并匹配样式Role   Role = "pasteAndMatchStyle" // 全选Role          Role = "selectAll" // 删除Role             Role = "delete" // 最小化Role           Role = "minimize" // 退出Role               Role = "quit" // 切换全屏Role       Role = "togglefullscreen" // 文件菜单Role         Role = "fileMenu" // 查看菜单Role           Role = "viewMenu" // 窗口菜单Role         Role = "windowMenu" // 隐藏Role               Role = "hide" // 隐藏其他Role         Role = "hideOthers" // 显示Role               Role = "unhide" // 置顶Role              Role = "front" // 放大Role               Role = "zoom" // 窗口子菜单Role      Role = "windowSubMenu" // 帮助子菜单Role    Role = "helpSubMenu" // 分隔符ItemRole      Role = "separatorItem" 
// 以上Go语言代码中的注释翻译成中文，主要定义了一系列与GUI界面操作相关的角色（Role）名称，如“关于”、“撤销”、“剪切”、“复制”、“粘贴”等，这些角色通常用于菜单栏、上下文菜单以及快捷键绑定等功能。
)

/*
// About 提供了一个具有“关于”角色的MenuItem

// ff:
func About() *MenuItem {
	return &MenuItem{
		Role: AboutRole,
	}
}

// Undo 提供了一个具有撤销角色的 MenuItem

// ff:
func Undo() *MenuItem {
	return &MenuItem{
		Role: UndoRole,
	}
}

// Redo 提供了一个具有 Redo 角色的 MenuItem

// ff:
func Redo() *MenuItem {
	return &MenuItem{
		Role: RedoRole,
	}
}

// Cut 提供了一个具有“Cut”角色的 MenuItem

// ff:
func Cut() *MenuItem {
	return &MenuItem{
		Role: CutRole,
	}
}

// Copy 提供了一个具有“复制”角色的 MenuItem

// ff:
func Copy() *MenuItem {
	return &MenuItem{
		Role: CopyRole,
	}
}

// Paste 提供了一个具有“粘贴”角色的 MenuItem

// ff:
func Paste() *MenuItem {
	return &MenuItem{
		Role: PasteRole,
	}
}

// PasteAndMatchStyle 提供了一个具有 PasteAndMatchStyle 角色的 MenuItem

// ff:
func PasteAndMatchStyle() *MenuItem {
	return &MenuItem{
		Role: PasteAndMatchStyleRole,
	}
}

// SelectAll 提供了一个具有 SelectAll 角色的 MenuItem

// ff:
func SelectAll() *MenuItem {
	return &MenuItem{
		Role: SelectAllRole,
	}
}

// Delete 为具有“删除”角色的MenuItem提供功能

// ff:
func Delete() *MenuItem {
	return &MenuItem{
		Role: DeleteRole,
	}
}

// Minimize 提供了一个具有最小化角色的MenuItem

// ff:
func Minimize() *MenuItem {
	return &MenuItem{
		Role: MinimizeRole,
	}
}

// Quit 为MenuItem提供了一个具有Quit角色的功能

// ff:
func Quit() *MenuItem {
	return &MenuItem{
		Role: QuitRole,
	}
}

// ToggleFullscreen 提供了一个具有 ToggleFullscreen 角色的 MenuItem

// ff:
func ToggleFullscreen() *MenuItem {
	return &MenuItem{
		Role: TogglefullscreenRole,
	}
}

// FileMenu 提供一个具有默认“文件”菜单（关闭/退出）的 MenuItem

// ff:
func FileMenu() *MenuItem {
	return &MenuItem{
		Role: FileMenuRole,
	}
}
*/

// EditMenu 提供一个具有默认“编辑”菜单（撤销、复制等）的 MenuItem。
func X创建菜单项并带编辑菜单() *MenuItem {
	return &MenuItem{
		X项角色: EditMenuRole,
	}
}

/*
// ViewMenu 提供一个具有默认“查看”菜单（如：重新加载、切换开发者工具等）的 MenuItem

// ff:
func ViewMenu() *MenuItem {
	return &MenuItem{
		Role: ViewMenuRole,
	}
}
*/

// WindowMenu 提供一个带有默认“窗口”菜单（最小化、缩放等）的 MenuItem。
// 在 MacOS 中，如果窗口无边框，则当前其中的所有选项将无法正常工作。
func X创建菜单项并带窗口菜单() *MenuItem {
	return &MenuItem{
		X项角色: WindowMenuRole,
	}
}

// 这些角色仅适用于Mac系统

// AppMenu 提供一个具有默认“应用”菜单（关于、服务等）的 MenuItem
func X创建菜单项并带应用菜单() *MenuItem {
	return &MenuItem{
		X项角色: AppMenuRole,
	}
}

/*
// Hide 提供了一个MenuItem，该MenuItem映射到隐藏动作。

// ff:
func Hide() *MenuItem {
	return &MenuItem{
		Role: HideRole,
	}
}

// HideOthers 提供了一个MenuItem，该MenuItem映射到 hideOtherApplications 动作。

// ff:
func HideOthers() *MenuItem {
	return &MenuItem{
		Role: HideOthersRole,
	}
}

// UnHide 提供了一个MenuItem，该MenuItem映射到unHideAllApplications动作。

// ff:
func UnHide() *MenuItem {
	return &MenuItem{
		Role: UnhideRole,
	}
}

// Front 提供了一个 MenuItem，该 MenuItem 与 arrangeInFront 动作相对应。

// ff:
func Front() *MenuItem {
	return &MenuItem{
		Role: FrontRole,
	}
}

// Zoom 提供了一个MenuItem，该MenuItem映射到执行缩放操作的performZoom动作。

// ff:
func Zoom() *MenuItem {
	return &MenuItem{
		Role: ZoomRole,
	}
}

// WindowSubMenu 提供一个具有“窗口”子菜单的 MenuItem。

// ff:
func WindowSubMenu() *MenuItem {
	return &MenuItem{
		Role: WindowSubMenuRole,
	}
}

// HelpSubMenu 提供了一个带有“帮助”子菜单的 MenuItem。

// ff:
func HelpSubMenu() *MenuItem {
	return &MenuItem{
		Role: HelpSubMenuRole,
	}
}
*/
