package main

import "fmt"

func main() {
	//Gán a bằng 10
	//Gán giá trị của a cho b
	//Sửa b thì giá trị của a không thay đổi
	a := 10
	b := a
	//b = 20

	fmt.Println(a, b)
	//Thay vì copy giá trị ta trỏ thẳng đến vùng nhớ của a
	//& Lấy địa chỉ của a
	//*Lấy giá trị từ địa chỉ
	c := 30
	d := &c
	fmt.Println(*d)

	//Sửa giá trị qua con trỏ
	e := 40
	f := &e
	*f = 70

	fmt.Println("f là ", *f, "e là", e)

}
