package main

import (
	"fmt"
)

// Определение структуры книги
type Book struct {
	Title  string
	Author string
	Genre  string
}

// Определение структуры читателя
type Reader struct {
	Name          string
	Age           int
	BorrowedBooks []Book
}

// Определение интерфейса LibraryItem
type LibraryItem interface {
	DisplayInfo() string
}

// Метод DisplayInfo для структуры Book
func (b Book) DisplayInfo() string {
	return fmt.Sprintf("Название: %s, Автор: %s, Жанр: %s", b.Title, b.Author, b.Genre)
}

// Метод DisplayInfo для структуры Reader
func (r Reader) DisplayInfo() string {
	return fmt.Sprintf("Имя: %s, Возраст: %d", r.Name, r.Age)
}

// Метод BorrowBook для структуры Reader
func (r *Reader) BorrowBook(book Book) {
	r.BorrowedBooks = append(r.BorrowedBooks, book)
}

func main() {
	// Создание книг
	book1 := Book{Title: "Мастер и Маргарита", Author: "Михаил Булгаков", Genre: "Роман"}
	book2 := Book{Title: "1984", Author: "Джордж Оруэлл", Genre: "Научная фантастика"}

	// Создание читателей
	reader1 := Reader{Name: "Иван", Age: 25}
	reader2 := Reader{Name: "Елена", Age: 30}

	// Запись книг в список взятых читателем книг
	reader1.BorrowBook(book1)
	reader2.BorrowBook(book2)

	// Вывод информации о книгах и читателях
	items := []LibraryItem{book1, book2, reader1, reader2}
	for _, item := range items {
		fmt.Println(item.DisplayInfo())
	}
}
