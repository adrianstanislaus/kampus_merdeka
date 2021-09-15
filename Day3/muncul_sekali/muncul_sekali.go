package main

import (
	"fmt"
	"strconv"
)

func munculSekali(angka string) []int {
	var muncul_sekali []int
	for _, v := range angka {
		muncul := 0
		for _, v2 := range angka {
			if v2 == v {
				muncul += 1
			}
		}
		if muncul == 1 {
			nilai, _ := strconv.Atoi(string(v))
			muncul_sekali = append(muncul_sekali, nilai)
		}
	}
	return muncul_sekali

}

func main() {
	fmt.Println(munculSekali("1234123"))
	fmt.Println(munculSekali("76523752"))
	fmt.Println(munculSekali("12345"))
	fmt.Println(munculSekali("1122334455"))
	fmt.Println(munculSekali("0872504"))
}
