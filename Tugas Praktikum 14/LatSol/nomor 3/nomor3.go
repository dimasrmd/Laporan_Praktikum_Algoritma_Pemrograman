package main

import "fmt"

func main() {
	const hasil1 = "merah"
	const hasil2 = "kuning"
	const hasil3 = "hijau"
	const hasil4 = "ungu"
	var jumlah int
	var masukan1, masukan2, masukan3, masukan4 string
	percobaan := 5
	jumlah = 0
	for i := 1; i <= percobaan; i++ {
		fmt.Scan(&masukan1, &masukan2, &masukan3, &masukan4)
		if masukan1 == hasil1 && masukan2 == hasil2 && masukan3 == hasil3 && masukan4 == hasil4 {
			jumlah++
		}
	}
	if jumlah == 5 {
		fmt.Print("BERHASIL: true")
	} else {
		fmt.Print("BERHASIL: false")
	}

}
