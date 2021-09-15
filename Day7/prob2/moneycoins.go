package main

import "fmt"

func moneyCoins(money int) []int {
	var koin = []int{}
	var pecahan = []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	for i := 0; money > 0; {
		if money-pecahan[i] >= 0 {
			money -= pecahan[i]
			koin = append(koin, pecahan[i])
		} else {
			i++
		}
	}
	return koin

}

func main() {
	fmt.Println(moneyCoins(123)) // [100 20 1 1 1]

	fmt.Println(moneyCoins(432)) // [200 200 20 10 1 1]

	fmt.Println(moneyCoins(543)) // [500, 20, 20, 1, 1, 1]

	fmt.Println(moneyCoins(7752)) // [5000, 2000, 500, 200, 50, 1, 1]

	fmt.Println(moneyCoins(15321)) // [10000 5000 200 100 20 1]
}
