package main
import (
	"fmt" 
	"math"
	)
func main() {
	var r float64
	fmt.Print("Masukkan panjang jari-jari lingkaran dalam satuan cm: ")
	fmt.Scanln(&r)
	luas :=math.Pi*r*r
	fmt.Printf("Luas lingkaran dengan jari-jari %.2f adalah %.2f\n", r, luas)
}