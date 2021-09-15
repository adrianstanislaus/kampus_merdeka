package main

import (
	"fmt"
	"math"
)

func primeX(number int) int {
	if number == 1 {
		return 2
	} else if number == 2 {
		return 3
	} else {
		var prima int = 5
		for i := 3; i < number; {
			prima += 2
			pembagi_terbesar := math.Floor(math.Sqrt(float64(prima)))
			for j := 3; j <= int(pembagi_terbesar); j += 2 {
				if (prima)%j == 0 {
					i--
					break
				}
			}
			i++

		}
		return prima
	}

}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}
