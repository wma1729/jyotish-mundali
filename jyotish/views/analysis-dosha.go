package views

import (
	"jyotish/analysis"
)

func GetMaanglikDosha(c *analysis.Chart) *analysis.MaanglikDosha {
	return c.EvaluateMaanglikDosha()
}

func GetKalaSarpaDosha(c *analysis.Chart) *analysis.KalaSarpaDosha {
	return c.EvaluateKalaSarpaDosha()
}

func GetDebilitatedGrahas(c *analysis.Chart) []analysis.DebilitatedGraha {
	return c.EvaluateDebilitatedGrahas()
}
