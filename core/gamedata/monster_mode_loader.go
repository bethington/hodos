package gamedata

import (
	d2txt "nostos/common/fileformats/txt"
)

// LoadMonModes loads monster records
func monsterModeLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(MonModes)

	for d.Next() {
		record := &MonModeRecord{
			Name:  d.String("name"),
			Token: d.String("token"),
			Code:  d.String("code"),
		}
		records[record.Name] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Debugf("Loaded %d MonMode records", len(records))

	r.Monster.Modes = records

	return nil
}
