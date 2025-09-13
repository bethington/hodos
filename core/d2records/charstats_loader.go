package d2records

import (
	"nostos/common/enum"
	"nostos/common/d2fileformats/d2txt"
)

// nolint:funlen // cant reduce
func charStatsLoader(r *RecordManager, d *d2txt.DataDictionary) error {
	records := make(CharStats)

	stringMap := map[string]enum.Hero{
		"Amazon":      enum.HeroAmazon,
		"Barbarian":   enum.HeroBarbarian,
		"Druid":       enum.HeroDruid,
		"Assassin":    enum.HeroAssassin,
		"Necromancer": enum.HeroNecromancer,
		"Paladin":     enum.HeroPaladin,
		"Sorceress":   enum.HeroSorceress,
	}

	tokenMap := map[string]enum.WeaponClass{
		"":    enum.WeaponClassNone,
		"hth": enum.WeaponClassHandToHand,
		"bow": enum.WeaponClassBow,
		"1hs": enum.WeaponClassOneHandSwing,
		"1ht": enum.WeaponClassOneHandThrust,
		"stf": enum.WeaponClassStaff,
		"2hs": enum.WeaponClassTwoHandSwing,
		"2ht": enum.WeaponClassTwoHandThrust,
		"xbw": enum.WeaponClassCrossbow,
		"1js": enum.WeaponClassLeftJabRightSwing,
		"1jt": enum.WeaponClassLeftJabRightThrust,
		"1ss": enum.WeaponClassLeftSwingRightSwing,
		"1st": enum.WeaponClassLeftSwingRightThrust,
		"ht1": enum.WeaponClassOneHandToHand,
		"ht2": enum.WeaponClassTwoHandToHand,
	}

	for d.Next() {
		record := &CharStatRecord{
			Class: stringMap[d.String("class")],

			InitStr:     d.Number("str"),
			InitDex:     d.Number("dex"),
			InitVit:     d.Number("vit"),
			InitEne:     d.Number("int"),
			InitStamina: d.Number("stamina"),

			ManaRegen:   d.Number("ManaRegen"),
			ToHitFactor: d.Number("ToHitFactor"),

			VelocityWalk:    d.Number("WalkVelocity"),
			VelocityRun:     d.Number("RunVelocity"),
			StaminaRunDrain: d.Number("RunDrain"),

			LifePerLevel:    d.Number("LifePerLevel"),
			ManaPerLevel:    d.Number("ManaPerLevel"),
			StaminaPerLevel: d.Number("StaminaPerLevel"),

			LifePerVit:    d.Number("LifePerVitality"),
			ManaPerEne:    d.Number("ManaPerMagic"),
			StaminaPerVit: d.Number("StaminaPerVitality"),

			StatPerLevel: d.Number("StatPerLevel"),
			BlockFactor:  d.Number("BlockFactor"),

			StartSkillBonus:   d.String("StartSkill"),
			SkillStrAll:       d.String("StrAllSkills"),
			SkillStrClassOnly: d.String("StrClassOnly"),

			BaseSkill: [10]string{
				d.String("Skill 1"),
				d.String("Skill 2"),
				d.String("Skill 3"),
				d.String("Skill 4"),
				d.String("Skill 5"),
				d.String("Skill 6"),
				d.String("Skill 7"),
				d.String("Skill 8"),
				d.String("Skill 9"),
				d.String("Skill 10"),
			},

			SkillStrTab: [3]string{
				d.String("StrSkillTab1"),
				d.String("StrSkillTab2"),
				d.String("StrSkillTab3"),
			},

			BaseWeaponClass: tokenMap[d.String("baseWClass")],

			StartItem: [10]string{
				d.String("item1"),
				d.String("item2"),
				d.String("item3"),
				d.String("item4"),
				d.String("item5"),
				d.String("item6"),
				d.String("item7"),
				d.String("item8"),
				d.String("item9"),
				d.String("item10"),
			},

			StartItemLocation: [10]string{
				d.String("item1loc"),
				d.String("item2loc"),
				d.String("item3loc"),
				d.String("item4loc"),
				d.String("item5loc"),
				d.String("item6loc"),
				d.String("item7loc"),
				d.String("item8loc"),
				d.String("item9loc"),
				d.String("item10loc"),
			},

			StartItemCount: [10]int{
				d.Number("item1count"),
				d.Number("item2count"),
				d.Number("item3count"),
				d.Number("item4count"),
				d.Number("item5count"),
				d.Number("item6count"),
				d.Number("item7count"),
				d.Number("item8count"),
				d.Number("item9count"),
				d.Number("item10count"),
			},
		}
		records[record.Class] = record
	}

	if d.Err != nil {
		return d.Err
	}

	r.Debugf("Loaded %d CharStat records", len(records))

	r.Character.Stats = records

	return nil
}
