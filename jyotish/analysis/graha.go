package analysis

import "math"

type GrahaAttributes struct {
	Name              string
	NaturalFriends    []string
	NaturalNeutrals   []string
	NaturalEnemies    []string
	TemporaryFriends  []string
	TemporaryEnemies  []string
	EffectiveFriends  []string
	EffectiveNeutrals []string
	EffectiveEnemies  []string
	Nature            string
	Combust           bool
}

func has(array []string, element string) bool {
	for _, elem := range array {
		if elem == element {
			return true
		}
	}
	return false
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
		if has(attr.TemporaryFriends, g) {
			attr.EffectiveFriends = append(attr.EffectiveFriends, g)
		} else if has(attr.TemporaryEnemies, g) {
			attr.EffectiveNeutrals = append(attr.EffectiveNeutrals, g)
		}
	}

	for _, g := range attr.NaturalNeutrals {
		if has(attr.TemporaryFriends, g) {
			attr.EffectiveFriends = append(attr.EffectiveFriends, g)
		} else if has(attr.TemporaryEnemies, g) {
			attr.EffectiveEnemies = append(attr.EffectiveEnemies, g)
		}
	}

	for _, g := range attr.NaturalEnemies {
		if has(attr.TemporaryFriends, g) {
			attr.EffectiveNeutrals = append(attr.EffectiveNeutrals, g)
		} else if has(attr.TemporaryEnemies, g) {
			attr.EffectiveEnemies = append(attr.EffectiveEnemies, g)
		}
	}
}

func (attr *GrahaAttributes) getNature(name string, chart *Chart) {
	switch name {
	case SUN:
		attr.Nature = MALEFIC

	case MOON:
		attr.Nature = MALEFIC

		i, b := chart.GetGrahaBhava(SUN)
		sun := b.GrahaByName(SUN)

		if chart.NthBhavaContainsGraha(i, 1, MOON) {
			moon := chart.GetNthBhava(i, 1).GrahaByName(MOON)
			if moon.Degree > sun.Degree {
				attr.Nature = BENEFIC
			}
		} else if chart.NthBhavaContainsGraha(i, 7, MOON) {
			moon := chart.GetNthBhava(i, 7).GrahaByName(MOON)
			if moon.Degree < sun.Degree {
				attr.Nature = BENEFIC
			}
		} else {
			for n := 2; n <= 6; n++ {
				if chart.NthBhavaContainsGraha(i, n, MOON) {
					attr.Nature = BENEFIC
					break
				}
			}
		}

	case MARS:
		attr.Nature = MALEFIC

	case MERCURY:
		_, b := chart.GetGrahaBhava(MERCURY)
		if b.ContainsGraha(SUN) || b.ContainsGraha(MARS) || b.ContainsGraha(SATURN) ||
			b.ContainsGraha(RAHU) || b.ContainsGraha(KETU) {
			attr.Nature = MALEFIC
		} else {
			attr.Nature = BENEFIC
		}

	case JUPITER:
		attr.Nature = BENEFIC

	case VENUS:
		attr.Nature = BENEFIC

	case SATURN:
		attr.Nature = MALEFIC

	case RAHU:
		attr.Nature = MALEFIC

	case KETU:
		attr.Nature = MALEFIC
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

func (attr *GrahaAttributes) Init(name string, chart *Chart) {
	attr.Name = name
	attr.getNaturalRelations(name)
	attr.getTemporalRelations(name, chart)
	attr.getEffectiveRelations(name, chart)
	attr.getNature(name, chart)
	attr.isCombust(name, chart)
}
