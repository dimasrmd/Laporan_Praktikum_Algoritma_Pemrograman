package main

import "fmt"

func main() {
	var kerucut, jarijari, tinggi float64
	phi := 3.14159265358979323846
	fmt.Print("Ada berapa kerucut? ")
	fmt.Scan(&kerucut)
	for i := 1; i <= int(kerucut); i++ {
		fmt.Printf("Jari-jari dan tinggi untuk kerucut ke-%.d: ", i)
		fmt.Scan(&jarijari, &tinggi)
		hasil := (phi * jarijari * jarijari * tinggi) / 3
		fmt.Printf("Volume kerucut ke-%d adalah %.14f\n", i, hasil)
	}
}