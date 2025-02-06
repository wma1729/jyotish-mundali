package analysis

import (
	"jyotish/constants"
	"jyotish/models"
	"math"
	"sort"
)

type GrahaAttributes struct {
	Relations GrahaRelations
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
	chart.findCombustGrahas()

	chart.EvaluateGrahaAttributes()

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
	if bn >= constants.MAX_BHAVA_NUM {
		bn -= constants.MAX_BHAVA_NUM
	}
	return &c.Bhavas[bn]
}

func (c *Chart) NthBhavaContainsGraha(i, n int, graha string) bool {
	b := c.GetNthBhava(i, n)
	return b.ContainsGraha(graha)
}

func evaluateCombustion(graha *GrahaLocCombust, distanceFromSun float32) {
	switch graha.Name {
	case constants.MERCURY:
		if graha.Retrograde {
			if distanceFromSun <= 12.0 {
				graha.Combust = true
			}
		} else if distanceFromSun <= 14.0 {
			graha.Combust = true
		}

	case constants.VENUS:
		if graha.Retrograde {
			if distanceFromSun <= 8.0 {
				graha.Combust = true
			}
		} else if distanceFromSun <= 10.0 {
			graha.Combust = true
		}

	case constants.MARS:
		if distanceFromSun <= 17.0 {
			graha.Combust = true
		}

	case constants.JUPITER:
		if distanceFromSun <= 11.0 {
			graha.Combust = true
		}

	case constants.SATURN:
		if distanceFromSun <= 15.0 {
			graha.Combust = true
		}
	}
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
	var sunDegree float32
	for _, graha := range c.Bhavas[sunIndex].Grahas {
		if graha.Name == constants.SUN {
			sunDegree = graha.Degree
		}
	}

	// get combustion of all grahas in the same bhava as SUN
	for _, graha := range c.Bhavas[sunIndex].Grahas {
		if graha.Name != constants.SUN {
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

func (c *Chart) EvaluateGrahaAttributes() {
	c.GrahasAttr = make([]GrahaAttributes, 9)
	for i, graha := range constants.GrahaNames {
		c.GrahasAttr[i].Relations.EvaluateGrahaRelations(graha, c)
	}
}
