package main

import (
	"fmt"
	"strings"
)

func operator() {

	var angkaPertama, angkaKedua int

	// Operasi Matematika

	fmt.Println("====================================")
	fmt.Println("Hasil Operasi Matematika")
	fmt.Println("====================================")
	fmt.Println("Masukkan angka pertama:")
	fmt.Scanln(&angkaPertama)
	fmt.Println("Masukkan angka kedua:")
	fmt.Scanln(&angkaKedua)
	fmt.Println("Masukkan operator (+, -, *, /):")
	var operator string
	fmt.Scanln(&operator)
	fmt.Println("Hasil:", operasiMatematika(angkaPertama, angkaKedua, operator))
	fmt.Println("besar", strings.ToUpper("besar"))
	fmt.Println("KECIL", strings.ToLower("KECIL"))
}

func operasiMatematika(angkaPertama, angkaKedua int, operator string) float64 {
	switch operator {
	case "+":
		return float64(angkaPertama + angkaKedua)
	case "-":
		return float64(angkaPertama - angkaKedua)
	case "*":
		return float64(angkaPertama * angkaKedua)
	case "/":
		if angkaKedua == 0 {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
			return 0
		}
		return float64(angkaPertama) / float64(angkaKedua)
	default:
		fmt.Println("Operator tidak valid.")
		return 0
	}
}
