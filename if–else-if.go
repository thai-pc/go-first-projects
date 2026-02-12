package main

import "fmt"

func main() {
	score := 8

	if score >= 9 {
		fmt.Println("Xuất sắc")
	} else if score >= 7 {
		fmt.Println("Khá")
	} else {
		fmt.Println("Trung bình")
	}

	//Cho phép tạo biến ngay trong if
	if age := 20; age >= 18 {
		fmt.Println("Adult")
	}
}
