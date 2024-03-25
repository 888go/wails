package keys

import (
	"strconv"
	"testing"
)

func TestStringify(t *testing.T) {

	const Windows = "windows"
	const Mac = "darwin"
	const Linux = "linux"
	tests := []struct {
		arg      *Accelerator
		want     string
		platform string
	}{
		// Single Keys
		{Key("a"), "A", Windows},
		{Key(""), "", Windows},
		{Key("?"), "?", Windows},
		{Key("a"), "A", Mac},
		{Key(""), "", Mac},
		{Key("?"), "?", Mac},
		{Key("a"), "A", Linux},
		{Key(""), "", Linux},
		{Key("?"), "?", Linux},

		// Single modifier
		{Control("a"), "Ctrl+A", Windows},
		{Control("a"), "Ctrl+A", Mac},
		{Control("a"), "Ctrl+A", Linux},
		{CmdOrCtrl("a"), "Ctrl+A", Windows},
		{CmdOrCtrl("a"), "Cmd+A", Mac},
		{CmdOrCtrl("a"), "Ctrl+A", Linux},
		{Shift("a"), "Shift+A", Windows},
		{Shift("a"), "Shift+A", Mac},
		{Shift("a"), "Shift+A", Linux},
		{OptionOrAlt("a"), "Alt+A", Windows},
		{OptionOrAlt("a"), "Option+A", Mac},
		{OptionOrAlt("a"), "Alt+A", Linux},
//{Super("a"), "Win+A", Windows},
// 在Windows系统下，组合键为"Win+A"，其中Super表示Windows键
//{Super("a"), "Cmd+A", Mac},
// 在Mac系统下，组合键为"Cmd+A"，其中Super在此处表示Command键
//{Super("a"), "Super+A", Linux},
// 在Linux系统下，组合键为"Super+A"，其中Super表示Linux系统的Super键（通常表现为窗口管理器定义的“超级键”，如Meta键或者Windows键）

		// 双重组合无重复
		{Combo("a", ControlKey, OptionOrAltKey), "Ctrl+Alt+A", Windows},
		{Combo("a", ControlKey, OptionOrAltKey), "Ctrl+Option+A", Mac},
		{Combo("a", ControlKey, OptionOrAltKey), "Ctrl+Alt+A", Linux},
		{Combo("a", CmdOrCtrlKey, OptionOrAltKey), "Ctrl+Alt+A", Windows},
		{Combo("a", CmdOrCtrlKey, OptionOrAltKey), "Cmd+Option+A", Mac},
		{Combo("a", CmdOrCtrlKey, OptionOrAltKey), "Ctrl+Alt+A", Linux},
		{Combo("a", ShiftKey, OptionOrAltKey), "Shift+Alt+A", Windows},
		{Combo("a", ShiftKey, OptionOrAltKey), "Shift+Option+A", Mac},
		{Combo("a", ShiftKey, OptionOrAltKey), "Shift+Alt+A", Linux},
// 在Windows系统下，组合键为"Win+Alt+A"，对应的按键组合是"a"、SuperKey（通常指Windows键）和OptionOrAltKey（即Alt键）
//{Combo("a", SuperKey, OptionOrAltKey), "Win+Alt+A", Windows},
// 在Mac系统下，组合键为"Cmd+Option+A"，对应的按键组合同样是"a"，但SuperKey此时代表Command键，OptionOrAltKey仍表示Option键（在Mac键盘上标为?）
//{Combo("a", SuperKey, OptionOrAltKey), "Cmd+Option+A", Mac},
// 在Linux系统下，组合键为"Super+Alt+A"，对应的按键组合依然是"a"，同时SuperKey在这里指的是Linux系统的Super键（有时也称作Win键），OptionOrAltKey依旧表示Alt键
//{Combo("a", SuperKey, OptionOrAltKey), "Super+Alt+A", Linux},

		// Combo duplicate
		{Combo("a", OptionOrAltKey, OptionOrAltKey), "Alt+A", Windows},
		{Combo("a", OptionOrAltKey, OptionOrAltKey), "Option+A", Mac},
		{Combo("a", OptionOrAltKey, OptionOrAltKey), "Alt+A", Linux},
//{组合键("a", OptionOrAlt键, Super键, OptionOrAlt键), "Alt+Win+A", Windows},
//{组合键("a", OptionOrAlt键, Super键, OptionOrAlt键), "Option+Cmd+A", Mac},
//{组合键("a", OptionOrAlt键, Super键, OptionOrAlt键), "Alt+Super+A", Linux},
// 翻译成中文：
// ```go
//{定义组合键("a", Option或Alt键, Super键, Option或Alt键), 对应快捷键为"Alt+Win+A", 适用于Windows系统},
//{定义组合键("a", Option或Alt键, Super键, Option或Alt键), 对应快捷键为"Option+Cmd+A", 适用于Mac系统},
//{定义组合键("a", Option或Alt键, Super键, Option或Alt键), 对应快捷键为"Alt+Super+A", 适用于Linux系统},
// 这段代码是在根据不同操作系统（Windows、Mac、Linux）定义键盘的组合键及其对应的快捷键表示。例如，在Windows系统中，同时按下Alt键、Windows键和字母A的组合快捷键可被表示为"Alt+Win+A"。
	}
	for index, tt := range tests {
		t.Run(strconv.Itoa(index), func(t *testing.T) {
			if got := Stringify(tt.arg, tt.platform); got != tt.want {
				t.Errorf("Stringify() = %v, want %v", got, tt.want)
			}
		})
	}
}
