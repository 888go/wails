
<原文开始>
// EventsOn registers a listener for the given event name. It returns a function to cancel the listener
<原文结束>

# <翻译开始>
// EventsOn 注册一个给定事件名称的监听器。它返回一个函数，用于取消该监听器
# <翻译结束>


<原文开始>
// EventsOff unregisters a listener for the given event name, optionally multiple listeners can be unregistered via `additionalEventNames`
<原文结束>

# <翻译开始>
// EventsOff 注销给定事件名称的监听器，可选地，可以通过 `additionalEventNames` 注销多个监听器
# <翻译结束>


<原文开始>
// EventsOnce registers a listener for the given event name. After the first callback, the
// listener is deleted. It returns a function to cancel the listener
<原文结束>

# <翻译开始>
// EventsOnce 为给定的事件名称注册一个监听器。在第一次回调之后，该监听器将被删除。它返回一个函数用于取消监听器
# <翻译结束>


<原文开始>
// EventsOnMultiple registers a listener for the given event name, that may be called a maximum of 'counter' times. It returns a function
// to cancel the listener
<原文结束>

# <翻译开始>
// EventsOnMultiple 注册一个给定事件名称的监听器，该监听器最多可以被调用 'counter' 次。它返回一个函数用于取消监听器
// ```go
// 注释翻译：
// EventsOnMultiple 函数用于注册针对指定事件名称的监听器，这个监听器最大可被触发 'counter' 次。
// 此函数返回一个用于撤销该监听器功能的函数
# <翻译结束>


<原文开始>
// EventsEmit pass through
<原文结束>

# <翻译开始>
// EventsEmit 传递通过
# <翻译结束>

