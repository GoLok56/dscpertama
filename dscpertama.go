package main

import "fmt"

func main() {
	var umurGebetan = 21

	if umurGebetan > 20 {
		fmt.Println("Udah mateng")
	} else if umurGebetan > 17 {
		fmt.Println("Dede gemes")
	} else {
		fmt.Println("Buset kamu pedopil?")
	}
}
