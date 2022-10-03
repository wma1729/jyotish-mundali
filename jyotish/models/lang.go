package models

import (
	"encoding/json"
	"io/ioutil"
	"jyotish/analysis"
	"log"
)

type Language struct {
	Action        string `json:"action"`
	Astrologer    string `json:"astrologer"`
	Attributes    string `json:"attributes"`
	Benefic       string `json:"benefic"`
	BirthDetails  string `json:"birthdetails"`
	ChartDetails  string `json:"chartdetails"`
	City          string `json:"city"`
	Combust       string `json:"combust"`
	Contact       string `json:"contact"`
	Country       string `json:"country"`
	CreateProfile string `json:"create-profile"`
	DateOfBirth   string `json:"dob"`
	DegreeInRashi string `json:"degree-in-rashi"`
	Description   string `json:"description"`
	Effective     string `json:"effective"`
	Enemies       string `json:"enemies"`
	English       string `json:"english"`
	FAQs          string `json:"faqs"`
	Forward       string `json:"forward"`
	Friends       string `json:"friends"`
	Graha         string `json:"graha"`
	Hindi         string `json:"hindi"`
	Home          string `json:"home"`
	Jupiter       string `json:"jupiter"`
	Ketu          string `json:"ketu"`
	KnowledgeBase string `json:"knowledgebase"`
	Lagna         string `json:"lagna"`
	Language      string `json:"language"`
	Logout        string `json:"logout"`
	Malefic       string `json:"malefic"`
	Mars          string `json:"mars"`
	Mercury       string `json:"mercury"`
	Moon          string `json:"moon"`
	Motion        string `json:"motion"`
	Name          string `json:"name"`
	Natural       string `json:"natural"`
	Nature        string `json:"nature"`
	Neutrals      string `json:"neutrals"`
	No            string `json:"no"`
	PlaceOfBirth  string `json:"pob"`
	Preferences   string `json:"preferences"`
	Profiles      string `json:"profiles"`
	Public        string `json:"public"`
	Rahu          string `json:"rahu"`
	RashiNumber   string `json:"rashi-number"`
	Remarks       string `json:"remarks"`
	Retrograde    string `json:"retrograde"`
	Saturn        string `json:"saturn"`
	Save          string `json:"save"`
	SiteAdmin     string `json:"siteadmin"`
	State         string `json:"state"`
	Sun           string `json:"sun"`
	Temporary     string `json:"temporary"`
	Venus         string `json:"venus"`
	Yes           string `json:"yes"`
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
