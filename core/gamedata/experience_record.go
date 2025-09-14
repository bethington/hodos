package gamedata

import "nostos/common/enum"

// ExperienceBreakpoints describes the required experience
// for each level for each character class
type ExperienceBreakpoints map[int]*ExperienceBreakpointRecord

// ExperienceMaxLevels defines the max character levels
type ExperienceMaxLevels map[enum.Hero]int

// ExperienceBreakpointRecord describes the experience points required to
// gain a level for all character classes
type ExperienceBreakpointRecord struct {
	Level           int
	HeroBreakpoints map[enum.Hero]int
	Ratio           int
}
