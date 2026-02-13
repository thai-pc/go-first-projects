package main

import "fmt"

type User struct {
	Name string
}

func sayHello(u User) {
	fmt.Println("Hello", u.Name)
}
func main() {
	user := User{
		Name: "Thái",
	}
	sayHello(user)
}
