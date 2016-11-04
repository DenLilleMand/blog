package controllers

import (
	"html/template"
	"net/http"
)

type PostController struct{}

func (postController *PostController) viewHandler(responseWriter http.ResponseWriter, request *http.Request, title string) {
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(responseWriter, request, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(responseWriter, "view", page)
}
