package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/hdadashi/jabama/config"
	"github.com/hdadashi/jabama/forms"
	"github.com/justinas/nosurf"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Data      struct {
		Name  string
		Lname string
		Email string
		Phone string
	}
	CSRF    string
	Flash   string
	Warning string
	Error   string
	IP      string
	Form    *forms.Form
}

var functions = template.FuncMap{}
var app *config.AppConfig
var pathToTemplates = "./blades/"

// AddDefaultData adds data for all templates
func AddDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.CSRF = nosurf.Token(r)
	return td
}

// NewRenderer sets the config for the template package
func NewRenderer(a *config.AppConfig) {
	app = a
}

func Renderer(w http.ResponseWriter, r *http.Request, tmpl string, tempData *TemplateData) error {

	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tmpl]
	if !ok {
		//log.Fatal("Could not get template from template cache")
		return errors.New("could not get template from cache")
	}
	buf := new(bytes.Buffer)

	tempData = AddDefaultData(tempData, r)

	t.Execute(buf, tempData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s*.page.html", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s*.layout.html", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s*.layout.html", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}

func Scream(e error) {
	if e != nil {
		panic(e)
	}
}
