package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var jumlahtetesan, jumlahtetesan_A, jumlahtetesan_B, jumlahtetesan_C, jumlahtetesan_D int
	var volume_A, volume_B, volume_C, volume_D, volume_x, volume_y, ukurantetesan float64
	ukurantetesan = 0.0001
	fmt.Scan(&jumlahtetesan)
	for i := 1; i <= jumlahtetesan; i++ {
		volume_x = rand.Float64()
		volume_y = rand.Float64()
		if volume_x <= 0.5 && volume_y <= 0.5 {
			jumlahtetesan_A++
		} else if volume_x > 0.5 && volume_y <= 0.5 {
			jumlahtetesan_B++
		} else if volume_x > 0.5 && volume_y > 0.5 {
			jumlahtetesan_C++
		} else if volume_x <= 0.5 && volume_y > 0.5 {
			jumlahtetesan_D++
		}
	}
	volume_A = float64(jumlahtetesan_A) * ukurantetesan
	volume_B = float64(jumlahtetesan_B) * ukurantetesan
	volume_C = float64(jumlahtetesan_C) * ukurantetesan
	volume_D = float64(jumlahtetesan_D) * ukurantetesan
	fmt.Printf("Volume A: %f milimeter\n", volume_A)
	fmt.Printf("Volume B: %f milimeter\n", volume_B)
	fmt.Printf("Volume C: %f milimeter\n", volume_C)
	fmt.Printf("Volume D: %f milimeter\n", volume_D)
}
