package main

import "fmt"

func main() {
	var angka int
	fmt.Scan(&angka)
	for angka > 0 {
		otp := angka % 10
		fmt.Println(otp)
		angka = angka / 10
	}
}
