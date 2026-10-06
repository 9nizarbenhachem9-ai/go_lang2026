package main

import (
	"fmt"
)

func variable() {

	fmt.Println("===============================")
	fmt.Println("Form Registration")
	fmt.Println("===============================")
	var Nama, Alamat, JenisKelamin string
	var NIK, Umur int
	var tinggiBadan, beratBadan float64
	fmt.Print("Masukkan Nama: ")
	fmt.Scanln(&Nama)
	fmt.Print("Masukkan NIK: ")
	fmt.Scanln(&NIK)
	fmt.Print("Masukkan Alamat: ")
	fmt.Scanln(&Alamat)
	fmt.Print("Masukkan Umur: ")
	fmt.Scanln(&Umur)
	fmt.Print("Masukkan Jenis Kelamin: ")
	fmt.Scanln(&JenisKelamin)
	fmt.Print("Masukkan Tinggi Badan: ")
	fmt.Scanln(&tinggiBadan)
	fmt.Print("Masukkan Berat Badan: ")
	fmt.Scanln(&beratBadan)

	fmt.Println("====================================")
	fmt.Println("Data Diri")
	fmt.Println("====================================")
	fmt.Println("Nama:", Nama)
	fmt.Println("NIK:", NIK)
	fmt.Println("Alamat:", Alamat)
	fmt.Println("Umur:", Umur)
	fmt.Println("Jenis Kelamin:", JenisKelamin)
	fmt.Println("Tinggi Badan:", tinggiBadan)
	fmt.Println("Berat Badan:", beratBadan)

}
