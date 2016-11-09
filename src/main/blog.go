package main

import (
	"fmt"
	"github.com/denlillemand/blog/src/controllers"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
)

func main() {
	templates := populateTemplates()
	controllers.RegisterTemplatesAndRoutes(templates)
	http.ListenAndServe(":3000", nil)
}

//@TODO: Make this method work recursively, so we can have a nested structure of templates. For math, programming, vim etc.
func populateTemplates() *template.Template {
	fmt.Print("")
	result := template.New("templates")

	basePath, _ := filepath.Abs("../../templates")
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathsRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
