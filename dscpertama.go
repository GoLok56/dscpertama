package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Sebelum diubah tanpa pointer", x)
	changeIntValue(x)
	fmt.Println("Sesudah diubah tanpa pointer", x)

	fmt.Println("Sebelum diubah dengan pointer", x)
	changeIntValuePointer(&x)
	fmt.Println("Sesudah diubah dengan pointer", x)
}

func changeIntValue(x int) {
	x = 2
}

func changeIntValuePointer(x *int) {
	*x = 2
}
