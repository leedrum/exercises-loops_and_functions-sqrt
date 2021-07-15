package main

import (
	"fmt"
)

var adjust = 0.00001

func Sqrt(x float64) float64 {
	z := 1.0
	result := 0.0

	for {
		result = (z*z - x) / (2 * z)

		if result < 0 {
			z += adjust
		} else {
			return z
		}

	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
