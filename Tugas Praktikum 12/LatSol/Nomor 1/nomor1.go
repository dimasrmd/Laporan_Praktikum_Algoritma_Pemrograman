package main

import "fmt"

func main() {
	const kunci1 = "Admin"
	const kunci2 = "Admin"
	var masukan1, masukan2 string
	var cek bool
	fmt.Scan(&masukan1, &masukan2)
	cek = masukan1 == kunci1 && masukan2 == kunci2
	percobaan := 0
	for cek == false {
		fmt.Scan(&masukan1, &masukan2)
		cek = masukan1 == kunci1 && masukan2 == kunci2
		percobaan++
	} 
	fmt.Print(percobaan, "percobaan gagal login")
}