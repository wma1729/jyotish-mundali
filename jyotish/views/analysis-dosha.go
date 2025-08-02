package views

import (
	"jyotish/analysis"
)

func GetMaanglikDosha(c *analysis.Chart) *analysis.MaanglikDosha {
	return c.EvaluateMaanglikDosha()
}
