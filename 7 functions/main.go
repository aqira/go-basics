package main

import "fmt"

//variable ini dapat di akses ke semua fungsi pada file ini
var correctUsername string = "admin"
var correctPassword string = "rahasia"

func main() {
	//function (fungsi) adalah cara agar kita dapat memisahkan beberapa
	//program yang kita telah buat agar lebih jelas, dan dapat digunakan berulang-ulang kali jika dibutuhkan.

	//struktur function sebagai berikut:

	//func [nama function]([(optional) parameter1], [(optional) parameter2]...)

	//nama function 	: digunakan biasanya untuk mendeskripsikan function sesuai dengan statementnya
	//parameter				: syarat untuk menggunakan function (berupa syarat tipe data yang digunakan pada function)
	//argument				: nilai yang dimasukan pada saat memanggil sebuah function yang ada syarat (parameter)-nya

	//function dapat menjalankan statement didalamnya dan tidak mengembalikan nilai apapun
	//jika function dibutuhkan untuk mengembalikan nilai, maka function akan "return" sebuah nilai sesuai dengan tipe datanya

	//pada contoh ini, kita akan mencoba untuk membuat logika untuk login, username dan password harus sesuai

	fmt.Println("Masukan username: ")
	var username string
	fmt.Scan(&username)

	fmt.Println("Masukan password: ")
	var password string
	fmt.Scan(&password)

	var loginSuccessful bool = isLoginAttemptSuccessful(username, password) //username dan password disini adalah argumen

	loginNotification(loginSuccessful)
}

//jika login berhasil, maka akan mengembalikan true, namun jika gagal maka false
func isLoginAttemptSuccessful(username string, password string) bool { //username dan password disini adalah parameter
	return username == correctUsername && password == correctPassword //mengembalikan nilai berupa bool
}

func loginNotification(loginSuccessful bool) {
	if loginSuccessful {
		fmt.Println("Selamat! anda berhasil login!")
	} else {
		fmt.Println("Mohon maaf, anda gagal login karena salah memasukan username atau password.")
	}
}
