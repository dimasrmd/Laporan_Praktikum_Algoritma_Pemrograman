package main

import "fmt"

func main() {
	var orang, motor int
	fmt.Scan(&orang)
	if orang%2 == 0 {
		motor = orang / 2
	}
	if orang%2 != 0 {
		motor = orang/2 + 1
	}
	fmt.Print(motor)
}
