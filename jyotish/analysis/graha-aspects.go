package analysis

import "jyotish/misc"

type AscpectAndDegree struct {
	Name   string
	Degree int
}

type GrahaAspects struct {
	Name     string
	Friends  []string
	Neutrals []string
	Enemies  []string
	Strength int
}

func (aspects *GrahaAspects) findAspectualStrength(grahaAttr *GrahaAttributes, aspectingGrahas []string) {
	for _, graha := range aspectingGrahas {
		if misc.StringSliceContains(grahaAttr.Relations.NaturalFriends, graha) {
			aspects.Friends = append(aspects.Friends, graha)
			aspects.Strength += 1
		} else if misc.StringSliceContains(grahaAttr.Relations.NaturalEnemies, graha) {
			aspects.Enemies = append(aspects.Enemies, graha)
			aspects.Strength -= 1
		} else {
			aspects.Neutrals = append(aspects.Neutrals, graha)
		}
	}

}

// Uses effective relations to determine aspect strength
func (aspects *GrahaAspects) EvaluateGrahaAspects(name string, chart *Chart) {
	aspects.Name = name
	aspects.Strength = 0.0
	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	ga := chart.GetGrahaAttributes(name)
	if ga == nil {
		return
	}

	aspects.Friends = make([]string, 0)
	aspects.Neutrals = make([]string, 0)
	aspects.Enemies = make([]string, 0)

	aspects.findAspectualStrength(ga, b.FullAspect)
}
