Dokumentasi basic tentang Go Programming Language ini dapat dibaca dari awal (1), hingga terakhir (10)

tujuan dari dokumentasi ini adalah untuk pembelajaran go bagi yang baru, ataupun yang sudah tidak asing dengan bahasa pemrograman
seperti Java, C#, Javascript, dst..

Pastikan anda telah install go pada https://go.dev/doc/install dan dapat melakukan command shell "go". 
Untuk memastikannya dapat dengan mengetik ini dalam shell anda:
go version


Cara menggunakan go pada shell:

1. setiap modul yang ada di dokumentasi ini terdapat prefix basics, jika anda ingin membuat mod sendiri, dapat dengan cara:
go mod init basics/[nama materi]

setelah itu akan dibuatkan file go.mod, yang digunakan untuk mengetahui spesifikasi dari modul project yang dibuat


2. Untuk setiap folder, anda dapat jalankan

dengan cara:
go run .

atau, dengan cara:
go run ./main.go


PS:
Perkembangan kedepannya dapat di tambahkan beberapa topik seperti conversion, time, interface, concurrency