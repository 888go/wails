package runtime

import (
	"context"
	"log"
	goruntime "runtime"

	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/internal/logger"
)

const contextError = `An invalid context was passed. This method requires the specific context given in the lifecycle hooks:
https://wails.io/docs/reference/runtime/intro`

func getFrontend(ctx context.Context) frontend.Frontend {
	if ctx == nil {
		pc, _, _, _ := goruntime.Caller(1)
		funcName := goruntime.FuncForPC(pc).Name()
		log.Fatalf("cannot call '%s': %s", funcName, contextError)
	}
	result := ctx.Value("frontend")
	if result != nil {
		return result.(frontend.Frontend)
	}
	pc, _, _, _ := goruntime.Caller(1)
	funcName := goruntime.FuncForPC(pc).Name()
	log.Fatalf("cannot call '%s': %s", funcName, contextError)
	return nil
}

func getLogger(ctx context.Context) *logger.Logger {
	if ctx == nil {
		pc, _, _, _ := goruntime.Caller(1)
		funcName := goruntime.FuncForPC(pc).Name()
		log.Fatalf("cannot call '%s': %s", funcName, contextError)
	}
	result := ctx.Value("logger")
	if result != nil {
		return result.(*logger.Logger)
	}
	pc, _, _, _ := goruntime.Caller(1)
	funcName := goruntime.FuncForPC(pc).Name()
	log.Fatalf("cannot call '%s': %s", funcName, contextError)
	return nil
}

func getEvents(ctx context.Context) frontend.Events {
	if ctx == nil {
		pc, _, _, _ := goruntime.Caller(1)
		funcName := goruntime.FuncForPC(pc).Name()
		log.Fatalf("cannot call '%s': %s", funcName, contextError)
	}
	result := ctx.Value("events")
	if result != nil {
		return result.(frontend.Events)
	}
	pc, _, _, _ := goruntime.Caller(1)
	funcName := goruntime.FuncForPC(pc).Name()
	log.Fatalf("cannot call '%s': %s", funcName, contextError)
	return nil
}

// Quit the application

// ff:退出
// ctx:上下文
func X退出(上下文 context.Context) {
	if 上下文 == nil {
		log.Fatalf("Error calling 'runtime.Quit': %s", contextError)
	}
	appFrontend := getFrontend(上下文)
	appFrontend.Quit()
}

// Hide the application

// ff:隐藏
// ctx:上下文
func X隐藏(上下文 context.Context) {
	if 上下文 == nil {
		log.Fatalf("Error calling 'runtime.Hide': %s", contextError)
	}
	appFrontend := getFrontend(上下文)
	appFrontend.Hide()
}

// 如果应用程序是隐藏的，则显示它

// ff:显示
// ctx:上下文
func X显示(上下文 context.Context) {
	if 上下文 == nil {
		log.Fatalf("Error calling 'runtime.Show': %s", contextError)
	}
	appFrontend := getFrontend(上下文)
	appFrontend.Show()
}

// EnvironmentInfo 包含有关环境的信息
type EnvironmentInfo struct {
	X构建类型 string `json:"buildType"` //hs:构建类型     
	X平台  string `json:"platform"` //hs:平台     
	X架构      string `json:"arch"` //hs:架构     
}

// Environment 返回关于环境的信息

// ff:取环境信息
// ctx:上下文
func X取环境信息(上下文 context.Context) EnvironmentInfo {
	var result EnvironmentInfo
	buildType := 上下文.Value("buildtype")
	if buildType != nil {
		result.X构建类型 = buildType.(string)
	}
	result.X平台 = goruntime.GOOS
	result.X架构 = goruntime.GOARCH
	return result
}
