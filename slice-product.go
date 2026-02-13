package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	pro := []Product{
		{"A", 1000},
		{"B", 2000},
		{"C", 3000},
	}

	for _, p := range pro {
		fmt.Println(p.Name, p.Price)
	}

	names := []string{"Thai", "Trung", "An"}

	for _, name := range names {
		fmt.Println("Hello", name)
	}
}
