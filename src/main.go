package controllers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"text/template"
)

var validGetPath = regexp.MustCompile("^/(frontpage|404)$")

func RegisterTemplatesAndRoutes(templates *template.Template) {
	postController := new(PostController)
	http.HandleFunc("/post/", makeHandler(postController.get, templates))
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, int, *template.Template), templates *template.Template) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		fmt.Print("Call back in makehandler was called")
		match := validGetPath.FindStringSubmatch(request.URL.Path)
		if match == nil {
			http.NotFound(responseWriter, request)
			return
		}
		fmt.Printf("%v", match)
		id, err := strconv.Atoi(match[2])
		if err == nil {
			fn(responseWriter, request, id, templates)
		} else {
			panic(err.Error())
		}
	}
}
