package main

import (
	"fmt"
	"math"
)

//bruteforce
func SimpleEquations(a, b, c int) {
	calculating := true
	var x, y, z int
	for x := 1; x*y*z <= b && calculating; x++ {
		for y := 1; x*y*z <= b && calculating; y++ {
			for z = 1; x*y*z <= b && calculating; z++ {
				if (x+y+z == a) && (x*y*z == b) && (math.Pow(float64(x), 2)+math.Pow(float64(y), 2)+math.Pow(float64(z), 2) == float64(c)) {
					calculating = false
				}
			}

		}

	}
	if calculating {
		fmt.Println("tidak ada solusi")
	} else {
		fmt.Println(x, y, z)
	}
}

func main() {
	SimpleEquations(1, 2, 3) // no solution

	SimpleEquations(6, 6, 14) // 1 2 3
}
