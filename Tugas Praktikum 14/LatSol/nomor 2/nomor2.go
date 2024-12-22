package main

import "fmt"

func main() {
	var angka, faktor int
	fmt.Scan(&angka)
	for i := 1; i <= angka; i++ {
		if angka%i==0 {
			faktor++
		}
	}
	if faktor == 2 {
		fmt.Print("Prima")
	} else {
		fmt.Print("Bukan Prima")
	}
}