package main

import "fmt"

func pairSum(arr []int, target int) []int {
	pairing := true
	i := 0
	j := len(arr) - 1
	var pair []int
	for pairing && (i != j) {
		if arr[i]+arr[j] > target {
			j--
		} else if arr[i]+arr[j] < target {
			i++
		} else {
			pair = []int{i, j}
			pairing = false
		}
	}
	return pair
}

func main() {
	fmt.Println(pairSum([]int{1, 2, 3, 4, 6}, 6))
	fmt.Println(pairSum([]int{2, 5, 9, 11}, 11))
	fmt.Println(pairSum([]int{1, 3, 5, 7}, 12))
	fmt.Println(pairSum([]int{1, 4, 6, 8}, 10))
	fmt.Println(pairSum([]int{1, 5, 6, 7}, 6))
}
