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

type BhavaLord struct {
	Name              string
	DistanceFromLagna int
	DistanceFromBhava int
}

type Bhava struct {
	Number             int
	RashiNum           int
	RashiLord          BhavaLord
	Grahas             []GrahaLocCombust
	FullAspect         []string
	ThreeQuarterAspect []string
	HalfAspect         []string
	QuarterAspect      []string
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
