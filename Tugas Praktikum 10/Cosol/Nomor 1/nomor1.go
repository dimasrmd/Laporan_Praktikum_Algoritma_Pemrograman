package main

import "fmt"

func main() {
	var umur int
	var kk bool
	fmt.Scan(&umur, &kk)
	if umur >= 17 && kk {
		fmt.Print("Bisa membuat KTP")
	} else {
		fmt.Print("Belum bisa membuat")
	}
}
