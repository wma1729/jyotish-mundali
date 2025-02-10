package analysis

import (
	"jyotish/constants"
	"jyotish/misc"
)

type GrahaStrength struct {
	Name            string
	Position        int
	Combust         bool
	Retrograde      bool
	FriendlyAspects []string
	NeutralAspects  []string
	InimicalAspects []string
	InKendra        bool
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
}
