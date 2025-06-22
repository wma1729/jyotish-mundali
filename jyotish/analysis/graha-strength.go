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
}

func (strength *GrahaStrength) findGrahaPosition(name string, chart *Chart) {
	_, bhava := chart.GetGrahaBhava(name)
	if bhava == nil {
		return
	}

	grahaBala := constants.GrahaBalaInRashiRulesMap[name]
	rashiLord := constants.RashiLordMap[bhava.RashiNum]
	rashLordGrahaAttr := chart.GetGrahaAttributes(rashiLord)
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

func (strength *GrahaStrength) findDirectionalStrength(name string, bhava *Bhava) {
	graha := bhava.GrahaByName(name)
	if graha == nil {
		return
	}

	degree := graha.Degree + 30.0*float64(bhava.Number-1)
	directionalStrength := float64(0.0)

	switch name {
	case constants.MERCURY:
	case constants.JUPITER:
		if degree > 180 {
			degree = degree - 180
			directionalStrength = degree / 180
		} else {
			directionalStrength = 1 - (degree / 180)
		}

	case constants.MOON:
	case constants.VENUS:
		if degree > 270 && degree < 360 {
			degree = degree - 270
			directionalStrength = degree / 180
		} else {
			degree = math.Abs(degree - 90)
			directionalStrength = 1 - (degree / 180)
		}

	case constants.SATURN:
		if degree >= 180 && degree < 360 {
			degree = degree - 180
		} else if degree >= 0 && degree < 180 {
			degree = 180 - degree
		}
		directionalStrength = 1 - (degree / 180)

	case constants.SUN:
	case constants.MARS:
		if degree >= 0 && degree < 90 {
			degree = degree + 90
		} else {
			degree = math.Abs(degree - 270)
		}
		directionalStrength = 1 - (degree / 180)
	}

	strength.DirectionalStrength = misc.RoundFloat(directionalStrength, 2)
}

func (strength *GrahaStrength) EvaluateGrahaStrength(name string, chart *Chart) {
	strength.Name = name

	_, bhava := chart.GetGrahaBhava(name)
	if bhava == nil {
		return
	}

	strength.Residence = bhava.Number
	strength.OwnerOf = chart.GetOwningBhavas(name)
	strength.findGrahaPosition(name, chart)
	strength.findCombustAndRetrograde(name, bhava)
	strength.findState(name, bhava)
	strength.findDirectionalStrength(name, bhava)
}
