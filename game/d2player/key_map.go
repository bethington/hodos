package d2player

import (
	"sync"

	"nostos/common/enum"
	d2asset "nostos/core/asset"
)

// KeyMap represents the key mappings of the game. Each game event
// can be associated to 2 different keys. A key of -1 means none
type KeyMap struct {
	mutex              sync.RWMutex
	mapping            map[enum.Key]enum.GameEvent
	controls           map[enum.GameEvent]*KeyBinding
	keyToStringMapping map[enum.Key]string
}

// KeyBindingType defines whether it's a primary or
// secondary binding
type KeyBindingType int

// Values defining the type of key binding
const (
	KeyBindingTypeNone KeyBindingType = iota
	KeyBindingTypePrimary
	KeyBindingTypeSecondary
)

// NewKeyMap returns a new instance of a KeyMap
func NewKeyMap(asset *d2asset.AssetManager) *KeyMap {
	return &KeyMap{
		mapping:            make(map[enum.Key]enum.GameEvent),
		controls:           make(map[enum.GameEvent]*KeyBinding),
		keyToStringMapping: getKeyStringMapping(asset),
	}
}

func getKeyStringMapping(assetManager *d2asset.AssetManager) map[enum.Key]string {
	return map[enum.Key]string{
		-1:                       assetManager.TranslateString("KeyNone"),
		enum.KeyTilde:          "~",
		enum.KeyHome:           assetManager.TranslateString("KeyHome"),
		enum.KeyControl:        assetManager.TranslateString("KeyControl"),
		enum.KeyShift:          assetManager.TranslateString("KeyShift"),
		enum.KeySpace:          assetManager.TranslateString("KeySpace"),
		enum.KeyAlt:            assetManager.TranslateString("KeyMenu"),
		enum.KeyTab:            assetManager.TranslateString("KeyTab"),
		enum.Key0:              "0",
		enum.Key1:              "1",
		enum.Key2:              "2",
		enum.Key3:              "3",
		enum.Key4:              "4",
		enum.Key5:              "5",
		enum.Key6:              "6",
		enum.Key7:              "7",
		enum.Key8:              "8",
		enum.Key9:              "9",
		enum.KeyA:              "A",
		enum.KeyB:              "B",
		enum.KeyC:              "C",
		enum.KeyD:              "D",
		enum.KeyE:              "E",
		enum.KeyF:              "F",
		enum.KeyG:              "G",
		enum.KeyH:              "H",
		enum.KeyI:              "I",
		enum.KeyJ:              "J",
		enum.KeyK:              "K",
		enum.KeyL:              "L",
		enum.KeyM:              "M",
		enum.KeyN:              "N",
		enum.KeyO:              "O",
		enum.KeyP:              "P",
		enum.KeyQ:              "Q",
		enum.KeyR:              "R",
		enum.KeyS:              "S",
		enum.KeyT:              "T",
		enum.KeyU:              "U",
		enum.KeyV:              "V",
		enum.KeyW:              "W",
		enum.KeyX:              "X",
		enum.KeyY:              "Y",
		enum.KeyZ:              "Z",
		enum.KeyF1:             "F1",
		enum.KeyF2:             "F2",
		enum.KeyF3:             "F3",
		enum.KeyF4:             "F4",
		enum.KeyF5:             "F5",
		enum.KeyF6:             "F6",
		enum.KeyF7:             "F7",
		enum.KeyF8:             "F8",
		enum.KeyF9:             "F9",
		enum.KeyF10:            "F10",
		enum.KeyF11:            "F11",
		enum.KeyF12:            "F12",
		enum.KeyKP0:            assetManager.TranslateString("KeyNumPad0"),
		enum.KeyKP1:            assetManager.TranslateString("KeyNumPad1"),
		enum.KeyKP2:            assetManager.TranslateString("KeyNumPad2"),
		enum.KeyKP3:            assetManager.TranslateString("KeyNumPad3"),
		enum.KeyKP4:            assetManager.TranslateString("KeyNumPad4"),
		enum.KeyKP5:            assetManager.TranslateString("KeyNumPad5"),
		enum.KeyKP6:            assetManager.TranslateString("KeyNumPad6"),
		enum.KeyKP7:            assetManager.TranslateString("KeyNumPad7"),
		enum.KeyKP8:            assetManager.TranslateString("KeyNumPad8"),
		enum.KeyKP9:            assetManager.TranslateString("KeyNumPad9"),
		enum.KeyPrintScreen:    assetManager.TranslateString("KeySnapshot"),
		enum.KeyRightBracket:   assetManager.TranslateString("KeyRBracket"),
		enum.KeyLeftBracket:    assetManager.TranslateString("KeyLBracket"),
		enum.KeyMouse3:         assetManager.TranslateString("KeyMButton"),
		enum.KeyMouse4:         assetManager.TranslateString("Key4Button"),
		enum.KeyMouse5:         assetManager.TranslateString("Key5Button"),
		enum.KeyMouseWheelUp:   assetManager.TranslateString("KeyWheelUp"),
		enum.KeyMouseWheelDown: assetManager.TranslateString("KeyWheelDown"),
	}
}
func (km *KeyMap) checkOverwrite(key enum.Key) (*KeyBinding, KeyBindingType) {
	var (
		overwrittenBinding     *KeyBinding
		overwrittenBindingType KeyBindingType
	)

	for _, binding := range km.controls {
		if binding.Primary == key {
			binding.Primary = -1
			overwrittenBinding = binding
			overwrittenBindingType = KeyBindingTypePrimary
		}

		if binding.Secondary == key {
			binding.Secondary = -1
			overwrittenBinding = binding
			overwrittenBindingType = KeyBindingTypeSecondary
		}
	}

	return overwrittenBinding, overwrittenBindingType
}

// SetPrimaryBinding binds the first key for gameEvent
func (km *KeyMap) SetPrimaryBinding(gameEvent enum.GameEvent, key enum.Key) (*KeyBinding, KeyBindingType) {
	if key == enum.KeyEscape {
		return nil, -1
	}

	km.mutex.Lock()
	defer km.mutex.Unlock()

	if km.controls[gameEvent] == nil {
		km.controls[gameEvent] = &KeyBinding{}
	}

	overwrittenBinding, overwrittenBindingType := km.checkOverwrite(key)

	currentKey := km.controls[gameEvent].Primary
	delete(km.mapping, currentKey)
	km.mapping[key] = gameEvent

	km.controls[gameEvent].Primary = key

	return overwrittenBinding, overwrittenBindingType
}

// SetSecondaryBinding binds the second key for gameEvent
func (km *KeyMap) SetSecondaryBinding(gameEvent enum.GameEvent, key enum.Key) (*KeyBinding, KeyBindingType) {
	if key == enum.KeyEscape {
		return nil, -1
	}

	km.mutex.Lock()
	defer km.mutex.Unlock()

	if km.controls[gameEvent] == nil {
		km.controls[gameEvent] = &KeyBinding{}
	}

	overwrittenBinding, overwrittenBindingType := km.checkOverwrite(key)

	currentKey := km.controls[gameEvent].Secondary
	delete(km.mapping, currentKey)
	km.mapping[key] = gameEvent

	if km.controls[gameEvent].Primary == key {
		km.controls[gameEvent].Primary = enum.Key(-1)
	}

	km.controls[gameEvent].Secondary = key

	return overwrittenBinding, overwrittenBindingType
}

func (km *KeyMap) getGameEvent(key enum.Key) enum.GameEvent {
	km.mutex.RLock()
	defer km.mutex.RUnlock()

	return km.mapping[key]
}

// GetKeysForGameEvent returns the bindings for a givent game event
func (km *KeyMap) GetKeysForGameEvent(gameEvent enum.GameEvent) *KeyBinding {
	km.mutex.RLock()
	defer km.mutex.RUnlock()

	return km.controls[gameEvent]
}

// GetBindingByKey returns the bindings for a givent game event
func (km *KeyMap) GetBindingByKey(key enum.Key) (*KeyBinding, enum.GameEvent, KeyBindingType) {
	km.mutex.RLock()
	defer km.mutex.RUnlock()

	for gameEvent, binding := range km.controls {
		if binding.Primary == key {
			return binding, gameEvent, KeyBindingTypePrimary
		}

		if binding.Secondary == key {
			return binding, gameEvent, KeyBindingTypeSecondary
		}
	}

	return nil, -1, -1
}

// KeyBinding holds the primary and secondary keys assigned to a GameEvent
type KeyBinding struct {
	Primary   enum.Key
	Secondary enum.Key
}

// IsEmpty checks if no keys are associated to the binding
func (b KeyBinding) IsEmpty() bool {
	return b.Primary == -1 && b.Secondary == -1
}

// ResetToDefault will reset the KeyMap to the default values
func (km *KeyMap) ResetToDefault() {
	defaultControls := map[enum.GameEvent]KeyBinding{
		enum.ToggleCharacterPanel: {enum.KeyA, enum.KeyC},
		enum.ToggleInventoryPanel: {enum.KeyB, enum.KeyI},
		enum.TogglePartyPanel:     {enum.KeyP, -1},
		enum.ToggleHirelingPanel:  {enum.KeyO, -1},
		enum.ToggleMessageLog:     {enum.KeyM, -1},
		enum.ToggleQuestLog:       {enum.KeyQ, -1},
		enum.ToggleHelpScreen:     {enum.KeyH, -1},

		enum.ToggleSkillTreePanel:     {enum.KeyT, -1},
		enum.ToggleRightSkillSelector: {enum.KeyS, -1},
		enum.UseSkill1:                {enum.KeyF1, -1},
		enum.UseSkill2:                {enum.KeyF2, -1},
		enum.UseSkill3:                {enum.KeyF3, -1},
		enum.UseSkill4:                {enum.KeyF4, -1},
		enum.UseSkill5:                {enum.KeyF5, -1},
		enum.UseSkill6:                {enum.KeyF6, -1},
		enum.UseSkill7:                {enum.KeyF7, -1},
		enum.UseSkill8:                {enum.KeyF8, -1},
		enum.UseSkill9:                {-1, -1},
		enum.UseSkill10:               {-1, -1},
		enum.UseSkill11:               {-1, -1},
		enum.UseSkill12:               {-1, -1},
		enum.UseSkill13:               {-1, -1},
		enum.UseSkill14:               {-1, -1},
		enum.UseSkill15:               {-1, -1},
		enum.UseSkill16:               {-1, -1},
		enum.SelectPreviousSkill:      {enum.KeyMouseWheelUp, -1},
		enum.SelectNextSkill:          {enum.KeyMouseWheelDown, -1},

		enum.ToggleBelts:  {enum.KeyTilde, -1},
		enum.UseBeltSlot1: {enum.Key1, -1},
		enum.UseBeltSlot2: {enum.Key2, -1},
		enum.UseBeltSlot3: {enum.Key3, -1},
		enum.UseBeltSlot4: {enum.Key4, -1},
		enum.SwapWeapons:  {enum.KeyW, -1},

		enum.ToggleChatBox:       {enum.KeyEnter, -1},
		enum.HoldRun:             {enum.KeyControl, -1},
		enum.ToggleRunWalk:       {enum.KeyR, -1},
		enum.HoldStandStill:      {enum.KeyShift, -1},
		enum.HoldShowGroundItems: {enum.KeyAlt, -1},
		enum.HoldShowPortraits:   {enum.KeyZ, -1},

		enum.ToggleAutomap:        {enum.KeyTab, -1},
		enum.CenterAutomap:        {enum.KeyHome, -1},
		enum.TogglePartyOnAutomap: {enum.KeyF11, -1},
		enum.ToggleNamesOnAutomap: {enum.KeyF12, -1},
		enum.ToggleMiniMap:        {enum.KeyV, -1},

		enum.SayHelp:         {enum.KeyKP0, -1},
		enum.SayFollowMe:     {enum.KeyKP1, -1},
		enum.SayThisIsForYou: {enum.KeyKP2, -1},
		enum.SayThanks:       {enum.KeyKP3, -1},
		enum.SaySorry:        {enum.KeyKP4, -1},
		enum.SayBye:          {enum.KeyKP5, -1},
		enum.SayNowYouDie:    {enum.KeyKP6, -1},
		enum.SayRetreat:      {enum.KeyKP7, -1},

		enum.TakeScreenShot: {enum.KeyPrintScreen, -1},
		enum.ClearScreen:    {enum.KeySpace, -1},
		enum.ClearMessages:  {enum.KeyN, -1},
	}

	for gameEvent, keys := range defaultControls {
		km.SetPrimaryBinding(gameEvent, keys.Primary)
		km.SetSecondaryBinding(gameEvent, keys.Secondary)
	}
}

// KeyToString returns a string representing the key
func (km *KeyMap) KeyToString(k enum.Key) string {
	return km.keyToStringMapping[k]
}

// GetDefaultKeyMap generates a KeyMap instance with the
// default values
func GetDefaultKeyMap(asset *d2asset.AssetManager) *KeyMap {
	keyMap := NewKeyMap(asset)
	keyMap.ResetToDefault()

	return keyMap
}
