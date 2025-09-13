package d2records

import (
	"nostos/common/d2fileformats/d2txt"

	"nostos/common/enum"
)

// LoadWeapons loads weapon records
func weaponsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCommonItems(d, enum.InventoryItemTypeWeapon)
	if err != nil {
		return err
	}

	r.Debugf("Loaded %d Weapon records", len(records))

	r.Item.Weapons = records

	return nil
}
