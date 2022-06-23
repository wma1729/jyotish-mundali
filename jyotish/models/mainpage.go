package models

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type MainPage struct {
	Home          string `json:"home"`
	Profiles      string `json:"profiles"`
	KnowledgeBase string `json:"knowledgebase"`
	FAQs          string `json:"faqs"`
	SiteAdmin     string `json:"siteadmin"`
	UserName      string
}

func (page *MainPage) Load(lang string, user string) error {
	fileName := fmt.Sprintf("lang/%s/mainpage.json", lang)
	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf("failed to open %s", fileName)
		return err
	}

	err = json.Unmarshal(fileContent, page)
	if err != nil {
		log.Printf("failed to unmarshal contents of %s", fileName)
		return err
	}

	page.UserName = user

	return nil
}

func (page *MainPage) Send(w http.ResponseWriter) error {
	tmplName := "main"
	tmpl := template.Must(template.New(tmplName).ParseFiles("templates/main.html", "templates/header.html", "templates/navbar.html"))
	return tmpl.ExecuteTemplate(w, tmplName, page)
}
