package main

import (
	"fmt"
)

func loop() { //increment
	for i := 0; i <= 5; i++ {
		fmt.Println("Perulangan ke-", i)
		for p := 0; p <= 5; p++ {
			fmt.Println("Perulangan ke-", p)
		}
		j := i * 5
		fmt.Println("Nilai j:", j)
	}
}

func loop2() { //decrement
	for y := 5; y >= 0; y-- {
		fmt.Println("Perulangan ke-", y)
	}
}

func loop3() { //break
	z := 0
	x := 1
	for z < 10 {
		fmt.Println("Perulangan ke-", z)
		z++
	}

	for z < 10 {
		fmt.Println("Perulangan ke-", z)
		z++
	}
	for x > 0 {
		fmt.Println("Perulangan ke-", x)
		x++
		if x == 1000 {
			break
		}
	}
}
