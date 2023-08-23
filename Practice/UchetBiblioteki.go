package main

import "fmt"

type book struct {
	Title  string
	Author string
	Genre  string
}
type reader struct {
	Name          string
	Age           int
	BorrowedBooks []book
}
type libraryItem interface {
	displayInfo() string
}

func (b book) displayInfo() string {
	return fmt.Sprintf("Название: %s, Автор: %s, Жанр: %s", b.Title, b.Author, b.Genre)
}

func (r reader) displayInfo() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d", r.Name, r.Age)
}

func (r *reader) borrowBook(b book) {
	r.BorrowedBooks = append(r.BorrowedBooks, b)
}

func main() {
	book1 := book{"Мастер и фломастер", "Мишка косолапый", "Фантастика"}
	book2 := book{"Мастер и карандаш", "Жираф длинношеей", "Классика"}

	reader1 := reader{Name: "Вася", Age: 22}
	reader2 := reader{Name: "Катя", Age: 36}

	reader1.borrowBook(book1)
	reader2.borrowBook(book2)

	items := []libraryItem{book1, book2, reader1, reader2}
	for _, item := range items {
		fmt.Println(item.displayInfo())
	}
}
