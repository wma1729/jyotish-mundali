package models

import (
	"jyotish/constants"
	"jyotish/docs"
)

type BhavaDescription struct {
	Number        int
	DisplayNumber string
	Name          string
	BodyParts     string
	Relations     string
	Meanings      string
	Significator  string
	Efforts       string
}

func loadBhavaDocumentation(lang string) (docs.Content, error) {
	var content docs.Content

	var err = docs.LoadFileContent("./docs/bhava-"+lang+".yaml", &content)

	return content, err
}

func GetBhavaDescription(lang string) ([]BhavaDescription, error) {
	var vocab *Language

	if lang == "en" {
		vocab = &EnglishVocab
	} else {
		vocab = &HindiVocab
	}

	content, err := loadBhavaDocumentation(lang)
	if err != nil {
		return nil, err
	}

	bhavaDescription := make([]BhavaDescription, 12)

	for _, subSection := range content.Section.SubSections {
		if subSection.Header == vocab.Bhava {
			for i := 0; i < constants.MAX_BHAVA_NUM; i++ {
				bhavaDescription[i].Number = i + 1
				bhavaDescription[i].DisplayNumber = subSection.Tables[0].Rows[i][0]
				bhavaDescription[i].Name = subSection.Tables[0].Rows[i][1]
				bhavaDescription[i].BodyParts = subSection.Tables[0].Rows[i][2]
				bhavaDescription[i].Relations = subSection.Tables[0].Rows[i][3]
				bhavaDescription[i].Meanings = subSection.Tables[0].Rows[i][4]
				bhavaDescription[i].Significator = subSection.Tables[0].Rows[i][5]
				bhavaDescription[i].Efforts = subSection.Tables[0].Rows[i][6]
			}
		}
	}

	return bhavaDescription, nil
}
