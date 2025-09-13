package interfaces

import "nostos/common/enum"

// InputManager manages an InputService
type InputManager interface {
	Advance(elapsedTime, currentTime float64) error
	BindHandlerWithPriority(InputEventHandler, enum.Priority) error
	BindHandler(h InputEventHandler) error
	UnbindHandler(handler InputEventHandler) error
}
