package main

import (
	"fmt"
)

//定义一个结构体,也就是java的类,但是没有多态
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	 Book1 := &Books
	 Book2 := &Books

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	// book 2 描述
	Book2.title = "Python 教程"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	bookInfo(&Book1)
	bookInfo(&Book2)

//struct是值拷贝，所以如果需要使用引用拷贝，就需要用到指针了
//在go中，可以使用(指针.属性）来取得一个类的字段
func bookInfo(book *Books) {
	fmt.Printf("Book  title : %s\n", book.title)
	fmt.Printf("Book  author : %s\n", book.author)
	fmt.Printf("Book  subject : %s\n", book.subject)
	fmt.Printf("Book  book_id : %d\n", book.book_id)
}

/*func makeSlice() {
	var numbers = make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(numbers), cap(numbers), numbers)
}
*/