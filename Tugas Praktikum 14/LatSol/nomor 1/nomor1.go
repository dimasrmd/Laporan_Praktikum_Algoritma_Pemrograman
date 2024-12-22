package main

import "fmt"

func main() {
	var angka, hasil int
	fmt.Scan(&angka)
	hasil = 0
	for i := 1; i <= angka; i++ {
		if i%2 != 0 {
			hasil++
		}
	}
	fmt.Printf("Terdapat %d bilangan ganjil.", hasil)
}