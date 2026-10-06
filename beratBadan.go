package main

import (
	"fmt"
)

func beratBadan() {
	var beratBadan float64
	fmt.Print("Masukkan berat badan: ")
	fmt.Scanln(&beratBadan)
	if beratBadan < 50 {
		fmt.Println("Berat badan kurang")
	} else if beratBadan >= 50 && beratBadan <= 70 {
		fmt.Println("Berat badan ideal")
	} else {
		fmt.Println("Berat badan berlebih")
	}
}
