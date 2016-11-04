package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html"))
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getTitle(responseWriter http.ResponseWriter, request *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(request.URL.Path)
	if m == nil {
		http.NotFound(responseWriter, request)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}

func editHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}
	renderTemplate(responseWriter, "edit", page)
}

func viewHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(responseWriter, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(responseWriter, "view", page)
}

func saveHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	body := request.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(responseWriter, request, "/view/"+title, http.StatusFound)
}

func renderTemplate(responseWriter http.ResponseWriter, templateName string, page *Page) {
	tmpl, err := template.ParseFiles(templateName + ".html")
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
	}
	err = tmpl.Execute(responseWriter, page)
	if err != nil {
		http.Error(responseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		match := validPath.FindStringSubmatch(request.URL.Path)
		if match == nil {
			http.NotFound(responseWriter, request)
			return
		}
		fmt.Printf("%v", match)
		fn(responseWriter, request, match[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.ListenAndServe(":3000", nil)
}
