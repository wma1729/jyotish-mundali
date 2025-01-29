package analysis

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"sort"
)

type GrahaDetails struct {
	Grahas []Graha
}

type Chart struct {
	Bhavas     []Bhava
	GrahasAttr []GrahaAttributes
}

func (gd *GrahaDetails) Value() (driver.Value, error) {
	return json.Marshal(gd)
}

func (gd *GrahaDetails) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("unexpected value type: expected []byte, found %T", value)
	}
	json.Unmarshal(b, gd)
	return nil
}

func (gd *GrahaDetails) GetLagnaRashi() int {
	for _, graha := range gd.Grahas {
		if graha.Name == LAGNA {
			return graha.RashiNum
		}
	}
	return -1
}

func GetChart(gd GrahaDetails) Chart {
	var bhavas [MAX_BHAVA_NUM]Bhava

	lagnaRashi := gd.GetLagnaRashi()

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
		for j := 0; j < len(gd.Grahas); j++ {
			if bhavas[i].RashiNum == gd.Grahas[j].RashiNum {
				bhavas[i].Grahas = append(bhavas[i].Grahas, gd.Grahas[j])
			}
		}
		sort.Slice(bhavas[i].Grahas, func(x, y int) bool {
			return bhavas[i].Grahas[x].Degree > bhavas[i].Grahas[y].Degree
		})
	}

	var chart Chart
	chart.Bhavas = bhavas[:]
	chart.EvaluateAspects()

	chart.GrahasAttr = make([]GrahaAttributes, 9)
	chart.GrahasAttr[0].Init(SUN, &chart)
	chart.GrahasAttr[1].Init(MOON, &chart)
	chart.GrahasAttr[2].Init(MARS, &chart)
	chart.GrahasAttr[3].Init(MERCURY, &chart)
	chart.GrahasAttr[4].Init(JUPITER, &chart)
	chart.GrahasAttr[5].Init(VENUS, &chart)
	chart.GrahasAttr[6].Init(SATURN, &chart)
	chart.GrahasAttr[7].Init(RAHU, &chart)
	chart.GrahasAttr[8].Init(KETU, &chart)

	chart.GrahasAttr[0].GetGrahaPosition(SUN, &chart)
	chart.GrahasAttr[1].GetGrahaPosition(MOON, &chart)
	chart.GrahasAttr[2].GetGrahaPosition(MARS, &chart)
	chart.GrahasAttr[3].GetGrahaPosition(MERCURY, &chart)
	chart.GrahasAttr[4].GetGrahaPosition(JUPITER, &chart)
	chart.GrahasAttr[5].GetGrahaPosition(VENUS, &chart)
	chart.GrahasAttr[6].GetGrahaPosition(SATURN, &chart)
	chart.GrahasAttr[7].GetGrahaPosition(RAHU, &chart)
	chart.GrahasAttr[8].GetGrahaPosition(KETU, &chart)

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
