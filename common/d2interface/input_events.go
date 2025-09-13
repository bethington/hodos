package d2interface

import "nostos/common/enum"

// HandlerEvent holds the qualifiers for a key or mouse event
type HandlerEvent interface {
	KeyMod() enum.KeyMod
	ButtonMod() enum.MouseButtonMod
	X() int
	Y() int
}

// KeyEvent represents an event associated with a keyboard key
type KeyEvent interface {
	HandlerEvent
	Key() enum.Key
	// Duration represents the number of frames this key has been pressed for
	Duration() int
}

// KeyCharsEvent represents an event associated with a keyboard character being pressed
type KeyCharsEvent interface {
	HandlerEvent
	Chars() []rune
}

// MouseEvent represents a mouse event
type MouseEvent interface {
	HandlerEvent
	Button() enum.MouseButton
}

// MouseMoveEvent represents a mouse movement event
type MouseMoveEvent interface {
	HandlerEvent
}
