package main

import (
	"github.com/pterm/pterm"
	"github.com/888go/wails/cmd/wails/flags"
	"github.com/888go/wails/cmd/wails/internal"
	"github.com/888go/wails/internal/colour"
	"github.com/888go/wails/internal/github"
)

func showReleaseNotes(f *flags.ShowReleaseNotes) error {
	if f.NoColour {
		pterm.DisableColor()
		colour.ColourEnabled = false
	}

	version := internal.Version
	if f.Version != "" {
		version = f.Version
	}

	app.PrintBanner()
	releaseNotes := github.GetReleaseNotes(version, f.NoColour)
	pterm.Println(releaseNotes)

	return nil
}
