package models

import (
	"sort"
)

func (chart *GrahaDetails) GetLagnaRashi() int {
	for _, graha := range chart.Grahas {
		if graha.Name == LAGNA {
			return graha.RashiNum
		}
	}
	return -1
}

func GetBhavasFromChart(chart GrahaDetails) []Bhava {
	var bhavas [MAX_BHAVA_NUM]Bhava

	lagnaRashi := chart.GetLagnaRashi()

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
		for j := 0; j < len(chart.Grahas); j++ {
			if bhavas[i].RashiNum == chart.Grahas[j].RashiNum {
				bhavas[i].Grahas = append(bhavas[i].Grahas, chart.Grahas[j])
			}
		}
		sort.Slice(bhavas[i].Grahas, func(x, y int) bool {
			return bhavas[i].Grahas[x].Degree > bhavas[i].Grahas[y].Degree
		})
	}

	return bhavas[:]
}
