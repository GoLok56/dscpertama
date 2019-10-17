package main

import "fmt"

func main() {
	hasilPenjumlahan := add(5, 7)
	fmt.Println(hasilPenjumlahan)

	printInfo("James Catalunya", 17)
}

func add(a, b int) int {
	return a + b
}

func printInfo(nama string, umur int) {
	fmt.Println("Nama saya", nama, "Umur saya ", umur)
}
