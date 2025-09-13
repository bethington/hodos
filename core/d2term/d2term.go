package d2term

import (
	"nostos/common/enum"
	"nostos/common/d2interface"
)

// New creates and initializes the terminal
func New(inputManager d2interface.InputManager) (*Terminal, error) {
	term, err := NewTerminal()
	if err != nil {
		return nil, err
	}

	if err := inputManager.BindHandlerWithPriority(term, enum.PriorityHigh); err != nil {
		return nil, err
	}

	return term, nil
}
