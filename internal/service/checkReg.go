package service

import (
	"app/internal/database"
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
)

func CheckReg(w http.ResponseWriter, r *http.Request) {
	var u database.Users
	login := r.FormValue("username")
	passwordUser := r.FormValue("password")

	hashpassword := md5.Sum([]byte(passwordUser))
	hash := hex.EncodeToString(hashpassword[:])

	result := u.SelectUser(login, hash)
	if result.Id == 0 {
		err := database.InsertUser(login, hash)

		if err == nil {
			http.Redirect(w, r, "http://localhost:8080/signIn", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "http://localhost:8080/signUp", http.StatusSeeOther)
		}
	} else {

		tmpl, err := template.ParseFiles("templates/html/signUp.html")

		if err != nil {
			log.Println("Ошибка обработки html")
			return
		}
		tmpl.ExecuteTemplate(w, "SignUp", "Такой пользователь уже есть!")
	}

}
