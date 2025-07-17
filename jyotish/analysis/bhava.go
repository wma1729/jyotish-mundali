package analysis

import (
	"jyotish/constants"
	"jyotish/models"
	"log"
)

type GrahaLocCombust struct {
	models.GrahaLoc
	Combust          bool
	CombustionExtent float64
}

type GrahaInfluenceRating struct {
	Value  int
	Rating int
	Notes  string
}

type GrahaInfluenceOnBhava struct {
	Name                  string
	AssociationWithBhava  []int
	Nature                GrahaInfluenceRating
	RelationWithBhavaLord GrahaInfluenceRating
	PositionInChart       GrahaInfluenceRating
	PositionalStrength    GrahaInfluenceRating
	OwnerOf               []GrahaInfluenceRating
	Combust               GrahaInfluenceRating
	Retrograde            GrahaInfluenceRating
	DirectionalStrength   GrahaInfluenceRating
	AspectualStrength     GrahaInfluenceRating
}

type Bhava struct {
	Number                     int
	RashiNum                   int
	RashiLord                  string
	Grahas                     []GrahaLocCombust
	FullAspect                 []string
	ThreeQuarterAspect         []string
	HalfAspect                 []string
	QuarterAspect              []string
	BhavaLordDistanceFromLagna GrahaInfluenceRating
	BhavaLordDistanceFromBhava GrahaInfluenceRating
	GrahasInfluence            []GrahaInfluenceOnBhava
}

func (b *Bhava) ContainsGraha(name string) bool {
	for _, g := range b.Grahas {
		if g.Name == name {
			return true
		}
	}
	return false
}

func (b *Bhava) GrahaByName(name string) *GrahaLocCombust {
	for _, g := range b.Grahas {
		if g.Name == name {
			return &g
		}
	}
	log.Printf("unable to find %s in bhava %d", name, b.Number)
	return nil
}

func (b *Bhava) GrahaDegree(name string) float64 {
	graha := b.GrahaByName(name)
	return graha.Degree
}

func (b *Bhava) IsRetrograde(name string) bool {
	graha := b.GrahaByName(name)
	if graha != nil {
		return graha.Retrograde
	}
	return false
}

func (b *Bhava) FindGrahasAssociations(c *Chart, name string, assoc int) {
	for i, gi := range b.GrahasInfluence {
		if gi.Name == name {
			b.GrahasInfluence[i].AssociationWithBhava =
				append(gi.AssociationWithBhava, assoc)
			return
		}
	}

	gi := GrahaInfluenceOnBhava{Name: name}
	gi.AssociationWithBhava = make([]int, 0)
	gi.AssociationWithBhava = append(gi.AssociationWithBhava, assoc)

	ga := c.GetGrahaAttributes(name)
	if ga == nil {
		log.Printf("unable to find attributes of graha %s", name)
		return
	}

	// Get natural nature

	gi.Nature.Value = ga.Nature.NaturalNature
	gi.Nature.Rating = ga.Nature.NaturalNature
	gi.Nature.Notes = ""

	// Is the graha friendly, inimical or neutral to bhava lord?

	if name == b.RashiLord && assoc != constants.BHAVA_OWNERSHIP {
		gi.RelationWithBhavaLord.Value = constants.BENEFIC
		gi.RelationWithBhavaLord.Rating = constants.BENEFIC
		gi.RelationWithBhavaLord.Notes = constants.BHAVA_LORD
	} else {
		gi.RelationWithBhavaLord.Value = constants.NEUTRAL
		gi.RelationWithBhavaLord.Rating = constants.NEUTRAL
		gi.RelationWithBhavaLord.Notes = ""
	}

	for _, g := range constants.GrahaBalaInRashiRulesMap[name].Friends {
		if g == b.RashiLord {
			gi.RelationWithBhavaLord.Value = constants.FRIEND
			gi.RelationWithBhavaLord.Rating = constants.BENEFIC
			break
		}
	}

	for _, g := range constants.GrahaBalaInRashiRulesMap[name].Enemies {
		if g == b.RashiLord {
			gi.RelationWithBhavaLord.Value = constants.ENEMY
			gi.RelationWithBhavaLord.Rating = constants.MALEFIC
			break
		}
	}

	// Get the position

	gi.PositionInChart.Value = ga.Strength.Residence
	if constants.IsGoodBhava(gi.PositionInChart.Value) {
		gi.PositionInChart.Rating = constants.BENEFIC
	} else if constants.IsBadBhava(gi.PositionInChart.Value) {
		gi.PositionInChart.Rating = constants.MALEFIC
	} else {
		gi.PositionInChart.Rating = constants.NEUTRAL
	}
	gi.PositionInChart.Notes = ""

	// Get the position strength

	gi.PositionalStrength.Value = ga.Strength.Position
	switch gi.PositionalStrength.Value {
	case constants.IN_EXALTATION_RASHI,
		constants.IN_MOOLTRIKONA_RASHI,
		constants.IN_OWN_RASHI,
		constants.IN_FRIENDLY_RASHI:
		gi.PositionalStrength.Rating = constants.BENEFIC

	case constants.IN_DEBILITATION_RASHI,
		constants.IN_INIMICAL_RASHI:
		gi.PositionalStrength.Rating = constants.MALEFIC

	default:
		gi.PositionalStrength.Rating = constants.NEUTRAL
	}
	gi.PositionalStrength.Notes = ""

	// Get the ownership

	gi.OwnerOf = make([]GrahaInfluenceRating, 0)
	ownerOf := c.GetOwningBhavas(name)
	for _, n := range ownerOf {
		gir := GrahaInfluenceRating{}
		gir.Value = n
		if constants.IsGoodBhava(n) || n == b.Number {
			gir.Rating = constants.BENEFIC
		} else if constants.IsBadBhava(n) {
			gir.Rating = constants.MALEFIC
		} else {
			gir.Rating = constants.NEUTRAL
		}
		gir.Notes = ""
		gi.OwnerOf = append(gi.OwnerOf, gir)
	}

	// Get the combustion

	if ga.Strength.Combust {
		gi.Combust.Value = 1
		gi.Combust.Rating = constants.MALEFIC
	} else {
		gi.Combust.Value = 0
		gi.Combust.Rating = constants.NEUTRAL
	}
	gi.Combust.Notes = ""

	// Get the retrogression

	if ga.Strength.Retrograde {
		gi.Retrograde.Value = 1
		if name == constants.RAHU || name == constants.KETU {
			gi.Retrograde.Rating = constants.NEUTRAL
		} else if ga.Nature.NaturalNature == constants.BENEFIC {
			gi.Retrograde.Rating = constants.BENEFIC
		} else {
			gi.Retrograde.Rating = constants.MALEFIC
		}
	} else {
		gi.Retrograde.Value = 0
		gi.Retrograde.Rating = constants.NEUTRAL
	}
	gi.Retrograde.Notes = ""

	// Get the directional strength

	if name == constants.RAHU || name == constants.KETU {
		gi.DirectionalStrength.Rating = constants.NEUTRAL
	} else {
		gi.DirectionalStrength.Value = int(ga.Strength.DirectionalStrength * 100)
		if ga.Strength.DirectionalStrength >= 0.75 {
			gi.DirectionalStrength.Rating = constants.BENEFIC
		} else if ga.Strength.DirectionalStrength <= 0.25 {
			gi.DirectionalStrength.Rating = constants.MALEFIC
		} else {
			gi.DirectionalStrength.Rating = constants.NEUTRAL
		}
	}
	gi.DirectionalStrength.Notes = ""

	gi.AspectualStrength.Value = ga.Strength.AspectualStrength
	if gi.AspectualStrength.Value > 0 {
		gi.AspectualStrength.Rating = constants.BENEFIC
	} else if gi.AspectualStrength.Value == 0 {
		gi.AspectualStrength.Rating = constants.NEUTRAL
	} else {
		gi.AspectualStrength.Rating = constants.MALEFIC
	}
	gi.AspectualStrength.Notes = ""

	b.GrahasInfluence = append(b.GrahasInfluence, gi)
}

func (b *Bhava) FindGrahasInfluence(c *Chart) {
	b.GrahasInfluence = make([]GrahaInfluenceOnBhava, 0)

	b.FindGrahasAssociations(c, b.RashiLord, constants.BHAVA_OWNERSHIP)

	for i := 0; i < constants.MAX_BHAVA_NUM; i++ {
		if c.Bhavas[i].ContainsGraha(b.RashiLord) {
			b.BhavaLordDistanceFromLagna.Value = i + 1
			var distanceFromBhava = i - (b.Number - 1)
			if distanceFromBhava < 0 {
				distanceFromBhava += constants.MAX_BHAVA_NUM
			}
			b.BhavaLordDistanceFromBhava.Value = distanceFromBhava + 1
			break
		}
	}

	switch b.BhavaLordDistanceFromLagna.Value {
	case 1, 5, 9:
		b.BhavaLordDistanceFromLagna.Rating = constants.BENEFIC

	case 6, 8, 12:
		b.BhavaLordDistanceFromLagna.Rating = constants.MALEFIC

	default:
		b.BhavaLordDistanceFromLagna.Rating = constants.NEUTRAL
	}
	b.BhavaLordDistanceFromLagna.Notes = constants.SUBJECTS_LIVING_BEING

	switch b.BhavaLordDistanceFromBhava.Value {
	case 1, 5, 9:
		b.BhavaLordDistanceFromBhava.Rating = constants.BENEFIC

	case 6, 8, 12:
		b.BhavaLordDistanceFromBhava.Rating = constants.MALEFIC

	default:
		b.BhavaLordDistanceFromBhava.Rating = constants.NEUTRAL
	}
	b.BhavaLordDistanceFromBhava.Notes = constants.SUBJECTS_NON_LIVING_BEING

	for _, g := range b.Grahas {
		if g.Name != constants.LAGNA {
			b.FindGrahasAssociations(c, g.Name, constants.BHAVA_PLACEMENT)
		}
	}

	for _, g := range b.FullAspect {
		b.FindGrahasAssociations(c, g, constants.BHAVA_ASPECT)
	}

	for _, g := range constants.BhavaKarakas[b.Number] {
		b.FindGrahasAssociations(c, g, constants.BHAVA_SIGNIFICATOR)
	}
}
