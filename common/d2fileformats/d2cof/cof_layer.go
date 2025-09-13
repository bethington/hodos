package d2cof

import "nostos/common/enum"

// CofLayer is a structure that represents a single layer in a COF file.
type CofLayer struct {
	Type        enum.CompositeType
	Shadow      byte
	Selectable  bool
	Transparent bool
	DrawEffect  enum.DrawEffect
	WeaponClass enum.WeaponClass
}
