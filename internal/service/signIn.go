package service

import (
	"html/template"
	"log"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/html/signIn.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "signIn", nil)
}
