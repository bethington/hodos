package d2records

import (
	"nostos/common/enum"
	"nostos/common/d2fileformats/d2txt"
)

/*	first column of experience.txt
	Level
	Amazon
	Sorceress
	Necromancer
	Paladin
	Barbarian
	Druid
	Assassin
	ExpRatio

	second row is special case, shows max levels

	MaxLvl
	99
	99
	99
	99
	99
	99
	99
	10

	the rest are the breakpoints records
*/

func experienceLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	breakpoints := make(ExperienceBreakpoints)

	d.Next() // move to the first row, the max level data

	// parse the max level data
	maxLevels := ExperienceMaxLevels{
		enum.HeroAmazon:      d.Number("Amazon"),
		enum.HeroBarbarian:   d.Number("Barbarian"),
		enum.HeroDruid:       d.Number("Druid"),
		enum.HeroAssassin:    d.Number("Assassin"),
		enum.HeroNecromancer: d.Number("Necromancer"),
		enum.HeroPaladin:     d.Number("Paladin"),
		enum.HeroSorceress:   d.Number("Sorceress"),
	}

	for d.Next() {
		record := &ExperienceBreakpointRecord{
			Level: d.Number("Level"),
			HeroBreakpoints: map[enum.Hero]int{
				enum.HeroAmazon:      d.Number("Amazon"),
				enum.HeroBarbarian:   d.Number("Barbarian"),
				enum.HeroDruid:       d.Number("Druid"),
				enum.HeroAssassin:    d.Number("Assassin"),
				enum.HeroNecromancer: d.Number("Necromancer"),
				enum.HeroPaladin:     d.Number("Paladin"),
				enum.HeroSorceress:   d.Number("Sorceress"),
			},
			Ratio: d.Number("ExpRatio"),
		}
		breakpoints[record.Level] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Debugf("Loaded %d ExperienceBreakpoint records", len(breakpoints))

	r.Character.MaxLevel = maxLevels
	r.Character.Experience = breakpoints

	return nil
}
