package main

import "fmt"

func main() {
	var perulangan, hasil int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&perulangan)
	for i := 1; i <= perulangan; i++ {
		hasil = hasil + i
		if i == perulangan {
			fmt.Printf("Maka hasilnya %.d", hasil)
		}
	}
}