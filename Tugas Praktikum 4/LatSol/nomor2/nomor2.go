package main

import (
	"fmt"
	"math"
)

func main() {
	var bmi, tinggibadan, beratbadan float64
	fmt.Print("Masukkan BMI anda: ")
	fmt.Scan(&bmi)
	fmt.Print("Masukkan Tinggi Badan anda: ")
	fmt.Scan(&tinggibadan)

	beratbadan = bmi * math.Pow(tinggibadan, 2)

	fmt.Printf("Maka Berat Badan Anda: %.0f kg", beratbadan)

}