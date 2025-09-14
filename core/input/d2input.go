package input

import (
	"errors"

	d2interface "nostos/common/interfaces"
)

var (
	// ErrHasReg shows the input system already has a registered handler
	ErrHasReg = errors.New("input system already has provided handler")
	// ErrNotReg shows the input system has no registered handler
	ErrNotReg = errors.New("input system does not have provided handler")
)

// Static checks to confirm struct conforms to interface
var _ d2interface.InputEventHandler = &HandlerEvent{}
var _ d2interface.KeyEvent = &KeyEvent{}
var _ d2interface.KeyCharsEvent = &KeyCharsEvent{}
var _ d2interface.MouseEvent = &MouseEvent{}
var _ d2interface.MouseMoveEvent = &MouseMoveEvent{}
