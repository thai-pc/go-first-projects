package main

import "fmt"

type Counter struct {
	Value int
}

// truy cập giá trị tại địa chỉ đó *
func increase(counter *Counter) {
	counter.Value++
}

func main() {
	value := Counter{
		Value: 1,
	}
	//lấy địa chỉ của biến &
	increase(&value)

	fmt.Println(value.Value)
}
