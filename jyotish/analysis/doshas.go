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

type KalaSarpaDosha struct {
	Present bool
	Type    int
	Axis    int
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

func (c *Chart) EvaluateKalaSarpaDosha() *KalaSarpaDosha {
	var rahuDegree, ketuDegree float64
	for _, ga := range c.GrahasAttr {
		switch ga.Name {
		case constants.RAHU:
			rahuDegree = ga.AbsoluteDegree
		case constants.KETU:
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

	count := 0

	for _, ga := range c.GrahasAttr {
		if ga.Name == constants.RAHU || ga.Name == constants.KETU {
			continue
		}

		if ga.AbsoluteDegree >= lower && ga.AbsoluteDegree <= upper {
			count++
		}
	}

	var dosha = new(KalaSarpaDosha)

	dosha.Present = count == 0 || count == 7
	if dosha.Present {
		if count == 7 {
			if upper == rahuDegree {
				dosha.Type = constants.BITING
			} else {
				dosha.Type = constants.STINGING
			}
		} else {
			if upper == rahuDegree {
				dosha.Type = constants.STINGING
			} else {
				dosha.Type = constants.BITING
			}
		}

		_, b := c.GetGrahaBhava(constants.RAHU)
		if b != nil {
			switch b.Number {
			case 6, 12:
				dosha.Axis = 6
			case 2, 8:
				dosha.Axis = 8
			}
		}
	}

	return dosha
}
