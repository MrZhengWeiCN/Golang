package main

import (
	"fmt"
)

type Human struct {
	Sex string
}

type person struct {
	Human //使用组合来代替继承
	Name  string
	Age   int
	//匿名结构
	Contact struct {
		Phone, City string
	}
}

func main() {
	a := person{Name: "joe", Age: 19, Human: Human{Sex: "male"}}
	a.Contact.City = "beijing"
	a.Contact.Phone = "123345"
	//a.Sex = "male"
	a.Human.Sex = "female"
	fmt.Println(a)
}
