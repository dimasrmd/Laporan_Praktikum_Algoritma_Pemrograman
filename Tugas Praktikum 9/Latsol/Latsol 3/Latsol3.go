package main

import "fmt"

func main() {
	var num1, num2 int
	var bukti1, bukti2 bool
	bukti1 = true
	bukti2 = true
	fmt.Scan(&num1, &num2)
	if num2%num1 != 0 {
		bukti1 = false
	}
	if num1%num2 != 0 {
		bukti2 = false
	}
	fmt.Println(bukti1)
	fmt.Print(bukti2)
}
