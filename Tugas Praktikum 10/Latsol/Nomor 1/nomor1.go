package main

import "fmt"

func main() {
	var berat, kg, gr, hargakg, hargagr, total int
	fmt.Print("Berat parsel (dalam gram): ")
	fmt.Scan(&berat)
	if berat > 10000 {
		kg = berat / 1000
		gr = berat % 1000
		hargakg = kg * 10000
		if gr >= 500 {
			hargagr = gr * 5
		} else if gr < 500 {
			hargagr = gr * 15
		}
		total = hargakg
	} else if berat <= 10000 {
		kg = berat / 1000
		gr = berat % 1000
		hargakg = kg * 10000
		if gr >= 500 {
			hargagr = gr * 5
		} else if gr < 500 {
			hargagr = gr * 15
		}
		total = hargakg + hargagr
	}
	fmt.Printf("Detail berat: %d kg + %d gram\n", kg, gr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", hargakg, hargagr)
	fmt.Print("Total biaya: Rp. ", total)
}
