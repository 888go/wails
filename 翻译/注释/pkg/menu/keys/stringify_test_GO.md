
<原文开始>
		//{Super("a"), "Win+A", Windows},
		//{Super("a"), "Cmd+A", Mac},
		//{Super("a"), "Super+A", Linux},
<原文结束>

# <翻译开始>
		//{Super("a"), "Win+A", Windows},
		// 在Windows系统下，组合键为"Win+A"，其中Super表示Windows键
		//{Super("a"), "Cmd+A", Mac},
		// 在Mac系统下，组合键为"Cmd+A"，其中Super在此处表示Command键
		//{Super("a"), "Super+A", Linux},
		// 在Linux系统下，组合键为"Super+A"，其中Super表示Linux系统的Super键（通常表现为窗口管理器定义的“超级键”，如Meta键或者Windows键）
# <翻译结束>


<原文开始>
// Dual Combo non duplicate
<原文结束>

# <翻译开始>
// 双重组合无重复
# <翻译结束>


<原文开始>
		//{Combo("a", SuperKey, OptionOrAltKey), "Win+Alt+A", Windows},
		//{Combo("a", SuperKey, OptionOrAltKey), "Cmd+Option+A", Mac},
		//{Combo("a", SuperKey, OptionOrAltKey), "Super+Alt+A", Linux},
<原文结束>

# <翻译开始>
		// 在Windows系统下，组合键为"Win+Alt+A"，对应的按键组合是"a"、SuperKey（通常指Windows键）和OptionOrAltKey（即Alt键）
		//{Combo("a", SuperKey, OptionOrAltKey), "Win+Alt+A", Windows},
		// 在Mac系统下，组合键为"Cmd+Option+A"，对应的按键组合同样是"a"，但SuperKey此时代表Command键，OptionOrAltKey仍表示Option键（在Mac键盘上标为?）
		//{Combo("a", SuperKey, OptionOrAltKey), "Cmd+Option+A", Mac},
		// 在Linux系统下，组合键为"Super+Alt+A"，对应的按键组合依然是"a"，同时SuperKey在这里指的是Linux系统的Super键（有时也称作Win键），OptionOrAltKey依旧表示Alt键
		//{Combo("a", SuperKey, OptionOrAltKey), "Super+Alt+A", Linux},
# <翻译结束>


<原文开始>
		//{Combo("a", OptionOrAltKey, SuperKey, OptionOrAltKey), "Alt+Win+A", Windows},
		//{Combo("a", OptionOrAltKey, SuperKey, OptionOrAltKey), "Option+Cmd+A", Mac},
		//{Combo("a", OptionOrAltKey, SuperKey, OptionOrAltKey), "Alt+Super+A", Linux},
<原文结束>

# <翻译开始>
		//{组合键("a", OptionOrAlt键, Super键, OptionOrAlt键), "Alt+Win+A", Windows},
		//{组合键("a", OptionOrAlt键, Super键, OptionOrAlt键), "Option+Cmd+A", Mac},
		//{组合键("a", OptionOrAlt键, Super键, OptionOrAlt键), "Alt+Super+A", Linux},
		// 翻译成中文：
		// ```go
		//{定义组合键("a", Option或Alt键, Super键, Option或Alt键), 对应快捷键为"Alt+Win+A", 适用于Windows系统},
		//{定义组合键("a", Option或Alt键, Super键, Option或Alt键), 对应快捷键为"Option+Cmd+A", 适用于Mac系统},
		//{定义组合键("a", Option或Alt键, Super键, Option或Alt键), 对应快捷键为"Alt+Super+A", 适用于Linux系统},
		// 这段代码是在根据不同操作系统（Windows、Mac、Linux）定义键盘的组合键及其对应的快捷键表示。例如，在Windows系统中，同时按下Alt键、Windows键和字母A的组合快捷键可被表示为"Alt+Win+A"。
# <翻译结束>

