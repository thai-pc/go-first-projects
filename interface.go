package main

import "fmt"

// Define interface
type Speaker interface {
	Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() {
	fmt.Println("Gâu gâu")
}

func (c Cat) Speak() {
	fmt.Println("Meo meo")
}

func makeSound(s Speaker) {
	s.Speak()
}

func main() {
	dog := Dog{}
	cat := Cat{}

	makeSound(dog)
	makeSound(cat)
}
