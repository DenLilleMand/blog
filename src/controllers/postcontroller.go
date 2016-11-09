package controllers

import (
	"fmt"
	"github.com/denlillemand/blog/src/models/repositories"
	"net/http"
	"text/template"
)

type PostController struct{}

func NewPostController() *PostController {
	return &PostController{}
}

func (postController *PostController) get(responseWriter http.ResponseWriter, request *http.Request, id int, templates *template.Template) {
	fmt.Print("")
	postRepository := repositories.NewPostRepository()

	post, err := postRepository.Get(id)
	fmt.Print(err)
	if err != nil {
		http.Redirect(responseWriter, request, "/404", http.StatusNotFound)
		return
	}
	template := templates.Lookup(post.Title + ".html")
	template.Execute(responseWriter, post)
}
