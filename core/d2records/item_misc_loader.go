package d2records

import (
	"nostos/common/d2fileformats/d2txt"

	"nostos/common/enum"
)

// LoadMiscItems loads ItemCommonRecords from misc.txt
func miscItemsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := loadCommonItems(d, enum.InventoryItemTypeItem)
	if err != nil {
		return err
	}

	r.Debugf("Loaded %d Misc Item records", len(records))

	r.Item.Misc = records

	return nil
}
