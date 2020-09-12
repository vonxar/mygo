package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	z_before := 0.0
	count := 0
	for i := 0; math.Abs(z-z_before) > 0.0000001; i++ {
		z_before = z
		z = z - (math.Pow(z, 2)-x)/(2*z)
		count += 1
	}
	fmt.Printf("(loop count:%d) ", count)
	return z
}

func main() {
	for i := 1.0; i < 10; i++ {
		fmt.Printf("Sqrt(%f) = ", i)
		fmt.Println(
			math.Sqrt(i),
			Sqrt(i),
		)
	}
}
