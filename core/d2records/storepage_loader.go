package d2records

import (
	d2txt "nostos/common/fileformats/txt"
)

func storePagesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(StorePages)

	for d.Next() {
		record := &StorePageRecord{
			StorePage: d.String("Store Page"),
			Code:      d.String("Code"),
		}
		records[record.StorePage] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Item.StorePages = records

	r.Debugf("Loaded %d StorePage records", len(records))

	return nil
}
