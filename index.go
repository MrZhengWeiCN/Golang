package main

import (
	"fmt"
	"unsafe"
)

type x, y int

type (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

/*
表示调用
在Go里面，函数定义的变量一定要用到。不然会报错
*/
func main() {
	pointer2pointer()
}

func hello() {
	fmt.Println("a.go is package mufunc.")
	var a = 1.5
	var b int = 2
	fmt.Println(a, b)
}

func test() {
	var avaliabled bool
	valid := false
	avaliabled = true
	fmt.Println(avaliabled, valid)
}

//语言变量
/*func value() {
	//要么使用var valueName valueType；要么使用valueName：= value
	i := "No Used Value" //如果不用会报错
	g, h := 123, "hello" //只能在函数体中出现
	fmt.Println(&g)      //地址指针的值
	fmt.Println(a, b, c, d, e, f, g, h, i)
}*/

//go常亮的用法
func cons_Value() {
	const (
		LEIGHT int = 10
		WEIGHT int = 5
		a          = 1
		b          = false
		c          = "str"
		d          = unsafe.Sizeof(c)
	)
	//常量表达式中，函数必须是内置函数，否则编译不过：
	fmt.Println(d)
	var area int
	area = LEIGHT * WEIGHT
	fmt.Printf("面积为：%d", area)
	println()
	println(a, b, c)
}

//其他运算符
func oprat_symbol() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("ptr 为 %d\n", ptr)
}

//循环,求素数
func prime_number() {
	var C, c int
	C = 1
A:
	for C < 100 {
		C++
		for c = 2; c < C; c++ {
			if C%c == 0 {
				goto A
			}
		}
		fmt.Println(C, "是素数")
	}

	for i := 1; i <= 9; i++ { // i 控制行，以及计算的最大值
		for j := i; j <= 9; j++ { // j 控制每行的计算个数
			fmt.Printf("%d*%d=%d   ", i, j, j*i)
		}
		fmt.Println("")
	}
}

func max(num1, num2 int) int {
	if num1 >= num2 {
		return num1
	} else {
		return num2
	}
}

//返回多个值
func swap(x, y string) (string, string) {
	return y, x
}

//引用传递
func swap_ref(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

//关于数组的
func array() {
	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
}

//数组的平均值

func getAverage(arr []int, size int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	//强制转型
	avg = float32(sum) / float32(size)

	return avg
}

//go里面的指针
func pointer() {
	var a int = 20
	var ip *int
	ip = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)
	fmt.Printf("*ip 变量的值: %d\n", *ip)
	var ptr *int
	if ptr != nil {
		fmt.Printf("ptr不是空指针 ")
	} else {
		fmt.Printf("ptr 的值为 : %x\n", ptr)
	}
}

//指针数组
func pointer_Array() {
	a := []int{1, 2, 3}
	const length int = 3
	var ptr [length]*int
	for i := 0; i < length; i++ {
		ptr[i] = &a[i] //ptr[i]每一个值都是地址
	}
	for i := 0; i < length; i++ {
		fmt.Printf("a[%d] = %d\n", i, *ptr[i]) //*ptr[i])这个地址指向的值
	}
}

//指针的指针
func pointer2pointer() {
	var a int = 3
	var ptr *int
	var pptr **int

	ptr = &a
	pptr = &ptr
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
}
