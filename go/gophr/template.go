package main

// var templates = template.Must(template.New("t").ParseGlob("templates/**/*.html"))

// var errorTemplate = `
// <html>
// <body>
// <h1>Error Rendering template %s</h1>
// <p>%s</p>
// </body>

// </html>
// `

// func RenderTemplate(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
// 	err := templates.ExecuteTemplate(w, name, data)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf(errorTemplate, name, err), http.StatusInternalServerError)
// 	}
// }
