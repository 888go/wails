
<原文开始>
// eventListener holds a callback function which is invoked when
// the event listened for is emitted. It has a counter which indicates
// how the total number of events it is interested in. A value of zero
// means it does not expire (default).
<原文结束>

# <翻译开始>
// eventListener 保存一个回调函数，当监听到相应的事件时会调用该函数。它有一个计数器，表示其感兴趣的总事件数量。值为零意味着它不会过期（默认情况）。
# <翻译结束>


<原文开始>
// Function to call with emitted event data
<原文结束>

# <翻译开始>
// 此函数用于在接收到发出的事件数据时调用
# <翻译结束>


<原文开始>
// The number of times this callback may be called. -1 = infinite
<原文结束>

# <翻译开始>
// 此回调函数可能被调用的次数。-1 = 无限次
# <翻译结束>


<原文开始>
// Flag to indicate that this listener should be deleted
<原文结束>

# <翻译开始>
// 标志，表示这个监听器应该被删除
# <翻译结束>


<原文开始>
// Events handles eventing
<原文结束>

# <翻译开始>
// Events 处理事件驱动功能
# <翻译结束>


<原文开始>
// NewEvents creates a new log subsystem
<原文结束>

# <翻译开始>
// NewEvents 创建一个新的日志子系统
# <翻译结束>


<原文开始>
// registerListener provides a means of subscribing to events of type "eventName"
<原文结束>

# <翻译开始>
// registerListener 提供了一种订阅 "eventName" 类型事件的方法
# <翻译结束>


<原文开始>
// Create new eventListener
<原文结束>

# <翻译开始>
// 创建新的事件监听器
# <翻译结束>


<原文开始>
// Append the new listener to the listeners slice
<原文结束>

# <翻译开始>
// 将新的监听器追加到listeners切片中
# <翻译结束>


<原文开始>
// unRegisterListener provides a means of unsubscribing to events of type "eventName"
<原文结束>

# <翻译开始>
// unRegisterListener 提供了一种取消订阅“eventName”类型事件的方法
# <翻译结束>


<原文开始>
// Notify backend for the given event name
<原文结束>

# <翻译开始>
// 为给定事件名称通知后端
# <翻译结束>


<原文开始>
// Get list of event listeners
<原文结束>

# <翻译开始>
// 获取事件监听器列表
# <翻译结束>


<原文开始>
// We have a dirty flag to indicate that there are items to delete
<原文结束>

# <翻译开始>
// 我们有一个脏标志（dirty flag）用于指示有待删除的项目
# <翻译结束>


<原文开始>
// Do we have items to delete?
<原文结束>

# <翻译开始>
// 我们是否有待删除的项目？
# <翻译结束>


<原文开始>
// Create a new Listeners slice
<原文结束>

# <翻译开始>
// 创建一个新的Listeners切片
# <翻译结束>


<原文开始>
// Iterate over current listeners
<原文结束>

# <翻译开始>
// 遍历当前监听器
# <翻译结束>


<原文开始>
// If we aren't deleting the listener, add it to the new list
<原文结束>

# <翻译开始>
// 如果我们没有删除监听器，则将其添加到新列表中
# <翻译结束>


<原文开始>
// Save new listeners or remove entry
<原文结束>

# <翻译开始>
// 保存新的监听器或移除条目
# <翻译结束>

