package main

import "fmt"

type Books struct {
	title string
	author string
	subject string
	book_id int
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