package main

import "fmt"

func main() {
	var bilangan int
	var pembuktian string
	pembuktian = "Bukan"
	fmt.Scan(&bilangan)
	if bilangan < 0 && bilangan%2 == 0 {
		pembuktian = "Genap Negatif"
	}
	fmt.Print(pembuktian)
}
