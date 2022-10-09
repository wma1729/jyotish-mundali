package docs

import (
	"log"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type TOCEntry struct {
	Name   string `yaml:"name"`
	Order  int    `yaml:"order"`
	Header string `yaml:"header"`
}

type HTMLTable struct {
	Caption string     `yaml:"caption"`
	Headers []string   `yaml:"headers"`
	Rows    [][]string `yaml:"rows"`
}

type SubSection struct {
	Header   string      `yaml:"header"`
	Content  []string    `yaml:"content"`
	Pictures []string    `yaml:"pictures"`
	Tables   []HTMLTable `yaml:"tables"`
}

type Section struct {
	Intro       []string     `yaml:"intro"`
	SubSections []SubSection `yaml:"sub-sections"`
	Remarks     []string     `yaml:"remarks"`
	Table       HTMLTable    `yaml:"table"`
}

type Content struct {
	TocEntry TOCEntry `yaml:"toc-entry"`
	Section  Section  `yaml:"section"`
}

func loadFileContent(fileName string, content *Content) error {
	f, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return err
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(content)
	if err != nil {
		log.Printf("failed to unmarshal %s: %s", fileName, err)
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
		if strings.HasSuffix(f.Name(), "-"+lang+".yaml") {
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
