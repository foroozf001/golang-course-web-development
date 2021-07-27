package render

import (
	"log"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("../../templates/" + tmpl)
	if err != nil {
		log.Println("error render.go:", err)
		return
	}
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		log.Println("error render.go:", err)
		return
	}
	log.Println("HTTP 200 OK")
}
