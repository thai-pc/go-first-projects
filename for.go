package main

import "fmt"

func main() {
	//for kiểu cơ bản
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//for kiểu while
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//For kiểu vô hạn dùng cho worker, server
	//for {
	//	fmt.Println("Running...")
	//}

	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("Break is", i)
	}

	//Continue
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Println("Continue is", i)
	}

	//

}
