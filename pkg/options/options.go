package options

import (
	"context"
	"html"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"

	"github.com/wailsapp/wails/v2/pkg/menu"

	"github.com/wailsapp/wails/v2/pkg/logger"
)

type WindowStartState int

const (
	Normal     WindowStartState = 0 //hs:常量_正常
	Maximised  WindowStartState = 1 //hs:常量_最大化
	Minimised  WindowStartState = 2 //hs:常量_最小化
	Fullscreen WindowStartState = 3 //hs:常量_全屏
)

type Experimental struct{}

// App 包含用于创建 App 的选项
type App struct {
	Title             string //hs:标题
	Width             int    //hs:宽度
	Height            int    //hs:高度
	DisableResize     bool   //hs:禁用调整大小
	Fullscreen        bool   //hs:全屏
	Frameless         bool   //hs:无边框
	MinWidth          int    //hs:最小宽度
	MinHeight         int    //hs:最小高度
	MaxWidth          int    //hs:最大宽度
	MaxHeight         int    //hs:最大高度
	StartHidden       bool   //hs:启动时隐藏窗口
	HideWindowOnClose bool   //hs:关闭时隐藏窗口
	AlwaysOnTop       bool   //hs:始终置顶
	// BackgroundColour 是窗口的背景颜色
	// 你可以使用 options.NewRGB 和 options.NewRGBA 函数来创建新的颜色
	BackgroundColour *RGBA //hs:背景颜色
	// 已弃用：请改用 AssetServer.Assets。
	Assets fs.FS //hs:Assets弃用
	// 已弃用：请改用 AssetServer.Handler。
	AssetsHandler http.Handler //hs:AssetsHandler弃用
	// AssetServer 配置应用所需的资源
	AssetServer        *assetserver.Options                     //hs:绑定http请求
	Menu               *menu.Menu                               //hs:菜单
	Logger             logger.Logger                            `json:"-"` //hs:日志
	LogLevel           logger.LogLevel                          //hs:日志级别
	LogLevelProduction logger.LogLevel                          //hs:生产日志级别
	OnStartup          func(ctx context.Context)                `json:"-"` //hs:绑定启动前函数
	OnDomReady         func(ctx context.Context)                `json:"-"` //hs:绑定DOM就绪函数
	OnShutdown         func(ctx context.Context)                `json:"-"` //hs:绑定应用退出函数
	OnBeforeClose      func(ctx context.Context) (prevent bool) `json:"-"` //hs:绑定应用关闭前函数
	Bind               []interface{}                            //hs:绑定调用方法
	EnumBind           []interface{}                            //hs:绑定常量枚举
	WindowStartState   WindowStartState                         //hs:窗口启动状态

	// ErrorFormatter 重写后端方法返回错误的格式化方式
	ErrorFormatter ErrorFormatter //hs:错误格式化

	// CSS属性，用于检测可拖动元素。默认值为 "--wails-draggable"
	CSSDragProperty string //hs:CSS拖动属性

	// CSSDragProperty必须拥有的CSS值才能被拖动，例如："drag"
	CSSDragValue string //hs:CSS拖动值

	// EnableDefaultContextMenu 在生产环境中启用浏览器的默认右键菜单
	// 在开发和调试版本中，此菜单已经默认启用
	EnableDefaultContextMenu bool //hs:右键菜单

	// EnableFraudulentWebsiteDetection 启用欺诈网站检测功能，该功能会扫描诸如恶意软件或网络钓鱼企图等欺诈内容。
	// 这些服务可能会从您的应用中发送信息，例如访问过的URL以及其他可能的内容到苹果和微软的云端服务。
	EnableFraudulentWebsiteDetection bool //hs:启用欺诈网站检测

	SingleInstanceLock *SingleInstanceLock //hs:单实例锁

	Windows *windows.Options //hs:Windows选项
	Mac     *mac.Options     //hs:Mac选项
	Linux   *linux.Options   //hs:Linux选项

	// Experimental options
	Experimental *Experimental //hs:Experimental实验性

	// 用于调试构建的调试选项。在生产构建中，这些选项将被忽略。
	Debug Debug //hs:调试选项
}

type ErrorFormatter func(error) any

type RGBA struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

// NewRGBA 通过给定的值创建一个新的 RGBA 结构体

// ff:创建RGBA
// a:
// b:
// g:
// r:
func NewRGBA(r, g, b, a uint8) *RGBA {
	return &RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// NewRGB 通过给定的值创建一个新的 RGBA 结构体，并将 Alpha 设置为 255

// ff:创建RGB
// b:
// g:
// r:
func NewRGB(r, g, b uint8) *RGBA {
	return &RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

// MergeDefaults 将为应用程序设置最小的默认值

// ff:
// appoptions:app选项
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

// ff:
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
