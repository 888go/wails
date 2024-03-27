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
	X常量_正常  WindowStartState = 0
	X常量_最大化 WindowStartState = 1
	X常量_最小化 WindowStartState = 2
	X常量_全屏  WindowStartState = 3
)

type Experimental struct{}

// App 包含用于创建 App 的选项
type App struct {
	X标题      string
	X宽度      int
	X高度      int
	X禁用调整大小  bool
	X全屏      bool
	X无边框     bool
	X最小宽度    int
	X最小高度    int
	X最大宽度    int
	X最大高度    int
	X启动时隐藏窗口 bool
	X关闭时隐藏窗口 bool
	X始终置顶    bool
	// BackgroundColour 是窗口的背景颜色
	// 你可以使用 options.NewRGB 和 options.NewRGBA 函数来创建新的颜色
	X背景颜色 *RGBA
	// 已弃用：请改用 AssetServer.Assets。
	Assets弃用 fs.FS
	// 已弃用：请改用 AssetServer.Handler。
	AssetsHandler弃用 http.Handler
	// AssetServer 配置应用所需的资源
	AssetServer        *assetserver.Options
	X菜单                *menu.Menu
	X日志记录器             logger.Logger `json:"-"`
	X日志级别              logger.LogLevel
	LogLevelProduction logger.LogLevel                          //hs:生产日志级别
	X启动前回调函数           func(ctx context.Context)                `json:"-"`
	DOM就绪回调函数          func(ctx context.Context)                `json:"-"`
	X应用退出回调函数          func(ctx context.Context)                `json:"-"`
	X应用关闭前回调函数         func(ctx context.Context) (prevent bool) `json:"-"`
	Bind               []interface{}
	EnumBind           []interface{}
	X窗口启动状态            WindowStartState

	// ErrorFormatter 重写后端方法返回错误的格式化方式
	ErrorFormatter ErrorFormatter

	// CSS属性，用于检测可拖动元素。默认值为 "--wails-draggable"
	CSS拖动属性 string

	// CSSDragProperty必须拥有的CSS值才能被拖动，例如："drag"
	CSS拖动值 string

	// EnableDefaultContextMenu 在生产环境中启用浏览器的默认右键菜单
	// 在开发和调试版本中，此菜单已经默认启用
	X右键菜单 bool

	// EnableFraudulentWebsiteDetection 启用欺诈网站检测功能，该功能会扫描诸如恶意软件或网络钓鱼企图等欺诈内容。
	// 这些服务可能会从您的应用中发送信息，例如访问过的URL以及其他可能的内容到苹果和微软的云端服务。
	X启用欺诈网站检测 bool

	X单实例锁 *SingleInstanceLock

	Windows选项 *windows.Options
	Mac选项     *mac.Options
	Linux选项   *linux.Options

	// Experimental options
	Experimental *Experimental

	// 用于调试构建的调试选项。在生产构建中，这些选项将被忽略。
	X调试选项 Debug
}

type ErrorFormatter func(error) any

type RGBA struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

// NewRGBA 通过给定的值创建一个新的 RGBA 结构体

// a:
// b:
// g:
// r:
func X创建RGBA(r, g, b, a uint8) *RGBA {
	return &RGBA{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

// NewRGB 通过给定的值创建一个新的 RGBA 结构体，并将 Alpha 设置为 255

// b:
// g:
// r:
func X创建RGB(r, g, b uint8) *RGBA {
	return &RGBA{
		R: r,
		G: g,
		B: b,
		A: 255,
	}
}

// MergeDefaults 将为应用程序设置最小的默认值

// ff:
func MergeDefaults(app选项 *App) {
	// Do set defaults
	if app选项.X宽度 <= 0 {
		app选项.X宽度 = 1024
	}
	if app选项.X高度 <= 0 {
		app选项.X高度 = 768
	}
	if app选项.X日志记录器 == nil {
		app选项.X日志记录器 = logger.X创建并按默认()
	}
	if app选项.X日志级别 == 0 {
		app选项.X日志级别 = logger.X常量_日志级别_信息
	}
	if app选项.LogLevelProduction == 0 {
		app选项.LogLevelProduction = logger.X常量_日志级别_错误
	}
	if app选项.CSS拖动属性 == "" {
		app选项.CSS拖动属性 = "--wails-draggable"
	}
	if app选项.CSS拖动值 == "" {
		app选项.CSS拖动值 = "drag"
	}
	if app选项.X背景颜色 == nil {
		app选项.X背景颜色 = &RGBA{
			R: 255,
			G: 255,
			B: 255,
			A: 255,
		}
	}

	// 确保max和min的有效性
	processMinMaxConstraints(app选项)

	// Default menus
	processMenus(app选项)

	// Process Drag Options
	processDragOptions(app选项)
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
		if appoptions.X菜单 == nil {
			items := []*menu.MenuItem{
				menu.X创建菜单项并带编辑菜单(),
			}
			if !appoptions.X无边框 {
				items = append(items, menu.X创建菜单项并带窗口菜单()) // 当前“窗口”菜单中的选项仅在非无边框模式下生效
			}

			appoptions.X菜单 = menu.X创建菜单并按菜单项(menu.X创建菜单项并带应用菜单(), items...)
		}
	}
}

func processMinMaxConstraints(appoptions *App) {
	if appoptions.X最小宽度 > 0 && appoptions.X最大宽度 > 0 {
		if appoptions.X最小宽度 > appoptions.X最大宽度 {
			appoptions.X最小宽度 = appoptions.X最大宽度
		}
	}
	if appoptions.X最小高度 > 0 && appoptions.X最大高度 > 0 {
		if appoptions.X最小高度 > appoptions.X最大高度 {
			appoptions.X最小高度 = appoptions.X最大高度
		}
	}
	// 确保当设置了最大值/最小值时，宽度和高度受到限制
	if appoptions.X宽度 < appoptions.X最小宽度 {
		appoptions.X宽度 = appoptions.X最小宽度
	}
	if appoptions.X最大宽度 > 0 && appoptions.X宽度 > appoptions.X最大宽度 {
		appoptions.X宽度 = appoptions.X最大宽度
	}
	if appoptions.X高度 < appoptions.X最小高度 {
		appoptions.X高度 = appoptions.X最小高度
	}
	if appoptions.X最大高度 > 0 && appoptions.X高度 > appoptions.X最大高度 {
		appoptions.X高度 = appoptions.X最大高度
	}
}

func processDragOptions(appoptions *App) {
	appoptions.CSS拖动属性 = html.EscapeString(appoptions.CSS拖动属性)
	appoptions.CSS拖动值 = html.EscapeString(appoptions.CSS拖动值)
}
