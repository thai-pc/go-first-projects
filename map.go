package main

import "fmt"

func main() {
	//string đầu là kiêu của key, string thứ 2 là kiểu của value
	user := map[string]string{
		"name": "Thai",
		"city": "HCM",
	}

	//map kiểu int

	scores := map[string]int{
		"thai": 10,
		"an":   5,
	}

	//Thêm/sửa dữ liệu
	user["email"] = "thai@gmail.com"
	user["name"] = "Trung"

	//Lấy giá trị từ map
	fmt.Println(user["name"])

	//Kiểm tra key có tồn tại không
	age, ok := scores["thai"]

	if ok {
		fmt.Println("Tuổi là", age)
	} else {
		fmt.Println("Không tìm thấy")
	}

	//Loop qua map
	for key, value := range scores {
		fmt.Println(key, value)
	}

	loginData := map[string]string{
		"username": "thai",
		"password": "123",
	}

	//Xóa phần tử ra khỏi map
	delete(loginData, "password")

	fmt.Println(loginData["username"])

	productPrices := map[string]int{
		"iphone":  2000,
		"samsung": 1500,
	}

	for key, value := range productPrices {
		fmt.Println(key, value)
	}

}
