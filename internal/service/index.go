package service

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/html/index.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}
