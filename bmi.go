package main

import (
	"fmt"
)

func bmi() {
	var beratBadan, tinggiBadan float64
	fmt.Println("====================================")
	fmt.Println("Perhitungan BMI")
	fmt.Println("====================================")
	fmt.Print("Masukkan berat badan (kg): ")
	fmt.Scanln(&beratBadan)
	fmt.Print("Masukkan tinggi badan (cm): ")
	fmt.Scanln(&tinggiBadan)

	bmi := beratBadan / ((tinggiBadan / 100) * (tinggiBadan / 100))
	fmt.Printf("BMI Anda: %.2f\n", bmi)
	switch {
	case bmi < 18.5:
		fmt.Println("__________")
		fmt.Println("|Kurus|")
	case bmi >= 18.5 && bmi < 25:
		fmt.Println("__________")
		fmt.Println("|Normal|")
	case bmi >= 25 && bmi < 30:
		fmt.Println("__________")
		fmt.Println("|Gemuk|")
	default:
		fmt.Println("__________")
		fmt.Println("|Obesitas|")
	}
}
