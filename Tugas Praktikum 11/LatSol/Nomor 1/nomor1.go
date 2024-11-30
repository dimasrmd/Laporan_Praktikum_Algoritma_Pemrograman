package main

import "fmt"

func main() {
	var ph float64
	fmt.Scan(&ph)
	switch {
	case ph >= 6.5 && ph <= 8.6:
		fmt.Print("Air Layak Minum")
	case (ph >= 0 && ph < 6.5) || (ph > 8.6 && ph <= 14):
		fmt.Print("Air Tidak Layak Minum")
	default:
		fmt.Print("Nilai pH tidak valid. Nilai pH harus antara 0 dan 14.")
	}
}
