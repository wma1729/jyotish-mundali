package analysis

import (
	"math"
	"sort"
)

type Chart struct {
	Bhavas []Bhava
}

func GetChart(gl GrahasLocation) Chart {
	var bhavas [MAX_BHAVA_NUM]Bhava

	lagnaRashi := gl.GetLagnaRashi()

	bhavas[0].Number = 1
	bhavas[0].RashiNum = lagnaRashi
	bhavas[0].RashiLord = RashiLordMap[bhavas[0].RashiNum]

	for i := 1; i < len(bhavas); i++ {
		lagnaRashi++
		if lagnaRashi > MAX_BHAVA_NUM {
			lagnaRashi = 1
		}
		bhavas[i].Number = i + 1
		bhavas[i].RashiNum = lagnaRashi
		bhavas[i].RashiLord = RashiLordMap[bhavas[i].RashiNum]
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
	chart.findCombustGrahas()

	return chart
}

func (c *Chart) GetGrahaBhava(name string) (int, *Bhava) {
	for i, b := range c.Bhavas {
		if b.ContainsGraha(name) {
			return i, &c.Bhavas[i]
		}
	}
	return -1, nil
}

func (c *Chart) GetNthBhava(i, n int) *Bhava {
	bn := i + n - 1
	if bn >= MAX_BHAVA_NUM {
		bn -= MAX_BHAVA_NUM
	}
	return &c.Bhavas[bn]
}

func (c *Chart) NthBhavaContainsGraha(i, n int, graha string) bool {
	b := c.GetNthBhava(i, n)
	return b.ContainsGraha(graha)
}

func (c *Chart) findCombustGrahas() {
	sunIndex, _ := c.GetGrahaBhava(SUN)
	prevIndex := sunIndex - 1
	if prevIndex < 0 {
		prevIndex = MAX_BHAVA_NUM - 1
	}
	nextIndex := sunIndex + 1
	if nextIndex == MAX_BHAVA_NUM {
		nextIndex = 0
	}

	// Get SUN's degree
	var sunDegree float32
	for _, graha := range c.Bhavas[sunIndex].Grahas {
		if graha.Name == SUN {
			sunDegree = graha.Degree
		}
	}

	// get combustion of all grahas in the same bhava as SUN
	for _, graha := range c.Bhavas[sunIndex].Grahas {
		if graha.Name != SUN {
			distance := math.Abs(float64(graha.Degree - sunDegree))
			evaluateCombustion(&graha, float32(distance))
		}
	}

	// get combustion of all grahas in the previous bhava of SUN
	for _, graha := range c.Bhavas[prevIndex].Grahas {
		distance := math.Abs(float64((graha.Degree - 30) - sunDegree))
		evaluateCombustion(&graha, float32(distance))
	}

	// get combustion of all grahas in the next bhava of SUN
	for _, graha := range c.Bhavas[nextIndex].Grahas {
		distance := math.Abs(float64((graha.Degree + 30) - sunDegree))
		evaluateCombustion(&graha, float32(distance))
	}
}

func evaluateCombustion(graha *GrahaLocCombust, distanceFromSun float32) {
	switch graha.Name {
	case MERCURY:
		if graha.Retrograde {
			if distanceFromSun <= 12.0 {
				graha.Combust = true
			}
		} else if distanceFromSun <= 14.0 {
			graha.Combust = true
		}

	case VENUS:
		if graha.Retrograde {
			if distanceFromSun <= 8.0 {
				graha.Combust = true
			}
		} else if distanceFromSun <= 10.0 {
			graha.Combust = true
		}

	case MARS:
		if distanceFromSun <= 17.0 {
			graha.Combust = true
		}

	case JUPITER:
		if distanceFromSun <= 11.0 {
			graha.Combust = true
		}

	case SATURN:
		if distanceFromSun <= 15.0 {
			graha.Combust = true
		}
	}
}

/*
func (c *Chart) GetEffectiveFriends(name string) []string {
	for _, ga := range c.GrahasAttr {
		if ga.Name == name {
			friends := make([]string, len(ga.EffectiveBestFriends))
			m := make(map[string]bool)

			for _, item := range ga.EffectiveBestFriends {
				m[item] = true
				friends = append(friends, item)
			}

			for _, item := range ga.EffectiveFriends {
				if _, ok := m[item]; !ok {
					friends = append(friends, item)
				}
			}

			return friends
		}
	}
	return nil
}

func (c *Chart) GetEffectiveNeutrals(name string) []string {
	for _, ga := range c.GrahasAttr {
		if ga.Name == name {
			return ga.EffectiveNeutrals
		}
	}
	return nil
}

func (c *Chart) GetEffectiveEnemies(name string) []string {
	for _, ga := range c.GrahasAttr {
		if ga.Name == name {
			enemies := make([]string, len(ga.EffectiveWorstEnemies))
			m := make(map[string]bool)

			for _, item := range ga.EffectiveWorstEnemies {
				m[item] = true
				enemies = append(enemies, item)
			}

			for _, item := range ga.EffectiveEnemies {
				if _, ok := m[item]; !ok {
					enemies = append(enemies, item)
				}
			}

			return enemies
		}
	}
	return nil
}

func (c *Chart) EvaluateAspects() {
	for i, b := range c.Bhavas {
		for _, g := range b.Grahas {
			if g.Name == LAGNA {
				continue
			}
			aspectedBhava := c.GetNthBhava(i, 3)
			if g.Name == SATURN {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.QuaterAspect = append(aspectedBhava.QuaterAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 4)
			if g.Name == MARS {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.ThreeQuaterAspect = append(aspectedBhava.ThreeQuaterAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 5)
			if g.Name == JUPITER {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.HalfAspect = append(aspectedBhava.HalfAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 7)
			aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)

			aspectedBhava = c.GetNthBhava(i, 8)
			if g.Name == MARS {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.ThreeQuaterAspect = append(aspectedBhava.ThreeQuaterAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 9)
			if g.Name == JUPITER {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.HalfAspect = append(aspectedBhava.HalfAspect, g.Name)
			}

			aspectedBhava = c.GetNthBhava(i, 10)
			if g.Name == SATURN {
				aspectedBhava.FullAspect = append(aspectedBhava.FullAspect, g.Name)
			} else {
				aspectedBhava.QuaterAspect = append(aspectedBhava.QuaterAspect, g.Name)
			}
		}
	}
}

*/
