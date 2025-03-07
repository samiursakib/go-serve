package services

import "fmt"

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserService struct {
	users []User
}

func (userService *UserService) InitUserService() {
	userService.users = []User{
		{"skdjde", "Alice Johnson", "alicej", "alice@example.com", "password123"},
		{"eirwdc", "Bob Smith", "bobsmith", "bob@example.com", "securePass!"},
		{"znchwr", "Charlie Brown", "charlieb", "charlie@example.com", "charliePass"},
	}
	fmt.Println(userService.users)
}

func (userService *UserService) getUsers() []User {
	fmt.Println(userService.users)
	return userService.users
}

func findUserById(id string, users *[]User) *User {
	for i := range *users {
		if (*users)[i].Id == id {
			return &(*users)[i]
		}
	}
	return nil
}
