package main

import "fmt"

func main() {
	var bilangan, pangkat, hasil int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&bilangan)
	fmt.Print("Masukkan pangkat: ")
	fmt.Scan(&pangkat)
	hasil = 1
	for i := 0; i < pangkat; i++ {
		hasil = hasil * bilangan
	}
	fmt.Println("Hasil: ", hasil)
}
