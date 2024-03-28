package assetserver

import (
	"bytes"
	"embed"
	"errors"
	"fmt"
	"io"
	iofs "io/fs"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

type Logger interface {
	Debug(message string, args ...interface{})
	Error(message string, args ...interface{})
}

//go:embed defaultindex.html
var defaultHTML []byte

const (
	indexHTML = "index.html"
)

type assetHandler struct {
	fs      iofs.FS
	handler http.Handler

	logger Logger

	retryMissingFiles bool
}


// ff:
// log:
// options:
func NewAssetHandler(options assetserver.Options, log Logger) (http.Handler, error) {
	vfs := options.Assets
	if vfs != nil {
		if _, err := vfs.Open("."); err != nil {
			return nil, err
		}

		subDir, err := FindPathToFile(vfs, indexHTML)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				msg := "no `index.html` could be found in your Assets fs.FS"
				if embedFs, isEmbedFs := vfs.(embed.FS); isEmbedFs {
					rootFolder, _ := FindEmbedRootPath(embedFs)
					msg += fmt.Sprintf(", please make sure the embedded directory '%s' is correct and contains your assets", rootFolder)
				}

				return nil, fmt.Errorf(msg)
			}

			return nil, err
		}

		vfs, err = iofs.Sub(vfs, path.Clean(subDir))
		if err != nil {
			return nil, err
		}
	}

	var result http.Handler = &assetHandler{
		fs:      vfs,
		handler: options.Handler,
		logger:  log,
	}

	if middleware := options.Middleware; middleware != nil {
		result = middleware(result)
	}

	return result, nil
}


// ff:
// req:
// rw:
func (d *assetHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	url := req.URL.Path
	handler := d.handler
	if strings.EqualFold(req.Method, http.MethodGet) {
		filename := path.Clean(strings.TrimPrefix(url, "/"))

		d.logDebug("Handling request '%s' (file='%s')", url, filename)
		if err := d.serveFSFile(rw, req, filename); err != nil {
			if os.IsNotExist(err) {
				if handler != nil {
					d.logDebug("File '%s' not found, serving '%s' by AssetHandler", filename, url)
					handler.ServeHTTP(rw, req)
					err = nil
				} else {
					rw.WriteHeader(http.StatusNotFound)
					err = nil
				}
			}

			if err != nil {
				d.logError("Unable to handle request '%s': %s", url, err)
				http.Error(rw, err.Error(), http.StatusInternalServerError)
			}
		}
	} else if handler != nil {
		d.logDebug("No GET request, serving '%s' by AssetHandler", url)
		handler.ServeHTTP(rw, req)
	} else {
		rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// serveFile 尝试从 fs.FS 加载文件并将其写入响应
func (d *assetHandler) serveFSFile(rw http.ResponseWriter, req *http.Request, filename string) error {
	if d.fs == nil {
		return os.ErrNotExist
	}

	file, err := d.fs.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	statInfo, err := file.Stat()
	if err != nil {
		return err
	}

	url := req.URL.Path
	isDirectoryPath := url == "" || url[len(url)-1] == '/'
	if statInfo.IsDir() {
		if !isDirectoryPath {
// 如果URL末尾通常没有斜杠，正常情况下应当执行http重定向，但在当前WebKit WebViews（macOS/Linux）环境下无法正常工作。
// 因此，我们将这种情况视为一个特定错误进行处理
			return fmt.Errorf("a directory has been requested without a trailing slash, please add a trailing slash to your request")
		}

		filename = path.Join(filename, indexHTML)

		file, err = d.fs.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		statInfo, err = file.Stat()
		if err != nil {
			return err
		}
	} else if isDirectoryPath {
		return fmt.Errorf("a file has been requested with a trailing slash, please remove the trailing slash from your request")
	}

	var buf [512]byte
	var n int
	if _, haveType := rw.Header()[HeaderContentType]; !haveType {
		// 通过嗅探前512个字节检测MimeType
		n, err = file.Read(buf[:])
		if err != nil && err != io.EOF {
			return err
		}

// 即使在 io.ReadSeeker 的情况下 http.ServeContent 会执行自定义 MimeType 探测，我们也进行自定义 MimeType 探测操作。我们希望在这两种情况下都具有一致的行为。
		if contentType := GetMimetype(filename, buf[:n]); contentType != "" {
			rw.Header().Set(HeaderContentType, contentType)
		}
	}

	if fileSeeker, _ := file.(io.ReadSeeker); fileSeeker != nil {
		if _, err := fileSeeker.Seek(0, io.SeekStart); err != nil {
			return fmt.Errorf("seeker can't seek")
		}

		http.ServeContent(rw, req, statInfo.Name(), statInfo.ModTime(), fileSeeker)
		return nil
	}

	size := strconv.FormatInt(statInfo.Size(), 10)
	rw.Header().Set(HeaderContentLength, size)

	// 写入前512字节，用于MimeType检测
	_, err = io.Copy(rw, bytes.NewReader(buf[:n]))
	if err != nil {
		return err
	}

	// 复制文件剩余内容
	_, err = io.Copy(rw, file)
	return err
}

func (d *assetHandler) logDebug(message string, args ...interface{}) {
	if d.logger != nil {
		d.logger.Debug("[AssetHandler] "+message, args...)
	}
}

func (d *assetHandler) logError(message string, args ...interface{}) {
	if d.logger != nil {
		d.logger.Error("[AssetHandler] "+message, args...)
	}
}
