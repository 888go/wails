//go:build darwin
// +build darwin

package darwin

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation -framework Cocoa -framework WebKit
#import <Foundation/Foundation.h>
#import "Application.h"
#import "CustomProtocol.h"
#import "WailsContext.h"

#include <stdlib.h>
*/
import "C"

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/url"
	"unsafe"

	"github.com/wailsapp/wails/v2/pkg/assetserver"
	"github.com/wailsapp/wails/v2/pkg/assetserver/webview"

	"github.com/wailsapp/wails/v2/internal/binding"
	"github.com/wailsapp/wails/v2/internal/frontend"
	"github.com/wailsapp/wails/v2/internal/frontend/runtime"
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)

const startURL = "wails://wails/"

var (
	messageBuffer        = make(chan string, 100)
	requestBuffer        = make(chan webview.Request, 100)
	callbackBuffer       = make(chan uint, 10)
	openFilepathBuffer   = make(chan string, 100)
	openUrlBuffer        = make(chan string, 100)
	secondInstanceBuffer = make(chan options.SecondInstanceData, 1)
)

type Frontend struct {
	// Context
	ctx context.Context

	frontendOptions *options.App
	logger          *logger.Logger
	debug           bool
	devtoolsEnabled bool

	// Assets
	assets   *assetserver.AssetServer
	startURL *url.URL

	// main window handle
	mainWindow *Window
	bindings   *binding.Bindings
	dispatcher frontend.Dispatcher
}


// ff:
func (f *Frontend) RunMainLoop() {
	C.RunMainLoop()
}


// ff:
func (f *Frontend) WindowClose() {
	C.ReleaseContext(f.mainWindow.context)
}


// ff:
// dispatcher:
// appBindings:
// myLogger:
// appoptions:
// ctx:
func NewFrontend(ctx context.Context, appoptions *options.App, myLogger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) *Frontend {
	result := &Frontend{
		frontendOptions: appoptions,
		logger:          myLogger,
		bindings:        appBindings,
		dispatcher:      dispatcher,
		ctx:             ctx,
	}
	result.startURL, _ = url.Parse(startURL)

	// 这个应当尽早初始化，以便处理首次实例启动
	C.StartCustomProtocolHandler()

	if _starturl, _ := ctx.Value("starturl").(*url.URL); _starturl != nil {
		result.startURL = _starturl
	} else {
		if port, _ := ctx.Value("assetserverport").(string); port != "" {
			result.startURL.Host = net.JoinHostPort(result.startURL.Host+".localhost", port)
		}

		var bindings string
		var err error
		if _obfuscated, _ := ctx.Value("obfuscated").(bool); !_obfuscated {
			bindings, err = appBindings.ToJSON()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			appBindings.DB().UpdateObfuscatedCallMap()
		}

		assets, err := assetserver.NewAssetServerMainPage(bindings, appoptions, ctx.Value("assetdir") != nil, myLogger, runtime.RuntimeAssetsBundle)
		if err != nil {
			log.Fatal(err)
		}
		assets.ExpectedWebViewHost = result.startURL.Host
		result.assets = assets

		go result.startRequestProcessor()
	}

	go result.startMessageProcessor()
	go result.startCallbackProcessor()
	go result.startFileOpenProcessor()
	go result.startUrlOpenProcessor()
	go result.startSecondInstanceProcessor()

	return result
}

func (f *Frontend) startFileOpenProcessor() {
	for filePath := range openFilepathBuffer {
		f.ProcessOpenFileEvent(filePath)
	}
}

func (f *Frontend) startUrlOpenProcessor() {
	for url := range openUrlBuffer {
		f.ProcessOpenUrlEvent(url)
	}
}

func (f *Frontend) startSecondInstanceProcessor() {
	for secondInstanceData := range secondInstanceBuffer {
		if f.frontendOptions.SingleInstanceLock != nil &&
			f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch != nil {
			f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch(secondInstanceData)
		}
	}
}

func (f *Frontend) startMessageProcessor() {
	for message := range messageBuffer {
		f.processMessage(message)
	}
}

func (f *Frontend) startRequestProcessor() {
	for request := range requestBuffer {
		f.assets.ServeWebViewRequest(request)
	}
}

func (f *Frontend) startCallbackProcessor() {
	for callback := range callbackBuffer {
		err := f.handleCallback(callback)
		if err != nil {
			println(err.Error())
		}
	}
}


// ff:
func (f *Frontend) WindowReload() {
	f.ExecJS("runtime.WindowReload();")
}


// ff:
func (f *Frontend) WindowReloadApp() {
	f.ExecJS(fmt.Sprintf("window.location.href = '%s';", f.startURL))
}


// ff:
func (f *Frontend) WindowSetSystemDefaultTheme() {
}


// ff:
func (f *Frontend) WindowSetLightTheme() {
}


// ff:
func (f *Frontend) WindowSetDarkTheme() {
}


// ff:
// ctx:
func (f *Frontend) Run(ctx context.Context) error {
	f.ctx = ctx

	if f.frontendOptions.SingleInstanceLock != nil {
		SetupSingleInstance(f.frontendOptions.SingleInstanceLock.UniqueId)
	}

	_debug := ctx.Value("debug")
	_devtoolsEnabled := ctx.Value("devtoolsEnabled")

	if _debug != nil {
		f.debug = _debug.(bool)
	}
	if _devtoolsEnabled != nil {
		f.devtoolsEnabled = _devtoolsEnabled.(bool)
	}

	mainWindow := NewWindow(f.frontendOptions, f.debug, f.devtoolsEnabled)
	f.mainWindow = mainWindow
	f.mainWindow.Center()

	go func() {
		if f.frontendOptions.OnStartup != nil {
			f.frontendOptions.OnStartup(f.ctx)
		}
	}()
	mainWindow.Run(f.startURL.String())
	return nil
}


// ff:
func (f *Frontend) WindowCenter() {
	f.mainWindow.Center()
}


// ff:
// onTop:
func (f *Frontend) WindowSetAlwaysOnTop(onTop bool) {
	f.mainWindow.SetAlwaysOnTop(onTop)
}


// ff:
// y:
// x:
func (f *Frontend) WindowSetPosition(x, y int) {
	f.mainWindow.SetPosition(x, y)
}


// ff:
func (f *Frontend) WindowGetPosition() (int, int) {
	return f.mainWindow.GetPosition()
}


// ff:
// height:
// width:
func (f *Frontend) WindowSetSize(width, height int) {
	f.mainWindow.SetSize(width, height)
}


// ff:
func (f *Frontend) WindowGetSize() (int, int) {
	return f.mainWindow.Size()
}


// ff:
// title:
func (f *Frontend) WindowSetTitle(title string) {
	f.mainWindow.SetTitle(title)
}


// ff:
func (f *Frontend) WindowFullscreen() {
	f.mainWindow.Fullscreen()
}


// ff:
func (f *Frontend) WindowUnfullscreen() {
	f.mainWindow.UnFullscreen()
}


// ff:
func (f *Frontend) WindowShow() {
	f.mainWindow.Show()
}


// ff:
func (f *Frontend) WindowHide() {
	f.mainWindow.Hide()
}


// ff:
func (f *Frontend) Show() {
	f.mainWindow.ShowApplication()
}


// ff:
func (f *Frontend) Hide() {
	f.mainWindow.HideApplication()
}


// ff:
func (f *Frontend) WindowMaximise() {
	f.mainWindow.Maximise()
}


// ff:
func (f *Frontend) WindowToggleMaximise() {
	f.mainWindow.ToggleMaximise()
}


// ff:
func (f *Frontend) WindowUnmaximise() {
	f.mainWindow.UnMaximise()
}


// ff:
func (f *Frontend) WindowMinimise() {
	f.mainWindow.Minimise()
}


// ff:
func (f *Frontend) WindowUnminimise() {
	f.mainWindow.UnMinimise()
}


// ff:
// height:
// width:
func (f *Frontend) WindowSetMinSize(width int, height int) {
	f.mainWindow.SetMinSize(width, height)
}


// ff:
// height:
// width:
func (f *Frontend) WindowSetMaxSize(width int, height int) {
	f.mainWindow.SetMaxSize(width, height)
}


// ff:
// col:
func (f *Frontend) WindowSetBackgroundColour(col *options.RGBA) {
	if col == nil {
		return
	}
	f.mainWindow.SetBackgroundColour(col.R, col.G, col.B, col.A)
}


// ff:
func (f *Frontend) ScreenGetAll() ([]frontend.Screen, error) {
	return GetAllScreens(f.mainWindow.context)
}


// ff:
func (f *Frontend) WindowIsMaximised() bool {
	return f.mainWindow.IsMaximised()
}


// ff:
func (f *Frontend) WindowIsMinimised() bool {
	return f.mainWindow.IsMinimised()
}


// ff:
func (f *Frontend) WindowIsNormal() bool {
	return f.mainWindow.IsNormal()
}


// ff:
func (f *Frontend) WindowIsFullscreen() bool {
	return f.mainWindow.IsFullScreen()
}


// ff:
func (f *Frontend) Quit() {
	if f.frontendOptions.OnBeforeClose != nil {
		go func() {
			if !f.frontendOptions.OnBeforeClose(f.ctx) {
				f.mainWindow.Quit()
			}
		}()
		return
	}
	f.mainWindow.Quit()
}


// ff:
func (f *Frontend) WindowPrint() {
	f.mainWindow.Print()
}

type EventNotify struct {
	Name string        `json:"name"`
	Data []interface{} `json:"data"`
}


// ff:
// data:
// name:
func (f *Frontend) Notify(name string, data ...interface{}) {
	notification := EventNotify{
		Name: name,
		Data: data,
	}
	payload, err := json.Marshal(notification)
	if err != nil {
		f.logger.Error(err.Error())
		return
	}
	f.ExecJS(`window.wails.EventsNotify('` + template.JSEscapeString(string(payload)) + `');`)
}

func (f *Frontend) processMessage(message string) {
	if message == "DomReady" {
		if f.frontendOptions.OnDomReady != nil {
			f.frontendOptions.OnDomReady(f.ctx)
		}
		return
	}

	if message == "runtime:ready" {
		cmd := fmt.Sprintf("window.wails.setCSSDragProperties('%s', '%s');", f.frontendOptions.CSSDragProperty, f.frontendOptions.CSSDragValue)
		f.ExecJS(cmd)
		return
	}

	if message == "wails:openInspector" {
		showInspector(f.mainWindow.context)
		return
	}

// 如果字符串message以"systemevent:"开头 {
//     f.processSystemEvent(message) //调用处理系统事件的方法
//     return //结束当前函数执行
// }

	go func() {
		result, err := f.dispatcher.ProcessMessage(message, f)
		if err != nil {
			f.logger.Error(err.Error())
			f.Callback(result)
			return
		}
		if result == "" {
			return
		}

		switch result[0] {
		case 'c':
			// 从方法调用返回的回调函数
			f.Callback(result[1:])
		default:
			f.logger.Info("Unknown message returned from dispatcher: %+v", result)
		}
	}()
}


// ff:
// filePath:
func (f *Frontend) ProcessOpenFileEvent(filePath string) {
	if f.frontendOptions.Mac != nil && f.frontendOptions.Mac.OnFileOpen != nil {
		f.frontendOptions.Mac.OnFileOpen(filePath)
	}
}


// ff:
// url:
func (f *Frontend) ProcessOpenUrlEvent(url string) {
	if f.frontendOptions.Mac != nil && f.frontendOptions.Mac.OnUrlOpen != nil {
		f.frontendOptions.Mac.OnUrlOpen(url)
	}
}


// ff:
// message:
func (f *Frontend) Callback(message string) {
	escaped, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	f.ExecJS(`window.wails.Callback(` + string(escaped) + `);`)
}


// ff:
// js:
func (f *Frontend) ExecJS(js string) {
	f.mainWindow.ExecJS(js)
}

// 对于函数 (f *Frontend) processSystemEvent(message string) {
// 将接收到的消息字符串message以":"为分隔符进行切割
// 如果切割后的子串数量不为2，则
// 记录错误日志，输出无效的系统消息，并直接返回
// 根据第二个子串（下标为1）的值进行判断并执行相应操作
// 若第二个子串为 "fullscreen"
// 则禁用主窗口大小的约束限制
// 若第二个子串为 "unfullscreen"
// 则启用主窗口大小的约束限制
// 若上述情况均不符合
// 记录错误日志，输出未知的系统消息
// }

//export processMessage
func processMessage(message *C.char) {
	goMessage := C.GoString(message)
	messageBuffer <- goMessage
}

//export processCallback
func processCallback(callbackID uint) {
	callbackBuffer <- callbackID
}

//export processURLRequest
// 导出processURLRequest函数（供C语言调用）
// 在Golang的cgo中，`//export`关键字用于声明一个Go函数，表示该函数可供C代码通过C ABI（应用程序二进制接口）进行调用。因此，这段注释翻译为：
// 声明导出函数processURLRequest，以便C代码能够调用
func processURLRequest(_ unsafe.Pointer, wkURLSchemeTask unsafe.Pointer) {
	requestBuffer <- webview.NewRequest(wkURLSchemeTask)
}

//export HandleOpenFile

// ff:
// filePath:
func HandleOpenFile(filePath *C.char) {
	goFilepath := C.GoString(filePath)
	openFilepathBuffer <- goFilepath
}

//export HandleCustomProtocol
// 导出HandleCustomProtocol函数，以便在C语言或其他外部环境中调用
// （该注释表明这个Go函数是用于被C语言或者其他需要与Go互操作的环境调用的，通过`//export`标记，Go构建工具会生成相应的导出符号，使得该函数可以在其他语言中被调用。）

// ff:
// url:
func HandleCustomProtocol(url *C.char) {
	goUrl := C.GoString(url)
	openUrlBuffer <- goUrl
}
