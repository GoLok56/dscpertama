package main

import "fmt"

func main() {
	var nilai [3]int
	nilai[0] = 80
	nilai[1] = 90
	nilai[2] = 40

	fmt.Println(nilai)

	warna := []string{"Hitam", "Merah", "Hijau"}
	for i, value := range warna {
		fmt.Println("Warna ke", i+1, "adalah", value)
	}
}
