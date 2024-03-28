package main

import (
	"os"

	"github.com/pterm/pterm"
	"github.com/888go/wails/cmd/wails/flags"
	"github.com/888go/wails/cmd/wails/internal/dev"
	"github.com/888go/wails/internal/colour"
	"github.com/888go/wails/pkg/clilogger"
)

func devApplication(f *flags.Dev) error {
	if f.NoColour {
		pterm.DisableColor()
		colour.ColourEnabled = false
	}

	quiet := f.Verbosity == flags.Quiet

	// Create logger
	logger := clilogger.X创建(os.Stdout)
	logger.X禁用日志(quiet)

	if quiet {
		pterm.DisableOutput()
	} else {
		app.PrintBanner()
	}

	err := f.Process()
	if err != nil {
		return err
	}

	return dev.Application(f, logger)
}
