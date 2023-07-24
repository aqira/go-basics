package main

import "fmt"

func main() {
	//if dan switch adalaha conditional statements

	// jika kondisi benar, maka akan dijalankan statementnya
	if true {
		fmt.Println("masuk, karena true")
	}

	//contoh if, else if, dan else
	var stubbornPerson bool = true
	var policeStandingBy bool = true

	if stubbornPerson {
		if policeStandingBy {
			fmt.Println("Di tilang pada saat lampu merah")
		} else if !policeStandingBy {
			fmt.Println("Resiko tertabrak / menabrak")
		}
	} else {
		fmt.Println("Berhenti pada saat lampu merah")
	}

	//conditional statement yang berikutnya juga terkhusus untuk satu statement yang diperiksa nilainya apakah sesuai atau tidak
	var trafficLightColor string = "merah"

	switch trafficLightColor {
	// case "red" || "merah":		--> case tidak bisa memeriksa lebih dari 1 kondisi
	case "merah": // sama seperti: if trafficLightColor == "merah" {...}
		fmt.Println("Berhenti")
	case "kuning":
		fmt.Println("Hati-hati")
	case "hijau":
		fmt.Println("Melaju")
	}
}
