package main

import "fmt"

type User struct {
	Name string
}

func (u User) sayHello() {
	fmt.Println("Hello", u.Name)
}

// Thay đổi giá trị trong go phải dùng biến con trỏ (pointer). Vì go truyền bản copy
func (u *User) changeName(newName string) {
	u.Name = newName
}

func main() {
	user := User{
		Name: "Thái",
	}
	user.changeName("Trung")

	user.sayHello()
}
