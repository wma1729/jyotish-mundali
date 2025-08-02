package analysis

import (
	"jyotish/constants"
	"jyotish/models"
	"log"
	"sort"
)

type GrahaLocCombust struct {
	models.GrahaLoc
	Combust          bool
	CombustionExtent float64
}

type GrahaInfluenceRating struct {
	Value  interface{}
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
	StateStrength         GrahaInfluenceRating
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
	BhavaKarakaInBhava         GrahaInfluenceRating
	GrahasOnEitherSide         GrahaInfluenceRating
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

func (b *Bhava) GrahaDegree(name string, absolute bool) float64 {
	graha := b.GrahaByName(name)
	if absolute {
		return float64(b.Number-1)*30.0 + graha.Degree
	} else {
		return graha.Degree
	}
}

func (b *Bhava) IsRetrograde(name string) bool {
	graha := b.GrahaByName(name)
	if graha != nil {
		return graha.Retrograde
	}
	return false
}

func (b *Bhava) FindGrahasInfluence(c *Chart, name string, assoc int) {
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
		return
	}

	// Get natural nature
	gi.Nature.Value = ga.Nature.NaturalNature
	gi.Nature.Rating = ga.Nature.NaturalNature

	// Is the graha friendly, inimical or neutral to bhava lord?
	if name == b.RashiLord && assoc != constants.BHAVA_OWNERSHIP {
		gi.RelationWithBhavaLord.Value = constants.BENEFIC
		gi.RelationWithBhavaLord.Rating = constants.BENEFIC
		gi.RelationWithBhavaLord.Notes = constants.BHAVA_LORD
	} else {
		gi.RelationWithBhavaLord.Value = constants.NEUTRAL
		gi.RelationWithBhavaLord.Rating = constants.NEUTRAL
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
	if constants.IsGoodBhava(ga.Strength.Residence) {
		gi.PositionInChart.Rating = constants.BENEFIC
	} else if constants.IsBadBhava(ga.Strength.Residence) {
		gi.PositionInChart.Rating = constants.MALEFIC
	} else {
		gi.PositionInChart.Rating = constants.NEUTRAL
	}

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

	gi.StateStrength.Value = ga.Strength.State
	switch ga.Strength.State {
	case constants.CHILD, constants.DEAD:
		gi.StateStrength.Rating = constants.MALEFIC
	case constants.YOUTH, constants.OLD:
		gi.StateStrength.Rating = constants.NEUTRAL
	case constants.ADULT:
		gi.StateStrength.Rating = constants.BENEFIC
	}

	// Get the directional strength
	if name == constants.RAHU || name == constants.KETU {
		gi.DirectionalStrength.Value = 0.50
		gi.DirectionalStrength.Rating = constants.NEUTRAL
	} else {
		gi.DirectionalStrength.Value = ga.Strength.DirectionalStrength
		if ga.Strength.DirectionalStrength >= 0.75 {
			gi.DirectionalStrength.Rating = constants.BENEFIC
		} else if ga.Strength.DirectionalStrength <= 0.25 {
			gi.DirectionalStrength.Rating = constants.MALEFIC
		} else {
			gi.DirectionalStrength.Rating = constants.NEUTRAL
		}
	}

	// Get the aspectual strength
	gi.AspectualStrength.Value = ga.Strength.AspectualStrength
	if ga.Strength.AspectualStrength > 0 {
		gi.AspectualStrength.Rating = constants.BENEFIC
	} else if ga.Strength.AspectualStrength == 0 {
		gi.AspectualStrength.Rating = constants.NEUTRAL
	} else {
		gi.AspectualStrength.Rating = constants.MALEFIC
	}

	b.GrahasInfluence = append(b.GrahasInfluence, gi)
}

func (b *Bhava) FindGrahasInfluenceBasedOnStrength(c *Chart) {
	b.GrahasInfluence = make([]GrahaInfluenceOnBhava, 0)

	b.FindGrahasInfluence(c, b.RashiLord, constants.BHAVA_OWNERSHIP)

	for _, g := range b.Grahas {
		if g.Name != constants.LAGNA {
			b.FindGrahasInfluence(c, g.Name, constants.BHAVA_PLACEMENT)
		}
	}

	for _, g := range b.FullAspect {
		b.FindGrahasInfluence(c, g, constants.BHAVA_ASPECT)
	}

	for _, g := range constants.BhavaKarakas[b.Number] {
		b.FindGrahasInfluence(c, g, constants.BHAVA_KARAKA)
	}
}

func (b *Bhava) FindGrahasInfluenceBasedOnPosition(c *Chart) {
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
}

func (b *Bhava) FindBhavaKarakaInfluence(c *Chart) {
	b.BhavaKarakaInBhava.Rating = constants.NEUTRAL

	switch b.Number {
	case 3:
		_, grahaBhava := c.GetGrahaBhava(constants.MARS)
		if grahaBhava.Number == 3 {
			b.BhavaKarakaInBhava.Value = constants.MARS
			b.BhavaKarakaInBhava.Rating = constants.MALEFIC
			b.BhavaKarakaInBhava.Notes = constants.SUBJECTS_YOUNGER_SIBLINGS
		}

	case 4:
		_, grahaBhava := c.GetGrahaBhava(constants.MOON)
		if grahaBhava.Number == 4 {
			ga := c.GetGrahaAttributes(constants.MOON)
			if ga != nil {
				if ga.Nature.NaturalNature == constants.MALEFIC {
					b.BhavaKarakaInBhava.Value = constants.MOON
					b.BhavaKarakaInBhava.Rating = constants.MALEFIC
					b.BhavaKarakaInBhava.Notes = constants.SUBJECTS_MOTHER
				}
			}
		}

	case 5:
		_, grahaBhava := c.GetGrahaBhava(constants.JUPITER)
		if grahaBhava.Number == 5 {
			b.BhavaKarakaInBhava.Value = constants.JUPITER
			b.BhavaKarakaInBhava.Rating = constants.MALEFIC
			b.BhavaKarakaInBhava.Notes = constants.SUBJECTS_CHILDREN
		}

	case 7:
		_, grahaBhava := c.GetGrahaBhava(constants.VENUS)
		if grahaBhava.Number == 7 {
			b.BhavaKarakaInBhava.Value = constants.VENUS
			b.BhavaKarakaInBhava.Rating = constants.MALEFIC
			b.BhavaKarakaInBhava.Notes = constants.SUBJECTS_SPOUSE
		}

	case 9:
		_, grahaBhava := c.GetGrahaBhava(constants.SUN)
		if grahaBhava.Number == 9 {
			b.BhavaKarakaInBhava.Value = constants.SUN
			b.BhavaKarakaInBhava.Rating = constants.MALEFIC
			b.BhavaKarakaInBhava.Notes = constants.SUBJECTS_FATHER
		}

	case 11:
		_, grahaBhava := c.GetGrahaBhava(constants.JUPITER)
		if grahaBhava.Number == 9 {
			b.BhavaKarakaInBhava.Value = constants.JUPITER
			b.BhavaKarakaInBhava.Rating = constants.MALEFIC
			b.BhavaKarakaInBhava.Notes = constants.SUBJECTS_ELDER_SIBLINGS
		}
	}
}

func (b *Bhava) FindGrahasOnEitherSide(c *Chart) {
	n := b.Number
	behind := n - 1
	ahead := n + 1

	switch n {
	case 1:
		behind = constants.MAX_BHAVA_NUM
	case 12:
		ahead = 1
	}

	behindGrahas := make([]GrahaLocCombust, 0)
	for _, grahaLocCombust := range c.Bhavas[behind-1].Grahas {
		if grahaLocCombust.Name != constants.LAGNA {
			behindGrahas = append(behindGrahas, grahaLocCombust)
		}
	}

	aheadGrahas := make([]GrahaLocCombust, 0)
	for _, grahaLocCombust := range c.Bhavas[ahead-1].Grahas {
		if grahaLocCombust.Name != constants.LAGNA {
			aheadGrahas = append(aheadGrahas, grahaLocCombust)
		}
	}

	b.GrahasOnEitherSide.Value = []string{}
	b.GrahasOnEitherSide.Rating = constants.NEUTRAL

	if len(behindGrahas) > 0 && len(aheadGrahas) > 0 {
		sort.Slice(behindGrahas, func(x, y int) bool {
			return behindGrahas[x].Degree > behindGrahas[y].Degree
		})

		sort.Slice(aheadGrahas, func(x, y int) bool {
			return aheadGrahas[x].Degree < aheadGrahas[y].Degree
		})

		closestBehindGrahaAttr := c.GetGrahaAttributes(behindGrahas[0].Name)
		if closestBehindGrahaAttr == nil {
			return
		}

		closestAheadGrahaAttr := c.GetGrahaAttributes(aheadGrahas[0].Name)
		if closestAheadGrahaAttr == nil {
			return
		}

		if closestBehindGrahaAttr.Nature.NaturalNature == constants.BENEFIC &&
			closestAheadGrahaAttr.Nature.NaturalNature == constants.BENEFIC {
			b.GrahasOnEitherSide.Value = []string{behindGrahas[0].Name, aheadGrahas[0].Name}
			b.GrahasOnEitherSide.Rating = constants.BENEFIC
		} else if closestBehindGrahaAttr.Nature.NaturalNature == constants.MALEFIC &&
			closestAheadGrahaAttr.Nature.NaturalNature == constants.MALEFIC {
			b.GrahasOnEitherSide.Value = []string{behindGrahas[0].Name, aheadGrahas[0].Name}
			b.GrahasOnEitherSide.Rating = constants.MALEFIC
		}
	}
}
