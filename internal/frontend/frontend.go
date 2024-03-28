package frontend

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
)

// FileFilter 定义了对话框的文件过滤器
type FileFilter struct {
	DisplayName string // 过滤器信息，例如："图片文件 (*.jpg, *.png)"
	Pattern     string // 以分号分隔的扩展名列表，例如: "*.jpg;*.png"
}

// OpenDialogOptions 包含了OpenDialogOptions运行时方法的选项参数
type OpenDialogOptions struct {
	DefaultDirectory           string
	DefaultFilename            string
	Title                      string
	Filters                    []FileFilter
	ShowHiddenFiles            bool
	CanCreateDirectories       bool
	ResolvesAliases            bool
	TreatPackagesAsDirectories bool
}

// SaveDialogOptions 包含了SaveDialog运行时方法的选项参数
type SaveDialogOptions struct {
	DefaultDirectory           string
	DefaultFilename            string
	Title                      string
	Filters                    []FileFilter
	ShowHiddenFiles            bool
	CanCreateDirectories       bool
	TreatPackagesAsDirectories bool
}

type DialogType string

const (
	InfoDialog     DialogType = "info"
	WarningDialog  DialogType = "warning"
	ErrorDialog    DialogType = "error"
	QuestionDialog DialogType = "question"
)

type Screen struct {
	IsCurrent bool `json:"isCurrent"`
	IsPrimary bool `json:"isPrimary"`

	// 已弃用：请使用 Size 和 PhysicalSize
	Width int `json:"width"`
	// 已弃用：请使用 Size 和 PhysicalSize
	Height int `json:"height"`

	// Size 是屏幕在逻辑像素空间中的尺寸，用于在 Wails 中设置尺寸时使用
	Size ScreenSize `json:"size"`
	// PhysicalSize 是屏幕的物理尺寸，单位为像素
	PhysicalSize ScreenSize `json:"physicalSize"`
}

type ScreenSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// MessageDialogOptions 包含了用于消息对话框（如Info、Warning等运行时方法）的选项。
type MessageDialogOptions struct {
	Type          DialogType
	Title         string
	Message       string
	Buttons       []string
	DefaultButton string
	CancelButton  string
	Icon          []byte
}

type Frontend interface {
	Run(ctx context.Context) error
	RunMainLoop()
	ExecJS(js string)
	Hide()
	Show()
	Quit()

	// Dialog
	OpenFileDialog(dialogOptions OpenDialogOptions) (string, error)
	OpenMultipleFilesDialog(dialogOptions OpenDialogOptions) ([]string, error)
	OpenDirectoryDialog(dialogOptions OpenDialogOptions) (string, error)
	SaveFileDialog(dialogOptions SaveDialogOptions) (string, error)
	MessageDialog(dialogOptions MessageDialogOptions) (string, error)

	// Window
	WindowSetTitle(title string)
	WindowShow()
	WindowHide()
	WindowCenter()
	WindowToggleMaximise()
	WindowMaximise()
	WindowUnmaximise()
	WindowMinimise()
	WindowUnminimise()
	WindowSetAlwaysOnTop(b bool)
	WindowSetPosition(x int, y int)
	WindowGetPosition() (int, int)
	WindowSetSize(width int, height int)
	WindowGetSize() (int, int)
	WindowSetMinSize(width int, height int)
	WindowSetMaxSize(width int, height int)
	WindowFullscreen()
	WindowUnfullscreen()
	WindowSetBackgroundColour(col *options.RGBA)
	WindowReload()
	WindowReloadApp()
	WindowSetSystemDefaultTheme()
	WindowSetLightTheme()
	WindowSetDarkTheme()
	WindowIsMaximised() bool
	WindowIsMinimised() bool
	WindowIsNormal() bool
	WindowIsFullscreen() bool
	WindowClose()
	WindowPrint()

	// Screen
	ScreenGetAll() ([]Screen, error)

	// Menus
	MenuSetApplicationMenu(menu *menu.Menu)
	MenuUpdateApplicationMenu()

	// Events
	Notify(name string, data ...interface{})

	// Browser
	BrowserOpenURL(url string)

	// Clipboard
	ClipboardGetText() (string, error)
	ClipboardSetText(text string) error
}
