package main

import "fmt"

func main() {
	var diskon, harga, hasil int
	fmt.Print("Masukkan Harga: Rp. ")
	fmt.Scan(&harga)
	fmt.Print("Masukkan Diskon dalam persen: ")
	fmt.Scan(&diskon)

	hasil = harga - (harga * diskon / 100)

	fmt.Printf("Harga setelah diberi diskon: Rp. %.d", hasil)
}
