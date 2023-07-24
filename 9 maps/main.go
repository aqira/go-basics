package main

import "fmt"

func main() {
	//Map digunakan jika kita membutuhkan key untuk pencariannya (bukan index)

	//dalam contoh ini kita akan membuat map dengan :
	//key string sebagai nama authornya
	//value string sebagai nama booknya
	books := map[string]string{
		"Uncle Bob":    "Clean Code",
		"J.K. Rowling":    "Harry Potter and the Sorcerer's Stone",
		// "Uncle Bob": "Clean Architecture",			//jika keynya sudah pernah digunakan, maka kita tidak dapat menggunakannya lagi
		"Jordan Peterson":   "12 Rules of Life",
		"Andy Hunt":     "The Pragmatic Programmer",
	}

	// Mencari buku dengan key, jika tidak ada maka akan kita akan menunjukkan error.
	book, found := books["Uncle Bob"]
	if found {
		fmt.Printf("Buku yang ditulis oleh Uncle Bob : %s \n", book)
	} else {
		fmt.Println("Penulis Uncle Bob belum terdaftar.")
	}

	// Cara untuk menambahkan kedalam hashmap
	books["J.R.R. Tolkien"] = "The Lord of the Rings"

	// Cara untuk menghapus buku yang sudah ada sebelumnya
	delete(books, "J.K. Rowling")

	// Iterating over the map
	fmt.Println("Semua buku yang ada:")
	for author, book := range books {
		fmt.Printf("%s: %s\n", author, book)
	}
}