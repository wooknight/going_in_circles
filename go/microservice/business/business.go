package business

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/wooknight/going_in_circles/go/microservice/config"
)

type address struct {
	contact_name string `json:"contactName"`
	street_name  string `json:"streetName"`
	street_name2 string `json:"streetName2"`
	city         string `json:"city"`
	state        string `json:"state"`
	zip          string `json:"zip"`
}
type image struct {
	file_name string `json:"fileName"`
	// fileData  image.image `json:"file"`
}
type business struct {
	business_name string    `json:"streetName"`
	logo_small    image     `json:"logoSmall"`
	logo_large    image     `json:"logoLarge"`
	hq            address   `json:"HQ"`
	branches      []address `json:"branches"`
}

var functions = template.FuncMap{}
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *config.TemplateData) {
	t, ok := app.TemplateCache[tmpl]
	if !ok {
		log.Fatal("Could not find template", tmpl)
	}
	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)
	if _, err := buf.WriteTo(w); err != nil {
		log.Fatal("Error writinging to w")
	}
}

func RenderTemplateTotal() (map[string]*template.Template, error) {
	myTempls := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myTempls, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myTempls, err
		}
		ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
		if err != nil {
			return myTempls, err
		}
		myTempls[name] = ts
	}
	return myTempls, nil
}
