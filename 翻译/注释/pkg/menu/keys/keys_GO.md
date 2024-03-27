
<原文开始>
// Modifier is actually a string
<原文结束>

# <翻译开始>
// Modifier 实际上是一个字符串
# <翻译结束>


<原文开始>
// CmdOrCtrlKey represents Command on Mac and Control on other platforms
<原文结束>

# <翻译开始>
// CmdOrCtrlKey 表示在 Mac 平台上代表 Command 键，在其他平台上代表 Control 键
# <翻译结束>


<原文开始>
// OptionOrAltKey represents Option on Mac and Alt on other platforms
<原文结束>

# <翻译开始>
// OptionOrAltKey 表示在 Mac 平台上代表 Option 键，在其他平台上代表 Alt 键
# <翻译结束>


<原文开始>
// ShiftKey represents the shift key on all systems
<原文结束>

# <翻译开始>
// ShiftKey 表示在所有系统上的 shift 键
# <翻译结束>


<原文开始>
	// SuperKey represents Command on Mac and the Windows key on the other platforms
	// SuperKey Modifier = "super"
	// ControlKey represents the control key on all systems
<原文结束>

# <翻译开始>
// SuperKey 表示在 Mac 上的 Command 键，在其他平台（如 Windows）上表示 Windows 键
// SuperKey Modifier = "super"
// ControlKey 代表在所有系统上的控制键
# <翻译结束>


<原文开始>
//"super":       SuperKey,
<原文结束>

# <翻译开始>
// "super":       超级键，
# <翻译结束>


<原文开始>
// Accelerator holds the keyboard shortcut for a menu item
<原文结束>

# <翻译开始>
// Accelerator 保存了菜单项的键盘快捷键
# <翻译结束>


<原文开始>
// Key creates a standard key Accelerator
<原文结束>

# <翻译开始>
// Key 创建一个标准的键Accelerator
# <翻译结束>


<原文开始>
// CmdOrCtrl creates a 'CmdOrCtrl' Accelerator
<原文结束>

# <翻译开始>
// CmdOrCtrl 创建一个 'CmdOrCtrl' 快捷键
# <翻译结束>


<原文开始>
// OptionOrAlt creates a 'OptionOrAlt' Accelerator
<原文结束>

# <翻译开始>
// OptionOrAlt 创建一个 'OptionOrAlt' 加速器
# <翻译结束>


<原文开始>
// Shift creates a 'Shift' Accelerator
<原文结束>

# <翻译开始>
// Shift 创建一个“Shift”加速器
# <翻译结束>


<原文开始>
// Control creates a 'Control' Accelerator
<原文结束>

# <翻译开始>
// Control 创建一个名为'Control'的加速器
# <翻译结束>


<原文开始>
//
//// Super creates a 'Super' Accelerator
//func Super(key string) *Accelerator {
//	return &Accelerator{
//		Key:       strings.ToLower(key),
//		Modifiers: []Modifier{SuperKey},
//	}
//}
<原文结束>

# <翻译开始>
// 
//// Super 函数用于创建一个 'Super' 加速器
//func Super(key string) *Accelerator {
//	// 将输入的 key 转换为小写并初始化 Accelerator 结构体实例
//	return &Accelerator{
//		Key:       strings.ToLower(key), // 设置 Key 字段为小写形式的 key
//		Modifiers: []Modifier{SuperKey}, // 设置 Modifiers 字段，包含 SuperKey 模块
//	}
//}
# <翻译结束>


<原文开始>
// Combo creates an Accelerator with multiple Modifiers
<原文结束>

# <翻译开始>
// Combo 创建一个带有多个修饰符的 Accelerator
# <翻译结束>

