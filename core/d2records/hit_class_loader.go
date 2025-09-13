package d2records

import (
	d2txt "nostos/common/fileformats/txt"
)

func hitClassLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(HitClasses)

	for d.Next() {
		record := &HitClassRecord{
			Name:  d.String("Hit Class"),
			Token: d.String("Code"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Animation.Token.HitClass = records

	r.Debugf("Loaded %d HitClass records", len(records))

	return nil
}
