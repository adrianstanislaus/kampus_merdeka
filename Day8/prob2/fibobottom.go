package main

import "fmt"

func fibo(n int) int {
	var fibo = []int{0, 1}
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		for i := 2; i <= n; i++ {
			fibo = append(fibo, fibo[i-1]+fibo[i-2])
		}
	}
	return fibo[n]

}

func main() {

	fmt.Println(fibo(0)) // 0

	fmt.Println(fibo(1)) // 1

	fmt.Println(fibo(2)) // 1

	fmt.Println(fibo(3)) // 2

	fmt.Println(fibo(5)) // 5

	fmt.Println(fibo(6)) // 8

	fmt.Println(fibo(7)) // 13

	fmt.Println(fibo(9)) // 13

	fmt.Println(fibo(10)) // 55

}
