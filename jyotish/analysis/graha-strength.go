package analysis

import (
	"fmt"
	"jyotish/constants"
	"jyotish/misc"
)

type GrahaStrength struct {
	Name                string
	Position            int
	Combust             bool
	Retrograde          bool
	FriendlyAspects     []string
	NeutralAspects      []string
	InimicalAspects     []string
	InKendra            bool
	State               int
	DirectionalStrength string
}

func (strength *GrahaStrength) findGrahaPosition(name string, bhava *Bhava) {
	grahaBala := constants.GrahaBalaInRashiRulesMap[name]
	rashiLord := constants.RashiLordMap[bhava.RashiNum]
	rashiLordRelations := constants.GrahaBalaInRashiRulesMap[rashiLord]

	if bhava.RashiNum == grahaBala.Exaltation.RashiNum {
		strength.Position = constants.IN_EXALTATION_RASHI
	} else if bhava.RashiNum == grahaBala.Debilitation.RashiNum {
		strength.Position = constants.IN_DEBILITATION_RASHI
	} else if bhava.RashiNum == grahaBala.Trinal.RashiNum {
		strength.Position = constants.IN_MOOLTRIKONA_RASHI
	} else if rashiLord == name {
		strength.Position = constants.IN_OWN_RASHI
	} else if misc.StringSliceContains(rashiLordRelations.Friends, name) {
		strength.Position = constants.IN_FRIENDLY_RASHI
	} else if misc.StringSliceContains(rashiLordRelations.Enemies, name) {
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

func (strength *GrahaStrength) findAspects(name string, bhava *Bhava) {
	grahaRelations := constants.GrahaBalaInRashiRulesMap[name]
	for _, graha := range bhava.FullAspect {
		if misc.StringSliceContains(grahaRelations.Friends, graha) {
			strength.FriendlyAspects = append(strength.FriendlyAspects, graha)
		} else if misc.StringSliceContains(grahaRelations.Enemies, graha) {
			strength.InimicalAspects = append(strength.InimicalAspects, graha)
		} else {
			strength.NeutralAspects = append(strength.NeutralAspects, graha)
		}
	}
}

func (strength *GrahaStrength) findState(name string, bhava *Bhava) {
	graha := bhava.GrahaByName(name)
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
	degree := graha.Degree + float32(30*(bhava.Number-1))
	directionalStrength := float32(0)

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
			degree = misc.AbsoluteDifference(degree, 90)
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
			degree = misc.AbsoluteDifference(degree, 270)
		}
		directionalStrength = 1 - (degree / 180)
	}

	strength.DirectionalStrength = fmt.Sprintf("%.2f", directionalStrength)
}

func (strength *GrahaStrength) EvaluateGrahaStrength(name string, chart *Chart) {
	strength.Name = name

	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	strength.findGrahaPosition(name, b)
	strength.findCombustAndRetrograde(name, b)
	strength.findAspects(name, b)
	strength.InKendra = b.Number == 4 || b.Number == 7 || b.Number == 10 || b.Number == 1
	strength.findState(name, b)
	strength.findDirectionalStrength(name, b)
}
