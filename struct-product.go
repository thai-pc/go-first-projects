package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	pro := Product{
		Name:  "Iphone",
		Price: 2000,
	}

	fmt.Println(pro.Name, "-", pro.Price)
}
