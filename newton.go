package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 0.0
	for z:= 1.0 ; z < 100; z++ {
		z = z - (z * z - x)/(2 * z)
		realAnswer := math.Sqrt(x)

		if z - realAnswer < 1 {
			fmt.Println("Ответ", z ,"похож на правду")
			return z
		}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
