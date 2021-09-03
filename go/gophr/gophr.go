package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

var errorTemplate = `<html>
<body>
<h1>Error Rendering template %s</h1>
<p>%s</p>
</body>
</html>
`

var layout = template.Must(template.New("layout.html").Funcs(layoutFuncs).ParseFiles("templates/layout.html"))

var layoutFuncs = template.FuncMap{"yield": func() (string, error) {
	return "", fmt.Errorf("Called inappropriately ")
}}

func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	funcs := template.FuncMap{"yield": func() (template.HTML, error) {
		buf := bytes.NewBuffer(nil)
		err := templates.ExecuteTemplate(buf, name, data)
		return template.HTML(buf.String()), err
	}}
	layoutClone, _ := layout.Clone()
	layoutClone.Funcs(funcs)
	err := layoutClone.Execute(w, data)
	if err != nil {
		http.Error(w, fmt.Sprintf(errorTemplate, name, err), http.StatusInternalServerError)
	}
}

func main() {
	// http.ListenAndServe(":3000", http.FileServer(http.Dir("assets/")))
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RenderTemplate(w, r, "index/home", nil)
	})
	mux.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets/"))))
	log.Fatal(http.ListenAndServe(":3000", mux))
}
