package main

import (
	"fmt"
)

func main() {
	//匿名函数
	/*a := func() {
		fmt.Println("No name function")
	}
	a()
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
	fmt.Println(f(3))*/
	A()
	B()
	C()
}

//不定长参数,只能是最后一个参数，不定长变参是值传递
func a(x ...int) {
	fmt.Println(x)
}

func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

//defer 类似于Java的finally
//类似于栈式的执行顺序
func A() {
	fmt.Println("Func A")
}

func B() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("Panic in B")
}

func C() {
	fmt.Print("Func C")
}
