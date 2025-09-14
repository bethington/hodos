package gamedata

import (
	d2txt "nostos/common/fileformats/txt"
)

func uniqueAppellationsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(UniqueAppellations)

	for d.Next() {
		record := &UniqueAppellationRecord{
			Name: d.String("Name"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Monster.Unique.Appellations = records

	r.Debugf("Loaded %d UniqueAppellation records", len(records))

	return nil
}
