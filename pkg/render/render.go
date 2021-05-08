package render

import (
	"bytes"
	"fmt"
	"github.com/cinguilherme/playground/pkg/config"
	"github.com/cinguilherme/playground/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{

}

var app *config.AppConfig
// NewTemplates sets the config for the template pkg
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UserCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}


	t, ok := tc[tmpl]
	if !ok {
		log.Fatal()
	}

	buff := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buff, td)

	_, err := buff.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template response writer")
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, _ := filepath.Glob("./templates/*.page.tmpl")

	for _, page := range pages {
		name := filepath.Base(page)
		ts, _ := template.New(name).Funcs(functions).ParseFiles(page)

		matches, _ := filepath.Glob("./templates/*.layout.tmpl")


		if len(matches) > 0 {
			ts, _ = ts.ParseGlob("./templates/*.layout.tmpl")
		}

		myCache[name] = ts
	}
	return myCache, nil
}
