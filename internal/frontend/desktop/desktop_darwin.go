//go:build darwin
// +build darwin

package desktop

import (
	"context"

	"github.com/888go/wails/internal/frontend/desktop/darwin"

	"github.com/888go/wails/internal/binding"
	"github.com/888go/wails/internal/frontend"

	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/pkg/options"
)

func NewFrontend(ctx context.Context, appoptions *options.App, logger *logger.Logger, appBindings *binding.Bindings, dispatcher frontend.Dispatcher) frontend.Frontend {
	return darwin.NewFrontend(ctx, appoptions, logger, appBindings, dispatcher)
}
