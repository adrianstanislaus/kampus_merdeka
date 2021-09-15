//Problem 4
package main

import (
	"fmt"
	"math"
)

func primeNumber(number int) string {
	if number <= 1 {
		return "Bukan bilangan prima"
	} else if (number > 2) && (number%2 == 0) {
		return "Bukan bilangan prima"
	} else {
		pembagi_terbesar := math.Floor(math.Sqrt(float64(number)))
		for i := 3; i <= int(pembagi_terbesar); i += 2 {
			if number%i == 0 {
				return "Bukan bilangan prima"
			}
		}
		return "Bilangan prima"
	}
}
func main() {
	fmt.Println(primeNumber(1000000007))
	fmt.Println(primeNumber(1500450271))
	fmt.Println(primeNumber(1000000000))
	fmt.Println(primeNumber(10000000019))
	fmt.Println(primeNumber(10000000033))

}
