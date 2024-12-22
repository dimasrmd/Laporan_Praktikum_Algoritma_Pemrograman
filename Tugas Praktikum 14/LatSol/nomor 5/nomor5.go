package main

import "fmt"

func main() {
	var kanan, kiri, total float64
	var cek, oleng bool
	cek = true
	for cek == true {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)
		total = kiri + kanan
		cek = total < 150 && kiri > 0 && kanan > 0
		if cek && kiri > kanan{
			oleng = kiri-kanan > 9
			fmt.Println("Sepeda motor pak Andi akan oleng: ", oleng)
		} else if cek && kiri < kanan{
			oleng = kanan-kiri > 9
			fmt.Println("Sepeda motor pak Andi akan oleng: ", oleng)
		}
	}
	fmt.Print("Proses selesai")
}