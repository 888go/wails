//go:build dev
// +build dev

// Package devserver 提供了一个基于Web的前端界面，以便于
// 在浏览器中运行Wails应用。
package devserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"sync"

	"github.com/888go/wails/pkg/assetserver"

	"github.com/888go/wails/internal/frontend/runtime"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/internal/menumanager"
	"github.com/888go/wails/pkg/options"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type Screen = frontend.Screen

type DevWebServer struct {
	server           *echo.Echo
	ctx              context.Context
	appoptions       *options.App
	logger           *logger.Logger
	appBindings      *binding.Bindings
	dispatcher       frontend.Dispatcher
	socketMutex      sync.Mutex
	websocketClients map[*websocket.Conn]*sync.Mutex
	menuManager      *menumanager.Manager
	starttime        string

	// Desktop frontend
	frontend.Frontend

	devServerAddr string
}

// ff:运行
// ctx:上下文
func (d *DevWebServer) Run(ctx context.Context) error {
	d.ctx = ctx

	d.server.GET("/wails/reload", d.handleReload)
	d.server.GET("/wails/ipc", d.handleIPCWebSocket)

	assetServerConfig, err := assetserver.BuildAssetServerConfig(d.appoptions)
	if err != nil {
		return err
	}

	var myLogger assetserver.Logger
	if _logger := ctx.Value("logger"); _logger != nil {
		myLogger = _logger.(*logger.Logger)
	}

	var wsHandler http.Handler

	_fronendDevServerURL, _ := ctx.Value("frontenddevserverurl").(string)
	if _fronendDevServerURL == "" {
		assetdir, _ := ctx.Value("assetdir").(string)
		d.server.GET("/wails/assetdir", func(c echo.Context) error {
			return c.String(http.StatusOK, assetdir)
		})

	} else {
		externalURL, err := url.Parse(_fronendDevServerURL)
		if err != nil {
			return err
		}

		// 目前在生产模式下不支持WebSockets，因此WebSocket连接是通过
		// 前端开发服务器（如Vite）建立的，目的是为了支持自动重载。
		// 因此，我们直接将WebSocket连接导向前端开发服务器，而不是返回一个未实现状态（NotImplementedStatus）。
		wsHandler = httputil.NewSingleHostReverseProxy(externalURL)
	}

	assetHandler, err := assetserver.NewAssetHandler(assetServerConfig, myLogger)
	if err != nil {
		log.Fatal(err)
	}

	// 设置内部开发服务器
	bindingsJSON, err := d.appBindings.ToJSON()
	if err != nil {
		log.Fatal(err)
	}

	assetServer, err := assetserver.NewDevAssetServer(assetHandler, bindingsJSON, ctx.Value("assetdir") != nil, myLogger, runtime.RuntimeAssetsBundle)
	if err != nil {
		log.Fatal(err)
	}

	d.server.Any("/*", func(c echo.Context) error {
		if c.IsWebSocket() {
			wsHandler.ServeHTTP(c.Response(), c.Request())
		} else {
			assetServer.ServeHTTP(c.Response(), c.Request())
		}
		return nil
	})

	if devServerAddr := d.devServerAddr; devServerAddr != "" {
		// Start server
		go func(server *echo.Echo, log *logger.Logger) {
			err := server.Start(devServerAddr)
			if err != nil {
				log.Error(err.Error())
			}
			d.LogDebug("Shutdown completed")
		}(d.server, d.logger)

		d.LogDebug("Serving DevServer at http://%s", devServerAddr)
	}

	// Launch desktop app
	err = d.Frontend.Run(ctx)

	return err
}

// ff:窗口重载
func (d *DevWebServer) WindowReload() {
	d.broadcast("reload")
	d.Frontend.WindowReload()
}

// ff:窗口重载应用程序前端
func (d *DevWebServer) WindowReloadApp() {
	d.broadcast("reloadapp")
	d.Frontend.WindowReloadApp()
}

// ff:
// data:
// name:
func (d *DevWebServer) Notify(name string, data ...interface{}) {
	d.notify(name, data...)
}

func (d *DevWebServer) handleReload(c echo.Context) error {
	d.WindowReload()
	return c.NoContent(http.StatusNoContent)
}

func (d *DevWebServer) handleReloadApp(c echo.Context) error {
	d.WindowReloadApp()
	return c.NoContent(http.StatusNoContent)
}

func (d *DevWebServer) handleIPCWebSocket(c echo.Context) error {
	websocket.Handler(func(c *websocket.Conn) {
		d.LogDebug(fmt.Sprintf("Websocket client %p connected", c))
		d.socketMutex.Lock()
		d.websocketClients[c] = &sync.Mutex{}
		locker := d.websocketClients[c]
		d.socketMutex.Unlock()

		defer func() {
			d.socketMutex.Lock()
			delete(d.websocketClients, c)
			d.socketMutex.Unlock()
			d.LogDebug(fmt.Sprintf("Websocket client %p disconnected", c))
		}()

		var msg string
		defer c.Close()
		for {
			if err := websocket.Message.Receive(c, &msg); err != nil {
				break
			}
			// 我们不支持在浏览器中拖拽
			if msg == "drag" {
				continue
			}

			// 通知其他浏览器关于"EventEmit"事件
			if len(msg) > 2 && strings.HasPrefix(string(msg), "EE") {
				d.notifyExcludingSender([]byte(msg), c)
			}

			// 将消息发送至调度程序以分发到前端
			result, err := d.dispatcher.ProcessMessage(string(msg), d)
			if err != nil {
				d.logger.Error(err.Error())
			}
			if result != "" {
				locker.Lock()
				if err = websocket.Message.Send(c, result); err != nil {
					locker.Unlock()
					break
				}
				locker.Unlock()
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}

// ff:日志调试
// args:参数
// message:消息
func (d *DevWebServer) LogDebug(message string, args ...interface{}) {
	d.logger.Debug("[DevWebServer] "+message, args...)
}

type EventNotify struct {
	Name string        `json:"name"`
	Data []interface{} `json:"data"`
}

func (d *DevWebServer) broadcast(message string) {
	d.socketMutex.Lock()
	defer d.socketMutex.Unlock()
	for client, locker := range d.websocketClients {
		go func(client *websocket.Conn, locker *sync.Mutex) {
			if client == nil {
				d.logger.Error("Lost connection to websocket server")
				return
			}
			locker.Lock()
			err := websocket.Message.Send(client, message)
			if err != nil {
				locker.Unlock()
				d.logger.Error(err.Error())
				return
			}
			locker.Unlock()
		}(client, locker)
	}
}

func (d *DevWebServer) notify(name string, data ...interface{}) {
	// Notify
	notification := EventNotify{
		Name: name,
		Data: data,
	}
	payload, err := json.Marshal(notification)
	if err != nil {
		d.logger.Error(err.Error())
		return
	}
	d.broadcast("n" + string(payload))
}

func (d *DevWebServer) broadcastExcludingSender(message string, sender *websocket.Conn) {
	d.socketMutex.Lock()
	defer d.socketMutex.Unlock()
	for client, locker := range d.websocketClients {
		go func(client *websocket.Conn, locker *sync.Mutex) {
			if client == sender {
				return
			}
			locker.Lock()
			err := websocket.Message.Send(client, message)
			if err != nil {
				locker.Unlock()
				d.logger.Error(err.Error())
				return
			}
			locker.Unlock()
		}(client, locker)
	}
}

func (d *DevWebServer) notifyExcludingSender(eventMessage []byte, sender *websocket.Conn) {
	message := "n" + string(eventMessage[2:])
	d.broadcastExcludingSender(message, sender)

	var notifyMessage EventNotify
	err := json.Unmarshal(eventMessage[2:], &notifyMessage)
	if err != nil {
		d.logger.Error(err.Error())
		return
	}
	d.Frontend.Notify(notifyMessage.Name, notifyMessage.Data...)
}

// ff:
// desktopFrontend:
// menuManager:
// dispatcher:
// appBindings:
// myLogger:
// appoptions:
// ctx:
func NewFrontend(ctx context.Context, appoptions *options.App, myLogger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher, menuManager *menumanager.Manager, desktopFrontend frontend.Frontend) *DevWebServer {
	result := &DevWebServer{
		ctx:              ctx,
		Frontend:         desktopFrontend,
		appoptions:       appoptions,
		logger:           myLogger,
		appBindings:      appBindings,
		dispatcher:       dispatcher,
		server:           echo.New(),
		menuManager:      menuManager,
		websocketClients: make(map[*websocket.Conn]*sync.Mutex),
	}

	result.devServerAddr, _ = ctx.Value("devserver").(string)
	result.server.HideBanner = true
	result.server.HidePort = true
	return result
}
