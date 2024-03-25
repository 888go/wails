package assetserver

import (
	"net/http"
	"path/filepath"
	"sync"

	"github.com/wailsapp/mimetype"
)

var (
	mimeCache = map[string]string{}
	mimeMutex sync.Mutex

// 这是Go语言标准库包"mime"中定义的，根据扩展名列出的内置MIME类型列表。
// Go语言标准库还会考虑从诸如'/etc/apache2/mime.types'等etc文件中加载MIME类型定义，但我们希望在所有平台上保持一致的行为，并且不依赖于任何外部文件。
	mimeTypesByExt = map[string]string{
		".avif": "image/avif",
		".css":  "text/css; charset=utf-8",
		".gif":  "image/gif",
		".htm":  "text/html; charset=utf-8",
		".html": "text/html; charset=utf-8",
		".jpeg": "image/jpeg",
		".jpg":  "image/jpeg",
		".js":   "text/javascript; charset=utf-8",
		".json": "application/json",
		".mjs":  "text/javascript; charset=utf-8",
		".pdf":  "application/pdf",
		".png":  "image/png",
		".svg":  "image/svg+xml",
		".wasm": "application/wasm",
		".webp": "image/webp",
		".xml":  "text/xml; charset=utf-8",
	}
)

func GetMimetype(filename string, data []byte) string {
	mimeMutex.Lock()
	defer mimeMutex.Unlock()

	result := mimeTypesByExt[filepath.Ext(filename)]
	if result != "" {
		return result
	}

	result = mimeCache[filename]
	if result != "" {
		return result
	}

	detect := mimetype.Detect(data)
	if detect == nil {
		result = http.DetectContentType(data)
	} else {
		result = detect.String()
	}

	if result == "" {
		result = "application/octet-stream"
	}

	mimeCache[filename] = result
	return result
}
