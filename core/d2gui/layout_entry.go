package d2gui

import (
	d2geom "nostos/common/geom"
	d2interface "nostos/common/interfaces"
)

type layoutEntry struct {
	widget widget

	x      int
	y      int
	width  int
	height int

	mouseOver bool
	mouseDown [3]bool
}

// IsIn layout entry, spc. of an event.
func (l *layoutEntry) IsIn(event d2interface.HandlerEvent) bool {
	sx, sy := l.widget.ScreenPos()
	rect := d2geom.Rectangle{Left: sx, Top: sy, Width: l.width, Height: l.height}

	return rect.IsInRect(event.X(), event.Y())
}
