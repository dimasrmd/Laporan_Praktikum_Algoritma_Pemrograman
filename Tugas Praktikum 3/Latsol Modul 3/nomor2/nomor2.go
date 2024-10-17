package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	const phi = 3.1415926535
	fmt.Print("Masukkan jari-jari: ")
	fmt.Scan(&r)
	volume := (4.0 / 3.0) * phi * math.Pow(r, 3)
	luas := 4 * phi * math.Pow(r, 2)
	fmt.Printf("Bola dengan jari-jari %.4f memiliki volume %.4f dan luas kulit %.4f", r, volume, luas)
}