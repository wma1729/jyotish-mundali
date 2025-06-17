package analysis

import (
	"jyotish/models"
	"log"
)

type GrahaLocCombust struct {
	models.GrahaLoc
	Combust          bool
	CombustionExtent float64
}

type GrahaPosition struct {
	Count    int
	Score    int
	Result   int
	Subjects int
}

type BhavaLordAttr struct {
	DistanceFromLagna GrahaPosition
	DistanceFromBhava GrahaPosition
}

type Bhava struct {
	Number             int
	RashiNum           int
	RashiLord          string
	Grahas             []GrahaLocCombust
	FullAspect         []string
	ThreeQuarterAspect []string
	HalfAspect         []string
	QuarterAspect      []string
	BhavaLord          BhavaLordAttr
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
