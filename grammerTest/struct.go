package main

import (
	"fmt"
)
/*
struct是值传递
嵌套时不允许同层拥有相同的属性，嵌套体和被嵌套体是可以的有相同的
 */
type Books struct {
	title string
	author string
	subject string
	book_id int
}

//嵌套型 使用起来很方便，就是初始化吗，麻烦点
type ChildBooks struct {
	Books
	age int
}

func main() {

	var book1 Books
	var book2 Books

	book1.title = "<<王阳明一切心法>>"
	book1.author = "熊逸"
	book1.subject = "心学的另一番解读"
	book1.book_id = 100

	book2.title = "程序员的数学"
	book2.author = "结成浩"
	book2.subject = "没有晦涩的公式，只有好玩的数学题"
	book2.book_id = 200

	printBooks(book1)
	printBooks(book2)

	printBook(&book1)
	printBook(&book2)

	p := &Books{title:"日诵",author:"老刘",subject:"每天一首诗伴你你整年",book_id:300}

	fmt.Println("Book title :",p.title,"Book author :",p.author)

	printBook(p)

	println("==================================")
	cb1 := ChildBooks{age:3,Books:Books{title:"ddddd",author:"aaa",book_id:400,subject:"bbbbb"}}
	cb1.title = "新世界"
	println("cb1.title = ",cb1.title)//嵌套后属性相当于直接在新的struct中了

}
func printBooks(book Books){
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}

func printBook(book *Books)  {
	fmt.Printf( "Book title : %s\n", book.title);
	fmt.Printf( "Book author : %s\n", book.author);
	fmt.Printf( "Book subject : %s\n", book.subject);
	fmt.Printf( "Book book_id : %d\n", book.book_id);
}