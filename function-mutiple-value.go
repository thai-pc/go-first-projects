package main

import "fmt"

func getUser() (string, int) {
	return "Thái", 25
}

func main() {
	//Bỏ qua giá trị cần dùng _
	//name, _ := getUser()
	//fmt.Println(name)
	name, age := getUser()
	fmt.Println(name, age)
}
