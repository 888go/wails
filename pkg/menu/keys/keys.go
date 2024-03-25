package keys

import (
	"fmt"
	"strings"
)

// Modifier 实际上是一个字符串
type Modifier string

const (
	// CmdOrCtrlKey 表示在 Mac 平台上代表 Command 键，在其他平台上代表 Control 键
	CmdOrCtrlKey Modifier = "cmdorctrl"
	// OptionOrAltKey 表示在 Mac 平台上代表 Option 键，在其他平台上代表 Alt 键
	OptionOrAltKey Modifier = "optionoralt"
	// ShiftKey 表示在所有系统上的 shift 键
	ShiftKey Modifier = "shift"
// SuperKey 表示在 Mac 上的 Command 键，在其他平台（如 Windows）上表示 Windows 键
// SuperKey Modifier = "super"
// ControlKey 代表在所有系统上的控制键
// 注释翻译：
// ```go
// SuperKey 代表在 Mac 操作系统中的 Command 键，在非 Mac 平台（如 Windows 系统）上则代表 Windows 键
// SuperKey 对应的修饰键名称为 "super"
// ControlKey 在所有操作系统中均表示控制键
	ControlKey Modifier = "ctrl"
)

var modifierMap = map[string]Modifier{
	"cmdorctrl":   CmdOrCtrlKey,
	"optionoralt": OptionOrAltKey,
	"shift":       ShiftKey,
	// "super":       超级键，
	"ctrl": ControlKey,
}

func parseModifier(text string) (*Modifier, error) {
	lowertext := strings.ToLower(text)
	result, valid := modifierMap[lowertext]
	if !valid {
		return nil, fmt.Errorf("'%s' is not a valid modifier", text)
	}

	return &result, nil
}

// Accelerator 保存了菜单项的键盘快捷键
type Accelerator struct {
	Key       string
	Modifiers []Modifier
}

// Key 创建一个标准的键Accelerator
func Key(key string) *Accelerator {
	return &Accelerator{
		Key: strings.ToLower(key),
	}
}

// CmdOrCtrl 创建一个 'CmdOrCtrl' 快捷键
func CmdOrCtrl(key string) *Accelerator {
	return &Accelerator{
		Key:       strings.ToLower(key),
		Modifiers: []Modifier{CmdOrCtrlKey},
	}
}

// OptionOrAlt 创建一个 'OptionOrAlt' 加速器
func OptionOrAlt(key string) *Accelerator {
	return &Accelerator{
		Key:       strings.ToLower(key),
		Modifiers: []Modifier{OptionOrAltKey},
	}
}

// Shift 创建一个“Shift”加速器
func Shift(key string) *Accelerator {
	return &Accelerator{
		Key:       strings.ToLower(key),
		Modifiers: []Modifier{ShiftKey},
	}
}

// Control 创建一个名为'Control'的加速器
func Control(key string) *Accelerator {
	return &Accelerator{
		Key:       strings.ToLower(key),
		Modifiers: []Modifier{ControlKey},
	}
}

// 
//// Super 函数用于创建一个 'Super' 加速器
//func Super(key string) *Accelerator {
//	// 将输入的 key 转换为小写并初始化 Accelerator 结构体实例
//	return &Accelerator{
//		Key:       strings.ToLower(key), // 设置 Key 字段为小写形式的 key
//		Modifiers: []Modifier{SuperKey}, // 设置 Modifiers 字段，包含 SuperKey 模块
//	}
//}

// Combo 创建一个带有多个修饰符的 Accelerator
func Combo(key string, modifier1 Modifier, modifier2 Modifier, rest ...Modifier) *Accelerator {
	result := &Accelerator{
		Key:       key,
		Modifiers: []Modifier{modifier1, modifier2},
	}
	result.Modifiers = append(result.Modifiers, rest...)
	return result
}
