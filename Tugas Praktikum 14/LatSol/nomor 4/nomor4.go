package main

import "fmt"

func main() {
	var bunga, hasil string
	i:=1
	for {
		fmt.Printf("Bunga %d: ", i)
		fmt.Scan(&bunga)
		if bunga == "SELESAI" {
			break
		}
		hasil += bunga + " - "
		i++
	}
	fmt.Printf("Pita: %s\n", hasil)
	fmt.Printf("Bunga: %d\n", i-1)
}