package main

import "fmt"

// Khai báo mảng trên ít dùng do phải biết trước số phần tử và không thêm phần tử vào nữa
//var numbers [3]int = [3]int{1, 2, 3}

//Thường dùng
//numbers := [3]int{1, 2, 3}

//Slice thêm được phần tử và không cần biết trước có bao nhiêu phần tử

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4)
	numbers = append(numbers, 5)

	list := []int{10, 20, 30, 40, 50}

	//Loop slice
	for i, v := range numbers {
		//i là index, v là value
		fmt.Println(i, v)
	}

	for _, v := range list {
		fmt.Println(v)
	}

}
