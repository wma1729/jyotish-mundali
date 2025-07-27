package views

import (
	"fmt"
	"html/template"
	"jyotish/analysis"
	"jyotish/constants"
	"jyotish/models"
	"strings"
)

func GetMaanglikDosha(c *analysis.Chart) *analysis.MaanglikDosha {
	return c.EvaluateMaanglikDosha()
}

func GetSeverity(sev int, lang string) template.HTML {
	sb := strings.Builder{}

	switch sev {
	case constants.MINIMUM:
		sb.WriteString(fmt.Sprintf(`<span class="good">%s`, models.GetSeverity(sev, lang)))
	case constants.AVERAGE:
		sb.WriteString(fmt.Sprintf(`<span class="neutral">%s`, models.GetSeverity(sev, lang)))
	case constants.MAXIMUM:
		sb.WriteString(fmt.Sprintf(`<span class="bad">%s`, models.GetSeverity(sev, lang)))
	}

	sb.WriteString("</span>")

	return template.HTML(sb.String())
}
