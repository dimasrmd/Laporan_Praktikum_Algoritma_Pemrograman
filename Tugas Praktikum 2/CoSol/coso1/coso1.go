package main

import "fmt"

func main() {
	var a, b, c, d, e int
	fmt.Print("Masukkan 5 angka yang akan dijumlahkan: ")
	fmt.Scan(&a, &b, &c, &d, &e)
	var hasil int
	hasil = a + b + c + d + e
	fmt.Print("Hasil dari penjumlahannya: ", hasil)
}