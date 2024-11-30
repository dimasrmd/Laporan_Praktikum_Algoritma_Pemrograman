package main

import "fmt"

func main() {
	var jam, harga float64
	var jenis_kendaraan string
	fmt.Scan(&jenis_kendaraan, &jam)
	switch jenis_kendaraan {
	case "Motor", "motor":
		if jam <= 1 && jam > 0 {
			jam = 1
			harga = jam * 2000
		} else if jam > 1 {
			harga = jam * 2000
		} else {
			fmt.Print("Masukkan jam yang sesuai.")
		}
		fmt.Printf("Rp. %0.f", harga)
	case "Mobil", "mobil":
		if jam <= 1 && jam > 0 {
			jam = 1
			harga = jam * 5000
		} else if jam > 1 {
			harga = jam * 5000
		} else {
			fmt.Print("Masukkan jam yang sesuai.")
		}
		fmt.Printf("Rp. %0.f", harga)
	case "Truk", "truk":
		if jam <= 1 && jam > 0 {
			jam = 1
			harga = jam * 8000
		} else if jam > 1 {
			harga = jam * 8000
		} else {
			fmt.Print("Masukkan jam yang sesuai.")
		}
		fmt.Printf("Rp. %0.f", harga)
	default:
		fmt.Print("Masukkan jenis kendaraan dengan benar.")
	}
}
