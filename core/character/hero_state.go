package character

import (
	"nostos/common/enum"
	d2inventory "nostos/core/equipment"
)

// HeroState stores the state of the player
type HeroState struct {
	HeroName   string                         `json:"heroName"`
	HeroType   enum.Hero                    `json:"heroType"`
	Act        int                            `json:"act"`
	FilePath   string                         `json:"-"`
	Equipment  d2inventory.CharacterEquipment `json:"equipment"`
	Stats      *HeroStatsState                `json:"stats"`
	Skills     map[int]*HeroSkill             `json:"skills"`
	X          float64                        `json:"x"`
	Y          float64                        `json:"y"`
	LeftSkill  int                            `json:"leftSkill"`
	RightSkill int                            `json:"rightSkill"`
	Gold       int                            `json:"Gold"`
	Difficulty enum.DifficultyType          `json:"difficulty"`
}
