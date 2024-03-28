//go:build (linux || darwin) && !bindings

package app

import (
	"github.com/wailsapp/wails/v2/internal/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
)


// ff:
// _:
// _:
func PreflightChecks(_ *options.App, _ *logger.Logger) error {
	return nil
}
