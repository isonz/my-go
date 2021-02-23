/**
结构体 - 可以为不同项定义不同的数据类型
*/
package main

import "fmt"

func main()  {
	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "Google", "Go 语言教程", 6495407})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Swift 语言", author: "Apple", subject: "Swift 语言教程", bookId: 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Java 语言", author: "Oracle"})


	var Book2 Books        /* 声明 Book2 为 Books 类型 */
	/* book 2 描述 */
	Book2.title = "Python 教程"
	Book2.author = "Guido van Rossum"
	Book2.subject = "Python 语言教程"
	Book2.bookId = 6495700
	/* 打印 Book2 信息 */
	printBook(Book2)
	/* 打印 Book2 信息 */
	printBook2(&Book2)

}


// 结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。
type Books struct {
	title string
	author string
	subject string
	bookId int
}

// 结构体作为函数参数
func printBook( book Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.bookId)
}

func printBook2( book *Books ) {
	fmt.Printf( "Book title : %s\n", book.title)
	fmt.Printf( "Book author : %s\n", book.author)
	fmt.Printf( "Book subject : %s\n", book.subject)
	fmt.Printf( "Book book_id : %d\n", book.bookId)
}