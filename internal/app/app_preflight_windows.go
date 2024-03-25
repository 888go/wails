//go:build windows && !bindings

package app

import (
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/internal/wv2installer"
	"github.com/888go/wails/pkg/options"
)


// ff:
// logger:
// options:
func PreflightChecks(options *options.App, logger *logger.Logger) error {

	_ = options

	// 处理 webview2 运行时的情况。我们可以通过在 `wails build` 命令中的 `webview2` 标志传递一种策略。
	// 这将决定 wv2runtime.Process 在缺乏有效运行时的情况下如何处理。
	installedVersion, err := wv2installer.Process(options)
	if installedVersion != "" {
		logger.Debug("已安装WebView2运行时版本'%s' 。 最低版本要求:%s.",
			installedVersion, wv2installer.MinimumRuntimeVersion)
	}
	if err != nil {
		return err
	}

	return nil
}
