package main

import "fmt"

func main() {
	var bilangan int
	var teks string
	fmt.Scan(&bilangan)
	teks = "Positif"
	if bilangan < 0 {
		teks = "negatif"
	} 
	fmt.Print(teks)
}