package main

import (
	"fmt"
	"sort"
)

func DragonOfLoowater(dragonHead, knightHeight []int) {
	index := 0
	jumlah := 0
	totalheight := 0
	if len(dragonHead) > len(knightHeight) {
		fmt.Println("Knight fall")
		return
	} else {
		sort.Ints(knightHeight)
		sort.Ints(dragonHead)
		for _, v := range dragonHead {
			nemu := false
			for j := index; j < len(knightHeight); j++ {
				if v <= knightHeight[j] {
					index = j + 1
					nemu = true
					jumlah += 1
					totalheight += knightHeight[j]
					break
				}
			}
			if !nemu {
				fmt.Println("Knight fall")
				return
			}
		}
	}
	fmt.Println(totalheight)

}

func main() {
	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4}) // 11

	DragonOfLoowater([]int{5, 10}, []int{5}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2}) // knight fall

	DragonOfLoowater([]int{7, 2}, []int{2, 1, 8, 5}) // 10
}
