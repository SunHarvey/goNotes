package main

import "fmt"

type Book struct {
    title string
    author string
}

func changeBook(book *Book) {
    book.title = "abc"
}

func main() {
    var book1 Book
    book1.author = "zhangsan"
    book1.title = "xiyouji"
    fmt.Printf("%v\n", book1)
    
    fmt.Println("----------------------")
    changeBook(&book1)
    fmt.Printf("%v\n", book1)

}

