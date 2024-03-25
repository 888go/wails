package dispatcher

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/888go/wails/pkg/runtime"

	"github.com/888go/wails/internal/frontend"
)

const systemCallPrefix = ":wails:"

type position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type size struct {
	W int `json:"w"`
	H int `json:"h"`
}

func (d *Dispatcher) processSystemCall(payload callMessage, sender frontend.Frontend) (interface{}, error) {
	// Strip prefix
	name := strings.TrimPrefix(payload.Name, systemCallPrefix)

	switch name {
	case "WindowGetPos":
		x, y := sender.X窗口取位置()
		return &position{x, y}, nil
	case "WindowGetSize":
		w, h := sender.X窗口取尺寸()
		return &size{w, h}, nil
	case "ScreenGetAll":
		return sender.X取屏幕信息()
	case "WindowIsMaximised":
		return sender.X窗口是否最大化(), nil
	case "WindowIsMinimised":
		return sender.X窗口是否最小化(), nil
	case "WindowIsNormal":
		return sender.X窗口是否为正常(), nil
	case "WindowIsFullscreen":
		return sender.X窗口是否全屏(), nil
	case "Environment":
		return runtime.X取环境信息(d.ctx), nil
	case "ClipboardGetText":
		t, err := sender.X剪贴板取文本()
		return t, err
	case "ClipboardSetText":
		if len(payload.Args) < 1 {
			return false, errors.New("empty argument, cannot set clipboard")
		}
		var arg string
		if err := json.Unmarshal(payload.Args[0], &arg); err != nil {
			return false, err
		}
		if err := sender.X剪贴板置文本(arg); err != nil {
			return false, err
		}
		return true, nil
	default:
		return nil, fmt.Errorf("unknown systemcall message: %s", payload.Name)
	}
}
