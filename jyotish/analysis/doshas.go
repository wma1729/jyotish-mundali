package analysis

import (
	"jyotish/constants"
	"jyotish/misc"
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

type RelativePosition struct {
	Name      string
	InKendra  bool
	Reference string
}

type DebilitatedGraha struct {
	Name                       string
	DebilitatedRashiLord       string
	ExaltedRashiLord           string
	RashiLordRelativePosition  []RelativePosition
	AspectedByExaltedRashiLord bool
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

func (c *Chart) evaluateDebilitatedGraha(ga *GrahaAttributes) DebilitatedGraha {
	dg := DebilitatedGraha{Name: ga.Name}

	_, b := c.GetGrahaBhava(ga.Name)
	if b != nil {
		dg.DebilitatedRashiLord = b.RashiLord
	}

	grahaBala := constants.GrahaBalaInRashiRulesMap[ga.Name]
	dg.ExaltedRashiLord = constants.RashiLordMap[grahaBala.Exaltation.RashiNum]
	dg.RashiLordRelativePosition = make([]RelativePosition, 0)

	var relPos RelativePosition
	var distance int

	distance = c.GetDistanceBetweenTwoGrahas(constants.LAGNA, dg.DebilitatedRashiLord)
	if distance == 1 || distance == 4 || distance == 7 || distance == 10 {
		relPos = RelativePosition{Name: dg.DebilitatedRashiLord, InKendra: true, Reference: constants.LAGNA}
		dg.RashiLordRelativePosition = append(dg.RashiLordRelativePosition, relPos)
	}

	distance = c.GetDistanceBetweenTwoGrahas(constants.MOON, dg.DebilitatedRashiLord)
	if distance == 1 || distance == 4 || distance == 7 || distance == 10 {
		relPos = RelativePosition{Name: dg.DebilitatedRashiLord, InKendra: true, Reference: constants.MOON}
		dg.RashiLordRelativePosition = append(dg.RashiLordRelativePosition, relPos)
	}

	distance = c.GetDistanceBetweenTwoGrahas(constants.LAGNA, dg.ExaltedRashiLord)
	if distance == 1 || distance == 4 || distance == 7 || distance == 10 {
		relPos = RelativePosition{Name: dg.ExaltedRashiLord, InKendra: true, Reference: constants.LAGNA}
		dg.RashiLordRelativePosition = append(dg.RashiLordRelativePosition, relPos)
	}

	distance = c.GetDistanceBetweenTwoGrahas(constants.MOON, dg.ExaltedRashiLord)
	if distance == 1 || distance == 4 || distance == 7 || distance == 10 {
		relPos = RelativePosition{Name: dg.ExaltedRashiLord, InKendra: true, Reference: constants.MOON}
		dg.RashiLordRelativePosition = append(dg.RashiLordRelativePosition, relPos)
	}

	distance = c.GetDistanceBetweenTwoGrahas(dg.DebilitatedRashiLord, dg.ExaltedRashiLord)
	if distance == 1 || distance == 4 || distance == 7 || distance == 10 {
		relPos = RelativePosition{Name: dg.DebilitatedRashiLord, InKendra: true, Reference: dg.ExaltedRashiLord}
		dg.RashiLordRelativePosition = append(dg.RashiLordRelativePosition, relPos)
	}

	var aspects []string = make([]string, 0)
	aspects = append(aspects, ga.Aspects.Friends...)
	aspects = append(aspects, ga.Aspects.Neutrals...)
	aspects = append(aspects, ga.Aspects.Enemies...)

	if misc.StringSliceContains(aspects, dg.ExaltedRashiLord) {
		dg.AspectedByExaltedRashiLord = true
	}

	return dg
}

func (c *Chart) EvaluateDebilitatedGrahas() []DebilitatedGraha {
	debilitatedGrahas := make([]DebilitatedGraha, 0)

	for _, ga := range c.GrahasAttr {
		if ga.Strength.Position == constants.IN_DEBILITATION_RASHI {
			dg := c.evaluateDebilitatedGraha(&ga)
			debilitatedGrahas = append(debilitatedGrahas, dg)
		}
	}

	return debilitatedGrahas
}
