package d2records

import (
	d2txt "nostos/common/fileformats/txt"
)

func lowQualityLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(LowQualities, 0)

	for d.Next() {
		record := &LowQualityRecord{
			Name: d.String("Hireling Description"),
		}

		records = append(records, record)
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Item.LowQualityPrefixes = records

	r.Debugf("Loaded %d LowQuality records", len(records))

	return nil
}
