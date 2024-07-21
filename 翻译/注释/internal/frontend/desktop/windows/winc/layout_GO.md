
<原文开始>
// Dockable component must satisfy interface to be docked.
<原文结束>

# <翻译开始>
// 可停靠组件必须满足接口要求才能被停靠。
# <翻译结束>


<原文开始>
// DockAllow is window, panel or other component that satisfies interface.
<原文结束>

# <翻译开始>
// DockAllow 是窗口、面板或其他满足接口的组件。
# <翻译结束>


<原文开始>
// Various layout managers
<原文结束>

# <翻译开始>
// 各种布局管理器
# <翻译结束>


<原文开始>
// DockState gets saved and loaded from json
<原文结束>

# <翻译开始>
// DockState 会从 json 中保存和加载
# <翻译结束>


<原文开始>
// Layout management for the child controls.
<原文结束>

# <翻译开始>
// 子控件的布局管理。
# <翻译结束>


<原文开始>
// SaveState of the layout. Only works for Docks with parent set to main form.
<原文结束>

# <翻译开始>
// 保存布局的状态。仅适用于父级设置为主窗口的停靠栏。
# <翻译结束>


<原文开始>
// LoadState of the layout. Only works for Docks with parent set to main form.
<原文结束>

# <翻译开始>
// 加载布局状态。仅适用于父级设置为主窗体的停靠栏。
# <翻译结束>


<原文开始>
	// if number of controls in the saved layout does not match
	// current number on screen - something changed and we do not reload
	// rest of control sizes from json
<原文结束>

# <翻译开始>
	// 如果保存的布局中的控件数量与当前屏幕上的控件数量不匹配
	// 意味着某些内容发生了变化，此时我们不重新加载
	// 来自json的其余控件大小信息
# <翻译结束>


<原文开始>
// SaveStateFile convenience function.
<原文结束>

# <翻译开始>
// SaveStateFile 便捷函数。
# <翻译结束>


<原文开始>
// LoadStateFile loads state ignores error if file is not found.
<原文结束>

# <翻译开始>
// LoadStateFile 加载状态文件，如果文件未找到，则忽略错误。
# <翻译结束>


<原文开始>
// if file is not found or not accessible ignore it
<原文结束>

# <翻译开始>
// 如果文件未找到或无法访问，则忽略它
# <翻译结束>


<原文开始>
// Update is called to resize child items based on layout directions.
<原文结束>

# <翻译开始>
// Update 在需要根据布局方向调整子项大小时被调用。
# <翻译结束>


<原文开始>
// Non visible controls do not preserve space.
<原文结束>

# <翻译开始>
// 非可见控件不保留空间。
# <翻译结束>


<原文开始>
//c.child.Invalidate(true)
<原文结束>

# <翻译开始>
// c.child.Invalidate(true) // 清除c.child的缓存，参数true表示递归清除子节点的缓存
# <翻译结束>

