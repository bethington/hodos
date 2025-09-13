package d2interface

import "nostos/common/enum"

// InputService represents an interface offering Keyboard and Mouse interactions.
type InputService interface {
	// CursorPosition returns a position of a mouse cursor relative to the game screen (window).
	CursorPosition() (x int, y int)
	// InputChars return "printable" runes read from the keyboard at the time update is called.
	InputChars() []rune
	// IsKeyPressed checks if the provided key is down.
	IsKeyPressed(key enum.Key) bool
	// IsKeyJustPressed checks if the provided key is just transitioned from up to down.
	IsKeyJustPressed(key enum.Key) bool
	// IsKeyJustReleased checks if the provided key is just transitioned from down to up.
	IsKeyJustReleased(key enum.Key) bool
	// IsMouseButtonPressed checks if the provided mouse button is down.
	IsMouseButtonPressed(button enum.MouseButton) bool
	// IsMouseButtonJustPressed checks if the provided mouse button is just transitioned from up to down.
	IsMouseButtonJustPressed(button enum.MouseButton) bool
	// IsMouseButtonJustReleased checks if the provided mouse button is just transitioned from down to up.
	IsMouseButtonJustReleased(button enum.MouseButton) bool
	// KeyPressDuration returns how long the key is pressed in frames.
	KeyPressDuration(key enum.Key) int
}
