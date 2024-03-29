package models

import (
	"encoding/json"
	"io/ioutil"
	"jyotish/analysis"
	"log"
)

type Language struct {
	Action            string `json:"action"`
	Aspect            string `json:"aspect"`
	Astrologer        string `json:"astrologer"`
	Attributes        string `json:"attributes"`
	Benefic           string `json:"benefic"`
	BestFriends       string `json:"best-friends"`
	Bhava             string `json:"bhava"`
	BirthDetails      string `json:"birthdetails"`
	Chart             string `json:"chart"`
	ChartDetails      string `json:"chartdetails"`
	City              string `json:"city"`
	Combust           string `json:"combust"`
	Contact           string `json:"contact"`
	Country           string `json:"country"`
	CreateProfile     string `json:"create-profile"`
	DateOfBirth       string `json:"dob"`
	Debilitated       string `json:"debilitated"`
	DegreeInRashi     string `json:"degree-in-rashi"`
	Description       string `json:"description"`
	Effective         string `json:"effective"`
	Enemies           string `json:"enemies"`
	EnemyRashi        string `json:"enemy-rashi"`
	English           string `json:"english"`
	Exalted           string `json:"exalted"`
	FAQs              string `json:"faqs"`
	Forward           string `json:"forward"`
	FriendlyRashi     string `json:"friendly-rashi"`
	Friends           string `json:"friends"`
	FullAspect        string `json:"full-aspect"`
	Graha             string `json:"graha"`
	HalfAspect        string `json:"half-aspect"`
	Hindi             string `json:"hindi"`
	Home              string `json:"home"`
	Jupiter           string `json:"jupiter"`
	Ketu              string `json:"ketu"`
	KnowledgeBase     string `json:"knowledgebase"`
	Lagna             string `json:"lagna"`
	Language          string `json:"language"`
	Logout            string `json:"logout"`
	Malefic           string `json:"malefic"`
	Mars              string `json:"mars"`
	Mercury           string `json:"mercury"`
	MoolTrikona       string `json:"mool-trikona"`
	Moon              string `json:"moon"`
	Motion            string `json:"motion"`
	Name              string `json:"name"`
	Natural           string `json:"natural"`
	Nature            string `json:"nature"`
	NeutralRashi      string `json:"neutral-rashi"`
	Neutrals          string `json:"neutrals"`
	No                string `json:"no"`
	OwnRashi          string `json:"own-rashi"`
	PlaceOfBirth      string `json:"pob"`
	Position          string `json:"position"`
	Preferences       string `json:"preferences"`
	Profiles          string `json:"profiles"`
	Public            string `json:"public"`
	QuaterAspect      string `json:"quater-aspect"`
	Rahu              string `json:"rahu"`
	RashiNumber       string `json:"rashi-number"`
	Remarks           string `json:"remarks"`
	Retrograde        string `json:"retrograde"`
	Saturn            string `json:"saturn"`
	Save              string `json:"save"`
	SiteAdmin         string `json:"siteadmin"`
	State             string `json:"state"`
	Sun               string `json:"sun"`
	Temporary         string `json:"temporary"`
	ThreeQuaterAspect string `json:"three-quater-aspect"`
	Venus             string `json:"venus"`
	WorstEnemies      string `json:"worst-enemies"`
	Yes               string `json:"yes"`
}

var EnglishVocab, HindiVocab Language

func loadFileContent(fileName string, language *Language) {
	fileContent, err := ioutil.ReadFile(fileName)
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
	case analysis.LAGNA:
		return vocab.Lagna
	case analysis.SUN:
		return vocab.Sun
	case analysis.MOON:
		return vocab.Moon
	case analysis.MARS:
		return vocab.Mars
	case analysis.MERCURY:
		return vocab.Mercury
	case analysis.JUPITER:
		return vocab.Jupiter
	case analysis.VENUS:
		return vocab.Venus
	case analysis.SATURN:
		return vocab.Saturn
	case analysis.RAHU:
		return vocab.Rahu
	case analysis.KETU:
		return vocab.Ketu
	}

	return ""
}

func GrahaNature(nature string, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch nature {
	case analysis.BENEFIC:
		return vocab.Benefic
	case analysis.MALEFIC:
		return vocab.Malefic
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

func GrahaPosition(position string, lang string) string {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	switch position {
	case analysis.RASHI_EXALTED:
		return vocab.Exalted
	case analysis.RASHI_DEBILITATED:
		return vocab.Debilitated
	case analysis.RASHI_MOOLTRIKONA:
		return vocab.MoolTrikona
	case analysis.RASHI_OWN:
		return vocab.OwnRashi
	case analysis.RASHI_FRIENDLY:
		return vocab.FriendlyRashi
	case analysis.RASHI_NEUTRAL:
		return vocab.NeutralRashi
	case analysis.RASHI_ENEMY:
		return vocab.EnemyRashi
	}

	return "-"
}
