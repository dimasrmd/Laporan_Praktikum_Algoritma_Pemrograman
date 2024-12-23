package main

import "fmt"

func main() {
	var n, pembagi, jumlah float64
	fmt.Scan(&n)
	for n != 9999 {
		jumlah += n
		pembagi += 1
		fmt.Scan(&n)
	}
	fmt.Print("Hasil: ", jumlah/pembagi)
}
