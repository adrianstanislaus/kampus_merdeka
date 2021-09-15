package main

import "fmt"

func MaxSequence(arr []int) int {
	i := 0
	j := len(arr) - 1
	var max int
	for i != j {
		var jumlah int
		for x := i; x <= j; x++ {
			jumlah += arr[x]
		}
		if jumlah > max {
			max = jumlah
		}
		if arr[i] < arr[j] {
			i++
		} else {
			j--
		}
	}
	return max
}

func main() {
	fmt.Println(MaxSequence([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 5, -6}))
	fmt.Println(MaxSequence([]int{-2, -3, 4, -1, -2, 1, 5, -3}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, -2, -3, 1, 6, -6}))
	fmt.Println(MaxSequence([]int{-2, -5, 6, 2, -3, 1, 6, -6}))
}
