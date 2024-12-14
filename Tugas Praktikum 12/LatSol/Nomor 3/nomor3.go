package main

import "fmt"

func main() {
	var angka1, angka2, i int
	fmt.Scan(&angka1, &angka2)
	for i = 0; angka1 >= angka2; i++ {
		angka1 = angka1 - angka2
	}
	fmt.Print(i)
}
