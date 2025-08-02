package analysis

import (
	"jyotish/constants"
	"jyotish/misc"
)

type GrahaAspects struct {
	Name     string
	Friends  []string
	Neutrals []string
	Enemies  []string
	Strength int
}

func (aspects *GrahaAspects) findAspectualStrength(grahaAttr *GrahaAttributes, aspectedGrahas []string) {
	aspectingGraha := grahaAttr.Name

	for _, aspectedGraha := range aspectedGrahas {
		if (aspectingGraha == constants.RAHU && aspectedGraha == constants.KETU) ||
			(aspectingGraha == constants.KETU && aspectedGraha == constants.RAHU) {
			continue
		}
		if misc.StringSliceContains(grahaAttr.Relations.NaturalFriends, aspectedGraha) {
			aspects.Friends = append(aspects.Friends, aspectedGraha)
			aspects.Strength += 1
		} else if misc.StringSliceContains(grahaAttr.Relations.NaturalEnemies, aspectedGraha) {
			aspects.Enemies = append(aspects.Enemies, aspectedGraha)
			aspects.Strength -= 1
		} else {
			aspects.Neutrals = append(aspects.Neutrals, aspectedGraha)
		}
	}

}

// Uses effective relations to determine aspect strength
func (aspects *GrahaAspects) EvaluateGrahaAspects(name string, chart *Chart) {
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
