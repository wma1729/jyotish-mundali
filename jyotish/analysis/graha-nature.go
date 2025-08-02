package analysis

import (
	"jyotish/constants"
	"math"
)

type GrahaNature struct {
	NaturalNature         int
	FunctionalNature      int
	FunctionalNatureScore int
}

func isMoonBenefic(chart *Chart) bool {
	i, b := chart.GetGrahaBhava(constants.SUN)
	sun := b.GrahaByName(constants.SUN)

	if chart.NthBhavaContainsGraha(i, 1, constants.MOON) {
		moon := chart.GetNthBhava(i, 1).GrahaByName(constants.MOON)
		if moon.Degree > sun.Degree {
			return true
		}
	} else if chart.NthBhavaContainsGraha(i, 7, constants.MOON) {
		moon := chart.GetNthBhava(i, 7).GrahaByName(constants.MOON)
		if moon.Degree < sun.Degree {
			return true
		}
	} else {
		for n := 2; n <= 6; n++ {
			if chart.NthBhavaContainsGraha(i, n, constants.MOON) {
				return true
			}
		}
	}
	return false
}

func isMercuryBenefic(chart *Chart) bool {
	var beneficDistance float64
	var maleficDistance float64
	var distance float64

	_, b := chart.GetGrahaBhava(constants.MERCURY)
	mercuryDegree := b.GrahaDegree(constants.MERCURY, false)
	benevolence := 0

	if b.ContainsGraha(constants.SUN) {
		degree := b.GrahaDegree(constants.SUN, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		maleficDistance = math.Min(maleficDistance, distance)
		benevolence--
	}
	if b.ContainsGraha(constants.MOON) {
		degree := b.GrahaDegree(constants.MOON, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		if isMoonBenefic(chart) {
			maleficDistance = math.Min(maleficDistance, distance)
			benevolence++
		} else {
			beneficDistance = math.Min(beneficDistance, distance)
			benevolence--
		}
	}
	if b.ContainsGraha(constants.MARS) {
		degree := b.GrahaDegree(constants.MARS, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		maleficDistance = math.Min(maleficDistance, distance)
		benevolence--
	}
	if b.ContainsGraha(constants.JUPITER) {
		degree := b.GrahaDegree(constants.JUPITER, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		beneficDistance = math.Min(beneficDistance, distance)
		benevolence++
	}
	if b.ContainsGraha(constants.VENUS) {
		degree := b.GrahaDegree(constants.VENUS, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		beneficDistance = math.Min(beneficDistance, distance)
		benevolence++
	}
	if b.ContainsGraha(constants.SATURN) {
		degree := b.GrahaDegree(constants.SATURN, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		maleficDistance = math.Min(maleficDistance, distance)
		benevolence--
	}
	if b.ContainsGraha(constants.RAHU) {
		degree := b.GrahaDegree(constants.RAHU, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		maleficDistance = math.Min(maleficDistance, distance)
		benevolence--
	}
	if b.ContainsGraha(constants.KETU) {
		degree := b.GrahaDegree(constants.KETU, false)
		distance = math.Abs(float64(degree - mercuryDegree))
		maleficDistance = math.Min(maleficDistance, distance)
		benevolence--
	}

	if benevolence > 0 {
		// More benefic planets
		return true
	} else if benevolence == 0 {
		// consider distance
		if beneficDistance <= maleficDistance {
			return true
		}
	}

	return false
}

func (nature *GrahaNature) findNaturalNature(name string, chart *Chart) {
	switch name {
	case constants.SUN:
		nature.NaturalNature = constants.MALEFIC

	case constants.MOON:
		nature.NaturalNature = constants.MALEFIC
		if isMoonBenefic(chart) {
			nature.NaturalNature = constants.BENEFIC
		}

	case constants.MARS:
		nature.NaturalNature = constants.MALEFIC

	case constants.MERCURY:
		nature.NaturalNature = constants.MALEFIC
		if isMercuryBenefic(chart) {
			nature.NaturalNature = constants.BENEFIC
		}

	case constants.JUPITER:
		nature.NaturalNature = constants.BENEFIC

	case constants.VENUS:
		nature.NaturalNature = constants.BENEFIC

	case constants.SATURN:
		nature.NaturalNature = constants.MALEFIC

	case constants.RAHU:
		nature.NaturalNature = constants.MALEFIC

	case constants.KETU:
		nature.NaturalNature = constants.MALEFIC
	}
}

func (nature *GrahaNature) findFunctionalNature(name string, chart *Chart) {
	nature.FunctionalNatureScore = 0

	bhavas := chart.GetOwningBhavas(name)
	for _, num := range bhavas {
		switch num {
		case 1:
			if name != constants.MOON {
				nature.FunctionalNatureScore += 2
			}

		case 2:
			if name != constants.SUN && name != constants.MOON {
				nature.FunctionalNatureScore += 1
			}

		case 3, 6, 11:
			nature.FunctionalNatureScore -= 1

		case 5, 9:
			nature.FunctionalNatureScore += 2

		case 8, 12:
			if name != constants.SUN && name != constants.MOON {
				nature.FunctionalNatureScore -= 1
			}

		}
	}

	// Override the rules if graha in its own rashi
	_, b := chart.GetGrahaBhava(name)
	if b != nil {
		if b.RashiLord == name {
			if nature.FunctionalNatureScore < 0 {
				nature.FunctionalNatureScore = 0
			}
			nature.FunctionalNatureScore += 1
		}
	}

	if nature.FunctionalNatureScore > 0 {
		nature.FunctionalNature = constants.BENEFIC
	} else if nature.FunctionalNatureScore == 0 {
		nature.FunctionalNature = constants.NEUTRAL
	} else {
		nature.FunctionalNature = constants.MALEFIC
	}
}

func (nature *GrahaNature) EvaluateGrahaNature(name string, chart *Chart) {
	nature.findNaturalNature(name, chart)
	nature.findFunctionalNature(name, chart)
}
