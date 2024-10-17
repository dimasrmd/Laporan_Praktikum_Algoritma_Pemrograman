package main

import "fmt"

func main() {
	var c float64
	fmt.Print("Masukkan temperatur dalam celcius: ")
	fmt.Scan(&c)
	f := ((9 * c) / 5) + 32
	r := (c * 4) / 5
	k := (f + 459.67) * 5 / 9

	var bulat int = int(k)
	fmt.Println("Derajat Fahrenheit: ", f)
	fmt.Println()
	fmt.Println("Temperatur Celcius: ", c)
	fmt.Println("Derajat Reamur: ", r)
	fmt.Println("Derajat Fahrenheit: ", f)
	fmt.Println("Derajat Kelvin: ", bulat)
}
