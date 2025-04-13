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
	Meaning       string
	Indicator     string
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
			var bhavaDesc BhavaDescription
			for i := 0; i < constants.MAX_BHAVA_NUM; i++ {
				bhavaDesc.Number = i + 1
				bhavaDesc.DisplayNumber = subSection.Tables[0].Rows[i][0]
				bhavaDesc.Name = subSection.Tables[0].Rows[i][1]
				bhavaDesc.BodyParts = subSection.Tables[0].Rows[i][2]
				bhavaDesc.Relations = subSection.Tables[0].Rows[i][3]
				bhavaDesc.Meaning = subSection.Tables[0].Rows[i][4]
				bhavaDesc.Indicator = subSection.Tables[0].Rows[i][5]
				bhavaDesc.Efforts = subSection.Tables[0].Rows[i][6]
			}
			bhavaDescription = append(bhavaDescription, bhavaDesc)
		}
	}

	return bhavaDescription, nil
}
