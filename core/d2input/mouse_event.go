package d2input

import "nostos/common/enum"

// MouseEvent represents a mouse event
type MouseEvent struct {
	HandlerEvent
	mouseButton enum.MouseButton
}

// KeyMod returns the key mod
func (e *MouseEvent) KeyMod() enum.KeyMod {
	return e.HandlerEvent.keyMod
}

// ButtonMod represents a button mod
func (e *MouseEvent) ButtonMod() enum.MouseButtonMod {
	return e.HandlerEvent.buttonMod
}

// X returns the event's X position
func (e *MouseEvent) X() int {
	return e.HandlerEvent.x
}

// Y returns the event's Y position
func (e *MouseEvent) Y() int {
	return e.HandlerEvent.y
}

// Button returns the mouse button
func (e *MouseEvent) Button() enum.MouseButton {
	return e.mouseButton
}
