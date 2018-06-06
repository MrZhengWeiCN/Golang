package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

//方法和函数是不同的
func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.Print()
	fmt.Println(b.Name)

	var work Work
	work.Increase(100)
	fmt.Println(work)
}

type Work int

func (work *Work) Increase(num int) {
	*work += Work(num)
}

//reciver，表示A的method
//类似于Java
//class A{
//	Print（）
//}
func (a *A) Print() {
	a.Name = "AAA"
	fmt.Println("A")
}

//表示A的method
func (b B) Print() {
	b.Name = "BB"
	fmt.Println("B")
}
