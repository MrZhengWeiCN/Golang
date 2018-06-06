package main

import "fmt"

func main() {
	//var a [2]int
	b := [...]int{19: 1}
	var p *[20]int = &b
	fmt.Println(p[0])
	sort()
}

func sort() {
	a := [...]int{5, 2, 6, 3, 9}
	fmt.Println(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i; j+1 < num; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}
	fmt.Println(a)
	for i := 0; i < num; i++ {
		for j := num - 1; j-1 >= 0; j-- {
			if a[j] > a[j-1] {
				temp := a[j]
				a[j] = a[j-1]
				a[j-1] = temp
			}
		}
	}
	fmt.Println(a)
}
