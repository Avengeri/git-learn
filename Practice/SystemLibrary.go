package main

import "fmt"

type libraryItem interface {
	Display()
}

func (b book) Display() {
	fmt.Printf("Книга: %s", b.Title)
	fmt.Printf("Автор: %s", b.Author)
	fmt.Printf("Год выпуска: %s", b.Year)
}

type book struct {
	Title  string
	Author string
	Year   int
}

func main() {

}
