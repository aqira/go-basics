package main

import (
	"basics/encapsulation/person"
	"fmt"
)

func main() {

	//encapsulation adalah pengaturan yang bisa kita lakukan jika beberapa function dapat kita panggil jika function tersebut
	//ada di package lain. (dengan cara menggunakan import ke package tersebut)

	//untuk memahami fungsi yang dapat di enkapsulasi,
	//kita perlu mengetahui konvensi penamaan naming convention:
	//PascalCase (huruf kapital di depannya, dan spasi ditandai dengan huruf kapital) seperti : FullName, ShowAll, IsAvailable, dst..
	//camelCase (huruf kecil di depannya, dan spasi ditandai dengan huruf kapital) seperti : fullName, showAll, isAvailable, dst...

	//jika kita menggunakan PascalCase, maka function tersebut dapat kita panggil pada package lain,
	//jika menggunakan camelCase, maka function hanya dapat di panggil pada package itu saja.

	//dapat di panggil, dan digunakan
	person.GetAll()

	for _, person := range person.GetAll() {
		fmt.Println(person)
	}

	//tidak dapat dipanggil
	// person.getByName("Jane")
}
