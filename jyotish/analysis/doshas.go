package analysis

import (
	"jyotish/constants"
)

type MarsPosition struct {
	BhavaNumber int
	Reference   string
	Strength    int
}

type MaanglikDosha struct {
	Present               bool
	Position              []MarsPosition
	MarsInOwnRashi        bool
	MarsAspectingOwnRashi bool
	ConjunctWithMoon      bool
	AspectedByMoon        bool
}

func isMaanglikDoshaFormed(distance int) bool {
	switch distance {
	case 1, 2, 4, 7, 8, 12:
		return true
	default:
		return false
	}
}

func (c *Chart) EvaluateMaanglikDosha() *MaanglikDosha {
	_, marsBhava := c.GetGrahaBhava(constants.MARS)
	if marsBhava == nil {
		return nil
	}

	// Distance from lagna
	distanceFromLagna := marsBhava.Number
	distanceFromMoon := c.GetDistanceBetweenTwoGrahas(constants.MOON, constants.MARS)
	distanceFromVenus := c.GetDistanceBetweenTwoGrahas(constants.VENUS, constants.MARS)

	isDoshaFormedFromLagna := isMaanglikDoshaFormed(marsBhava.Number)
	isDoshaFormedFromMoon := isMaanglikDoshaFormed(distanceFromMoon)
	isDoshaFormedFromVenus := isMaanglikDoshaFormed(distanceFromVenus)

	var dosha *MaanglikDosha = new(MaanglikDosha)

	if isDoshaFormedFromLagna || isDoshaFormedFromMoon || isDoshaFormedFromVenus {
		dosha.Present = true
	} else {
		return dosha
	}

	dosha.Position = make([]MarsPosition, 0)
	if isDoshaFormedFromLagna {
		dosha.Position = append(dosha.Position, MarsPosition{distanceFromLagna, constants.LAGNA, constants.FULL})
	}
	if isDoshaFormedFromMoon {
		dosha.Position = append(dosha.Position, MarsPosition{distanceFromMoon, constants.MOON, constants.HALF})
	}
	if isDoshaFormedFromVenus {
		dosha.Position = append(dosha.Position, MarsPosition{distanceFromVenus, constants.VENUS, constants.QUARTER})
	}

	_, moonBhava := c.GetGrahaBhava(constants.MOON)
	if moonBhava == nil {
		return nil
	}

	dosha.MarsInOwnRashi = marsBhava.RashiLord == constants.MARS
	dosha.ConjunctWithMoon = marsBhava.Number == moonBhava.Number
	dosha.AspectedByMoon = c.IsGrahaAspectedBy(constants.MARS, constants.MOON)

	bhavasAspectedByMars := c.GetBhavasAspectedBy(constants.MARS)
	for _, n := range bhavasAspectedByMars {
		if c.Bhavas[n-1].RashiLord == constants.MARS {
			dosha.MarsAspectingOwnRashi = true
		}
	}

	return dosha
}

/*

func (c *Chart) EvaluateKalaSarpaDosha() {
	var rahuDegree, ketuDegree float64
	for _, ga := range c.GrahasAttr {
		if ga.Name == constants.RAHU {
			rahuDegree = ga.AbsoluteDegree
		} else if ga.Name == constants.KETU {
			ketuDegree = ga.AbsoluteDegree
		}
	}

	var lower, upper float64
	if rahuDegree < ketuDegree {
		lower = rahuDegree
		upper = ketuDegree
	} else {
		lower = ketuDegree
		upper = rahuDegree
	}

	isDoshaFormed := true

	for _, ga := range c.GrahasAttr {
		if ga.Name == constants.RAHU || ga.Name == constants.KETU {
			continue
		}
	}
}

*/
