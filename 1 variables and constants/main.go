package main

import "fmt"

func main() {

	//line of code ini berfungsi sebagai komen. dengan cara menggunakan double slash (//)

	//Pertama, kita bisa mencoba untuk mencetak dengan menggunakan package fmt
	//cara mencetak salah satunya menggunakan Println function
	fmt.Println("Hello World")

	//variable (var) berfungsi untuk menyimpan value (nilai) yang anda butuhkan jika ingin menggunakan nilai tersebut di lain waktu.

	//deklarasi variabel
	var name = "John Doe"
	var age = 65

	//nilai konstan (const) berfungsi seperti var, namun tidak dapat berubah nilainya jika sudah di set
	const gender = "Laki-laki"
	age = 24 									//var dapat melakukan reassignment
	//gender = "Perempuan" 		//const akan error jika mencoba reassingment

	fmt.Println("Nama saya adalah", name, "berumur", age, "saya seorang", gender)
}
