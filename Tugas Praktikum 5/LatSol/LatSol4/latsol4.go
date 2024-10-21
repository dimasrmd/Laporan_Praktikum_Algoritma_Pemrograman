package main

import "fmt"

func main() {
	var angka int
	fmt.Print("Masukkan angka yang ingin di faktorialkan: ")
	fmt.Scan(&angka)
	if angka == 0 {
		fmt.Println("Hasil: 1")
	} else {
		for i := angka - 1; i >= 1; i-- {
			angka = angka * i
		}
		fmt.Printf("Hasil: %.d", angka)
	}
}