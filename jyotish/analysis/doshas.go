package analysis

import (
	"jyotish/constants"
)

type MarsPosition struct {
	BhavaNumber int
	Reference   string
}

type MaanglikDosha struct {
	Present               bool
	Position              []MarsPosition
	Severity              int
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
		dosha.Position = append(dosha.Position, MarsPosition{distanceFromLagna, constants.LAGNA})
	}
	if isDoshaFormedFromMoon {
		dosha.Position = append(dosha.Position, MarsPosition{distanceFromMoon, constants.MOON})
	}
	if isDoshaFormedFromVenus {
		dosha.Position = append(dosha.Position, MarsPosition{distanceFromVenus, constants.VENUS})
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

	dosha.Severity = constants.MAXIMUM

	if dosha.ConjunctWithMoon {
		dosha.Severity = constants.MINIMUM
	} else if dosha.AspectedByMoon {
		switch len(dosha.Position) {
		case 1:
			if dosha.MarsInOwnRashi || dosha.MarsAspectingOwnRashi {
				dosha.Severity = constants.MINIMUM
			} else {
				dosha.Severity = constants.AVERAGE
			}
		case 2:
			if dosha.MarsInOwnRashi || dosha.MarsAspectingOwnRashi {
				dosha.Severity = constants.AVERAGE
			}
		}
	}

	return dosha
}
