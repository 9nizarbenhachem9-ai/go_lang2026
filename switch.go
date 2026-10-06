package main

import (
	"fmt"
)

func switchExample() {
	var hari int
	fmt.Print("Masukkan hari (1-7): ")
	fmt.Scanln(&hari)

	switch hari {
	case 1:
		fmt.Println("hari ini adalah hari Senin	")
	case 2:
		fmt.Println("hari ini adalah hari Selasa")
	case 3:
		fmt.Println("hari ini adalah hari Rabu")
	case 4:
		fmt.Println("hari ini adalah hari Kamis")
	case 5:
		fmt.Println("hari ini adalah hari Jumat")
	case 6:
		fmt.Println("hari ini adalah hari Sabtu")

	default:
		fmt.Println("hari ini adalah hari Minggu")
	}

}

func switchExample2() {
	fmt.Println("====================================")
	fmt.Println("Hasil Penilaian")
	fmt.Println("====================================")
	var nilai int
	fmt.Println("Masukkan nilai (0-100): ")
	fmt.Scanln(&nilai)

	switch {
	case nilai == 100:
		fmt.Println("Nilai A+")
	case nilai >= 90:
		fmt.Println("Nilai A")
	case nilai >= 85:
		fmt.Println("Nilai A-")
	case nilai >= 80:
		fmt.Println("Nilai B")
	case nilai >= 70:
		fmt.Println("Nilai C")
	case nilai >= 60:
		fmt.Println("Nilai D")
	default:
		fmt.Println("Nilai E")
	}

}
