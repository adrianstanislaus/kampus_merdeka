package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var max int
	var maxcard int = -1
	for i := range cards {
		for j := 0; j <= 1; j++ {
			if cards[i][j] == deck[0] || cards[i][j] == deck[1] {
				if cards[i][0]+cards[i][1] > max {
					max = cards[i][0] + cards[i][1]
					maxcard = i
				}

			}
		}
	}
	if maxcard <= -1 {
		return "tutup kartu"
	} else {
		return cards[maxcard]
	}
}

func main() {
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{[]int{6, 5}, []int{3, 3}, []int{3, 4}, []int{2, 1}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
}
