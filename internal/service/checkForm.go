package service

import (
	"app/internal/database"
	"crypto/md5"
	"encoding/hex"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

func CheckForm(w http.ResponseWriter, r *http.Request) {
	var u database.Users
	login := r.FormValue("username")
	passwordUser := r.FormValue("password")

	hashpassword := md5.Sum([]byte(passwordUser))
	hash := hex.EncodeToString(hashpassword[:])

	result := u.SelectUser(login, hash)

	if result.Id == 0 {
		tmpl, err := template.ParseFiles("templates/html/signIn.html")

		if err != nil {
			log.Println("Ошибка обработки html")
			return
		}
		tmpl.ExecuteTemplate(w, "signIn", "Не правильный логин или")
		return
	} else {
		secretKey := []byte("your-secret-key")

		// Создайте новые Claims с информацией, которую хотите добавить
		payload := jwt.MapClaims{
			"sub": strconv.Itoa(result.Id),
			"exp": time.Now().Add(time.Minute * 50).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		tokenString, _ := token.SignedString(secretKey)
		livingTime := time.Minute * 60
		expiration := time.Now().Add(livingTime)

		cookie := http.Cookie{Name: "jwt-token", Value: url.QueryEscape(tokenString), Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "http://localhost:8080/index", http.StatusSeeOther)
	}

	// }

}
