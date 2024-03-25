package frontend

import (
	"context"

	"github.com/888go/wails/pkg/menu"
	"github.com/888go/wails/pkg/options"
)

// FileFilter 定义了对话框的文件过滤器
type FileFilter struct {
	X显示名称  string // 过滤器信息，例如："图片文件 (*.jpg, *.png)"
	X扩展名列表 string // 以分号分隔的扩展名列表，例如: "*.jpg;*.png"
}

// OpenDialogOptions 包含了OpenDialogOptions运行时方法的选项参数
type OpenDialogOptions struct {
	X默认目录     string
	X默认文件名    string
	X标题       string
	X过滤器      []FileFilter
	X显示隐藏文件   bool
	X是否可创建目录  bool
	X是否解析别名   bool
	X是否将包视为目录 bool
}

// SaveDialogOptions 包含了SaveDialog运行时方法的选项参数
type SaveDialogOptions struct {
	X默认目录     string
	X默认文件名    string
	X标题       string
	X过滤器      []FileFilter
	X显示隐藏文件   bool
	X是否可创建目录  bool
	X是否将包视为目录 bool
}

type DialogType string

const (
	X常量_对话框_信息 DialogType = "info"
	X常量_对话框_警告 DialogType = "warning"
	X常量_对话框_错误 DialogType = "error"
	X常量_对话框_问题 DialogType = "question"
)

type Screen struct {
	IsCurrent bool `json:"isCurrent"`
	IsPrimary bool `json:"isPrimary"`

	// 已弃用：请使用 Size 和 PhysicalSize
	Width int `json:"width"`
	// 已弃用：请使用 Size 和 PhysicalSize
	Height int `json:"height"`

	// Size 是屏幕在逻辑像素空间中的尺寸，用于在 Wails 中设置尺寸时使用
	X大小 ScreenSize `json:"size"`
	// PhysicalSize 是屏幕的物理尺寸，单位为像素
	PhysicalSize ScreenSize `json:"physicalSize"`
}

type ScreenSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

// MessageDialogOptions 包含了用于消息对话框（如Info、Warning等运行时方法）的选项。
type MessageDialogOptions struct {
	X对话框类型 DialogType
	X标题    string
	X消息    string
	X按钮s   []string
	X默认按钮  string
	X取消按钮  string
	X图标    []byte
}

type Frontend interface {
	X运行(ctx context.Context) error
	RunMainLoop()
	X窗口执行JS(js string)
	X隐藏()
	X显示()
	X退出()

	// Dialog
	X对话框选择文件(dialogOptions OpenDialogOptions) (string, error)
	X对话框多选文件(dialogOptions OpenDialogOptions) ([]string, error)
	X对话框选择目录(dialogOptions OpenDialogOptions) (string, error)
	X对话框保存文件(dialogOptions SaveDialogOptions) (string, error)
	X对话框弹出消息(dialogOptions MessageDialogOptions) (string, error)

	// Window
	X窗口设置标题(title string)
	X窗口显示()
	X窗口隐藏()
	X窗口居中()
	X窗口最大化切换()
	X窗口最大化()
	X窗口取消最大化()
	X窗口最小化()
	X窗口取消最小化()
	X窗口设置置顶(b bool)
	X窗口设置位置(x int, y int)
	X窗口取位置() (int, int)
	X窗口设置尺寸(width int, height int)
	X窗口取尺寸() (int, int)
	X窗口设置最小尺寸(width int, height int)
	X窗口设置最大尺寸(width int, height int)
	X窗口设置全屏()
	X窗口取消全屏()
	X窗口设置背景色(col *options.RGBA)
	X窗口重载()
	X窗口重载应用程序前端()
	X窗口设置系统默认主题()
	X窗口设置浅色主题()
	X窗口设置深色主题()
	X窗口是否最大化() bool
	X窗口是否最小化() bool
	X窗口是否为正常() bool
	X窗口是否全屏() bool
	WindowClose()
	X窗口打开打印对话框()

	// Screen
	X取屏幕信息() ([]Screen, error)

	// Menus
	X菜单设置(menu *menu.Menu)
	X菜单更新()

	// Events
	Notify(name string, data ...interface{})

	// Browser
	X默认浏览器打开url(url string)

	// Clipboard
	X剪贴板取文本() (string, error)
	X剪贴板置文本(text string) error
}
