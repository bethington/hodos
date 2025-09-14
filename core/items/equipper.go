package items

// Equipper is an interface for something that can equip items
type Equipper interface {
	EquippedItems() []Item
	CarriedItems() []Item
}
