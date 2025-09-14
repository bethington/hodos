package input

import "nostos/common/enum"

// KeyEvent represents key events
type KeyEvent struct {
	HandlerEvent
	key enum.Key
	// Duration represents the number of frames this key has been pressed for
	duration int
}

// Key returns the key
func (e *KeyEvent) Key() enum.Key {
	return e.key
}

// Duration returns the duration
func (e *KeyEvent) Duration() int {
	return e.duration
}
