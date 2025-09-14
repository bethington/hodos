package gamedata

import (
	d2txt "nostos/common/fileformats/txt"

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
