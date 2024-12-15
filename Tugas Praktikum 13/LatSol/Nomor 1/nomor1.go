package main

import "fmt"

func main() {
	var angka, hasil int
	fmt.Scan(&angka)
	for angka > 0 {
		hasil++
		angka = angka/10
	}
	fmt.Print(hasil)
}