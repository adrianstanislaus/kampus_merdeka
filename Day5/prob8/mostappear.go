package main

import "fmt"

func mergeSort(items []pair) []pair {
	if len(items) < 2 {
		return items
	}
	first := mergeSort(items[:len(items)/2])
	second := mergeSort(items[len(items)/2:])
	return merge(first, second)
}

func merge(a []pair, b []pair) []pair {
	final := []pair{}
	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		if a[i].appear < b[j].appear {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		final = append(final, a[i])
	}
	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}

type pair struct {
	item   string
	appear int
}

func MostAppearItem(items []string) []pair {
	var answers = map[string]int{}
	var hasil []pair
	for _, v := range items {
		answers[v] += 1
	}
	for i, v := range answers {
		hasil = append(hasil, pair{i, v})
	}
	return mergeSort(hasil)
}

func main() {
	fmt.Println(MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	fmt.Println(MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	fmt.Println(MostAppearItem([]string{"football", "basketball", "tenis"}))
}
