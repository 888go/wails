package runtime

import (
	"context"
)

// EventsOn 注册一个给定事件名称的监听器。它返回一个函数，用于取消该监听器

// ff:绑定事件
// callback:回调函数
// optionalData:可选数据
// eventName:事件名称
// ctx:上下文
func X绑定事件(上下文 context.Context, 事件名称 string, 回调函数 func(可选数据 ...interface{})) func() {
	events := getEvents(上下文)
	return events.On(事件名称, 回调函数)
}

// EventsOff 注销给定事件名称的监听器，可选地，可以通过 `additionalEventNames` 注销多个监听器

// ff:移除事件
// additionalEventNames:移除事件名称
// eventName:事件名称
// ctx:上下文
func X移除事件(上下文 context.Context, 事件名称 string, 移除事件名称 ...string) {
	events := getEvents(上下文)
	events.Off(事件名称)

	if len(移除事件名称) > 0 {
		for _, eventName := range 移除事件名称 {
			events.Off(eventName)
		}
	}
}

// EventsOff 注销给定事件名称的监听器，可选地，可以通过 `additionalEventNames` 注销多个监听器

// ff:移除所有事件
// ctx:上下文
func X移除所有事件(上下文 context.Context) {
	events := getEvents(上下文)
	events.OffAll()
}

// EventsOnce 为给定的事件名称注册一个监听器。在第一次回调之后，该监听器将被删除。它返回一个函数用于取消监听器

// ff:绑定单次事件
// callback:回调函数
// optionalData:可选数据
// eventName:事件名称
// ctx:上下文
func X绑定单次事件(上下文 context.Context, 事件名称 string, 回调函数 func(可选数据 ...interface{})) func() {
	events := getEvents(上下文)
	return events.Once(事件名称, 回调函数)
}

// EventsOnMultiple 注册一个给定事件名称的监听器，该监听器最多可以被调用 'counter' 次。它返回一个函数用于取消监听器
// ```go
// 注释翻译：
// EventsOnMultiple 函数用于注册针对指定事件名称的监听器，这个监听器最大可被触发 'counter' 次。
// 此函数返回一个用于撤销该监听器功能的函数

// ff:绑定N次事件
// counter:次数
// callback:回调函数
// optionalData:可选数据
// eventName:事件名称
// ctx:上下文
func X绑定N次事件(上下文 context.Context, 事件名称 string, 回调函数 func(可选数据 ...interface{}), 次数 int) func() {
	events := getEvents(上下文)
	return events.OnMultiple(事件名称, 回调函数, 次数)
}

// EventsEmit 传递通过

// ff:触发指定事件
// optionalData:可选数据
// eventName:事件名称
// ctx:上下文
func X触发指定事件(上下文 context.Context, 事件名称 string, 可选数据 ...interface{}) {
	events := getEvents(上下文)
	events.Emit(事件名称, 可选数据...)
}
