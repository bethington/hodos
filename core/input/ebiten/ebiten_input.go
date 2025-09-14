// Package ebiten provides graphics and input API to develop a 2D game.
package input

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"nostos/common/enum"
)

var (
	//nolint:gochecknoglobals // This is a constant in all but by name, no constant map in go
	keyToEbiten = map[enum.Key]ebiten.Key{
		enum.Key0:            ebiten.Key0,
		enum.Key1:            ebiten.Key1,
		enum.Key2:            ebiten.Key2,
		enum.Key3:            ebiten.Key3,
		enum.Key4:            ebiten.Key4,
		enum.Key5:            ebiten.Key5,
		enum.Key6:            ebiten.Key6,
		enum.Key7:            ebiten.Key7,
		enum.Key8:            ebiten.Key8,
		enum.Key9:            ebiten.Key9,
		enum.KeyA:            ebiten.KeyA,
		enum.KeyB:            ebiten.KeyB,
		enum.KeyC:            ebiten.KeyC,
		enum.KeyD:            ebiten.KeyD,
		enum.KeyE:            ebiten.KeyE,
		enum.KeyF:            ebiten.KeyF,
		enum.KeyG:            ebiten.KeyG,
		enum.KeyH:            ebiten.KeyH,
		enum.KeyI:            ebiten.KeyI,
		enum.KeyJ:            ebiten.KeyJ,
		enum.KeyK:            ebiten.KeyK,
		enum.KeyL:            ebiten.KeyL,
		enum.KeyM:            ebiten.KeyM,
		enum.KeyN:            ebiten.KeyN,
		enum.KeyO:            ebiten.KeyO,
		enum.KeyP:            ebiten.KeyP,
		enum.KeyQ:            ebiten.KeyQ,
		enum.KeyR:            ebiten.KeyR,
		enum.KeyS:            ebiten.KeyS,
		enum.KeyT:            ebiten.KeyT,
		enum.KeyU:            ebiten.KeyU,
		enum.KeyV:            ebiten.KeyV,
		enum.KeyW:            ebiten.KeyW,
		enum.KeyX:            ebiten.KeyX,
		enum.KeyY:            ebiten.KeyY,
		enum.KeyZ:            ebiten.KeyZ,
		enum.KeyApostrophe:   ebiten.KeyApostrophe,
		enum.KeyBackslash:    ebiten.KeyBackslash,
		enum.KeyBackspace:    ebiten.KeyBackspace,
		enum.KeyCapsLock:     ebiten.KeyCapsLock,
		enum.KeyComma:        ebiten.KeyComma,
		enum.KeyDelete:       ebiten.KeyDelete,
		enum.KeyDown:         ebiten.KeyDown,
		enum.KeyEnd:          ebiten.KeyEnd,
		enum.KeyEnter:        ebiten.KeyEnter,
		enum.KeyEqual:        ebiten.KeyEqual,
		enum.KeyEscape:       ebiten.KeyEscape,
		enum.KeyF1:           ebiten.KeyF1,
		enum.KeyF2:           ebiten.KeyF2,
		enum.KeyF3:           ebiten.KeyF3,
		enum.KeyF4:           ebiten.KeyF4,
		enum.KeyF5:           ebiten.KeyF5,
		enum.KeyF6:           ebiten.KeyF6,
		enum.KeyF7:           ebiten.KeyF7,
		enum.KeyF8:           ebiten.KeyF8,
		enum.KeyF9:           ebiten.KeyF9,
		enum.KeyF10:          ebiten.KeyF10,
		enum.KeyF11:          ebiten.KeyF11,
		enum.KeyF12:          ebiten.KeyF12,
		enum.KeyGraveAccent:  ebiten.KeyGraveAccent,
		enum.KeyHome:         ebiten.KeyHome,
		enum.KeyInsert:       ebiten.KeyInsert,
		enum.KeyKP0:          ebiten.KeyKP0,
		enum.KeyKP1:          ebiten.KeyKP1,
		enum.KeyKP2:          ebiten.KeyKP2,
		enum.KeyKP3:          ebiten.KeyKP3,
		enum.KeyKP4:          ebiten.KeyKP4,
		enum.KeyKP5:          ebiten.KeyKP5,
		enum.KeyKP6:          ebiten.KeyKP6,
		enum.KeyKP7:          ebiten.KeyKP7,
		enum.KeyKP8:          ebiten.KeyKP8,
		enum.KeyKP9:          ebiten.KeyKP9,
		enum.KeyKPAdd:        ebiten.KeyKPAdd,
		enum.KeyKPDecimal:    ebiten.KeyKPDecimal,
		enum.KeyKPDivide:     ebiten.KeyKPDivide,
		enum.KeyKPEnter:      ebiten.KeyKPEnter,
		enum.KeyKPEqual:      ebiten.KeyKPEqual,
		enum.KeyKPMultiply:   ebiten.KeyKPMultiply,
		enum.KeyKPSubtract:   ebiten.KeyKPSubtract,
		enum.KeyLeft:         ebiten.KeyLeft,
		enum.KeyLeftBracket:  ebiten.KeyLeftBracket,
		enum.KeyMenu:         ebiten.KeyMenu,
		enum.KeyMinus:        ebiten.KeyMinus,
		enum.KeyNumLock:      ebiten.KeyNumLock,
		enum.KeyPageDown:     ebiten.KeyPageDown,
		enum.KeyPageUp:       ebiten.KeyPageUp,
		enum.KeyPause:        ebiten.KeyPause,
		enum.KeyPeriod:       ebiten.KeyPeriod,
		enum.KeyPrintScreen:  ebiten.KeyPrintScreen,
		enum.KeyRight:        ebiten.KeyRight,
		enum.KeyRightBracket: ebiten.KeyRightBracket,
		enum.KeyScrollLock:   ebiten.KeyScrollLock,
		enum.KeySemicolon:    ebiten.KeySemicolon,
		enum.KeySlash:        ebiten.KeySlash,
		enum.KeySpace:        ebiten.KeySpace,
		enum.KeyTab:          ebiten.KeyTab,
		enum.KeyUp:           ebiten.KeyUp,
		enum.KeyAlt:          ebiten.KeyAlt,
		enum.KeyControl:      ebiten.KeyControl,
		enum.KeyShift:        ebiten.KeyShift,
	}
	//nolint:gochecknoglobals // This is a constant in all but by name, no constant map in go
	mouseButtonToEbiten = map[enum.MouseButton]ebiten.MouseButton{
		enum.MouseButtonLeft:   ebiten.MouseButtonLeft,
		enum.MouseButtonMiddle: ebiten.MouseButtonMiddle,
		enum.MouseButtonRight:  ebiten.MouseButtonRight,
	}
)

// InputService provides an abstraction on ebiten to support handling input events
type InputService struct{}

// CursorPosition returns a position of a mouse cursor relative to the game screen (window).
func (is InputService) CursorPosition() (x, y int) {
	return ebiten.CursorPosition()
}

// InputChars return "printable" runes read from the keyboard at the time update is called.
func (is InputService) InputChars() []rune {
	return ebiten.InputChars()
}

// IsKeyPressed checks if the provided key is down.
func (is InputService) IsKeyPressed(key enum.Key) bool {
	return ebiten.IsKeyPressed(keyToEbiten[key])
}

// IsKeyJustPressed checks if the provided key is just transitioned from up to down.
func (is InputService) IsKeyJustPressed(key enum.Key) bool {
	return inpututil.IsKeyJustPressed(keyToEbiten[key])
}

// IsKeyJustReleased checks if the provided key is just transitioned from down to up.
func (is InputService) IsKeyJustReleased(key enum.Key) bool {
	return inpututil.IsKeyJustReleased(keyToEbiten[key])
}

// IsMouseButtonPressed checks if the provided mouse button is down.
func (is InputService) IsMouseButtonPressed(button enum.MouseButton) bool {
	return ebiten.IsMouseButtonPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustPressed checks if the provided mouse button is just transitioned from up to down.
func (is InputService) IsMouseButtonJustPressed(button enum.MouseButton) bool {
	return inpututil.IsMouseButtonJustPressed(mouseButtonToEbiten[button])
}

// IsMouseButtonJustReleased checks if the provided mouse button is just transitioned from down to up.
func (is InputService) IsMouseButtonJustReleased(button enum.MouseButton) bool {
	return inpututil.IsMouseButtonJustReleased(mouseButtonToEbiten[button])
}

// KeyPressDuration returns how long the key is pressed in frames.
func (is InputService) KeyPressDuration(key enum.Key) int {
	return inpututil.KeyPressDuration(keyToEbiten[key])
}
