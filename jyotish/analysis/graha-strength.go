package analysis

import (
	"jyotish/constants"
	"jyotish/misc"
	"log"
	"math"
)

type GrahaStrength struct {
	Name                string
	Position            int
	Combust             bool
	Retrograde          bool
	Residence           int
	OwnerOf             []int
	State               int
	DirectionalStrength float64
	AspectualStrength   int
	CumulativeStrength  int
}

func (strength *GrahaStrength) findGrahaPosition(name string, chart *Chart) {
	_, bhava := chart.GetGrahaBhava(name)
	if bhava == nil {
		return
	}

	grahaBala := constants.GrahaBalaInRashiRulesMap[name]
	rashiLord := constants.RashiLordMap[bhava.RashiNum]
	rashLordGrahaAttr := chart.GetGrahaAttributes(rashiLord)
	if rashLordGrahaAttr == nil {
		return
	}
	rashiLordRelations := rashLordGrahaAttr.Relations

	if name == constants.RAHU {
		log.Print(grahaBala)
		log.Print(rashiLordRelations)
	}

	if bhava.RashiNum == grahaBala.Exaltation.RashiNum {
		strength.Position = constants.IN_EXALTATION_RASHI
	} else if bhava.RashiNum == grahaBala.Debilitation.RashiNum {
		strength.Position = constants.IN_DEBILITATION_RASHI
	} else if bhava.RashiNum == grahaBala.Trinal.RashiNum {
		strength.Position = constants.IN_MOOLTRIKONA_RASHI
	} else if rashiLord == name {
		strength.Position = constants.IN_OWN_RASHI
	} else if misc.StringSliceContains(rashiLordRelations.NaturalFriends, name) {
		strength.Position = constants.IN_FRIENDLY_RASHI
	} else if misc.StringSliceContains(rashiLordRelations.NaturalEnemies, name) {
		strength.Position = constants.IN_INIMICAL_RASHI
	} else {
		strength.Position = constants.IN_NEUTRAL_RASHI
	}
}

func (strength *GrahaStrength) findCombustAndRetrograde(name string, bhava *Bhava) {
	for _, g := range bhava.Grahas {
		if g.Name == name {
			strength.Combust = g.Combust
			strength.Retrograde = g.Retrograde
			return
		}
	}
}

func (strength *GrahaStrength) findState(name string, bhava *Bhava) {
	graha := bhava.GrahaByName(name)
	if graha == nil {
		return
	}

	degree := graha.Degree
	if bhava.Number%2 == 0 {
		if degree < 6 {
			strength.State = constants.DEAD
		} else if degree >= 6 && degree < 12 {
			strength.State = constants.OLD
		} else if degree >= 12 && degree < 18 {
			strength.State = constants.ADULT
		} else if degree >= 18 && degree < 24 {
			strength.State = constants.YOUTH
		} else {
			strength.State = constants.CHILD
		}
	} else {
		if degree < 6 {
			strength.State = constants.CHILD
		} else if degree >= 6 && degree < 12 {
			strength.State = constants.YOUTH
		} else if degree >= 12 && degree < 18 {
			strength.State = constants.ADULT
		} else if degree >= 18 && degree < 24 {
			strength.State = constants.OLD
		} else {
			strength.State = constants.DEAD
		}
	}
}

func (strength *GrahaStrength) findDirectionalStrength(name string, bhava *Bhava, ascendent float64) {
	nadir := ascendent + 90
	descendent := nadir + 90
	zenith := descendent + 90
	var diff float64

	graha := bhava.GrahaByName(name)
	if graha == nil {
		return
	}

	degree := graha.Degree + 30.0*float64(bhava.Number-1)

	switch name {
	case constants.MERCURY, constants.JUPITER:
		diff = math.Abs(degree - descendent)

	case constants.MOON, constants.VENUS:
		diff = math.Abs(degree - zenith)

	case constants.SATURN:
		diff = math.Abs(degree - ascendent)

	case constants.SUN, constants.MARS:
		diff = math.Abs(degree - nadir)
	}

	if diff > 180 {
		diff = 360 - diff
	}

	diff /= (3 * 60.0)

	strength.DirectionalStrength = misc.RoundFloat(diff, 2)
}

func (strength *GrahaStrength) findAspectualStrength(name string, chart *Chart) {
	ga := chart.GetGrahaAttributes(name)
	if ga != nil {
		strength.AspectualStrength = ga.Aspects.Strength
	}
}

func (strength *GrahaStrength) EvaluateGrahaStrength(name string, chart *Chart) {
	strength.Name = name

	_, bhava := chart.GetGrahaBhava(name)
	if bhava == nil {
		return
	}

	var ascendent float64

	for _, g := range chart.Bhavas[0].Grahas {
		if g.Name == constants.LAGNA {
			ascendent = g.Degree
			break
		}
	}

	strength.Residence = bhava.Number
	strength.OwnerOf = chart.GetOwningBhavas(name)
	strength.findGrahaPosition(name, chart)
	strength.findCombustAndRetrograde(name, bhava)
	strength.findState(name, bhava)
	strength.findDirectionalStrength(name, bhava, ascendent)
	strength.findAspectualStrength(name, chart)
}
