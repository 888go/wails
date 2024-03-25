package options

import (
	"context"
	"html"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/888go/wails/pkg/options/assetserver"
	"github.com/888go/wails/pkg/options/linux"
	"github.com/888go/wails/pkg/options/mac"
	"github.com/888go/wails/pkg/options/windows"

	"github.com/888go/wails/pkg/menu"

	"github.com/888go/wails/pkg/logger"
)

type WindowStartState int

const (
	Normal     WindowStartState = 0
	Maximised  WindowStartState = 1
	Minimised  WindowStartState = 2
	Fullscreen WindowStartState = 3
)

type Experimental struct{}

// App 包含用于创建 App 的选项
type App struct {
	Title             string
	Width             int
	Height            int
	DisableResize     bool
	Fullscreen        bool
	Frameless         bool
	MinWidth          int
	MinHeight         int
	MaxWidth          int
	MaxHeight         int
	StartHidden       bool
	HideWindowOnClose bool
	AlwaysOnTop       bool
// BackgroundColour 是窗口的背景颜色
// 你可以使用 options.NewRGB 和 options.NewRGBA 函数来创建新的颜色
	BackgroundColour *RGBA
	// 已弃用：请改用 AssetServer.Assets。
	Assets fs.FS
	// 已弃用：请改用 AssetServer.Handler。
	AssetsHandler http.Handler
	// AssetServer 配置应用所需的资源
	AssetServer        *assetserver.Options
	Menu               *menu.Menu
	Logger             logger.Logger `json:"-"`
	LogLevel           logger.LogLevel
	LogLevelProduction logger.LogLevel
	OnStartup          func(ctx context.Context)                `json:"-"`
	OnDomReady         func(ctx context.Context)                `json:"-"`
	OnShutdown         func(ctx context.Context)                `json:"-"`
	OnBeforeClose      func(ctx context.Context) (prevent bool) `json:"-"`
	Bind               []interface{}
	EnumBind           []interface{}
	WindowStartState   WindowStartState

	// ErrorFormatter 重写后端方法返回错误的格式化方式
	ErrorFormatter ErrorFormatter

	// CSS属性，用于检测可拖动元素。默认值为 "--wails-draggable"
	CSSDragProperty string

	// CSSDragProperty必须拥有的CSS值才能被拖动，例如："drag"
	CSSDragValue string

// EnableDefaultContextMenu 在生产环境中启用浏览器的默认右键菜单
// 在开发和调试版本中，此菜单已经默认启用
	EnableDefaultContextMenu bool

// EnableFraudulentWebsiteDetection 启用欺诈网站检测功能，该功能会扫描诸如恶意软件或网络钓鱼企图等欺诈内容。
// 这些服务可能会从您的应用中发送信息，例如访问过的URL以及其他可能的内容到苹果和微软的云端服务。
	EnableFraudulentWebsiteDetection bool

	SingleInstanceLock *SingleInstanceLock

	Windows *windows.Options
	Mac     *mac.Options
	Linux   *linux.Options

	// Experimental options
	Experimental *Experimental

	// 用于调试构建的调试选项。在生产构建中，这些选项将被忽略。
	Debug Debug
}

type ErrorFormatter func(error) any

type RGBA struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

// NewRGBA 通过给定的值创建一个新的 RGBA 结构体
func NewRGBA(r, g, b, a uint8) *RGBA {
	return &RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// NewRGB 通过给定的值创建一个新的 RGBA 结构体，并将 Alpha 设置为 255
func NewRGB(r, g, b uint8) *RGBA {
	return &RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

// MergeDefaults 将为应用程序设置最小的默认值
func MergeDefaults(appoptions *App) {
	// Do set defaults
	if appoptions.Width <= 0 {
		appoptions.Width = 1024
	}
	if appoptions.Height <= 0 {
		appoptions.Height = 768
	}
	if appoptions.Logger == nil {
		appoptions.Logger = logger.NewDefaultLogger()
	}
	if appoptions.LogLevel == 0 {
		appoptions.LogLevel = logger.INFO
	}
	if appoptions.LogLevelProduction == 0 {
		appoptions.LogLevelProduction = logger.ERROR
	}
	if appoptions.CSSDragProperty == "" {
		appoptions.CSSDragProperty = "--wails-draggable"
	}
	if appoptions.CSSDragValue == "" {
		appoptions.CSSDragValue = "drag"
	}
	if appoptions.BackgroundColour == nil {
		appoptions.BackgroundColour = &RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
	}

	// 确保max和min的有效性
	processMinMaxConstraints(appoptions)

	// Default menus
	processMenus(appoptions)

	// Process Drag Options
	processDragOptions(appoptions)
}

type SingleInstanceLock struct {
	// uniqueId 是用于在实例之间设置消息传递的唯一标识符
	UniqueId               string
	OnSecondInstanceLaunch func(secondInstanceData SecondInstanceData)
}

type SecondInstanceData struct {
	Args             []string
	WorkingDirectory string
}

func NewSecondInstanceData() (*SecondInstanceData, error) {
	ex, err := os.Executable()
	if err != nil {
		return nil, err
	}
	workingDirectory := filepath.Dir(ex)

	return &SecondInstanceData{
		Args:             os.Args[1:],
		WorkingDirectory: workingDirectory,
	}, nil
}

func processMenus(appoptions *App) {
	switch runtime.GOOS {
	case "darwin":
		if appoptions.Menu == nil {
			items := []*menu.MenuItem{
				menu.EditMenu(),
			}
			if !appoptions.Frameless {
				items = append(items, menu.WindowMenu()) // 当前“窗口”菜单中的选项仅在非无边框模式下生效
			}

			appoptions.Menu = menu.NewMenuFromItems(menu.AppMenu(), items...)
		}
	}
}

func processMinMaxConstraints(appoptions *App) {
	if appoptions.MinWidth > 0 && appoptions.MaxWidth > 0 {
		if appoptions.MinWidth > appoptions.MaxWidth {
			appoptions.MinWidth = appoptions.MaxWidth
		}
	}
	if appoptions.MinHeight > 0 && appoptions.MaxHeight > 0 {
		if appoptions.MinHeight > appoptions.MaxHeight {
			appoptions.MinHeight = appoptions.MaxHeight
		}
	}
	// 确保当设置了最大值/最小值时，宽度和高度受到限制
	if appoptions.Width < appoptions.MinWidth {
		appoptions.Width = appoptions.MinWidth
	}
	if appoptions.MaxWidth > 0 && appoptions.Width > appoptions.MaxWidth {
		appoptions.Width = appoptions.MaxWidth
	}
	if appoptions.Height < appoptions.MinHeight {
		appoptions.Height = appoptions.MinHeight
	}
	if appoptions.MaxHeight > 0 && appoptions.Height > appoptions.MaxHeight {
		appoptions.Height = appoptions.MaxHeight
	}
}

func processDragOptions(appoptions *App) {
	appoptions.CSSDragProperty = html.EscapeString(appoptions.CSSDragProperty)
	appoptions.CSSDragValue = html.EscapeString(appoptions.CSSDragValue)
}
