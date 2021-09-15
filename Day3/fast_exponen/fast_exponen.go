package main

import "fmt"

func pangkat(base, pangkat int) int {
	var eksponen int = 1
	if pangkat%2 == 0 {
		for i := 1; i <= pangkat/2; i++ {
			eksponen *= base * base
		}
	} else {
		for i := 1; i <= pangkat/2; i++ {
			eksponen *= base * base
		}
		eksponen *= base
	}
	return eksponen
}

func main() {
	fmt.Println(pangkat(2, 3))
	fmt.Println(pangkat(7, 2))
	fmt.Println(pangkat(10, 5))
	fmt.Println(pangkat(17, 6))
	fmt.Println(pangkat(5, 3))
}
