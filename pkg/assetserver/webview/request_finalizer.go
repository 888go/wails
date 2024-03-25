package webview

import (
	"runtime"
	"sync/atomic"
)

var _ Request = &requestFinalizer{}

type requestFinalizer struct {
	Request
	closed int32
}

// newRequestFinalizer 返回一个带有运行时终结器的请求，确保即使尚未被显式关闭，也能在终结器中进行关闭。
// 同时，它还确保包装请求的 Close() 方法仅被调用一次。
func newRequestFinalizer(r Request) Request {
	rf := &requestFinalizer{Request: r}
	// 确保异步释放，因为它可能会阻塞终结器 goroutine 较长时间
	runtime.SetFinalizer(rf, func(obj *requestFinalizer) { rf.close(true) })
	return rf
}

func (r *requestFinalizer) X关闭() error {
	return r.close(false)
}

func (r *requestFinalizer) close(asyncRelease bool) error {
	if atomic.CompareAndSwapInt32(&r.closed, 0, 1) {
		runtime.SetFinalizer(r, nil)
		if asyncRelease {
			go r.Request.X关闭()
			return nil
		} else {
			return r.Request.X关闭()
		}
	}
	return nil
}
