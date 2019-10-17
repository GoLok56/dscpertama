package main

import "fmt"

var umur int = 17

func main() {
	const kampus string = "Universitas Komputer Indonesia"
	nama := "James Catalunya"

	var (
		semester = 3
		ipk      float32
	)
	ipk = 6.81

	fmt.Println("Saya", nama, "kuliah di", kampus, "semester", semester, "Saya berumur", umur, "tahun.")
	fmt.Println("Ipk saya adalah ", ipk)
}
