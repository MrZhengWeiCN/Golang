package main

import (
	"fmt"
)

func main() {

	var sum int = 17
	var count int = 5
	var mean float32

	mean = float32(sum) / float32(count) //数据类型强制转换
	fmt.Printf("mean 的值为: %f\n", mean)

	for i := 1; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}

}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
