package assetserver

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"golang.org/x/net/html"
	"html/template"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const (
	runtimeJSPath = "/wails/runtime.js"
	ipcJSPath     = "/wails/ipc.js"
	runtimePath   = "/wails/runtime"
)

type RuntimeAssets interface {
	DesktopIPC() []byte
	WebsocketIPC() []byte
	RuntimeDesktopJS() []byte
}

type RuntimeHandler interface {
	HandleRuntimeCall(w http.ResponseWriter, r *http.Request)
}

type AssetServer struct {
	handler   http.Handler
	runtimeJS []byte
	ipcJS     func(*http.Request) []byte

	logger  Logger
	runtime RuntimeAssets

	servingFromDisk     bool
	appendSpinnerToBody bool

	// Use http based runtime
	runtimeHandler RuntimeHandler

	// plugin scripts
	pluginScripts map[string]string

	assetServerWebView
}


// ff:
// runtime:
// logger:
// servingFromDisk:
// options:
// bindingsJSON:
func NewAssetServerMainPage(bindingsJSON string, options *options.App, servingFromDisk bool, logger Logger, runtime RuntimeAssets) (*AssetServer, error) {
	assetOptions, err := BuildAssetServerConfig(options)
	if err != nil {
		return nil, err
	}
	return NewAssetServer(bindingsJSON, assetOptions, servingFromDisk, logger, runtime)
}


// ff:
// runtime:
// logger:
// servingFromDisk:
// options:
// bindingsJSON:
func NewAssetServer(bindingsJSON string, options assetserver.Options, servingFromDisk bool, logger Logger, runtime RuntimeAssets) (*AssetServer, error) {
	handler, err := NewAssetHandler(options, logger)
	if err != nil {
		return nil, err
	}

	return NewAssetServerWithHandler(handler, bindingsJSON, servingFromDisk, logger, runtime)
}


// ff:
// runtime:
// logger:
// servingFromDisk:
// bindingsJSON:
// handler:
func NewAssetServerWithHandler(handler http.Handler, bindingsJSON string, servingFromDisk bool, logger Logger, runtime RuntimeAssets) (*AssetServer, error) {

	var buffer bytes.Buffer
	if bindingsJSON != "" {
		escapedBindingsJSON := template.JSEscapeString(bindingsJSON)
		buffer.WriteString(`window.wailsbindings='` + escapedBindingsJSON + `';` + "\n")
	}
	buffer.Write(runtime.RuntimeDesktopJS())

	result := &AssetServer{
		handler:   handler,
		runtimeJS: buffer.Bytes(),

// 检查我们是否已指定一个目录用于提供静态资源。
// 如果是这样，这意味着我们处于开发模式，且直接从磁盘提供资源。
// 我们通过`servingFromDisk`标志来表明这一点，以确保在开发模式下请求不会被缓存。
		servingFromDisk: servingFromDisk,
		logger:          logger,
		runtime:         runtime,
	}

	return result, nil
}


// ff:
// handler:
func (d *AssetServer) UseRuntimeHandler(handler RuntimeHandler) {
	d.runtimeHandler = handler
}


// ff:
// script:
// pluginName:
func (d *AssetServer) AddPluginScript(pluginName string, script string) {
	if d.pluginScripts == nil {
		d.pluginScripts = make(map[string]string)
	}
	pluginName = strings.ReplaceAll(pluginName, "/", "_")
	pluginName = html.EscapeString(pluginName)
	pluginScriptName := fmt.Sprintf("/plugin_%s_%d.js", pluginName, rand.Intn(100000))
	d.pluginScripts[pluginScriptName] = script
}


// ff:
// req:
// rw:
func (d *AssetServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if isWebSocket(req) {
		// WebSockets 不被 AssetServer 支持
		rw.WriteHeader(http.StatusNotImplemented)
		return
	}

	if d.servingFromDisk {
		rw.Header().Add(HeaderCacheControl, "no-cache")
	}

	handler := d.handler
	if req.Method != http.MethodGet {
		handler.ServeHTTP(rw, req)
		return
	}

	path := req.URL.Path
	if path == runtimeJSPath {
		d.writeBlob(rw, path, d.runtimeJS)
	} else if path == runtimePath && d.runtimeHandler != nil {
		d.runtimeHandler.HandleRuntimeCall(rw, req)
	} else if path == ipcJSPath {
		content := d.runtime.DesktopIPC()
		if d.ipcJS != nil {
			content = d.ipcJS(req)
		}
		d.writeBlob(rw, path, content)

	} else if script, ok := d.pluginScripts[path]; ok {
		d.writeBlob(rw, path, []byte(script))
	} else if d.isRuntimeInjectionMatch(path) {
		recorder := &bodyRecorder{
			ResponseWriter: rw,
			doRecord: func(code int, h http.Header) bool {
				if code == http.StatusNotFound {
					return true
				}

				if code != http.StatusOK {
					return false
				}

				return strings.Contains(h.Get(HeaderContentType), "text/html")
			},
		}

		handler.ServeHTTP(recorder, req)

		body := recorder.Body()
		if body == nil {
			// 已经流式传输了主体且未被记录，我们已经完成
			return
		}

		code := recorder.Code()
		switch code {
		case http.StatusOK:
			content, err := d.processIndexHTML(body.Bytes())
			if err != nil {
				d.serveError(rw, err, "Unable to processIndexHTML")
				return
			}
			d.writeBlob(rw, indexHTML, content)

		case http.StatusNotFound:
			d.writeBlob(rw, indexHTML, defaultHTML)

		default:
			rw.WriteHeader(code)

		}

	} else {
		handler.ServeHTTP(rw, req)
	}
}

func (d *AssetServer) processIndexHTML(indexHTML []byte) ([]byte, error) {
	htmlNode, err := getHTMLNode(indexHTML)
	if err != nil {
		return nil, err
	}

	if d.appendSpinnerToBody {
		err = appendSpinnerToBody(htmlNode)
		if err != nil {
			return nil, err
		}
	}

	if err := insertScriptInHead(htmlNode, runtimeJSPath); err != nil {
		return nil, err
	}

	if err := insertScriptInHead(htmlNode, ipcJSPath); err != nil {
		return nil, err
	}

	// Inject plugins
	for scriptName := range d.pluginScripts {
		if err := insertScriptInHead(htmlNode, scriptName); err != nil {
			return nil, err
		}
	}

	var buffer bytes.Buffer
	err = html.Render(&buffer, htmlNode)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (d *AssetServer) writeBlob(rw http.ResponseWriter, filename string, blob []byte) {
	err := serveFile(rw, filename, blob)
	if err != nil {
		d.serveError(rw, err, "Unable to write content %s", filename)
	}
}

func (d *AssetServer) serveError(rw http.ResponseWriter, err error, msg string, args ...interface{}) {
	args = append(args, err)
	d.logError(msg+": %s", args...)
	rw.WriteHeader(http.StatusInternalServerError)
}

func (d *AssetServer) logDebug(message string, args ...interface{}) {
	if d.logger != nil {
		d.logger.Debug("[AssetServer] "+message, args...)
	}
}

func (d *AssetServer) logError(message string, args ...interface{}) {
	if d.logger != nil {
		d.logger.Error("[AssetServer] "+message, args...)
	}
}

func (AssetServer) isRuntimeInjectionMatch(path string) bool {
	if path == "" {
		path = "/"
	}

	return strings.HasSuffix(path, "/") ||
		strings.HasSuffix(path, "/"+indexHTML)
}
