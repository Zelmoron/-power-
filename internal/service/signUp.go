package service

import (
	"html/template"
	"log"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("templates/html/signUp.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "SignUp", nil)
}
