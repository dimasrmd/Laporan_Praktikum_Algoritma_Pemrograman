package main

import (
	"fmt"
	"math"
)

func main() {
	var x1, x2, x3, y1, y2, y3 float64
	fmt.Print("Masukkan koordinat ke-1 dalam (x, y): ")
	fmt.Scan(&x1, &y1)
	fmt.Print("Masukkan koordinat ke-2 dalam (x, y): ")
	fmt.Scan(&x2, &y2)
	fmt.Print("Masukkan koordinat ke-3 dalam (x, y): ")
	fmt.Scan(&x3, &y3)

	AB := math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
	BC := math.Sqrt(math.Pow(x2-x3, 2) + math.Pow(y2-y3, 2))
	CA := math.Sqrt(math.Pow(x3-x1, 2) + math.Pow(y3-y1, 2))
	//jadi disini math.Sqrt digunakan untuk operasi akar
	//sedangkan math.Pow digunakan untuk operasi perpangkatan

	sisiTerpanjang := math.Max(AB, math.Max(BC, CA))
	//ini math.Max digunakan untuk mmebandingkan nilai tertinggi antara dua perbandingan

	fmt.Print(sisiTerpanjang)
}
