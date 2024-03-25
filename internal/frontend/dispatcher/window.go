package dispatcher

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/pkg/options"
)

func (d *Dispatcher) mustAtoI(input string) int {
	result, err := strconv.Atoi(input)
	if err != nil {
		d.log.Error("cannot convert %s to integer!", input)
	}
	return result
}

func (d *Dispatcher) processWindowMessage(message string, sender frontend.Frontend) (string, error) {
	if len(message) < 2 {
		return "", errors.New("Invalid Window Message: " + message)
	}

	switch message[1] {
	case 'A':
		switch message[2:] {
		case "SDT":
			go sender.X窗口设置系统默认主题()
		case "LT":
			go sender.X窗口设置浅色主题()
		case "DT":
			go sender.X窗口设置深色主题()
		case "TP:0", "TP:1":
			if message[2:] == "TP:0" {
				go sender.X窗口设置置顶(false)
			} else if message[2:] == "TP:1" {
				go sender.X窗口设置置顶(true)
			}
		}
	case 'c':
		go sender.X窗口居中()
	case 'T':
		title := message[2:]
		go sender.X窗口设置标题(title)
	case 'F':
		go sender.X窗口设置全屏()
	case 'f':
		go sender.X窗口取消全屏()
	case 's':
		parts := strings.Split(message[3:], ":")
		w := d.mustAtoI(parts[0])
		h := d.mustAtoI(parts[1])
		go sender.X窗口设置尺寸(w, h)
	case 'p':
		parts := strings.Split(message[3:], ":")
		x := d.mustAtoI(parts[0])
		y := d.mustAtoI(parts[1])
		go sender.X窗口设置位置(x, y)
	case 'H':
		go sender.X窗口隐藏()
	case 'S':
		go sender.X窗口显示()
	case 'R':
		go sender.X窗口重载应用程序前端()
	case 'r':
		var rgba options.RGBA
		err := json.Unmarshal([]byte(message[3:]), &rgba)
		if err != nil {
			return "", err
		}
		go sender.X窗口设置背景色(&rgba)
	case 'M':
		go sender.X窗口最大化()
	case 't':
		go sender.X窗口最大化切换()
	case 'U':
		go sender.X窗口取消最大化()
	case 'm':
		go sender.X窗口最小化()
	case 'u':
		go sender.X窗口取消最小化()
	case 'Z':
		parts := strings.Split(message[3:], ":")
		w := d.mustAtoI(parts[0])
		h := d.mustAtoI(parts[1])
		go sender.X窗口设置最大尺寸(w, h)
	case 'z':
		parts := strings.Split(message[3:], ":")
		w := d.mustAtoI(parts[0])
		h := d.mustAtoI(parts[1])
		go sender.X窗口设置最小尺寸(w, h)
	default:
		d.log.Error("unknown Window message: %s", message)
	}

	return "", nil
}
