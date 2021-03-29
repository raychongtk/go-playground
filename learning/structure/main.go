package main

import "fmt"

type User struct {
	userId   int
	username string
	password string
	email    string
}

func (user User) login() bool {
	return user.username == "ray" && user.password == "123456"
}

func main() {
	user := User{userId: 1, username: "ray", password: "123456", email: "ray@gmail.com"}

	fmt.Println("id =", user.userId)
	fmt.Println("username =", user.username)
	fmt.Println("email =", user.email)
	fmt.Println(user.login())
}
