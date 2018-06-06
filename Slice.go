package main

import "fmt"

/*
Array和slice的关系就是Array和ArrayList的关系
*/
func main() {

	//var s1 []int //slice不需要指定数组的大小，和数组的区别
	//s1 := make([]int, 3, 6) //比较正式的写法
	/*fmt.Printf("%p\n", s1)
	fmt.Println(len(s1), cap(s1))
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 4, 5, 6)
	fmt.Printf("%v %p\n", s1, s1)*/

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	s3 := s1[0:6]
	copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s3)

	/*numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers)
	fmt.Println(numbers)
	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	/*fmt.Println("numbers[1:4] ==", numbers[1:4])
	/* 默认下限为 0*/
	/*fmt.Println("numbers[:3] ==", numbers[:3])
	fmt.Println("numbers[3:] ==", numbers[3:])

	numbers1 := make([]int, 0, 5)
	printSlice(numbers1)

	numbers2 := numbers[:2]
	printSlice(numbers2)

	numbers3 := numbers[2:5]
	printSlice(numbers3)*/

	var numbers []int
	printSlice(numbers)
	numbers = append(numbers, 0)
	printSlice(numbers)
	numbers = append(numbers, 1)
	printSlice(numbers)
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	numbers_one := make([]int, len(numbers), cap(numbers)*2)
	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers_one, numbers)
	printSlice(numbers_one)

}

func printSlice(x []int) {
	//cap类似于hashMap的cap，会自动扩容，而且cap = cap<<1
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
