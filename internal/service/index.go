package service

import (
	"app/internal/database"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("sub").(string)
	var u database.Users
	id, err := strconv.Atoi(user)

	Login := u.SelecLogin(id)

	tmpl, err := template.ParseFiles("templates/html/index.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "index", Login)
}
