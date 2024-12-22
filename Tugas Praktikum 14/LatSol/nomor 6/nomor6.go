package main

import "fmt"

func main() {
	var k int
	var hasil, K float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	K = 0
	hasil = 1
	for i := 0; i <= k; i++ {
		hasil *= ((4*K+2)*(4*K+2))/((4*K+1)*(4*K+3))
		K++
	}
	fmt.Printf("Nilai akar 2 = %.10f",hasil)

}