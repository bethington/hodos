package ui

import (
	d2interface "nostos/common/interfaces"
	d2util "nostos/common/util"
	d2asset "nostos/core/asset"
)

// CursorButton represents a mouse button
type CursorButton uint8

const (
	logPrefix = "UI Manager"
)

const (
	// CursorButtonLeft represents the left mouse button
	CursorButtonLeft CursorButton = 1
	// CursorButtonRight represents the right mouse button
	CursorButtonRight CursorButton = 2
)

// HorizontalAlign type, determines alignment along x-axis within a layout
type HorizontalAlign int

// Horizontal alignment types
const (
	HorizontalAlignLeft HorizontalAlign = iota
	HorizontalAlignCenter
	HorizontalAlignRight
)

// NewUIManager creates a UIManager instance with the given input and audio provider
func NewUIManager(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	input d2interface.InputManager,
	l d2util.LogLevel,
	audio d2interface.AudioProvider,
) *UIManager {
	ui := &UIManager{
		asset:        asset,
		renderer:     renderer,
		inputManager: input,
		audio:        audio,
	}

	ui.Logger = d2util.NewLogger()
	ui.Logger.SetPrefix(logPrefix)
	ui.Logger.SetLevel(l)

	return ui
}
