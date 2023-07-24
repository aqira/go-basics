package main

import "fmt"

func main() {
	//menyimpan nilai lebih dari 1 kita dapat menggunakan array atau slice

	//array harus ditentukan batas nilainya
	var hobbies [3]string
	hobbies[0] = "berenang" //nilai pertama adalah nilai
	hobbies[1] = "musik"
	hobbies[2] = "jalan-jalan"
	//hobbies[3] = "tidur"		nilai ini tidak bisa dimasukkan, karena diluar batas [3]string (yaitu 0,1,2)

	fmt.Println(hobbies)

	//berikutnya, Slice menyimpan nilai dengan tidak ada batas
	var favouriteFoods []string
	favouriteFoods = append(favouriteFoods, "Pizza")
	favouriteFoods = append(favouriteFoods, "Burger")
	favouriteFoods = append(favouriteFoods, "Sushi")
	favouriteFoods = append(favouriteFoods, "Nasi Goreng")
	favouriteFoods = append(favouriteFoods, "Ayam Bakar")
	favouriteFoods = append(favouriteFoods, "Coto Makasar")
	favouriteFoods = append(favouriteFoods, "Batagor")
	
	fmt.Println(favouriteFoods)
}
