package gamedata

import (
	d2txt "nostos/common/fileformats/txt"
)

func objectModesLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(ObjectModes)

	for d.Next() {
		record := &ObjectModeRecord{
			Name:  d.String("Name"),
			Token: d.String("Token"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Object.Modes = records

	r.Debugf("Loaded %d ObjectMode records", len(records))

	return nil
}
