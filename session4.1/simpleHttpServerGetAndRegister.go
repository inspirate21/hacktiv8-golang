package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
}

type userSvc struct {
	data []*User
}

func (uS *userSvc) RegisterUser(u *User) string {
	uS.data = append(uS.data, u)
	return u.Name + " berhasil didaftarkan"
}

func (uS *userSvc) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	if r.Method == http.MethodPost {
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		uS.RegisterUser(&user)

		w.Write([]byte(user.Name + " berhasil didaftarkan"))

	} else {
		w.Write([]byte("can't use " + r.Method + " http method"))
	}

}

func (uS *userSvc) GetUser() []*User {

	return uS.data
}

func (uS *userSvc) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	users := uS.GetUser()
	data, _ := json.Marshal(users)
	w.Write(data)
}

func NewUserService(users []*User) UserInterface {
	return &userSvc{
		data: users,
	}
}

type UserInterface interface {
	RegisterUser(u *User) string
	GetUser() []*User
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	GetUserHandler(w http.ResponseWriter, r *http.Request)
}

func main() {
	var data []*User
	userSvc := NewUserService(data)

	// http.HandleFunc("/register", userSvc.RegisterHandler)
	// http.HandleFunc("/user", userSvc.GetUserHandler)
	// http.ListenAndServe("localhost:8080", nil)

	r := mux.NewRouter()
	r.HandleFunc("/register", userSvc.RegisterHandler)
	r.HandleFunc("/user", userSvc.GetUserHandler)
	r.HandleFunc("/user/{id}", userSvc.GetUserHandler)
	srv := &http.Server{
		Handler: r,
		Addr:    "localhost:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	_ = srv.ListenAndServe()

}
