package d2inventory

import "nostos/common/enum"

// InventoryItem defines the functionality of an inventory item
type InventoryItem interface {
	// GetInventoryItemName returns the name of this inventory item
	GetInventoryItemName() string
	// GetInventoryItemType returns the type of item this is
	GetInventoryItemType() enum.InventoryItemType
	// GetInventoryGridSize returns the width/height grid size of this inventory item
	GetInventoryGridSize() (int, int)
	// Returns the item code
	GetItemCode() string
	// Serializes the object for transport
	Serialize() []byte
}
