package main

//struct ngang class
import "fmt"

type User struct {
	Name string
	Age  int
}

// Khai báo struct không giá trị

func main() {
	//user := User{
	//	Name: "Thái",
	//	Age:  25,
	//}

	// Khai báo struct không giá trị Go có zero value
	var user User
	fmt.Println(user.Name) //""
	fmt.Println(user.Age)  //0

	fmt.Println(user)
	//Truy cập dữ liệu trong struct
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
