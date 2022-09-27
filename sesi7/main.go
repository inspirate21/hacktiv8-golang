package main

import (
	"fmt"
	"math/rand"
	"sesi7/config"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := config.ConnectPostgres()
	if err != nil {
		panic(err)
	}

	db := config.GetPostgres()
	if db == nil {
		panic("db not connected")
	}

	fmt.Println("db connected !")
	defer db.Close()

	CreateUser(2)
}

type User struct {
	ID        int
	Name      string
	Email     string
	Contact   string
	CreatedAt time.Time
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func CreateUser(len int) {
	db := config.GetPostgres()
	var users []User
	for i := 0; i < len; i++ {
		users = append(users, User{
			Name:      RandStringRunes(10),
			Email:     fmt.Sprintf("%s@gmail.com", RandStringRunes(20)),
			Contact:   RandStringRunes(10),
			CreatedAt: time.Now(),
		})
	}

	query := `
		INSERT INTO users (
			name, email, contact, created_at
		) VALUES ($1, $2, $3, $4)
	`

	stmt, err := db.Prepare(query)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	for _, u := range users {
		_, err := stmt.Exec(u.Name, u.Email, u.Contact, u.CreatedAt)
		if err != nil {
			fmt.Println("error :", err.Error())
		}

	}

	fmt.Println("created", len, "users")
}
