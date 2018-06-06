package main

import (
	"fmt"
)

//在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 值。
func main() {
	nums := []int{2, 3, 4}
	sum := 0
	//Java里面增强的for循环
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range也可以用在map的键值对上。返回key和value
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Printf("%s -> %s\n", key, value)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
