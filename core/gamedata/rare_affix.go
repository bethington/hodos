package gamedata

// RareItemAffix described both rare prefixes and suffixes
type RareItemAffix struct {
	Name          string
	IncludedTypes []string
	ExcludedTypes []string
}
