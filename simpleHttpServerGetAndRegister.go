package main

import (
	"encoding/json"
	"net/http"
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
	http.HandleFunc("/register", userSvc.RegisterHandler)
	http.HandleFunc("/user", userSvc.GetUserHandler)
	http.ListenAndServe("localhost:8080", nil)

}
