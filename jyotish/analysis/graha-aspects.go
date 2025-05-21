package analysis

import "jyotish/misc"

type GrahaAspects struct {
	Name               string
	FullAspect         []string
	ThreeQuarterAspect []string
	HalfAspect         []string
	QuarterAspect      []string
	Strength           float64
}

func (aspects *GrahaAspects) EvaluateGrahaAspects(name string, chart *Chart) {
	aspects.Name = name
	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	aspects.FullAspect = b.FullAspect
	aspects.ThreeQuarterAspect = b.ThreeQuarterAspect
	aspects.HalfAspect = b.HalfAspect
	aspects.QuarterAspect = b.QuarterAspect

	ga := chart.GetGrahaAttributes(name)
	if ga == nil {
		aspects.Strength = 0.0
		return
	}

	friendlyGrahas := append(ga.Relations.EffectiveBestFriends, ga.Relations.EffectiveFriends...)
	enemyGrahas := append(ga.Relations.EffectiveWorstEnemies, ga.Relations.EffectiveEnemies...)

	for _, graha := range aspects.FullAspect {
		if misc.StringSliceContains(friendlyGrahas, graha) {
			aspects.Strength += 1.0
		} else if misc.StringSliceContains(enemyGrahas, graha) {
			aspects.Strength += 0.5
		}
	}

	for _, graha := range aspects.ThreeQuarterAspect {
		if misc.StringSliceContains(friendlyGrahas, graha) {
			aspects.Strength += 0.875
		} else if misc.StringSliceContains(enemyGrahas, graha) {
			aspects.Strength += 0.375
		}
	}

	for _, graha := range aspects.HalfAspect {
		if misc.StringSliceContains(friendlyGrahas, graha) {
			aspects.Strength += 0.75
		} else if misc.StringSliceContains(enemyGrahas, graha) {
			aspects.Strength += 0.25
		}
	}

	for _, graha := range aspects.QuarterAspect {
		if misc.StringSliceContains(friendlyGrahas, graha) {
			aspects.Strength += 0.625
		} else if misc.StringSliceContains(enemyGrahas, graha) {
			aspects.Strength += 0.125
		}
	}

	aspects.Strength = misc.RoundFloat(aspects.Strength/8, 2)
}
