package mytemplate

import (
	"html/template"
	"net/http"

	"log"
)

func PrintDB(w http.ResponseWriter, str string) {
	tmpl, err := template.New("spark").Parse("<h1>Hello {{.}}</h1>\n")
	if err != nil {
		log.Printf("Hit a problem %v\n", err)
	}
	err = tmpl.Execute(w, "World"+str)
	if err != nil {
		log.Printf("Encountered an error while printing : %v", err)
	}
}
