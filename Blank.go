package main

import (
	"fmt"
	"strconv"
)

//_表示空白符号

func if_block() {
	a := 10
	if a := 1; a > 0 {
		fmt.Println(a)
		a++
	}
	fmt.Print(a)
}

func for_Loop() {
	a := 1
	for a <= 5 {
		fmt.Println(a)
		a++
	}
	fmt.Println("over")
}

func switch_Block() {

	switch a := 1; {
	case a >= 0:
		fmt.Println("a==0")
		fallthrough
	case a >= 1:
		fmt.Println("a==1")
	case a >= 2:
		fmt.Println("a==2")
	default:
		fmt.Println("a==default")
	}
}

func break_Lable() {
LABEL1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABEL1
			}
		}
	}
	fmt.Println("跳出循环")
}

func main() {
	/*a, _, c, d := 1, 2, 3, 4
	fmt.Println(a)
	fmt.Println(c)
	fmt.Println(d)

	var flot float32 = 0.9
	trans := int(flot)
	fmt.Print(trans)*/
	var a int = 65
	b := string(a)
	c := strconv.Itoa(a) //数字转string
	/*d := strconv.Atoi(c)*/
	fmt.Println(b)
	fmt.Println(c)
	/*fmt.Print(d)*/
	//if_block()
	for_Loop()
	switch_Block()
	break_Lable()
	goto_Block()
}

func goto_Block() {
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABEL1 //标签放在goto下面
			}
		}
	}
LABEL1:
	fmt.Print("跳出循环")
}
