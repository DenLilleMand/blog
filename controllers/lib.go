package lib

import (
	"errors"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getTitle(responseWriter http.ResponseWriter, request *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(request.URL.Path)
	if m == nil {
		http.NotFound(responseWriter, request)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil
}
