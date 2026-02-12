package main

import "fmt"

func main() {
	username := "thai"
	password := "123456"

	if username == "thai" && password == "123456" {
		fmt.Println("Login success")
	} else {
		fmt.Println("Login failed")
	}

	if age := 15; age > 15 {
		fmt.Println("Trưởng thành")
	}
	fmt.Println("Trẻ em")
}
