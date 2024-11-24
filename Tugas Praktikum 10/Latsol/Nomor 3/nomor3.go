package main

import "fmt"

func main() {
	var bilangan, jumlahfaktor int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)
	jumlahfaktor = 0
	fmt.Print("Faktor: ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Printf("%d ", i)
			jumlahfaktor = jumlahfaktor + 1
		}
	}
	fmt.Println("")
	if jumlahfaktor == 2 {
		fmt.Print("Prima: true")
	} else if jumlahfaktor != 2 {
		fmt.Print("Prima: false")
	}
}
