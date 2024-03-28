package signal

import (
	"os"
	gosignal "os/signal"
	"sync"
	"syscall"
)

var signalChannel = make(chan os.Signal, 2)

var (
	callbacks []func()
	lock      sync.Mutex
)


// ff:
// callback:
func OnShutdown(callback func()) {
	lock.Lock()
	defer lock.Unlock()
	callbacks = append(callbacks, callback)
}

// 启动信号管理器

// ff:
func Start() {
	// Hook into interrupts
	gosignal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

// 启动信号监听器，并等待取消信号或外部信号的到来
	go func() {
		<-signalChannel
		println("")
		println("Ctrl+C detected. Shutting down...")
		for _, callback := range callbacks {
			callback()
		}
	}()
}
