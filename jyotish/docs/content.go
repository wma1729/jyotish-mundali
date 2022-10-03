package docs

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type TOCEntry struct {
	Name   string `json:"name"`
	Order  int    `json:"order"`
	Header string `json:"header"`
}

type HTMLTable struct {
	Caption string     `json:"caption"`
	Headers []string   `json:"headers"`
	Rows    [][]string `json:"rows"`
}

type SubSection struct {
	Header   string      `json:"header"`
	Content  []string    `json:"content"`
	Pictures []string    `json:"pictures"`
	Tables   []HTMLTable `json:"tables"`
}

type Section struct {
	Intro       []string     `json:"intro"`
	SubSections []SubSection `json:"sub-sections"`
	Remarks     []string     `json:"remarks"`
}

type Content struct {
	TocEntry TOCEntry `json:"toc-entry"`
	Section  Section  `json:"section"`
}

func loadFileContent(fileName string, content *Content) error {
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s: %s", fileName, err)
		return err
	}

	err = json.Unmarshal(fileContent, content)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s: %s", fileName, err)
		return err
	}

	return nil
}

func loadDocumentation(lang string) ([]Content, error) {
	files, err := os.ReadDir("./docs")
	if err != nil {
		log.Printf("failed to reader dir ./docs: %s", err)
		return nil, err
	}

	contents := []Content{}

	for _, f := range files {
		if strings.HasSuffix(f.Name(), "-"+lang+".json") {
			var c Content
			if loadFileContent("./docs/"+f.Name(), &c) == nil {
				contents = append(contents, c)
			}
		}
	}

	return contents, nil
}

func arrangeDocumentation(contents []Content) {
	sort.Slice(contents, func(i, j int) bool {
		return contents[i].TocEntry.Order < contents[j].TocEntry.Order
	})
}

func PrepareDocumentation(lang string) ([]Content, error) {
	contents, err := loadDocumentation(lang)
	if err != nil {
		return nil, err
	}

	arrangeDocumentation(contents)
	return contents, nil
}
