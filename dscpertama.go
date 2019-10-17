package main

import "fmt"

func main() {
	var nilai = []int{80, 20, 10, 40, 90, 50, 70}
	var totalNilai = 0
	for _, value := range nilai {
		totalNilai = totalNilai + value
	}
	var banyakNilai = len(nilai)
	var ip = totalNilai / banyakNilai
	fmt.Println("Ip saya adalah", ip)

	fmt.Println("\nNilai dari for loop auto inc:")
	for i := 0; i < banyakNilai; i++ {
		fmt.Println(nilai[i])
	}

	fmt.Println("\nNilai dari for loop manual inc:")
	i := 1
	for i < banyakNilai {
		fmt.Println(nilai[i])
		i++
	}

}
