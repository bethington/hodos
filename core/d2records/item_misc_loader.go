package d2records

import (
	"nostos/common/d2fileformats/d2txt"

	"nostos/common/d2enum"
)

// LoadMiscItems loads ItemCommonRecords from misc.txt
func miscItemsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCommonItems(d, d2enum.InventoryItemTypeItem)
	if err != nil {
		return err
	}

	r.Debugf("Loaded %d Misc Item records", len(records))

	r.Item.Misc = records

	return nil
}
