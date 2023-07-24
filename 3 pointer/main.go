package main

import "fmt"

func main() {
	//pada Go, kita akan menggunakan pointer untuk mengubah nilai sebuah variable
	var name string = "John"
	fmt.Println("Nama lama : ", name)

	//sekarang kita ingin mengubah name dengan input user menggunakan Scan
	//&[var name]				berfungsi untuk memberikan alamat variable.
	fmt.Scan(&name)

	fmt.Println("Nama baru : ", name)


	//jika kita ingin mengetahui alamat memori dari sebuah variable:
	fmt.Println("Alamat variable name : ", &name)

	//jika sebuah variable ingin menyimpan sebuah alamat memori
	var addr *string = &name		//addr menunjuk alamat variable name

	fmt.Println("Alamat variable addr : ", addr)
	fmt.Println("Cara mengakses nilai name dari variable addr : ", *addr)
}
