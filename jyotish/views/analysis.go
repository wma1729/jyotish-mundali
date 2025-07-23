package views

import (
	"fmt"
	"html/template"
	"jyotish/analysis"
	"jyotish/constants"
	"jyotish/models"
	"net/http"
	"strings"
)

type AnalysisPage struct {
	MainPage
	Chart     analysis.Chart
	BhavaDesc []models.BhavaDescription
}

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

func GetBhava(chart *analysis.Chart, bhavaNum int) analysis.Bhava {
	return chart.Bhavas[bhavaNum-1]
}

func GetGrahaName(graha string, lang string) string {
	return models.GrahaName(graha, lang)
}

func ListOfGrahas(grahas []string, lang string) string {
	sb := strings.Builder{}
	first := true
	for _, g := range grahas {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(models.GrahaName(g, lang))
		first = false
	}
	return sb.String()
}

func ListOfIntegers(numbers []int) string {
	sb := strings.Builder{}
	first := true
	for _, n := range numbers {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%d", n))
		first = false
	}
	return sb.String()
}

func ListOfAspectingGrahas(aspectingGrahas []analysis.AscpectAndDegree, lang string) string {
	sb := strings.Builder{}
	first := true
	for _, ag := range aspectingGrahas {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%s (%d)", models.GrahaName(ag.Name, lang), ag.Degree))
		first = false
	}
	return sb.String()
}

func GetInfluenceOnBhava(assoc []int, lang string) string {
	sb := strings.Builder{}
	first := true
	for _, a := range assoc {
		if !first {
			sb.WriteString(", ")
		}
		sb.WriteString(models.GetInfluenceOnBhava(a, lang))
		first = false
	}
	return sb.String()
}

func GetInfluenceRating(category string, gir analysis.GrahaInfluenceRating, lang string) template.HTML {
	var vocab *models.Language

	if lang == "en" {
		vocab = &models.EnglishVocab
	} else {
		vocab = &models.HindiVocab
	}

	sb := strings.Builder{}

	switch category {
	case "bhava-karaka":
		if gir.Rating == constants.MALEFIC {
			sb.WriteString(fmt.Sprintf(`<span class="bad">%s`,
				models.GrahaName(gir.Value.(string), lang)))
		} else {
			sb.WriteString(`<span class="neutral">-`)
		}
	case "distance", "position", "owner", "aspectual-strength":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="good">%d`, gir.Value.(int)))
		case constants.MALEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="bad">%d`, gir.Value.(int)))
		default:
			sb.WriteString(fmt.Sprintf(`<span class="neutral">%d`, gir.Value.(int)))
		}
	case "nature":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="good">%s`,
					models.GrahaNature(gir.Value.(int), lang)))
		case constants.MALEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="bad">%s`,
					models.GrahaNature(gir.Value.(int), lang)))
		default:
			sb.WriteString(
				fmt.Sprintf(`<span class="neutral">%s`,
					models.GrahaNature(gir.Value.(int), lang)))
		}
	case "relation":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="good">%s`, vocab.Friendly))
		case constants.MALEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="bad">%s`, vocab.Inimical))
		default:
			sb.WriteString(fmt.Sprintf(`<span class="neutral">%s`, vocab.Neutral))
		}
	case "position-strength":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="good">%s`,
					models.GrahaPosition(gir.Value.(int), lang)))
		case constants.MALEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="bad">%s`,
					models.GrahaPosition(gir.Value.(int), lang)))
		default:
			sb.WriteString(
				fmt.Sprintf(`<span class="neutral">%s`,
					models.GrahaPosition(gir.Value.(int), lang)))
		}
	case "combust":
		if gir.Rating == constants.MALEFIC {
			sb.WriteString(
				fmt.Sprintf(`<span class="bad">%s`, vocab.Yes))
		} else {
			sb.WriteString(
				fmt.Sprintf(`<span class="neutral">%s`, vocab.No))
		}
	case "retrograde":
		var motion string
		if gir.Value == 1 {
			motion = vocab.Retrograde
		} else {
			motion = vocab.Forward
		}

		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="good">%s`, motion))
		case constants.MALEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="bad">%s`, motion))
		default:
			sb.WriteString(fmt.Sprintf(`<span class="neutral">%s`, motion))
		}
	case "direction-strength":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="good">%.2f`, gir.Value.(float64)))
		case constants.MALEFIC:
			sb.WriteString(fmt.Sprintf(`<span class="bad">%.2f`, gir.Value.(float64)))
		default:
			sb.WriteString(fmt.Sprintf(`<span class="neutral">%.2f`, gir.Value.(float64)))
		}
	case "state-strength":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="good">%s`,
					models.GrahaState(gir.Value.(int), lang)))
		case constants.MALEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="bad">%s`,
					models.GrahaState(gir.Value.(int), lang)))
		default:
			sb.WriteString(
				fmt.Sprintf(`<span class="neutral">%s`,
					models.GrahaState(gir.Value.(int), lang)))
		}
	case "bhava-flanked-by":
		switch gir.Rating {
		case constants.BENEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="good">%s`,
					ListOfGrahas(gir.Value.([]string), lang)))
		case constants.MALEFIC:
			sb.WriteString(
				fmt.Sprintf(`<span class="bad">%s`,
					ListOfGrahas(gir.Value.([]string), lang)))
		}
	}

	switch gir.Notes {
	case constants.SUBJECTS_CHILDREN:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsChildren))
	case constants.SUBJECTS_ELDER_SIBLINGS:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsElderSiblings))
	case constants.SUBJECTS_FATHER:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsFather))
	case constants.SUBJECTS_LIVING_BEING:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsLiving))
	case constants.SUBJECTS_MOTHER:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsMother))
	case constants.SUBJECTS_NON_LIVING_BEING:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsNonLiving))
	case constants.SUBJECTS_SPOUSE:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsSpouse))
	case constants.SUBJECTS_YOUNGER_SIBLINGS:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.SubjectsYoungerSiblings))
	case constants.BHAVA_LORD:
		sb.WriteString(fmt.Sprintf(" - %s", vocab.OwnRashi))
	}

	sb.WriteString("</span>")

	return template.HTML(sb.String())
}

func GetRashiName(number int, lang string) string {
	return fmt.Sprintf("%s (%d)", models.RashiName(number, lang), number)
}

func GetAnalysisPage(user *models.User, chart analysis.Chart) (*AnalysisPage, error) {
	var page AnalysisPage

	if user.Lang == "en" {
		page.Vocab = &models.EnglishVocab
	} else {
		page.Vocab = &models.HindiVocab
	}

	page.User = user
	page.Chart = chart
	page.BhavaDesc, _ = models.GetBhavaDescription(user.Lang)

	return &page, nil

}

func (page *AnalysisPage) Send(w http.ResponseWriter) error {
	tmplName := "analysis"
	tmpl := template.Must(template.New(tmplName).Funcs(
		template.FuncMap{
			"GetBhava":              GetBhava,
			"GetGrahaNameForChart":  GetGrahaNameForChart,
			"GetGrahaName":          GetGrahaName,
			"ListOfAspectingGrahas": ListOfAspectingGrahas,
			"ListOfGrahas":          ListOfGrahas,
			"ListOfIntegers":        ListOfIntegers,
			"GrahaNature":           models.GrahaNature,
			"GrahaMotion":           models.GrahaMotion,
			"YesOrNo":               models.YesOrNo,
			"GrahaPosition":         models.GrahaPosition,
			"GrahaState":            models.GrahaState,
			"GetRashiName":          GetRashiName,
			"GetInfluenceRating":    GetInfluenceRating,
			"GetInfluenceOnBhava":   GetInfluenceOnBhava,
		}).ParseFiles(
		"templates/analysis.html",
		"templates/lagna-chart.html",
		"templates/grahas-relations.html",
		"templates/grahas-aspects.html",
		"templates/grahas-nature.html",
		"templates/grahas-strength.html",
		"templates/bhavas.html",
		"templates/header.html",
		"templates/navbar.html",
		"templates/footer.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
