package keys

import (
	"fmt"
	"strings"
)

// Modifier 实际上是一个字符串
type Modifier string

const (
	// CmdOrCtrlKey 表示在 Mac 平台上代表 Command 键，在其他平台上代表 Control 键
	X常量_组合键_Cmd或Ctrl键 Modifier = "cmdorctrl" //hs:常量_组合键_Cmd或Ctrl键     
	// OptionOrAltKey 表示在 Mac 平台上代表 Option 键，在其他平台上代表 Alt 键
	X常量_组合键_Option或Alt键 Modifier = "optionoralt" //hs:常量_组合键_Option或Alt键     
	// ShiftKey 表示在所有系统上的 shift 键
	X常量_组合键_Shift键 Modifier = "shift" //hs:常量_组合键_Shift键     
// SuperKey 表示在 Mac 上的 Command 键，在其他平台（如 Windows）上表示 Windows 键
// SuperKey Modifier = "super"
// ControlKey 代表在所有系统上的控制键
	X常量_组合键_Ctrl键 Modifier = "ctrl" //hs:常量_组合键_Ctrl键     
)

var modifierMap = map[string]Modifier{
	"cmdorctrl":   X常量_组合键_Cmd或Ctrl键,
	"optionoralt": X常量_组合键_Option或Alt键,
	"shift":       X常量_组合键_Shift键,
	// "super":       超级键，
	"ctrl": X常量_组合键_Ctrl键,
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
	X名称       string //hs:名称     
	X修饰符 []Modifier //hs:修饰符     
}

// Key 创建一个标准的键Accelerator

// ff:按键
// key:按键字符
func X按键(按键字符 string) *Accelerator {
	return &Accelerator{
		X名称: strings.ToLower(按键字符),
	}
}

// CmdOrCtrl 创建一个 'CmdOrCtrl' 快捷键

// ff:组合按键Cmd或Ctrl
// key:按键字符
func X组合按键Cmd或Ctrl(按键字符 string) *Accelerator {
	return &Accelerator{
		X名称:       strings.ToLower(按键字符),
		X修饰符: []Modifier{X常量_组合键_Cmd或Ctrl键},
	}
}

// OptionOrAlt 创建一个 'OptionOrAlt' 加速器

// ff:组合按键Option或Alt键
// key:按键字符
func X组合按键Option或Alt键(按键字符 string) *Accelerator {
	return &Accelerator{
		X名称:       strings.ToLower(按键字符),
		X修饰符: []Modifier{X常量_组合键_Option或Alt键},
	}
}

// Shift 创建一个“Shift”加速器

// ff:组合按键Shift
// key:按键字符
func X组合按键Shift(按键字符 string) *Accelerator {
	return &Accelerator{
		X名称:       strings.ToLower(按键字符),
		X修饰符: []Modifier{X常量_组合键_Shift键},
	}
}

// Control 创建一个名为'Control'的加速器

// ff:组合按键Ctrl键
// key:按键字符
func X组合按键Ctrl键(按键字符 string) *Accelerator {
	return &Accelerator{
		X名称:       strings.ToLower(按键字符),
		X修饰符: []Modifier{X常量_组合键_Ctrl键},
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

// ff:组合按键
// rest:组合键s
// modifier2:组合键2
// modifier1:组合键1
// key:按键字符
func X组合按键(按键字符 string, 组合键1 Modifier, 组合键2 Modifier, 组合键s ...Modifier) *Accelerator {
	result := &Accelerator{
		X名称:       按键字符,
		X修饰符: []Modifier{组合键1, 组合键2},
	}
	result.X修饰符 = append(result.X修饰符, 组合键s...)
	return result
}
