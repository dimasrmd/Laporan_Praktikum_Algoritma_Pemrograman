package main

import "fmt"

func main() {
	var target, masukan, temp int
	fmt.Scan(&target)
	urutan := 0
	cek := false
	for !cek {
		fmt.Scan(&masukan)
		temp = temp + masukan
		urutan++
		fmt.Printf("Donatur %d: menyumbang %d. Total terkumpul: %d\n", urutan, masukan, temp)
		cek = temp >= target
	}
	fmt.Printf("Targer tercapai! Total donasi: %d dari %d donatur.", temp, urutan)
}
