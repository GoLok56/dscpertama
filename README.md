# Catatan
Ada 3 cara untuk menjalankan program dalam bahasa pemrograman Go. Pertama bisa langsung menggunakan perintah `go run`. Kedua bisa membuat file executable terlebih dahulu menggunakan perintah `go build` install `go install`. Apa perbedaan dari ketiga perintah tersebut?

## go run
Dengan menggunakan perintah `go run`, go akan melakukan kompilasi terhadap kode yang kita tulis dan **langsung** menjalankannya. Sehingga hasil dari kode yang kita tulis bisa langsung dilihat hasilnya.

## go build
Perintah ini akan melakukan kompilasi dan menghasilkan file executable **di folder** kita memanggil perintah ini.

## go install
Tidak berbeda dengan perintah `go build`, hanya saja file executable akan disimpan di folder `$GOPATH/bin`.
