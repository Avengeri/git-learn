package main

import (
	"fmt"
)

// Определяем интерфейс LibraryItem
type LibraryItem interface {
	Display()
}

// Создаем структуру Book
type Book struct {
	Title  string
	Author string
	Year   int
}

// Метод Display для Book
func (b Book) Display() {
	fmt.Printf("Книга: %s\n", b.Title)
	fmt.Printf("Автор: %s\n", b.Author)
	fmt.Printf("Год выпуска: %d\n", b.Year)
}

// Создаем структуры для разных типов книг
type FictionBook struct {
	Book
	Genre string
}

func (fb FictionBook) Display() {
	fb.Book.Display()
	fmt.Printf("Жанр: %s\n", fb.Genre)
}

type ScientificBook struct {
	Book
	ISBN string
}

func (sb ScientificBook) Display() {
	sb.Book.Display()
	fmt.Printf("ISBN: %s\n", sb.ISBN)
}

type ChildrenBook struct {
	Book
	AgeRestriction int
}

func (cb ChildrenBook) Display() {
	cb.Book.Display()
	fmt.Printf("Возрастное ограничение: %d+\n", cb.AgeRestriction)
}

// Создаем структуру Library
type Library struct {
	Books []LibraryItem
}

func (lib *Library) AddBook(book LibraryItem) {
	lib.Books = append(lib.Books, book)
}

func (lib *Library) RemoveBook(index int) {
	if index >= 0 && index < len(lib.Books) {
		lib.Books = append(lib.Books[:index], lib.Books[index+1:]...)
	}
}

func (lib Library) DisplayBooks() {
	fmt.Println("Список книг в библиотеке:")
	for _, item := range lib.Books {
		item.Display()
		fmt.Println("--------------")
	}
}

func main() {
	library := Library{}

	// Создаем книги разных типов и добавляем их в библиотеку
	fictionBook := FictionBook{Book{"Война и мир", "Л. Толстой", 1869}, "Роман"}
	scientificBook := ScientificBook{Book{"Go в действии", "W. Kennedy, B. Ketelsen", 2015}, "9781617291784"}
	childrenBook := ChildrenBook{Book{"Приключения Тома Сойера", "М. Твен", 1876}, 12}

	library.AddBook(fictionBook)
	library.AddBook(scientificBook)
	library.AddBook(childrenBook)

	// Выводим список книг в библиотеке
	library.DisplayBooks()
}
