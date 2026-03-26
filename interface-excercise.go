package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(r Shape) {
	fmt.Println(r.Area())
}

func main() {
	rect := Rectangle{
		5, 10,
	}

	printArea(rect)
}
