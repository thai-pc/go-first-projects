package main

import "fmt"

func login(username, pasword string) (bool, error) {
	if username == "" || pasword == "" {
		return false, fmt.Errorf("username or password is empty")
	}

	if username == "Thái" && pasword == "123456" {
		return true, nil
	}

	return false, fmt.Errorf("username or password is wrong")
}

func checkAge(age int) string {
	if age >= 18 {
		return "Trưởng thành"
	}
	return "Còn nhỏ"
}

func main() {
	ok, err := login("Thái", "123456")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Login", ok)

	result := checkAge(18)

	fmt.Println("CheckAge", result)
}
