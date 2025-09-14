package console

import (
	"nostos/common/enum"
	d2interface "nostos/common/interfaces"
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
