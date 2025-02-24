package main

import "fmt"

type Book struct {
	Name        string
	Author      string
	Year        int
	IsAvailable bool
}
type Library map[string]Book

func (lib Library) AddBook(name string, author string, year int) {
	book := Book{name, author, year, true}
	lib[book.Name] = book
}
func (lib Library) Issue(name string) {
	if book, exists := SearchByName(lib, name); exists {
		book.IsAvailable = false
		lib[name] = book
	}
}
func (lib Library) Return(name string) {
	if book, exists := SearchByName(lib, name); exists {
		book.IsAvailable = true
		lib[name] = book
	}
}
func SearchByName(lib Library, name string) (Book, bool) {
	book, exists := lib[name]
	return book, exists
}
func (lib Library) PrintState() {
	fmt.Printf("Books count: %v\n", len(lib))
	for _, b := range lib {
		fmt.Printf("%q, %v, %v, %v\n", b.Name, b.Author, b.Year, b.IsAvailable)
	}
	fmt.Println()
}

func main() {
	lib := make(Library)
	lib.PrintState()
	lib.AddBook("1984", "Orwell", 1949)
	lib.PrintState()
	lib.AddBook("Brave New World", "Huxley", 1931)
	lib.PrintState()
	book, _ := SearchByName(lib, "1984")
	lib.Issue(book.Name)
	lib.PrintState()
	lib.Return(book.Name)
	lib.PrintState()
}
