package analysis

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

type GrahaDetais struct {
	Grahas []Graha
}

type Chart struct {
	Bhavas     []Bhava
	GrahasAttr []GrahaAttributes
}

func (gd *GrahaDetais) Value() (driver.Value, error) {
	return json.Marshal(gd)
}

func (gd *GrahaDetais) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprintf("unexpected value type: expected []byte, found %T", value))
	}
	json.Unmarshal(b, gd)
	return nil
}

func (gd *GrahaDetais) GetLagnaRashi() int {
	for _, graha := range gd.Grahas {
		if graha.Name == LAGNA {
			return graha.RashiNum
		}
	}
	return -1
}

func GetChart(gd GrahaDetais) Chart {
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
