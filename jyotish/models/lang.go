package models

import (
	"encoding/json"
	"jyotish/constants"
	"log"
	"os"
)

type Language struct {
	Action                   string `json:"action"`
	Adult                    string `json:"adult"`
	Analysis                 string `json:"analysis"`
	Aries                    string `json:"aries"`
	Aquarius                 string `json:"aquarius"`
	Aspect                   string `json:"aspect"`
	Astrologer               string `json:"astrologer"`
	Attributes               string `json:"attributes"`
	Average                  string `json:"average"`
	Benefic                  string `json:"benefic"`
	BestFriends              string `json:"best-friends"`
	Bhava                    string `json:"bhava"`
	BhavaAspect              string `json:"bhava-aspect"`
	BhavaPlacement           string `json:"bhava-placement"`
	BhavaOwnership           string `json:"bhava-ownership"`
	BhavaKaraka              string `json:"bhava-karaka"`
	BirthDetails             string `json:"birthdetails"`
	BodyParts                string `json:"body-parts"`
	Cancer                   string `json:"cancer"`
	Capricorn                string `json:"capricorn"`
	Chart                    string `json:"chart"`
	ChartDetails             string `json:"chartdetails"`
	Child                    string `json:"child"`
	City                     string `json:"city"`
	Combust                  string `json:"combust"`
	CombustAbbr              string `json:"combust-abbr"`
	Conjunction              string `json:"conjunction"`
	Contact                  string `json:"contact"`
	Country                  string `json:"country"`
	CreateProfile            string `json:"create-profile"`
	DateOfBirth              string `json:"dob"`
	Dead                     string `json:"dead"`
	Debilitated              string `json:"debilitated"`
	DegreeInRashi            string `json:"degree-in-rashi"`
	Description              string `json:"description"`
	DirectMeanings           string `json:"direct-meanings"`
	Directional              string `json:"directional"`
	DistanceFromBhava        string `json:"distance-from-bhava"`
	DistanceFromLagna        string `json:"distance-from-lagna"`
	Effective                string `json:"effective"`
	Efforts                  string `json:"efforts"`
	Enemies                  string `json:"enemies"`
	EnemyRashi               string `json:"enemy-rashi"`
	English                  string `json:"english"`
	Exalted                  string `json:"exalted"`
	FAQs                     string `json:"faqs"`
	Forward                  string `json:"forward"`
	FiveFold                 string `json:"five-fold"`
	Flaws                    string `json:"flaws"`
	Friendly                 string `jsons:"friendly"`
	FriendlyRashi            string `json:"friendly-rashi"`
	Friends                  string `json:"friends"`
	FullAspect               string `json:"full-aspect"`
	Functional               string `json:"functional"`
	Gemini                   string `json:"gemini"`
	Graha                    string `json:"graha"`
	GrahasOnBothSidesOfBhava string `json:"grahas-on-both-sides-of-bhava"`
	GrahaState               string `json:"graha-state"`
	HalfAspect               string `json:"half-aspect"`
	Hindi                    string `json:"hindi"`
	Home                     string `json:"home"`
	Influence                string `json:"influence"`
	Inimical                 string `json:"inimical"`
	InKendra                 string `json:"in-kendra"`
	Jupiter                  string `json:"jupiter"`
	Ketu                     string `json:"ketu"`
	KnowledgeBase            string `json:"knowledgebase"`
	Lagna                    string `json:"lagna"`
	Language                 string `json:"language"`
	Leo                      string `json:"leo"`
	Libra                    string `json:"libra"`
	Logout                   string `json:"logout"`
	Lord                     string `json:"lord"`
	Maanglik                 string `json:"maanglik"`
	Malefic                  string `json:"malefic"`
	Mars                     string `json:"mars"`
	Maximum                  string `json:"maximum"`
	Mercury                  string `json:"mercury"`
	Minimum                  string `json:"minimum"`
	MoolTrikona              string `json:"mool-trikona"`
	Moon                     string `json:"moon"`
	Motion                   string `json:"motion"`
	Name                     string `json:"name"`
	Natural                  string `json:"natural"`
	Nature                   string `json:"nature"`
	Neutral                  string `json:"neutral"`
	NeutralRashi             string `json:"neutral-rashi"`
	No                       string `json:"no"`
	Number1                  string `json:"num-1"`
	Number2                  string `json:"num-2"`
	Number3                  string `json:"num-3"`
	Number4                  string `json:"num-4"`
	Number5                  string `json:"num-5"`
	Number6                  string `json:"num-6"`
	Number7                  string `json:"num-7"`
	Number8                  string `json:"num-8"`
	Number9                  string `json:"num-9"`
	Number10                 string `json:"num-10"`
	Number11                 string `json:"num-11"`
	Number12                 string `json:"num-12"`
	Old                      string `json:"old"`
	OwnRashi                 string `json:"own-rashi"`
	Pisces                   string `json:"pisces"`
	PlaceOfBirth             string `json:"pob"`
	Position                 string `json:"position"`
	Preferences              string `json:"preferences"`
	Profiles                 string `json:"profiles"`
	Public                   string `json:"public"`
	QuarterAspect            string `json:"quarter-aspect"`
	Rahu                     string `json:"rahu"`
	Rashi                    string `json:"rashi"`
	RashiNumber              string `json:"rashi-number"`
	Relations                string `json:"relations"`
	Remarks                  string `json:"remarks"`
	Retrograde               string `json:"retrograde"`
	RetrogradeAbbr           string `json:"retrograde-abbr"`
	Sagittarius              string `json:"sagittarius"`
	Saturn                   string `json:"saturn"`
	Save                     string `json:"save"`
	Scorpio                  string `json:"scorpio"`
	Severity                 string `json:"severity"`
	SiteAdmin                string `json:"siteadmin"`
	State                    string `json:"state"`
	Strength                 string `json:"strength"`
	SubjectsChildren         string `json:"subjects-children"`
	SubjectsElderSiblings    string `json:"subjects-elder-siblings"`
	SubjectsFather           string `json:"subjects-father"`
	SubjectsLiving           string `json:"subjects-living"`
	SubjectsMother           string `json:"subjects-mother"`
	SubjectsNonLiving        string `json:"subjects-non-living"`
	SubjectsSpouse           string `json:"subjects-spouse"`
	SubjectsYoungerSiblings  string `json:"subjects-younger-siblings"`
	Sun                      string `json:"sun"`
	Taurus                   string `json:"taurus"`
	Temporary                string `json:"temporary"`
	ThreeQuarterAspect       string `json:"three-quarter-aspect"`
	Venus                    string `json:"venus"`
	Virgo                    string `json:"virgo"`
	WorstEnemies             string `json:"worst-enemies"`
	Yes                      string `json:"yes"`
	Youth                    string `json:"youth"`
}

var EnglishVocab, HindiVocab Language

func loadFileContent(fileName string, language *Language) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed to open %s", fileName)
	}

	err = json.Unmarshal(fileContent, language)
	if err != nil {
		log.Fatalf("failed to unmarshal contents of %s", fileName)
	}
}

func init() {
	loadFileContent("lang/en.json", &EnglishVocab)
	loadFileContent("lang/hi.json", &HindiVocab)
}

func GrahaName(graha string, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch graha {
	case constants.LAGNA:
		return vocab.Lagna
	case constants.SUN:
		return vocab.Sun
	case constants.MOON:
		return vocab.Moon
	case constants.MARS:
		return vocab.Mars
	case constants.MERCURY:
		return vocab.Mercury
	case constants.JUPITER:
		return vocab.Jupiter
	case constants.VENUS:
		return vocab.Venus
	case constants.SATURN:
		return vocab.Saturn
	case constants.RAHU:
		return vocab.Rahu
	case constants.KETU:
		return vocab.Ketu
	}

	return ""
}

func GrahaNature(nature int, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch nature {
	case constants.BENEFIC:
		return vocab.Benefic
	case constants.MALEFIC:
		return vocab.Malefic
	case constants.NEUTRAL:
		return vocab.Neutral
	}

	return ""
}

func GrahaMotion(retrograde bool, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	if retrograde {
		return vocab.Retrograde
	}
	return vocab.Forward
}

func YesOrNo(flag bool, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	if flag {
		return vocab.Yes
	}
	return vocab.No
}

func GrahaPosition(position int, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch position {
	case constants.IN_EXALTATION_RASHI:
		return vocab.Exalted
	case constants.IN_DEBILITATION_RASHI:
		return vocab.Debilitated
	case constants.IN_MOOLTRIKONA_RASHI:
		return vocab.MoolTrikona
	case constants.IN_OWN_RASHI:
		return vocab.OwnRashi
	case constants.IN_FRIENDLY_RASHI:
		return vocab.FriendlyRashi
	case constants.IN_NEUTRAL_RASHI:
		return vocab.NeutralRashi
	case constants.IN_INIMICAL_RASHI:
		return vocab.EnemyRashi
	}

	return "-"
}

func GrahaState(state int, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch state {
	case constants.CHILD:
		return vocab.Child
	case constants.YOUTH:
		return vocab.Youth
	case constants.ADULT:
		return vocab.Adult
	case constants.OLD:
		return vocab.Old
	case constants.DEAD:
		return vocab.Dead
	}

	return "-"
}

func RashiName(number int, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch number {
	case 1:
		return vocab.Aries
	case 2:
		return vocab.Taurus
	case 3:
		return vocab.Gemini
	case 4:
		return vocab.Cancer
	case 5:
		return vocab.Leo
	case 6:
		return vocab.Virgo
	case 7:
		return vocab.Libra
	case 8:
		return vocab.Scorpio
	case 9:
		return vocab.Sagittarius
	case 10:
		return vocab.Capricorn
	case 11:
		return vocab.Aquarius
	case 12:
		return vocab.Pisces
	}

	return "-"
}

func GetInfluenceOnBhava(assoc int, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch assoc {
	case constants.BHAVA_ASPECT:
		return vocab.BhavaAspect
	case constants.BHAVA_PLACEMENT:
		return vocab.BhavaPlacement
	case constants.BHAVA_OWNERSHIP:
		return vocab.BhavaOwnership
	case constants.BHAVA_KARAKA:
		return vocab.BhavaKaraka
	default:
		return "-"
	}
}

func GetSeverity(sev int, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch sev {
	case constants.MINIMUM:
		return vocab.Minimum
	case constants.AVERAGE:
		return vocab.Average
	case constants.MAXIMUM:
		return vocab.Maximum
	default:
		return "-"
	}
}
