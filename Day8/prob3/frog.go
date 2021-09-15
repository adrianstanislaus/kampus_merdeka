package main

import "fmt"

func abs(a, b int) int {
	hasil := a - b
	if hasil < 0 {
		hasil = b - a
	}
	return hasil
}

func Frog(jumps []int) int {
	var min map[int]int = map[int]int{}
	min[0] = 0
	min[1] = abs(jumps[1], jumps[0])
	for i := 2; i < len(jumps); i++ {
		lompat1 := abs(jumps[i], jumps[i-1]) + min[i-1]
		lompat2 := abs(jumps[i], jumps[i-2]) + min[i-2]
		if lompat1 < lompat2 {
			min[i] = lompat1
		} else {
			min[i] = lompat2
		}
	}
	return min[len(jumps)-1]
}

func main() {

	fmt.Println(Frog([]int{10, 30, 40, 20})) // 30

	fmt.Println(Frog([]int{30, 10, 60, 10})) // 40

}
