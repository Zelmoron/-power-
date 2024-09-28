package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Users struct {
	Id       int
	Login    string
	Password string
}

func ConnectDb() (*pgxpool.Pool, error) {
	db, err := pgxpool.Connect(context.Background(), "postgres://zelmoron:132313Igor@localhost:5432/sportsite")

	if err != nil {
		log.Println("Error with connection")
		return db, err
	}

	return db, err
}

func InsertUser(login string, password string) error {
	db, err := ConnectDb()
	defer fmt.Println("тест defer")
	defer db.Close()
	if err != nil {
		return err
	}
	result, err := db.Exec(context.Background(), "INSERT INTO users(Login,Password) VALUES($1,$2)", login, password)

	if err != nil {
		log.Println("Ошибка Insert при регистрации пользователя")
		log.Println(result)
		return err
	}

	return nil

}

func (u Users) SelectUser(login string, password string) Users {
	db, err := ConnectDb()

	defer db.Close()
	if err != nil {

		return Users{}
	}

	err = db.QueryRow(context.Background(), "SELECT Id,Login,Password FROM users WHERE Login = $1 and Password = $2", login, password).Scan(&u.Id, &u.Login, &u.Password)
	if err != nil {

		return Users{}
	}

	return u
}

func (u Users) SelecLogin(Id int) Users {
	db, err := ConnectDb()

	defer db.Close()

	if err != nil {
		return Users{}
	}
	err = db.QueryRow(context.Background(), "SELECT Login FROM users WHERE Id = $1", Id).Scan(&u.Login)
	if err != nil {

		return Users{}
	}

	return u
}
