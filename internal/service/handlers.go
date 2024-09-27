package service

import (
	"app/internal/utils"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString, err := utils.ReadCookie("jwt-token", r)

		if err != nil {
			http.Redirect(w, r, "http://localhost:8080/signIn", http.StatusSeeOther)
			return
		}
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})
		if err != nil || !token.Valid {
			log.Printf("Error parsing token: %v", err)
			http.Redirect(w, r, "http://localhost:8080/signIn", http.StatusSeeOther)

		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("Error getting claims from token")
			http.Redirect(w, r, "http://localhost:8080/signIn", http.StatusSeeOther)

		}

		roles, ok := claims["sub"].(string)
		if !ok {
			log.Printf("Roles not found in token")
			http.Redirect(w, r, "http://localhost:8080/signIn", http.StatusSeeOther)

		}

		ctx := context.WithValue(r.Context(), "sub", roles)
		r = r.WithContext(ctx)

		// Передача запроса на следующий обработчик
		next.ServeHTTP(w, r)

	})
}

func Handlers() {
	rtr := mux.NewRouter()

	//подключил свои стили и js
	http.Handle("/templates/",
		http.StripPrefix("/templates", http.FileServer(http.Dir("./templates/"))))

	rtr.HandleFunc("/signIn", SignIn).Methods("GET") // обработчик страницы авторизации
	// rtr.HandleFunc("/signUp", service.SignUp).Methods("GET")        // обработчик страницы регистрации
	http.Handle("/index", jwtMiddleware(http.HandlerFunc(handleIndex))) // обработчик главной страницы
	// Приватный подмаршрутизатор
	rtr.HandleFunc("/checkform", CheckForm).Methods("POST") // обработчик для проверки пароля и логина страницы

	http.Handle("/", rtr) // все обраоботчики через роутер

	// privateRouter.HandleFunc("/index", handleIndex)

}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("sub").(string)

	fmt.Println(userId)
	tmpl, err := template.ParseFiles("templates/html/index.html")

	if err != nil {
		log.Println("Ошибка обработки html")
		return
	}
	tmpl.ExecuteTemplate(w, "index", nil)

}
