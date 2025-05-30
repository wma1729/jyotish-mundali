package analysis

import (
	"jyotish/constants"
	"jyotish/misc"
	"log"
	"math"
)

type GrahaStrength struct {
	Name                  string
	Position              int
	PositionalStrength    float64
	Combust               bool
	CombustionStrength    float64
	Retrograde            bool
	RetrogressionStrength float64
	Residence             int
	ResidenceStrength     float64
	State                 int
	StateStrength         float64
	DirectionalStrength   float64
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
		strength.PositionalStrength = 0.9
	} else if bhava.RashiNum == grahaBala.Debilitation.RashiNum {
		strength.Position = constants.IN_DEBILITATION_RASHI
		strength.PositionalStrength = -0.9
	} else if bhava.RashiNum == grahaBala.Trinal.RashiNum {
		strength.Position = constants.IN_MOOLTRIKONA_RASHI
		strength.PositionalStrength = 0.8
	} else if rashiLord == name {
		strength.Position = constants.IN_OWN_RASHI
		strength.PositionalStrength = 0.6
	} else if misc.StringSliceContains(rashiLordRelations.EffectiveBestFriends, name) {
		strength.Position = constants.IN_FRIENDLY_RASHI
		strength.PositionalStrength = 0.4
	} else if misc.StringSliceContains(rashiLordRelations.EffectiveFriends, name) {
		strength.Position = constants.IN_FRIENDLY_RASHI
		strength.PositionalStrength = 0.2
	} else if misc.StringSliceContains(rashiLordRelations.EffectiveEnemies, name) {
		strength.Position = constants.IN_INIMICAL_RASHI
		strength.PositionalStrength = -0.2
	} else if misc.StringSliceContains(rashiLordRelations.EffectiveWorstEnemies, name) {
		strength.Position = constants.IN_INIMICAL_RASHI
		strength.PositionalStrength = -0.4
	} else {
		strength.Position = constants.IN_NEUTRAL_RASHI
		strength.PositionalStrength = 0.0
	}
}

func (strength *GrahaStrength) findCombustAndRetrograde(name string, bhava *Bhava) {
	for _, g := range bhava.Grahas {
		if g.Name == name {
			strength.Combust = g.Combust
			if g.Combust {
				combustionStrength := misc.RoundFloat((1.0-g.CombustionExtent)/4.0, 2)
				if strength.PositionalStrength >= 0.0 {
					strength.CombustionStrength = -1.0 * combustionStrength
				} else {
					strength.CombustionStrength = combustionStrength
				}
			} else {
				strength.CombustionStrength = 0.0
			}

			strength.Retrograde = g.Retrograde
			if g.Retrograde && g.Name != constants.RAHU && g.Name != constants.KETU {
				if strength.Position == constants.IN_DEBILITATION_RASHI {
					strength.RetrogressionStrength = -0.25
				} else {
					strength.RetrogressionStrength = 0.25
				}
			} else {
				strength.RetrogressionStrength = 0.0
			}
			return
		}
	}
}

func (strength *GrahaStrength) findResidenceStrength(name string, bhava *Bhava) {
	strength.Residence = bhava.Number
	residenceStrength := float64(0.0)

	switch bhava.Number {
	case 9:
		residenceStrength = 1.0
	case 5:
		residenceStrength = 5.0 / 6.0
	case 1:
		residenceStrength = 4.0 / 6.0
	case 10:
		residenceStrength = 3.0 / 6.0
	case 4:
		residenceStrength = 2.0 / 6.0
	case 7:
		residenceStrength = 1.0 / 6.0
	}

	strength.ResidenceStrength = misc.RoundFloat(residenceStrength, 2)
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
			strength.StateStrength = 0.25
		} else if degree >= 6 && degree < 12 {
			strength.State = constants.OLD
			strength.StateStrength = 0.50
		} else if degree >= 12 && degree < 18 {
			strength.State = constants.ADULT
			strength.StateStrength = 1.0
		} else if degree >= 18 && degree < 24 {
			strength.State = constants.YOUTH
			strength.StateStrength = 0.50
		} else {
			strength.State = constants.CHILD
			strength.StateStrength = 0.25
		}
	} else {
		if degree < 6 {
			strength.State = constants.CHILD
			strength.StateStrength = 0.25
		} else if degree >= 6 && degree < 12 {
			strength.State = constants.YOUTH
			strength.StateStrength = 0.50
		} else if degree >= 12 && degree < 18 {
			strength.State = constants.ADULT
			strength.StateStrength = 1.0
		} else if degree >= 18 && degree < 24 {
			strength.State = constants.OLD
			strength.StateStrength = 0.50
		} else {
			strength.State = constants.DEAD
			strength.StateStrength = 0.25
		}
	}

	strength.StateStrength = misc.RoundFloat(strength.StateStrength, 2)
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

	_, b := chart.GetGrahaBhava(name)
	if b == nil {
		return
	}

	strength.findGrahaPosition(name, chart)
	strength.findCombustAndRetrograde(name, b)
	strength.findResidenceStrength(name, b)
	strength.findState(name, b)
	strength.findDirectionalStrength(name, b)
}
