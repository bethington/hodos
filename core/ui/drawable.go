package ui

import (
	d2interface "nostos/common/interfaces"
)

// Drawable represents an instance that can be drawn
type Drawable interface {
	Render(target d2interface.Surface)
	Advance(elapsed float64) error
	GetSize() (width, height int)
	SetPosition(x, y int)
	GetPosition() (x, y int)
	OffsetPosition(xo, yo int)
	GetVisible() bool
	SetVisible(visible bool)
	SetRenderPriority(priority RenderPriority)
	GetRenderPriority() (priority RenderPriority)
}
