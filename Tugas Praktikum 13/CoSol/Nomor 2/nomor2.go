package main

import "fmt"

func main() {
	var angka int
	var continueloop bool
	for continueloop=true; continueloop; {
		fmt.Scan(&angka)
		continueloop=angka <= 0
	}
	fmt.Printf("%d adalah bilangan bulat positif\n", angka)
}