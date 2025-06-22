package analysis

import "jyotish/misc"

type AscpectAndDegree struct {
	Name   string
	Degree int
}

type GrahaAspects struct {
	Name     string
	Friends  []AscpectAndDegree
	Neutrals []AscpectAndDegree
	Enemies  []AscpectAndDegree
}

func (aspects *GrahaAspects) findAspectualStrength(grahaAttr *GrahaAttributes, aspectingGrahas []string, degree int) {
	for _, graha := range aspectingGrahas {
		if misc.StringSliceContains(grahaAttr.Relations.NaturalFriends, graha) {
			aspects.Friends = append(aspects.Friends, AscpectAndDegree{graha, degree})
		} else if misc.StringSliceContains(grahaAttr.Relations.NaturalEnemies, graha) {
			aspects.Enemies = append(aspects.Enemies, AscpectAndDegree{graha, degree})
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
		return
	}

	aspects.Friends = make([]AscpectAndDegree, 0)
	aspects.Neutrals = make([]AscpectAndDegree, 0)
	aspects.Enemies = make([]AscpectAndDegree, 0)

	aspects.findAspectualStrength(ga, b.FullAspect, 100)
	aspects.findAspectualStrength(ga, b.ThreeQuarterAspect, 75)
	aspects.findAspectualStrength(ga, b.HalfAspect, 50)
	aspects.findAspectualStrength(ga, b.QuarterAspect, 25)

}
