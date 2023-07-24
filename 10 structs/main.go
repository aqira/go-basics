package main

import "fmt"

// Menggunakan struct, kita dapat gunakan jika suatu saat valuenya tidak cukup jika hanya 1 tipe data saja.
type BookInfo struct {
	Name      string
	Genre     string
	Published int
}

func main() {
	books := map[string]BookInfo{
		"Uncle Bob":       BookInfo{Name: "Clean Code", Genre: "Programming", Published: 1977},
		"J.K. Rowling":    BookInfo{Name: "Harry Potter and the Sorcerer's Stone", Genre: "Fantasy", Published: 1997},
		"Jordan Peterson": BookInfo{Name: "12 Rules of Life", Genre: "Self-Help", Published: 1934},
		"Andy Hunt":       BookInfo{Name: "The Pragmatic Programmer", Genre: "Programming", Published: 1999},
	}

	book, found := books["Uncle Bob"]
	if found {
		fmt.Printf("Buku yang ditulis oleh Uncle Bob : %s \n", book.Name)
		fmt.Printf("Genre: %s\n", book.Genre)
		fmt.Printf("Di Publish pada tahun: %d\n", book.Published)
	} else {
		fmt.Println("Penulis Uncle Bob belum terdaftar.")
	}

	// Adding a new book to the map
	books["J.R.R. Tolkien"] = BookInfo{Name: "The Lord of the Rings", Genre: "Fantasy", Published: 1954}

	// Deleting a book from the map
	delete(books, "J.K. Rowling")

	// Iterating over the map
	fmt.Println("Semua buku yang ada:")
	for author, book := range books {
		fmt.Printf("Author: %s\n", author)
		fmt.Printf("Nama Buku: %s\n", book.Name)
		fmt.Printf("Genre: %s\n", book.Genre)
		fmt.Printf("Published: %d\n", book.Published)
		fmt.Println("--------------------")
	}
}
