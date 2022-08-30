package main

import "fmt"

type User struct {
	Name string
}

type userSvc struct {
	data []*User
}

func (uS *userSvc) RegisterUser(u *User) string {
	uS.data = append(uS.data, u)
	return u.Name + " berhasil didaftarkan"
}

func (uS *userSvc) GetUser() []*User {

	return uS.data
}

func NewUserService(users []*User) UserInterface {
	return &userSvc{
		data: users,
	}
}

type UserInterface interface {
	RegisterUser(u *User) string
	GetUser() []*User
}

func main() {

	var data []*User
	userSvc := NewUserService(data)
	names := []string{"giva", "ucup", "edi"}

	for _, name := range names {
		resp := userSvc.RegisterUser(&User{Name: name})
		fmt.Println(resp)

	}

	GetUser := userSvc.GetUser()
	fmt.Println("-----------Hasil get user-------------")
	for _, value := range GetUser {
		fmt.Println(value.Name)
	}

}
