package main

import "fmt"

func FindMinAndMax(arr []int) string {
	min := arr[0]
	max := arr[0]
	var imin, imax int
	for i := 0; i < (len(arr) - 1); i++ {
		if arr[i+1] < min {
			min = arr[i+1]
			imin = i + 1
		}
		if arr[i+1] > max {
			max = arr[i+1]
			imax = i + 1
		}
	}
	return ("min: " + fmt.Sprint(min) + " index: " + fmt.Sprint(imin) + " max: " + fmt.Sprint(max) + " index: " + fmt.Sprint(imax))
}

func main() {
	fmt.Println(FindMinAndMax([]int{2, -5, -4, 22, 7, 7}))
	fmt.Println(FindMinAndMax([]int{-1, 5, 6, 4, 2, 18}))
	fmt.Println(FindMinAndMax([]int{-2, 5, -7, 4, 7, -20}))
}
