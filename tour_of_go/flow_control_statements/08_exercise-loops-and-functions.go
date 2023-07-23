package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		if tmp-z < 0.00000000001 && tmp-z > -0.00000000001 {
			return z
		}
		z = tmp
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
