package main

import "fmt"

func main() {
	//tipe data sangat berguna untuk memastikan masukan tidak salah, 
	//contohnya kita ingin nama kita sebagai nomor, dan umur tidak sebagai "delapan", yang seharusnya 8 (agar dapat di hitung)
	
	//string			: menambahkan huruf
	var name string = "John Doe"

	//contoh salah:
	// var age string = 24

	//int					: nomor yang dapat menjadi negatif dan positif
	//uint				: nomor yang hanya dapat menjadi positif
	var age uint = 65
	var remainingMoney int = -10
	//contoh salah:
	// var age string = 8

	//bool				: kondisi true (jika benar) atau false (jika salah)
	var isAwake bool = true

	//nilai konstan (const) berfungsi seperti var, namun tidak dapat berubah nilainya jika sudah di set
	const gender = "Laki-laki"
	age = 24 									//var dapat melakukan reassignment
	//gender = "Perempuan" 		//const akan error jika mencoba reassingment


	fmt.Println("Nama saya adalah", name, "berumur", age, "saya adalah", gender)
	fmt.Println("Apakah saya sedang bangun?", isAwake)
	fmt.Println("Saya memiliki uang sebanyak: ", remainingMoney)
}
