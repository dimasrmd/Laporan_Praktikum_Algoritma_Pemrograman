package main

import "fmt"

func main() {
	var bilangan, j int
	fmt.Scan(&bilangan)
	for j = 1; j <= bilangan; j++ {
		if bilangan%j == 0 {
			fmt.Print(j, " ")
		}
	}
}
