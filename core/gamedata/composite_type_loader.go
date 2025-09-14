package gamedata

import (
	d2txt "nostos/common/fileformats/txt"
)

func compositeTypeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(CompositeTypes)

	for d.Next() {
		record := &CompositeTypeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Animation.Token.Composite = records

	r.Debugf("Loaded %d CompositeType records", len(records))

	return nil
}
