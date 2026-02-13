package main

import "fmt"

type User struct {
	ID       int
	Username string
	Email    string
	IsActive bool
}

func main() {
	user := User{
		ID:       1,
		Username: "thai",
		Email:    "thai@gmail.com",
		IsActive: true,
	}

	fmt.Println(user.Username)
}
