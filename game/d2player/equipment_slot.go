package d2player

import (
	"nostos/common/enum"
	"nostos/core/d2records"
)

// EquipmentSlot represents an equipment slot for a player
type EquipmentSlot struct {
	item   InventoryItem
	x      int
	y      int
	width  int
	height int
}

func genEquipmentSlotsMap(record *d2records.InventoryRecord) map[enum.EquippedSlot]EquipmentSlot {
	slotMap := map[enum.EquippedSlot]EquipmentSlot{}

	slots := []enum.EquippedSlot{
		enum.EquippedSlotHead,
		enum.EquippedSlotTorso,
		enum.EquippedSlotLegs,
		enum.EquippedSlotRightArm,
		enum.EquippedSlotLeftArm,
		enum.EquippedSlotLeftHand,
		enum.EquippedSlotRightHand,
		enum.EquippedSlotNeck,
		enum.EquippedSlotBelt,
		enum.EquippedSlotGloves,
	}

	for _, slot := range slots {
		box := record.Slots[slot]
		equipmentSlot := EquipmentSlot{
			nil,
			box.Left,
			box.Bottom + cellPadding,
			box.Width,
			box.Height,
		}
		slotMap[slot] = equipmentSlot
	}

	return slotMap
}
