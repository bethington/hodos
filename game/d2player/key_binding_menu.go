package d2player

import (
	"image/color"

	"nostos/common/enum"
	d2interface "nostos/common/interfaces"
	d2util "nostos/common/util"
	d2asset "nostos/core/asset"
	d2gui "nostos/core/gui"
	"nostos/core/d2ui"
)

const (
	selectionBackgroundColor = 0x000000d0
	defaultGameEvent         = -1

	keyBindingMenuWidth  = 620
	keyBindingMenuHeight = 375
	keyBindingMenuX      = 90
	keyBindingMenuY      = 75

	keyBindingMenuPaddingX    = 17
	keyBindingSettingPaddingY = 19

	keyBindingMenuHeaderHeight  = 24
	keyBindingMenuHeaderSpacer1 = 131
	keyBindingMenuHeaderSpacer2 = 86

	keyBindingMenuBindingSpacerBetween   = 25
	keyBindingMenuBindingSpacerLeft      = 17
	keyBindingMenuBindingDescWidth       = 190
	keyBindingMenuBindingDescHeight      = 0
	keyBindingMenuBindingPrimaryWidth    = 190
	keyBindingMenuBindingPrimaryHeight   = 0
	keyBindingMenuBindingSecondaryWidth  = 90
	keyBindingMenuBindingSecondaryHeight = 0
)

type bindingChange struct {
	target    *KeyBinding
	primary   enum.Key
	secondary enum.Key
}

// NewKeyBindingMenu generates a new instance of the "Configure Keys"
// menu found in the options
func NewKeyBindingMenu(
	asset *d2asset.AssetManager,
	renderer d2interface.Renderer,
	ui *d2ui.UIManager,
	guiManager *d2gui.GuiManager,
	keyMap *KeyMap,
	l d2util.LogLevel,
	escapeMenu *EscapeMenu,
) *KeyBindingMenu {
	mainLayout := d2gui.CreateLayout(renderer, d2gui.PositionTypeAbsolute, asset)
	contentLayout := mainLayout.AddLayout(d2gui.PositionTypeAbsolute)

	ret := &KeyBindingMenu{
		keyMap:           keyMap,
		asset:            asset,
		ui:               ui,
		guiManager:       guiManager,
		renderer:         renderer,
		mainLayout:       mainLayout,
		contentLayout:    contentLayout,
		bindingLayouts:   []*bindingLayout{},
		changesToBeSaved: make(map[enum.GameEvent]*bindingChange),
		escapeMenu:       escapeMenu,
	}

	ret.Logger = d2util.NewLogger()
	ret.Logger.SetLevel(l)
	ret.Logger.SetPrefix(logPrefix)

	ret.Box = d2gui.NewBox(
		asset, renderer, ui, ret.mainLayout,
		keyBindingMenuWidth, keyBindingMenuHeight,
		keyBindingMenuX, keyBindingMenuY, l, "",
	)

	ret.Box.SetPadding(keyBindingMenuPaddingX, keyBindingSettingPaddingY)

	ret.Box.SetOptions([]*d2gui.LabelButton{
		d2gui.NewLabelButton(0, 0, "Cancel", d2util.Color(d2gui.ColorRed), d2util.LogLevelDefault, func() {
			if err := ret.onCancelClicked(); err != nil {
				ret.Errorf("error while clicking option Cancel: %v", err.Error())
			}
		}),
		d2gui.NewLabelButton(0, 0, "Default", d2util.Color(d2gui.ColorBlue), d2util.LogLevelDefault, func() {
			if err := ret.onDefaultClicked(); err != nil {
				ret.Errorf("error while clicking option Default: %v", err)
			}
		}),
		d2gui.NewLabelButton(0, 0, "Accept", d2util.Color(d2gui.ColorGreen), d2util.LogLevelDefault, func() {
			if err := ret.onAcceptClicked(); err != nil {
				ret.Errorf("error while clicking option Accept: %v", err)
			}
		}),
	})

	return ret
}

// KeyBindingMenu represents the menu to view/edit the
// key bindings
type KeyBindingMenu struct {
	*d2gui.Box

	asset      *d2asset.AssetManager
	renderer   d2interface.Renderer
	ui         *d2ui.UIManager
	guiManager *d2gui.GuiManager
	keyMap     *KeyMap
	escapeMenu *EscapeMenu

	mainLayout       *d2gui.Layout
	contentLayout    *d2gui.Layout
	scrollbar        *d2gui.LayoutScrollbar
	bindingLayouts   []*bindingLayout
	changesToBeSaved map[enum.GameEvent]*bindingChange

	isAwaitingKeyDown          bool
	currentBindingModifierType KeyBindingType
	currentBindingModifier     enum.GameEvent
	currentBindingLayout       *bindingLayout
	lastBindingLayout          *bindingLayout

	*d2util.Logger
}

// Close will disable the render of the menu and clear
// the current selection
func (menu *KeyBindingMenu) Close() error {
	menu.Box.Close()

	if err := menu.clearSelection(); err != nil {
		return err
	}

	return nil
}

// Load will setup the layouts of the menu
func (menu *KeyBindingMenu) Load() error {
	if err := menu.Box.Load(); err != nil {
		return err
	}

	mainLayoutW, mainLayoutH := menu.mainLayout.GetSize()

	headerLayout := menu.contentLayout.AddLayout(d2gui.PositionTypeHorizontal)
	headerLayout.SetSize(mainLayoutW, keyBindingMenuHeaderHeight)

	if _, err := headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgFunction"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(d2gui.ColorBrown),
	); err != nil {
		return err
	}

	headerLayout.AddSpacerStatic(keyBindingMenuHeaderSpacer1, keyBindingMenuHeaderHeight)

	if _, err := headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgPrimaryKey"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(d2gui.ColorBrown),
	); err != nil {
		return err
	}

	headerLayout.AddSpacerStatic(keyBindingMenuHeaderSpacer2, 1)

	if _, err := headerLayout.AddLabelWithColor(
		menu.asset.TranslateString("CfgSecondaryKey"),
		d2gui.FontStyleFormal11Units,
		d2util.Color(d2gui.ColorBrown),
	); err != nil {
		return err
	}

	headerLayout.SetVerticalAlign(d2gui.VerticalAlignMiddle)

	bindingWrapper := menu.contentLayout.AddLayout(d2gui.PositionTypeAbsolute)
	bindingWrapper.SetPosition(0, keyBindingMenuHeaderHeight)
	bindingWrapper.SetSize(mainLayoutW, mainLayoutH-keyBindingMenuHeaderHeight)

	bindingLayout := menu.generateLayout()

	menu.Box.GetLayout().AdjustEntryPlacement()
	menu.mainLayout.AdjustEntryPlacement()
	menu.contentLayout.AdjustEntryPlacement()

	menu.scrollbar = d2gui.NewLayoutScrollbar(bindingWrapper, bindingLayout)

	if err := menu.scrollbar.Load(menu.ui); err != nil {
		return err
	}

	bindingWrapper.AddLayoutFromSource(bindingLayout)
	bindingWrapper.AdjustEntryPlacement()

	return nil
}

type keyBindingSetting struct {
	label     string
	gameEvent enum.GameEvent
}

func (menu *KeyBindingMenu) getBindingGroups() [][]keyBindingSetting {
	return [][]keyBindingSetting{
		{
			{menu.asset.TranslateString("CfgCharacter"), enum.ToggleCharacterPanel},
			{menu.asset.TranslateString("CfgInventory"), enum.ToggleInventoryPanel},
			{menu.asset.TranslateString("CfgParty"), enum.TogglePartyPanel},
			{menu.asset.TranslateString("Cfghireling"), enum.ToggleHirelingPanel},
			{menu.asset.TranslateString("CfgMessageLog"), enum.ToggleMessageLog},
			{menu.asset.TranslateString("CfgQuestLog"), enum.ToggleQuestLog},
			{menu.asset.TranslateString("CfgHelp"), enum.ToggleHelpScreen},
		},
		{
			{menu.asset.TranslateString("CfgSkillTree"), enum.ToggleSkillTreePanel},
			{menu.asset.TranslateString("CfgSkillPick"), enum.ToggleRightSkillSelector},
			{menu.asset.TranslateString("CfgSkill1"), enum.UseSkill1},
			{menu.asset.TranslateString("CfgSkill2"), enum.UseSkill2},
			{menu.asset.TranslateString("CfgSkill3"), enum.UseSkill3},
			{menu.asset.TranslateString("CfgSkill4"), enum.UseSkill4},
			{menu.asset.TranslateString("CfgSkill5"), enum.UseSkill5},
			{menu.asset.TranslateString("CfgSkill6"), enum.UseSkill6},
			{menu.asset.TranslateString("CfgSkill7"), enum.UseSkill7},
			{menu.asset.TranslateString("CfgSkill8"), enum.UseSkill8},
			{menu.asset.TranslateString("CfgSkill9"), enum.UseSkill9},
			{menu.asset.TranslateString("CfgSkill10"), enum.UseSkill10},
			{menu.asset.TranslateString("CfgSkill11"), enum.UseSkill11},
			{menu.asset.TranslateString("CfgSkill12"), enum.UseSkill12},
			{menu.asset.TranslateString("CfgSkill13"), enum.UseSkill13},
			{menu.asset.TranslateString("CfgSkill14"), enum.UseSkill14},
			{menu.asset.TranslateString("CfgSkill15"), enum.UseSkill15},
			{menu.asset.TranslateString("CfgSkill16"), enum.UseSkill16},
			{menu.asset.TranslateString("Cfgskillup"), enum.SelectPreviousSkill},
			{menu.asset.TranslateString("Cfgskilldown"), enum.SelectNextSkill},
		},
		{
			{menu.asset.TranslateString("CfgBeltShow"), enum.ToggleBelts},
			{menu.asset.TranslateString("CfgBelt1"), enum.UseBeltSlot1},
			{menu.asset.TranslateString("CfgBelt2"), enum.UseBeltSlot2},
			{menu.asset.TranslateString("CfgBelt3"), enum.UseBeltSlot3},
			{menu.asset.TranslateString("CfgBelt4"), enum.UseBeltSlot4},
			{menu.asset.TranslateString("Cfgswapweapons"), enum.SwapWeapons},
		},
		{
			{menu.asset.TranslateString("CfgChat"), enum.ToggleChatBox},
			{menu.asset.TranslateString("CfgRun"), enum.HoldRun},
			{menu.asset.TranslateString("CfgRunLock"), enum.ToggleRunWalk},
			{menu.asset.TranslateString("CfgStandStill"), enum.HoldStandStill},
			{menu.asset.TranslateString("CfgShowItems"), enum.HoldShowGroundItems},
			{menu.asset.TranslateString("CfgTogglePortraits"), enum.HoldShowPortraits},
		},
		{
			{menu.asset.TranslateString("CfgAutoMap"), enum.ToggleAutomap},
			{menu.asset.TranslateString("CfgAutoMapCenter"), enum.CenterAutomap},
			{menu.asset.TranslateString("CfgAutoMapParty"), enum.TogglePartyOnAutomap},
			{menu.asset.TranslateString("CfgAutoMapNames"), enum.ToggleNamesOnAutomap},
			{menu.asset.TranslateString("CfgToggleminimap"), enum.ToggleMiniMap},
		},
		{
			{menu.asset.TranslateString("CfgSay0"), enum.SayHelp},
			{menu.asset.TranslateString("CfgSay1"), enum.SayFollowMe},
			{menu.asset.TranslateString("CfgSay2"), enum.SayThisIsForYou},
			{menu.asset.TranslateString("CfgSay3"), enum.SayThanks},
			{menu.asset.TranslateString("CfgSay4"), enum.SaySorry},
			{menu.asset.TranslateString("CfgSay5"), enum.SayBye},
			{menu.asset.TranslateString("CfgSay6"), enum.SayNowYouDie},
			{menu.asset.TranslateString("CfgSay7"), enum.SayNowYouDie},
		},
		{
			{menu.asset.TranslateString("CfgSnapshot"), enum.TakeScreenShot},
			{menu.asset.TranslateString("CfgClearScreen"), enum.ClearScreen},
			{menu.asset.TranslateString("Cfgcleartextmsg"), enum.ClearMessages},
		},
	}
}

func (menu *KeyBindingMenu) generateLayout() *d2gui.Layout {
	groups := menu.getBindingGroups()

	wrapper := d2gui.CreateLayout(menu.renderer, d2gui.PositionTypeAbsolute, menu.asset)
	layout := wrapper.AddLayout(d2gui.PositionTypeVertical)

	for i, settingsGroup := range groups {
		groupLayout := layout.AddLayout(d2gui.PositionTypeVertical)

		for _, setting := range settingsGroup {
			bl := bindingLayout{}

			settingLayout := groupLayout.AddLayout(d2gui.PositionTypeHorizontal)
			settingLayout.AddSpacerStatic(keyBindingMenuBindingSpacerLeft, 0)
			descLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
			descLabelWrapper.SetSize(keyBindingMenuBindingDescWidth, keyBindingMenuBindingDescHeight)

			descLabel, _ := descLabelWrapper.AddLabel(setting.label, d2gui.FontStyleFormal11Units)
			descLabel.SetHoverColor(d2util.Color(d2gui.ColorBlue))

			bl.wrapperLayout = settingLayout
			bl.descLabel = descLabel
			bl.descLayout = descLabelWrapper

			if binding := menu.keyMap.GetKeysForGameEvent(setting.gameEvent); binding != nil {
				primaryStr := menu.keyMap.KeyToString(binding.Primary)
				secondaryStr := menu.keyMap.KeyToString(binding.Secondary)
				primaryCol := menu.getKeyColor(binding.Primary)
				secondaryCol := menu.getKeyColor(binding.Secondary)

				if binding.IsEmpty() {
					primaryCol = d2util.Color(d2gui.ColorRed)
					secondaryCol = d2util.Color(d2gui.ColorRed)
				}

				primaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				primaryKeyLabelWrapper.SetSize(keyBindingMenuBindingPrimaryWidth, keyBindingMenuBindingPrimaryHeight)
				primaryLabel, _ := primaryKeyLabelWrapper.AddLabelWithColor(primaryStr, d2gui.FontStyleFormal11Units, primaryCol)
				primaryLabel.SetHoverColor(d2util.Color(d2gui.ColorBlue))

				bl.primaryLabel = primaryLabel
				bl.primaryLayout = primaryKeyLabelWrapper
				bl.gameEvent = setting.gameEvent

				secondaryKeyLabelWrapper := settingLayout.AddLayout(d2gui.PositionTypeAbsolute)
				secondaryKeyLabelWrapper.SetSize(keyBindingMenuBindingSecondaryWidth, keyBindingMenuBindingSecondaryHeight)
				secondaryLabel, _ := secondaryKeyLabelWrapper.AddLabelWithColor(secondaryStr, d2gui.FontStyleFormal11Units, secondaryCol)
				secondaryLabel.SetHoverColor(d2util.Color(d2gui.ColorBlue))

				bl.secondaryLabel = secondaryLabel
				bl.secondaryLayout = secondaryKeyLabelWrapper
				bl.binding = binding
			}

			menu.bindingLayouts = append(menu.bindingLayouts, &bl)
		}

		if i < len(groups)-1 {
			layout.AddSpacerStatic(0, keyBindingMenuBindingSpacerBetween)
		}
	}

	return wrapper
}

func (menu *KeyBindingMenu) getKeyColor(key enum.Key) color.RGBA {
	switch key {
	case -1:
		return d2util.Color(d2gui.ColorGrey)
	default:
		return d2util.Color(d2gui.ColorBrown)
	}
}

func (menu *KeyBindingMenu) setSelection(bl *bindingLayout, bindingType KeyBindingType, gameEvent enum.GameEvent) error {
	if menu.currentBindingLayout != nil {
		menu.lastBindingLayout = menu.currentBindingLayout
		if err := menu.currentBindingLayout.Reset(); err != nil {
			return err
		}
	}

	menu.currentBindingModifier = gameEvent
	menu.currentBindingLayout = bl

	if bindingType == KeyBindingTypePrimary {
		menu.currentBindingLayout.primaryLabel.SetIsBlinking(true)
	} else if bindingType == KeyBindingTypeSecondary {
		menu.currentBindingLayout.secondaryLabel.SetIsBlinking(true)
	}

	menu.currentBindingModifierType = bindingType
	menu.isAwaitingKeyDown = true

	if err := bl.descLabel.SetIsHovered(true); err != nil {
		return err
	}

	if err := bl.primaryLabel.SetIsHovered(true); err != nil {
		return err
	}

	if err := bl.secondaryLabel.SetIsHovered(true); err != nil {
		return err
	}

	return nil
}

func (menu *KeyBindingMenu) onMouseButtonDown(event d2interface.MouseEvent) error {
	if !menu.IsOpen() {
		return nil
	}

	menu.Box.OnMouseButtonDown(event)

	if menu.scrollbar != nil {
		if menu.scrollbar.IsInSliderRect(event.X(), event.Y()) {
			menu.scrollbar.SetSliderClicked(true)
			menu.scrollbar.OnSliderMouseClick(event)

			return nil
		}

		if menu.scrollbar.IsInArrowUpRect(event.X(), event.Y()) {
			if !menu.scrollbar.IsArrowUpClicked() {
				menu.scrollbar.SetArrowUpClicked(true)
			}

			menu.scrollbar.OnArrowUpClick()

			return nil
		}

		if menu.scrollbar.IsInArrowDownRect(event.X(), event.Y()) {
			if !menu.scrollbar.IsArrowDownClicked() {
				menu.scrollbar.SetArrowDownClicked(true)
			}

			menu.scrollbar.OnArrowDownClick()

			return nil
		}
	}

	for _, bl := range menu.bindingLayouts {
		gameEvent, typ := bl.GetPointedLayoutAndLabel(event.X(), event.Y())

		if gameEvent != -1 {
			if err := menu.setSelection(bl, typ, gameEvent); err != nil {
				return err
			}

			break
		} else if menu.currentBindingLayout != nil {
			if err := menu.clearSelection(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (menu *KeyBindingMenu) onMouseMove(event d2interface.MouseMoveEvent) {
	if !menu.IsOpen() {
		return
	}

	if menu.scrollbar != nil && menu.scrollbar.IsSliderClicked() {
		menu.scrollbar.OnMouseMove(event)
	}
}

func (menu *KeyBindingMenu) onMouseButtonUp() {
	if !menu.IsOpen() {
		return
	}

	if menu.scrollbar != nil {
		menu.scrollbar.SetSliderClicked(false)
		menu.scrollbar.SetArrowDownClicked(false)
		menu.scrollbar.SetArrowUpClicked(false)
	}
}

func (menu *KeyBindingMenu) getPendingChangeByKey(key enum.Key) (*bindingChange, *KeyBinding, enum.GameEvent, KeyBindingType) {
	var (
		existingBinding *KeyBinding
		gameEvent       enum.GameEvent
		bindingType     KeyBindingType
	)

	for ge, existingChange := range menu.changesToBeSaved {
		if existingChange.primary == key {
			bindingType = KeyBindingTypePrimary
		} else if existingChange.secondary == key {
			bindingType = KeyBindingTypeSecondary
		}

		if bindingType != -1 {
			existingBinding = existingChange.target
			gameEvent = ge

			return existingChange, existingBinding, gameEvent, bindingType
		}
	}

	return nil, nil, -1, KeyBindingTypeNone
}

func (menu *KeyBindingMenu) saveKeyChange(key enum.Key) error {
	changeExisting, existingBinding, gameEvent, bindingType := menu.getPendingChangeByKey(key)

	if changeExisting == nil {
		existingBinding, gameEvent, bindingType = menu.keyMap.GetBindingByKey(key)
	}

	if existingBinding != nil && changeExisting == nil {
		changeExisting = &bindingChange{
			target:    existingBinding,
			primary:   existingBinding.Primary,
			secondary: existingBinding.Secondary,
		}

		menu.changesToBeSaved[gameEvent] = changeExisting
	}

	changeCurrent := menu.changesToBeSaved[menu.currentBindingLayout.gameEvent]
	if changeCurrent == nil {
		changeCurrent = &bindingChange{
			target:    menu.currentBindingLayout.binding,
			primary:   menu.currentBindingLayout.binding.Primary,
			secondary: menu.currentBindingLayout.binding.Secondary,
		}

		menu.changesToBeSaved[menu.currentBindingLayout.gameEvent] = changeCurrent
	}

	switch menu.currentBindingModifierType {
	case KeyBindingTypePrimary:
		changeCurrent.primary = key
	case KeyBindingTypeSecondary:
		changeCurrent.secondary = key
	}

	if changeExisting != nil {
		if bindingType == KeyBindingTypePrimary {
			changeExisting.primary = -1
		}

		if bindingType == KeyBindingTypeSecondary {
			changeExisting.secondary = -1
		}
	}

	if err := menu.setBindingLabels(
		changeCurrent.primary,
		changeCurrent.secondary,
		menu.currentBindingLayout,
	); err != nil {
		return err
	}

	if changeExisting != nil {
		for _, bindingLayout := range menu.bindingLayouts {
			if bindingLayout.binding == changeExisting.target {
				if err := menu.setBindingLabels(changeExisting.primary, changeExisting.secondary, bindingLayout); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (menu *KeyBindingMenu) setBindingLabels(primary, secondary enum.Key, bl *bindingLayout) error {
	noneStr := menu.keyMap.KeyToString(-1)

	if primary != -1 {
		if err := bl.SetPrimaryBindingTextAndColor(menu.keyMap.KeyToString(primary), d2util.Color(d2gui.ColorBrown)); err != nil {
			return err
		}
	} else {
		if err := bl.SetPrimaryBindingTextAndColor(noneStr, d2util.Color(d2gui.ColorGrey)); err != nil {
			return err
		}
	}

	if secondary != -1 {
		if err := bl.SetSecondaryBindingTextAndColor(menu.keyMap.KeyToString(secondary), d2util.Color(d2gui.ColorBrown)); err != nil {
			return err
		}
	} else {
		if err := bl.SetSecondaryBindingTextAndColor(noneStr, d2util.Color(d2gui.ColorGrey)); err != nil {
			return err
		}
	}

	if primary == -1 && secondary == -1 {
		if err := bl.primaryLabel.SetColor(d2util.Color(d2gui.ColorRed)); err != nil {
			return err
		}

		if err := bl.secondaryLabel.SetColor(d2util.Color(d2gui.ColorRed)); err != nil {
			return err
		}
	}

	return nil
}

func (menu *KeyBindingMenu) onCancelClicked() error {
	for gameEvent := range menu.changesToBeSaved {
		for _, bindingLayout := range menu.bindingLayouts {
			if bindingLayout.gameEvent == gameEvent {
				if err := menu.setBindingLabels(bindingLayout.binding.Primary, bindingLayout.binding.Secondary, bindingLayout); err != nil {
					return err
				}
			}
		}
	}

	menu.changesToBeSaved = make(map[enum.GameEvent]*bindingChange)
	if menu.currentBindingLayout != nil {
		if err := menu.clearSelection(); err != nil {
			return err
		}
	}

	if err := menu.Close(); err != nil {
		return err
	}

	menu.escapeMenu.showLayout(optionsLayoutID)

	return nil
}

func (menu *KeyBindingMenu) reload() error {
	for _, bl := range menu.bindingLayouts {
		if bl.binding != nil {
			if err := menu.setBindingLabels(bl.binding.Primary, bl.binding.Secondary, bl); err != nil {
				return err
			}
		}
	}

	return nil
}

func (menu *KeyBindingMenu) clearSelection() error {
	if menu.currentBindingLayout != nil {
		if err := menu.currentBindingLayout.Reset(); err != nil {
			return err
		}

		menu.lastBindingLayout = menu.currentBindingLayout
		menu.currentBindingLayout = nil
		menu.currentBindingModifier = -1
		menu.currentBindingModifierType = -1
	}

	return nil
}

func (menu *KeyBindingMenu) onDefaultClicked() error {
	menu.keyMap.ResetToDefault()

	if err := menu.reload(); err != nil {
		return err
	}

	menu.changesToBeSaved = make(map[enum.GameEvent]*bindingChange)

	return menu.clearSelection()
}

func (menu *KeyBindingMenu) onAcceptClicked() error {
	for gameEvent, change := range menu.changesToBeSaved {
		menu.keyMap.SetPrimaryBinding(gameEvent, change.primary)
		menu.keyMap.SetSecondaryBinding(gameEvent, change.secondary)
	}

	menu.changesToBeSaved = make(map[enum.GameEvent]*bindingChange)

	return menu.clearSelection()
}

// OnKeyDown will assign the new key to the selected binding if any
func (menu *KeyBindingMenu) OnKeyDown(event d2interface.KeyEvent) error {
	if menu.isAwaitingKeyDown {
		key := event.Key()

		if key == enum.KeyEscape {
			if menu.currentBindingLayout != nil {
				menu.lastBindingLayout = menu.currentBindingLayout

				if err := menu.currentBindingLayout.Reset(); err != nil {
					return err
				}

				if err := menu.clearSelection(); err != nil {
					return err
				}
			}
		} else {
			if err := menu.saveKeyChange(key); err != nil {
				return err
			}
		}

		menu.isAwaitingKeyDown = false
	}

	return nil
}

// Advance computes the state of the elements of the menu overtime
func (menu *KeyBindingMenu) Advance(elapsed float64) error {
	if menu.scrollbar != nil {
		if err := menu.scrollbar.Advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}

// Render draws the different element of the menu on the target surface
func (menu *KeyBindingMenu) Render(target d2interface.Surface) error {
	if menu.IsOpen() {
		if err := menu.Box.Render(target); err != nil {
			return err
		}

		if menu.scrollbar != nil {
			menu.scrollbar.Render(target)
		}

		if menu.currentBindingLayout != nil {
			x, y := menu.currentBindingLayout.wrapperLayout.Sx, menu.currentBindingLayout.wrapperLayout.Sy
			w, h := menu.currentBindingLayout.wrapperLayout.GetSize()

			target.PushTranslation(x, y)
			target.DrawRect(w, h, d2util.Color(selectionBackgroundColor))
			target.Pop()
		}
	}

	return nil
}
