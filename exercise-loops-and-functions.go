package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	t,z := 0.0, 1.0
	for {
		z, t = z - (z*z - x)/(2*z), z
//		fmt.Printf("z is %v & t is %v. \n", z, t)
		if math.Abs(z - t) < 1e-8 {
			break
		}
	}
	return z
}

func main() {
	i := 2.
	fmt.Println(Sqrt(i))
	fmt.Println(Sqrt(i) == math.Sqrt(i))
}
