package gamedata

import (
	d2txt "nostos/common/fileformats/txt"
)

func bodyLocationsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(BodyLocations)

	for d.Next() {
		location := &BodyLocationRecord{
			Name: d.String("Name"),
			Code: d.String("Code"),
		}
		records[location.Code] = location
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Debugf("Loaded %d BodyLocation records", len(records))

	r.BodyLocations = records

	return nil
}
