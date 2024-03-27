package runtime

import (
	"sync"

	"github.com/samber/lo"
	"github.com/888go/wails/internal/frontend"
)

type Logger interface {
	Trace(format string, v ...interface{})
}

// eventListener 保存一个回调函数，当监听到相应的事件时会调用该函数。它有一个计数器，表示其感兴趣的总事件数量。值为零意味着它不会过期（默认情况）。
type eventListener struct {
	callback func(...interface{}) // 此函数用于在接收到发出的事件数据时调用
	counter  int                  // 此回调函数可能被调用的次数。-1 = 无限次
	delete   bool                 // 标志，表示这个监听器应该被删除
}

// Events 处理事件驱动功能
type Events struct {
	log      Logger
	frontend []frontend.Frontend

	// Go event listeners
	listeners  map[string][]*eventListener
	notifyLock sync.RWMutex
}


// ff:
// data:
// name:
// sender:
func (e *Events) Notify(sender frontend.Frontend, name string, data ...interface{}) {
	e.notifyBackend(name, data...)
	for _, thisFrontend := range e.frontend {
		if thisFrontend == sender {
			continue
		}
		thisFrontend.Notify(name, data...)
	}
}


// ff:
// callback:
// eventName:
func (e *Events) On(eventName string, callback func(...interface{})) func() {
	return e.registerListener(eventName, callback, -1)
}


// ff:
// counter:
// callback:
// eventName:
func (e *Events) OnMultiple(eventName string, callback func(...interface{}), counter int) func() {
	return e.registerListener(eventName, callback, counter)
}


// ff:
// callback:
// eventName:
func (e *Events) Once(eventName string, callback func(...interface{})) func() {
	return e.registerListener(eventName, callback, 1)
}


// ff:
// data:
// eventName:
func (e *Events) Emit(eventName string, data ...interface{}) {
	e.notifyBackend(eventName, data...)
	for _, thisFrontend := range e.frontend {
		thisFrontend.Notify(eventName, data...)
	}
}


// ff:
// eventName:
func (e *Events) Off(eventName string) {
	e.unRegisterListener(eventName)
}


// ff:
func (e *Events) OffAll() {
	e.notifyLock.Lock()
	for eventName := range e.listeners {
		delete(e.listeners, eventName)
	}
	e.notifyLock.Unlock()
}

// NewEvents 创建一个新的日志子系统

// ff:
// log:
func NewEvents(log Logger) *Events {
	result := &Events{
		log:       log,
		listeners: make(map[string][]*eventListener),
	}
	return result
}

// registerListener 提供了一种订阅 "eventName" 类型事件的方法
func (e *Events) registerListener(eventName string, callback func(...interface{}), counter int) func() {
	// 创建新的事件监听器
	thisListener := &eventListener{
		callback: callback,
		counter:  counter,
		delete:   false,
	}
	e.notifyLock.Lock()
	// 将新的监听器追加到listeners切片中
	e.listeners[eventName] = append(e.listeners[eventName], thisListener)
	e.notifyLock.Unlock()
	return func() {
		e.notifyLock.Lock()
		defer e.notifyLock.Unlock()

		if _, ok := e.listeners[eventName]; !ok {
			return
		}
		e.listeners[eventName] = lo.Filter(e.listeners[eventName], func(l *eventListener, i int) bool {
			return l != thisListener
		})
	}
}

// unRegisterListener 提供了一种取消订阅“eventName”类型事件的方法
func (e *Events) unRegisterListener(eventName string) {
	e.notifyLock.Lock()
	// Clear the listeners
	delete(e.listeners, eventName)
	e.notifyLock.Unlock()
}

// 为给定事件名称通知后端
func (e *Events) notifyBackend(eventName string, data ...interface{}) {
	e.notifyLock.Lock()
	defer e.notifyLock.Unlock()

	// 获取事件监听器列表
	listeners := e.listeners[eventName]
	if listeners == nil {
		e.log.Trace("No listeners for event '%s'", eventName)
		return
	}

	// 我们有一个脏标志（dirty flag）用于指示有待删除的项目
	itemsToDelete := false

	// Callback in goroutine
	for _, listener := range listeners {
		if listener.counter > 0 {
			listener.counter--
		}
		go listener.callback(data...)

		if listener.counter == 0 {
			listener.delete = true
			itemsToDelete = true
		}
	}

	// 我们是否有待删除的项目？
	if itemsToDelete {

		// 创建一个新的Listeners切片
		var newListeners []*eventListener

		// 遍历当前监听器
		for _, listener := range listeners {
			// 如果我们没有删除监听器，则将其添加到新列表中
			if !listener.delete {
				newListeners = append(newListeners, listener)
			}
		}

		// 保存新的监听器或移除条目
		if len(newListeners) > 0 {
			e.listeners[eventName] = newListeners
		} else {
			delete(e.listeners, eventName)
		}
	}
}


// ff:
// appFrontend:
func (e *Events) AddFrontend(appFrontend frontend.Frontend) {
	e.frontend = append(e.frontend, appFrontend)
}
