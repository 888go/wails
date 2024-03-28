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
func EventsOn(ctx context.Context, eventName string, callback func(optionalData ...interface{})) func() {
	events := getEvents(ctx)
	return events.On(eventName, callback)
}

// EventsOff 注销给定事件名称的监听器，可选地，可以通过 `additionalEventNames` 注销多个监听器

// ff:移除事件
// additionalEventNames:移除事件名称
// eventName:事件名称
// ctx:上下文
func EventsOff(ctx context.Context, eventName string, additionalEventNames ...string) {
	events := getEvents(ctx)
	events.Off(eventName)

	if len(additionalEventNames) > 0 {
		for _, eventName := range additionalEventNames {
			events.Off(eventName)
		}
	}
}

// EventsOff 注销给定事件名称的监听器，可选地，可以通过 `additionalEventNames` 注销多个监听器

// ff:移除所有事件
// ctx:上下文
func EventsOffAll(ctx context.Context) {
	events := getEvents(ctx)
	events.OffAll()
}

// EventsOnce 为给定的事件名称注册一个监听器。在第一次回调之后，该监听器将被删除。它返回一个函数用于取消监听器

// ff:绑定单次事件
// callback:回调函数
// optionalData:可选数据
// eventName:事件名称
// ctx:上下文
func EventsOnce(ctx context.Context, eventName string, callback func(optionalData ...interface{})) func() {
	events := getEvents(ctx)
	return events.Once(eventName, callback)
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
func EventsOnMultiple(ctx context.Context, eventName string, callback func(optionalData ...interface{}), counter int) func() {
	events := getEvents(ctx)
	return events.OnMultiple(eventName, callback, counter)
}

// EventsEmit 传递通过

// ff:触发指定事件
// optionalData:可选数据
// eventName:事件名称
// ctx:上下文
func EventsEmit(ctx context.Context, eventName string, optionalData ...interface{}) {
	events := getEvents(ctx)
	events.Emit(eventName, optionalData...)
}
