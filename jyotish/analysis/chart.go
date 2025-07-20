package analysis

import (
	"jyotish/constants"
	"jyotish/models"
	"log"
	"math"
	"sort"
)

type GrahaAttributes struct {
	Relations GrahaRelations
	Aspects   GrahaAspects
	Strength  GrahaStrength
	Nature    GrahaNature
}

type Chart struct {
	Bhavas     []Bhava
	GrahasAttr []GrahaAttributes
}

func GetChart(gl models.GrahasLocation) Chart {
	var bhavas [constants.MAX_BHAVA_NUM]Bhava

	lagnaRashi := gl.GetLagnaRashi()

	bhavas[0].Number = 1
	bhavas[0].RashiNum = lagnaRashi
	bhavas[0].RashiLord = constants.RashiLordMap[bhavas[0].RashiNum]

	for i := 1; i < len(bhavas); i++ {
		lagnaRashi++
		if lagnaRashi > constants.MAX_BHAVA_NUM {
			lagnaRashi = 1
		}
		bhavas[i].Number = i + 1
		bhavas[i].RashiNum = lagnaRashi
		bhavas[i].RashiLord = constants.RashiLordMap[bhavas[i].RashiNum]
	}

	for i := 0; i < len(bhavas); i++ {
		for j := 0; j < len(gl.Grahas); j++ {
			if bhavas[i].RashiNum == gl.Grahas[j].RashiNum {
				var grahaLocationState GrahaLocCombust
				grahaLocationState.Name = gl.Grahas[j].Name
				grahaLocationState.RashiNum = gl.Grahas[j].RashiNum
				grahaLocationState.Degree = gl.Grahas[j].Degree
				grahaLocationState.Retrograde = gl.Grahas[j].Retrograde
				bhavas[i].Grahas = append(bhavas[i].Grahas, grahaLocationState)
			}
		}
		sort.Slice(bhavas[i].Grahas, func(x, y int) bool {
			return bhavas[i].Grahas[x].Degree > bhavas[i].Grahas[y].Degree
		})
	}

	var chart Chart
	chart.Bhavas = bhavas[:]
	chart.GrahasAttr = make([]GrahaAttributes, 9)

	chart.findCombustGrahas()
	chart.findAspectsOnBhavas()
	chart.EvaluateGrahaRelations()
	chart.EvaluateGrahaAspects()
	chart.EvaluateGrahaNature()
	chart.EvaluateGrahaStrength()

	for i := 0; i < len(chart.Bhavas); i++ {
		chart.Bhavas[i].FindGrahasInfluenceBasedOnPosition(&chart)
		chart.Bhavas[i].FindGrahasInfluenceBasedOnStrength(&chart)
		chart.Bhavas[i].FindBhavaKarakaInfluence(&chart)
	}

	return chart
}

func (c *Chart) GetGrahaBhava(name string) (int, *Bhava) {
	for i, b := range c.Bhavas {
		if b.ContainsGraha(name) {
			return i, &c.Bhavas[i]
		}
	}
	log.Printf("unable to find bhava where %s is placed in", name)
	return -1, nil
}

func (c *Chart) GetNthBhava(i, n int) *Bhava {
	bn := i + n - 1
	if bn >= constants.MAX_BHAVA_NUM {
		bn -= constants.MAX_BHAVA_NUM
	}
	return &c.Bhavas[bn]
}

func (c *Chart) NthBhavaContainsGraha(i, n int, graha string) bool {
	b := c.GetNthBhava(i, n)
	return b.ContainsGraha(graha)
}

func (c *Chart) GetGrahaAttributes(name string) *GrahaAttributes {
	for _, grahaAttr := range c.GrahasAttr {
		if grahaAttr.Relations.Name == name {
			return &grahaAttr
		}
	}
	log.Printf("unable to find attributes of %s", name)
	return nil
}

func (c *Chart) GetOwningBhavas(name string) []int {
	bhavas := make([]int, 0)
	for _, b := range c.Bhavas {
		if b.RashiLord == name {
			bhavas = append(bhavas, b.Number)
		}
	}
	return bhavas
}

func isCombust(graha string, retrograde bool, distanceFromSun float64) (bool, float64) {
	d := distanceFromSun
	switch graha {
	case constants.MERCURY:
		if retrograde {
			if d <= 12.0 {
				return true, (12.0 - d) / 12.0
			}
		} else if d <= 14.0 {
			return true, (14.0 - d) / 14.0
		}

	case constants.VENUS:
		if retrograde {
			if d <= 8.0 {
				return true, (8.0 - d) / 8.0
			}
		} else if d <= 10.0 {
			return true, (10.0 - d) / 10.0
		}

	case constants.MARS:
		if d <= 17.0 {
			return true, (17.0 - d) / 17.0
		}

	case constants.JUPITER:
		if d <= 11.0 {
			return true, (11.0 - d) / 11.0
		}

	case constants.SATURN:
		if d <= 15.0 {
			return true, (15.0 - d) / 15.0
		}
	}

	return false, 0.0
}

func (c *Chart) findCombustGrahas() {
	sunIndex, _ := c.GetGrahaBhava(constants.SUN)
	prevIndex := sunIndex - 1
	if prevIndex < 0 {
		prevIndex = constants.MAX_BHAVA_NUM - 1
	}
	nextIndex := sunIndex + 1
	if nextIndex == constants.MAX_BHAVA_NUM {
		nextIndex = 0
	}

	// Get SUN's degree
	var sunDegree float64
	for _, graha := range c.Bhavas[sunIndex].Grahas {
		if graha.Name == constants.SUN {
			sunDegree = graha.Degree
		}
	}

	// get combustion of all grahas in the same bhava as SUN
	for _, graha := range c.Bhavas[sunIndex].Grahas {
		if graha.Name != constants.SUN {
			distance := math.Abs(graha.Degree - sunDegree)
			graha.Combust, graha.CombustionExtent = isCombust(graha.Name, graha.Retrograde, distance)
		}
	}

	// get combustion of all grahas in the previous bhava of SUN
	for _, graha := range c.Bhavas[prevIndex].Grahas {
		distance := math.Abs((graha.Degree - 30) - sunDegree)
		graha.Combust, graha.CombustionExtent = isCombust(graha.Name, graha.Retrograde, distance)
	}

	// get combustion of all grahas in the next bhava of SUN
	for _, graha := range c.Bhavas[nextIndex].Grahas {
		distance := math.Abs((graha.Degree + 30) - sunDegree)
		graha.Combust, graha.CombustionExtent = isCombust(graha.Name, graha.Retrograde, distance)
	}
}

func (c *Chart) findAspectsOnBhavas() {
	for i, b := range c.Bhavas {
		for _, g := range b.Grahas {
			if g.Name == constants.LAGNA {
				continue
			}

			aspectedBhava := c.GetNthBhava(i, 3)
			if g.Name == constants.SATURN {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.QuarterAspect = append(aspectedBhava.QuarterAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 4)
			if g.Name == constants.MARS {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.ThreeQuarterAspect = append(aspectedBhava.ThreeQuarterAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 5)
			if g.Name == constants.JUPITER {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.HalfAspect = append(aspectedBhava.HalfAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 7)
			aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)

			aspectedBhava = c.GetNthBhava(i, 8)
			if g.Name == constants.MARS {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.ThreeQuarterAspect = append(aspectedBhava.ThreeQuarterAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 9)
			if g.Name == constants.JUPITER {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.HalfAspect = append(aspectedBhava.HalfAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 10)
			if g.Name == constants.SATURN {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.QuarterAspect = append(aspectedBhava.QuarterAspect, g.Name)
			}
		}
	}
}

func (c *Chart) EvaluateGrahaRelations() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Relations.EvaluateGrahaRelations(graha, c)
	}
}

func (c *Chart) EvaluateGrahaAspects() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Aspects.EvaluateGrahaAspects(graha, c)
	}
}

func (c *Chart) EvaluateGrahaNature() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Nature.EvaluateGrahaNature(graha, c)
	}
}

func (c *Chart) EvaluateGrahaStrength() {
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Strength.EvaluateGrahaStrength(graha, c)
	}
}
