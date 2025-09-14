package gamedata

import d2txt "nostos/common/fileformats/txt"

// recordLoader represents something that can load a data dictionary and
// handles placing it in the record manager exports
type recordLoader func(r *RecordManager, dictionary *d2txt.DataDictionary) error
