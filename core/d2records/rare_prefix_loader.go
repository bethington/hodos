package d2records

import (
	d2txt "nostos/common/fileformats/txt"
)

func rareItemPrefixLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records, err := rareItemAffixLoader(d)
	if err != nil {
		return err
	}

	r.Item.Rare.Prefix = records

	r.Debugf("Loaded %d RarePrefix records", len(records))

	return nil
}
