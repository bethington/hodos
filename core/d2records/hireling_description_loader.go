package d2records

import (
	d2txt "nostos/common/fileformats/txt"
)

func hirelingDescriptionLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(HirelingDescriptions)

	for d.Next() {
		record := &HirelingDescriptionRecord{
			Name:  d.String("Hireling Description"),
			Token: d.String("Code"),
		}

		records[record.Name] = record
	}

	if d.Err != nil {
		panic(d.Err)
	}

	r.Hireling.Descriptions = records

	r.Debugf("Loaded %d HirelingDescription records", len(records))

	return nil
}
