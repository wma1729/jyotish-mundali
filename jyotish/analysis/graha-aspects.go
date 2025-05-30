package analysis

import "jyotish/misc"

type AscpectAndDegree struct {
	Name   string
	Degree int
}

type GrahaAspects struct {
	Name         string
	BestFriends  []AscpectAndDegree
	Friends      []AscpectAndDegree
	Neutrals     []AscpectAndDegree
	Enemies      []AscpectAndDegree
	WorstEnemies []AscpectAndDegree
	Strength     float64
}

func (aspects *GrahaAspects) findAspectualStrength(grahaAttr *GrahaAttributes, aspectingGrahas []string, degree int) {
	for _, graha := range aspectingGrahas {
		if misc.StringSliceContains(grahaAttr.Relations.EffectiveBestFriends, graha) {
			aspects.BestFriends = append(aspects.BestFriends, AscpectAndDegree{graha, degree})
			aspects.Strength += float64(degree) * 8.0 / 800.0
		} else if misc.StringSliceContains(grahaAttr.Relations.EffectiveFriends, graha) {
			aspects.Friends = append(aspects.Friends, AscpectAndDegree{graha, degree})
			aspects.Strength += float64(degree) * 7.0 / 800.00
		} else if misc.StringSliceContains(grahaAttr.Relations.EffectiveEnemies, graha) {
			aspects.Enemies = append(aspects.Enemies, AscpectAndDegree{graha, degree})
			aspects.Strength -= float64(degree) * 7.0 / 800.00
		} else if misc.StringSliceContains(grahaAttr.Relations.EffectiveWorstEnemies, graha) {
			aspects.WorstEnemies = append(aspects.WorstEnemies, AscpectAndDegree{graha, degree})
			aspects.Strength -= float64(degree) * 8.0 / 800.00
		} else {
			aspects.Neutrals = append(aspects.Neutrals, AscpectAndDegree{graha, degree})
		}
	}

}

// Uses effective relations to determine aspect strength
func (aspects *GrahaAspects) EvaluateGrahaAspects(name string, chart *Chart) {
	aspects.Name = name
	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	ga := chart.GetGrahaAttributes(name)
	if ga == nil {
		aspects.Strength = 0.0
		return
	}

	aspects.BestFriends = make([]AscpectAndDegree, 0)
	aspects.Friends = make([]AscpectAndDegree, 0)
	aspects.Neutrals = make([]AscpectAndDegree, 0)
	aspects.Enemies = make([]AscpectAndDegree, 0)
	aspects.WorstEnemies = make([]AscpectAndDegree, 0)

	aspects.findAspectualStrength(ga, b.FullAspect, 100)
	aspects.findAspectualStrength(ga, b.ThreeQuarterAspect, 75)
	aspects.findAspectualStrength(ga, b.HalfAspect, 50)
	aspects.findAspectualStrength(ga, b.QuarterAspect, 25)

	aspects.Strength = misc.RoundFloat(aspects.Strength, 2)
}
