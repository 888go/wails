//go:build (linux || darwin) && !bindings

package app

import (
	"github.com/888go/wails/internal/logger"
	"github.com/888go/wails/pkg/options"
)


// ff:
// _:
// _:
func PreflightChecks(_ *options.App, _ *logger.Logger) error {
	return nil
}
