package main

import "fmt"

type User struct {
	Name string
}

func main() {
	users := []User{
		{Name: "Võ"},
		{Name: "Đông"},
		{Name: "Thái"},
	}

	var products []string
	products = append(products, "A")
	products = append(products, "B")
	products = append(products, "C")

	for _, user := range users {
		fmt.Println(user.Name)
	}

	for _, product := range products {
		fmt.Println(product)
	}

	//lấy một phần tử theo index
	fmt.Println(products[0])
	//Lấy một đoạn của slice
	fmt.Println(products[1:3])
}
