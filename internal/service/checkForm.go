package service

import (
	"net/http"
	"net/url"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CheckForm(w http.ResponseWriter, r *http.Request) {

	name := r.FormValue("username")

	if name == "admin" {
		secretKey := []byte("your-secret-key")

		// Создайте новые Claims с информацией, которую хотите добавить
		payload := jwt.MapClaims{
			"sub": name,
			"exp": time.Now().Add(time.Second * 5).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
		tokenString, _ := token.SignedString(secretKey)
		livingTime := time.Minute * 60
		expiration := time.Now().Add(livingTime)

		cookie := http.Cookie{Name: "jwt-token", Value: url.QueryEscape(tokenString), Expires: expiration}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "http://localhost:8080/index", http.StatusSeeOther)
	}

}
