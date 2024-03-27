//go:build windows
// +build windows

package desktop

import (
	"context"
	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend"
	"github.com/888go/wails/internal/frontend/desktop/windows"
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/pkg/options"
)


// ff:
// dispatcher:
// appBindings:
// logger:
// appoptions:
// ctx:
func NewFrontend(ctx context.Context, appoptions *options.App, logger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) frontend.Frontend {
	return windows.NewFrontend(ctx, appoptions, logger, appBindings, dispatcher)
}
