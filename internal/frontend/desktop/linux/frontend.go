//go:build linux
// +build linux

package linux

/*
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0

#include "gtk/gtk.h"
#include "webkit2/webkit2.h"

// 代码来源：https://github.com/rainycape/magick
#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <string.h>

static void fix_signal(int signum)
{
    struct sigaction st;

    if (sigaction(signum, NULL, &st) < 0) {
        goto fix_signal_error;
    }
    st.sa_flags |= SA_ONSTACK;
    if (sigaction(signum, &st,  NULL) < 0) {
        goto fix_signal_error;
    }
    return;
fix_signal_error:
        fprintf(stderr, "error fixing handler for signal %d, please "
                "report this issue to "
                "https://github.com/wailsapp/wails: %s\n",
                signum, strerror(errno));
}

static void install_signal_handlers()
{
#if defined(SIGCHLD)
    fix_signal(SIGCHLD);
#endif
#if defined(SIGHUP)
    fix_signal(SIGHUP);
#endif
#if defined(SIGINT)
    fix_signal(SIGINT);
#endif
#if defined(SIGQUIT)
    fix_signal(SIGQUIT);
#endif
#if defined(SIGABRT)
    fix_signal(SIGABRT);
#endif
#if defined(SIGFPE)
    fix_signal(SIGFPE);
#endif
#if defined(SIGTERM)
    fix_signal(SIGTERM);
#endif
#if defined(SIGBUS)
    fix_signal(SIGBUS);
#endif
#if defined(SIGSEGV)
    fix_signal(SIGSEGV);
#endif
#if defined(SIGXCPU)
    fix_signal(SIGXCPU);
#endif
#if defined(SIGXFSZ)
    fix_signal(SIGXFSZ);
#endif
}

*/
import "C"
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"
	"text/template"
	"unsafe"

	"github.com/888go/wails/pkg/assetserver"
	"github.com/888go/wails/pkg/assetserver/webview"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend"
	wailsruntime "github.com/888go/wails/internal/frontend/runtime"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/pkg/options"
)

var initOnce = sync.Once{}

const startURL = "wails://wails/"

var secondInstanceBuffer = make(chan options.SecondInstanceData, 1)

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
	C.gtk_main()
}


// ff:
func (f *Frontend) WindowClose() {
	f.mainWindow.Destroy()
}


// ff:
// dispatcher:
// appBindings:
// myLogger:
// appoptions:
// ctx:
func NewFrontend(ctx context.Context, appoptions *options.App, myLogger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) *Frontend {
	initOnce.Do(func() {
		runtime.LockOSThread()

		// 如果当前未设置GDK_BACKEND，并且XDG_SESSION_TYPE也未设置、未指定或为x11，则设置GDK_BACKEND为x11，以防止出现警告
		if os.Getenv("GDK_BACKEND") == "" && (os.Getenv("XDG_SESSION_TYPE") == "" || os.Getenv("XDG_SESSION_TYPE") == "unspecified" || os.Getenv("XDG_SESSION_TYPE") == "x11") {
			_ = os.Setenv("GDK_BACKEND", "x11")
		}

		if ok := C.gtk_init_check(nil, nil); ok != 1 {
			panic(errors.New("failed to init GTK"))
		}
	})

	result := &Frontend{
		frontendOptions: appoptions,
		logger:          myLogger,
		bindings:        appBindings,
		dispatcher:      dispatcher,
		ctx:             ctx,
	}
	result.startURL, _ = url.Parse(startURL)

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
		assets, err := assetserver.NewAssetServerMainPage(bindings, appoptions, ctx.Value("assetdir") != nil, myLogger, wailsruntime.RuntimeAssetsBundle)
		if err != nil {
			log.Fatal(err)
		}
		result.assets = assets

		go result.startRequestProcessor()
	}

	go result.startMessageProcessor()

	var _debug = ctx.Value("debug")
	var _devtoolsEnabled = ctx.Value("devtoolsEnabled")

	if _debug != nil {
		result.debug = _debug.(bool)
	}
	if _devtoolsEnabled != nil {
		result.devtoolsEnabled = _devtoolsEnabled.(bool)
	}

	result.mainWindow = NewWindow(appoptions, result.debug, result.devtoolsEnabled)

	C.install_signal_handlers()

	if appoptions.Linux != nil && appoptions.Linux.ProgramName != "" {
		prgname := C.CString(appoptions.Linux.ProgramName)
		C.g_set_prgname(prgname)
		C.free(unsafe.Pointer(prgname))
	}

	go result.startSecondInstanceProcessor()

	return result
}

func (f *Frontend) startMessageProcessor() {
	for message := range messageBuffer {
		f.processMessage(message)
	}
}


// ff:
func (f *Frontend) WindowReload() {
	f.ExecJS("runtime.WindowReload();")
}


// ff:
func (f *Frontend) WindowSetSystemDefaultTheme() {
	return
}


// ff:
func (f *Frontend) WindowSetLightTheme() {
	return
}


// ff:
func (f *Frontend) WindowSetDarkTheme() {
	return
}


// ff:
// ctx:
func (f *Frontend) Run(ctx context.Context) error {
	f.ctx = ctx

	go func() {
		if f.frontendOptions.OnStartup != nil {
			f.frontendOptions.OnStartup(f.ctx)
		}
	}()

	if f.frontendOptions.SingleInstanceLock != nil {
		SetupSingleInstance(f.frontendOptions.SingleInstanceLock.UniqueId)
	}

	f.mainWindow.Run(f.startURL.String())

	return nil
}


// ff:
func (f *Frontend) WindowCenter() {
	f.mainWindow.Center()
}


// ff:
// b:
func (f *Frontend) WindowSetAlwaysOnTop(b bool) {
	f.mainWindow.SetKeepAbove(b)
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
	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = false;")
	}
	f.mainWindow.Fullscreen()
}


// ff:
func (f *Frontend) WindowUnfullscreen() {
	if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
		f.ExecJS("window.wails.flags.enableResize = true;")
	}
	f.mainWindow.UnFullscreen()
}


// ff:
func (f *Frontend) WindowReloadApp() {
	f.ExecJS(fmt.Sprintf("window.location.href = '%s';", f.startURL))
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
	f.mainWindow.Show()
}


// ff:
func (f *Frontend) Hide() {
	f.mainWindow.Hide()
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
func (f *Frontend) ScreenGetAll() ([]Screen, error) {
	return GetAllScreens(f.mainWindow.asGTKWindow())
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
	f.ExecJS("window.print();")
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
	f.mainWindow.ExecJS(`window.wails.EventsNotify('` + template.JSEscapeString(string(payload)) + `');`)
}

var edgeMap = map[string]uintptr{
	"n-resize":  C.GDK_WINDOW_EDGE_NORTH,
	"ne-resize": C.GDK_WINDOW_EDGE_NORTH_EAST,
	"e-resize":  C.GDK_WINDOW_EDGE_EAST,
	"se-resize": C.GDK_WINDOW_EDGE_SOUTH_EAST,
	"s-resize":  C.GDK_WINDOW_EDGE_SOUTH,
	"sw-resize": C.GDK_WINDOW_EDGE_SOUTH_WEST,
	"w-resize":  C.GDK_WINDOW_EDGE_WEST,
	"nw-resize": C.GDK_WINDOW_EDGE_NORTH_WEST,
}

func (f *Frontend) processMessage(message string) {
	if message == "DomReady" {
		if f.frontendOptions.OnDomReady != nil {
			f.frontendOptions.OnDomReady(f.ctx)
		}
		return
	}

	if message == "drag" {
		if !f.mainWindow.IsFullScreen() {
			f.startDrag()
		}
		return
	}

	if message == "wails:showInspector" {
		f.mainWindow.ShowInspector()
		return
	}

	if strings.HasPrefix(message, "resize:") {
		if !f.mainWindow.IsFullScreen() {
			sl := strings.Split(message, ":")
			if len(sl) != 2 {
				f.logger.Info("Unknown message returned from dispatcher: %+v", message)
				return
			}
			edge := edgeMap[sl[1]]
			err := f.startResize(edge)
			if err != nil {
				f.logger.Error(err.Error())
			}
		}
		return
	}

	if message == "runtime:ready" {
		cmd := fmt.Sprintf(
			"window.wails.setCSSDragProperties('%s', '%s');\n"+
				"window.wails.flags.deferDragToMouseMove = true;", f.frontendOptions.CSSDragProperty, f.frontendOptions.CSSDragValue)
		f.ExecJS(cmd)

		if f.frontendOptions.Frameless && f.frontendOptions.DisableResize == false {
			f.ExecJS("window.wails.flags.enableResize = true;")
		}
		return
	}

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
// message:
func (f *Frontend) Callback(message string) {
	escaped, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	f.ExecJS(`window.wails.Callback(` + string(escaped) + `);`)
}

func (f *Frontend) startDrag() {
	f.mainWindow.StartDrag()
}

func (f *Frontend) startResize(edge uintptr) error {
	f.mainWindow.StartResize(edge)
	return nil
}


// ff:
// js:
func (f *Frontend) ExecJS(js string) {
	f.mainWindow.ExecJS(js)
}

var messageBuffer = make(chan string, 100)

//export processMessage
func processMessage(message *C.char) {
	goMessage := C.GoString(message)
	messageBuffer <- goMessage
}

var requestBuffer = make(chan webview.Request, 100)

func (f *Frontend) startRequestProcessor() {
	for request := range requestBuffer {
		f.assets.ServeWebViewRequest(request)
	}
}

//export processURLRequest
// 导出processURLRequest函数（供C语言调用）
// 在Golang的cgo中，`//export`关键字用于声明一个Go函数，表示该函数可供C代码通过C ABI（应用程序二进制接口）进行调用。因此，这段注释翻译为：
// 声明导出函数processURLRequest，以便C代码能够调用
func processURLRequest(request unsafe.Pointer) {
	requestBuffer <- webview.NewRequest(request)
}

func (f *Frontend) startSecondInstanceProcessor() {
	for secondInstanceData := range secondInstanceBuffer {
		if f.frontendOptions.SingleInstanceLock != nil &&
			f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch != nil {
			f.frontendOptions.SingleInstanceLock.OnSecondInstanceLaunch(secondInstanceData)
		}
	}
}
