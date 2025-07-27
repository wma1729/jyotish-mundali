package views

import (
	"fmt"
	"html/template"
	"jyotish/analysis"
	"jyotish/models"
	"strings"
)

func GetGrahaNameForChart(graha analysis.GrahaLocCombust, lang string) template.HTML {
	vocab := &models.EnglishVocab
	if lang != "en" {
		vocab = &models.HindiVocab
	}

	sb := strings.Builder{}
	sb.WriteString(models.GrahaName(graha.Name, lang))
	sb.WriteString("<sup>")
	sb.WriteString(fmt.Sprintf("%.f", graha.Degree))
	sb.WriteString("</sup>")

	closeParenthesis, addComma := false, false
	if graha.Combust || graha.Retrograde {
		sb.WriteByte('(')
		closeParenthesis = true
	}
	if graha.Combust {
		sb.WriteString(vocab.CombustAbbr)
		addComma = true
	}
	if graha.Retrograde {
		if addComma {
			sb.WriteByte(',')
		}
		sb.WriteString(vocab.RetrogradeAbbr)
	}
	if closeParenthesis {
		sb.WriteByte(')')
	}

	return template.HTML(sb.String())
}
