package analysis

import (
	"jyotish/constants"
	"jyotish/misc"
)

type GrahaRelations struct {
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
}

// Populate natural relations for the given graha directly from the fixed rules.
// This is same for all charts.
func (relations *GrahaRelations) populateNaturalRelations(name string) {
	relations.NaturalFriends = constants.GrahaBalaInRashiRulesMap[name].Friends
	relations.NaturalNeutrals = constants.GrahaBalaInRashiRulesMap[name].Neutrals
	relations.NaturalEnemies = constants.GrahaBalaInRashiRulesMap[name].Enemies
}

// Add grahas in the nth bhava from ith bhava to a set excluding Rahu, Ketu, Lagna
// (not a graha but treated like one in the implementation), and itself.
func addGrahasInNthBhava(grahas map[string]bool, name string, c, n int, chart *Chart) {
	bhava := chart.GetNthBhava(c, n)
	for _, g := range bhava.Grahas {
		if g.Name != constants.RAHU && g.Name != constants.KETU && g.Name != constants.LAGNA && g.Name != name {
			grahas[g.Name] = true
		}
	}
}

// Populate temporal relations for the given graha in a given chart.
// Grahas in 2nd, 3rd, 4th, 10th, 11th, 12th are friends and the rest are enemies.
func (relations *GrahaRelations) populateTemporalRelations(name string, chart *Chart) {
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
		relations.TemporaryFriends = append(relations.TemporaryFriends, key)
	}

	for key := range enemies {
		relations.TemporaryEnemies = append(relations.TemporaryEnemies, key)
	}
}

// Populate effective relations for the given graha.
//
// natural friend + temporal friend = best friend
//
// natural friend + temporal enemy = netural
//
// natural neutral + temporal friend = friend
//
// natural neutral + temporay enemy = enemy
//
// natural enemy + temporal friend = neutral
//
// natural enemy + temporal enemy = worst enemy
func (relations *GrahaRelations) populateEffectiveRelations(name string) {
	for _, g := range relations.NaturalFriends {
		if misc.StringSliceContains(relations.TemporaryFriends, g) {
			relations.EffectiveBestFriends = append(relations.EffectiveBestFriends, g)
		} else if misc.StringSliceContains(relations.TemporaryEnemies, g) {
			relations.EffectiveNeutrals = append(relations.EffectiveNeutrals, g)
		}
	}

	for _, g := range relations.NaturalNeutrals {
		if misc.StringSliceContains(relations.TemporaryFriends, g) {
			relations.EffectiveFriends = append(relations.EffectiveFriends, g)
		} else if misc.StringSliceContains(relations.TemporaryEnemies, g) {
			relations.EffectiveEnemies = append(relations.EffectiveEnemies, g)
		}
	}

	for _, g := range relations.NaturalEnemies {
		if misc.StringSliceContains(relations.TemporaryFriends, g) {
			relations.EffectiveNeutrals = append(relations.EffectiveNeutrals, g)
		} else if misc.StringSliceContains(relations.TemporaryEnemies, g) {
			relations.EffectiveWorstEnemies = append(relations.EffectiveWorstEnemies, g)
		}
	}
}

// Evaluate graha relations for the given graha in a given chart.
//
// This finds natural, temporal, and effective relations.
func (relations *GrahaRelations) EvaluateGrahaRelations(name string, chart *Chart) {
	relations.Name = name
	relations.populateNaturalRelations(name)
	relations.populateTemporalRelations(name, chart)
	relations.populateEffectiveRelations(name)
}
