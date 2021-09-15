package main

import "fmt"

func arrayMerge(arrayA, arrayB []string) []string {
	if len(arrayA) == 0 {
		return arrayB
	} else {
		for i := 0; i < len(arrayB); i++ {
			beda := true
			for j := 0; j < len(arrayA) && beda; j++ {
				if arrayB[i] == arrayA[j] {
					beda = false
				}
			}
			if beda {
				arrayA = append(arrayA, arrayB[i])
			}
		}
	}
	return arrayA
}

func main() {
	fmt.Println(arrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(arrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(arrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(arrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(arrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(arrayMerge([]string{}, []string{}))

}
