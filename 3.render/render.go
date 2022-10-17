package render

import (
	"net/http"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Data      map[string]interface{}
	CSRF      string
	Flash     string
	Warning   string
	Error     string
	IP        string
}

var functions = template.FuncMap{}

func Renderer(w http.ResponseWriter, r *http.Request, tmpl string, tempData *TemplateData) {

	pages, err := filepath.Glob("./*.page.html")
	Scream(err)

	//parsing pages and layouts
	for _, page := range pages {

		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		Scream(err)
		ts, _ = ts.ParseGlob("./*.layout.html")
		if name == tmpl {
			ts.Execute(w, tempData)
		}
	}

	Scream(err)
}

func Scream(e error) {
	if e != nil {
		panic(e)
	}
}
