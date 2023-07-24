package main

import "fmt"

func main() {
	//for digunakan untuk perulangan. Perulangan akan sesuai dengan ketentuan yang diberikan

	var hobbies [3]string

	// struktur for sebagai berikut:
	// 1. [variable perulangan]; 
	// 2. [periksa apakah kondisi sudah false, jika tidak false maka lakukan statement];
	// 3. [variable perulangan ditambah setalah statement telah dijalankan]
	for index := 0; index < len(hobbies); index++ {
		fmt.Scan(&hobbies[index])
	}

	//berikutnya, dengan data yang kita miliki, kita akan mencetak seluruh hobby dari awal hingga terakhir

	//index : untuk menunjukan posisi hobby pada variable hobbies
	//hobby : nilai yang telah dimasukkan
	for index, hobby := range hobbies {
		fmt.Println("hobby ke",index + 1, "anda adalah", hobby)
	}
}
