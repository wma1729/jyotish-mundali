package analysis

import (
	"jyotish/misc"
	"math"
)

type GrahaAttributes struct {
	Name                  string
	NaturalFriends        []string
	NaturalNeutrals       []string
	NaturalEnemies        []string
	TemporaryFriends      []string
	TemporaryEnemies      []string
	EffectiveBestFriends  []string
	EffectiveFriends      []string
	EffectiveNeutrals     []string
	EffectiveEnemies      []string
	EffectiveWorstEnemies []string
	NaturalNature         string
	Retrograde            bool
	Combust               bool
	Position              string
}

func addGrahasInNthBhava(grahas map[string]bool, name string, c, n int, chart *Chart) {
	bhava := chart.GetNthBhava(c, n)
	for _, g := range bhava.Grahas {
		if g.Name != RAHU && g.Name != KETU && g.Name != LAGNA && g.Name != name {
			grahas[g.Name] = true
		}
	}
}

func (attr *GrahaAttributes) getNaturalRelations(name string) {
	attr.NaturalFriends = GrahaAttrMap[name].Friends
	attr.NaturalNeutrals = GrahaAttrMap[name].Neutrals
	attr.NaturalEnemies = GrahaAttrMap[name].Enemies
}

func (attr *GrahaAttributes) getTemporalRelations(name string, chart *Chart) {
	friends := make(map[string]bool)
	enemies := make(map[string]bool)

	i, _ := chart.GetGrahaBhava(name)
	addGrahasInNthBhava(enemies, name, i, 1, chart)
	addGrahasInNthBhava(friends, name, i, 2, chart)
	addGrahasInNthBhava(friends, name, i, 3, chart)
	addGrahasInNthBhava(friends, name, i, 4, chart)
	addGrahasInNthBhava(enemies, name, i, 5, chart)
	addGrahasInNthBhava(enemies, name, i, 6, chart)
	addGrahasInNthBhava(enemies, name, i, 7, chart)
	addGrahasInNthBhava(enemies, name, i, 8, chart)
	addGrahasInNthBhava(enemies, name, i, 9, chart)
	addGrahasInNthBhava(friends, name, i, 10, chart)
	addGrahasInNthBhava(friends, name, i, 11, chart)
	addGrahasInNthBhava(friends, name, i, 12, chart)

	for key := range friends {
		attr.TemporaryFriends = append(attr.TemporaryFriends, key)
	}

	for key := range enemies {
		attr.TemporaryEnemies = append(attr.TemporaryEnemies, key)
	}
}

func (attr *GrahaAttributes) getEffectiveRelations(name string, chart *Chart) {
	for _, g := range attr.NaturalFriends {
		if misc.StringSliceContains(attr.TemporaryFriends, g) {
			attr.EffectiveBestFriends = append(attr.EffectiveBestFriends, g)
		} else if misc.StringSliceContains(attr.TemporaryEnemies, g) {
			attr.EffectiveNeutrals = append(attr.EffectiveNeutrals, g)
		}
	}

	for _, g := range attr.NaturalNeutrals {
		if misc.StringSliceContains(attr.TemporaryFriends, g) {
			attr.EffectiveFriends = append(attr.EffectiveFriends, g)
		} else if misc.StringSliceContains(attr.TemporaryEnemies, g) {
			attr.EffectiveEnemies = append(attr.EffectiveEnemies, g)
		}
	}

	for _, g := range attr.NaturalEnemies {
		if misc.StringSliceContains(attr.TemporaryFriends, g) {
			attr.EffectiveNeutrals = append(attr.EffectiveNeutrals, g)
		} else if misc.StringSliceContains(attr.TemporaryEnemies, g) {
			attr.EffectiveWorstEnemies = append(attr.EffectiveWorstEnemies, g)
		}
	}
}

func (attr *GrahaAttributes) getNaturalNature(name string, chart *Chart) {
	switch name {
	case SUN:
		attr.NaturalNature = MALEFIC

	case MOON:
		attr.NaturalNature = MALEFIC

		i, b := chart.GetGrahaBhava(SUN)
		sun := b.GrahaByName(SUN)

		if chart.NthBhavaContainsGraha(i, 1, MOON) {
			moon := chart.GetNthBhava(i, 1).GrahaByName(MOON)
			if moon.Degree > sun.Degree {
				attr.NaturalNature = BENEFIC
			}
		} else if chart.NthBhavaContainsGraha(i, 7, MOON) {
			moon := chart.GetNthBhava(i, 7).GrahaByName(MOON)
			if moon.Degree < sun.Degree {
				attr.NaturalNature = BENEFIC
			}
		} else {
			for n := 2; n <= 6; n++ {
				if chart.NthBhavaContainsGraha(i, n, MOON) {
					attr.NaturalNature = BENEFIC
					break
				}
			}
		}

	case MARS:
		attr.NaturalNature = MALEFIC

	case MERCURY:
		_, b := chart.GetGrahaBhava(MERCURY)
		if b.ContainsGraha(SUN) || b.ContainsGraha(MARS) || b.ContainsGraha(SATURN) ||
			b.ContainsGraha(RAHU) || b.ContainsGraha(KETU) {
			attr.NaturalNature = MALEFIC
		} else {
			attr.NaturalNature = BENEFIC
		}

	case JUPITER:
		attr.NaturalNature = BENEFIC

	case VENUS:
		attr.NaturalNature = BENEFIC

	case SATURN:
		attr.NaturalNature = MALEFIC

	case RAHU:
		attr.NaturalNature = MALEFIC

	case KETU:
		attr.NaturalNature = MALEFIC
	}
}

func (attr *GrahaAttributes) isCombust(name string, chart *Chart) {
	if name == SUN {
		attr.Combust = false
		return
	}

	_, b := chart.GetGrahaBhava(name)
	if !b.ContainsGraha(SUN) {
		attr.Combust = false
		return
	}

	attr.Combust = false
	inputGraha := b.GrahaByName(name)
	sun := b.GrahaByName(SUN)

	difference := math.Abs(float64(inputGraha.Degree - sun.Degree))

	switch inputGraha.Name {
	case MERCURY:
		if inputGraha.Retrograde {
			if difference <= 12.0 {
				attr.Combust = true
			}
		} else if difference <= 14.0 {
			attr.Combust = true
		}

	case VENUS:
		if inputGraha.Retrograde {
			if difference <= 8.0 {
				attr.Combust = true
			}
		} else if difference <= 10.0 {
			attr.Combust = true
		}

	case MARS:
		if difference <= 17.0 {
			attr.Combust = true
		}

	case JUPITER:
		if difference <= 11.0 {
			attr.Combust = true
		}

	case SATURN:
		if difference <= 15.0 {
			attr.Combust = true
		}
	}
}

func (attr *GrahaAttributes) GetGrahaPosition(name string, chart *Chart) {
	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	g := b.GrahaByName(name)
	if g == nil {
		return
	}

	ga := GrahaAttrMap[name]

	// Accurate (take degrees in account)
	if g.RashiNum == ga.Exaltation.RashiNum &&
		g.Degree >= float32(ga.Exaltation.Min) &&
		g.Degree <= float32(ga.Exaltation.Max) {
		attr.Position = RASHI_EXALTED
	} else if g.RashiNum == ga.Debilitation.RashiNum &&
		g.Degree >= float32(ga.Debilitation.Min) &&
		g.Degree <= float32(ga.Debilitation.Max) {
		attr.Position = RASHI_DEBILITATED
	} else if g.RashiNum == ga.Trinal.RashiNum &&
		g.Degree >= float32(ga.Trinal.Min) &&
		g.Degree <= float32(ga.Trinal.Max) {
		attr.Position = RASHI_MOOLTRIKONA
	}

	// Rough (just look at the rashi as a whole)
	if attr.Position == "" {
		if g.RashiNum == ga.Exaltation.RashiNum {
			attr.Position = RASHI_EXALTED
		} else if g.RashiNum == ga.Debilitation.RashiNum {
			attr.Position = RASHI_DEBILITATED
		} else if misc.IntSliceContains(ga.Owner, g.RashiNum) {
			attr.Position = RASHI_OWN
		}
	}

	// Friendly/neutral/unfriendly
	if attr.Position == "" {
		rashiLord := RashiLordMap[g.RashiNum]
		if misc.StringSliceContains(chart.GetEffectiveFriends(rashiLord), name) {
			attr.Position = RASHI_FRIENDLY
		} else if misc.StringSliceContains(chart.GetEffectiveEnemies(rashiLord), name) {
			attr.Position = RASHI_ENEMY
		} else {
			attr.Position = RASHI_NEUTRAL
		}
	}
}

func (attr *GrahaAttributes) Init(name string, chart *Chart) {
	attr.Name = name
	attr.getNaturalRelations(name)
	attr.getTemporalRelations(name, chart)
	attr.getEffectiveRelations(name, chart)
	attr.getNaturalNature(name, chart)
	attr.isCombust(name, chart)
	attr.Retrograde = false
	_, b := chart.GetGrahaBhava(name)
	if b != nil {
		attr.Retrograde = b.IsRetrograde(name)
	}
}
