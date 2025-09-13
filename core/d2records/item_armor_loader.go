package d2records

import (
	"nostos/common/d2fileformats/d2txt"

	"nostos/common/d2enum"
)

func armorLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	if r.Item.Armors != nil {
		return nil // already loaded
	}

	records, err := loadCommonItems(d, d2enum.InventoryItemTypeArmor)
	if err != nil {
		return err
	}

	r.Debugf("Loaded %d Armor Item records", len(records))

	r.Item.Armors = records

	return nil
}
