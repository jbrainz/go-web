package render

import (
	"fmt"
	"html/template"
	"net/http"
)

// RenderTemplate renders and html template on the frontend
func RenderTemplate(w http.ResponseWriter, html string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + html)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template", err)
		return
	}
}
