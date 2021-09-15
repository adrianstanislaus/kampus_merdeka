package main

import (
	"fmt"
	"strings"
)

func Compare(a, b string) string {
	if len(a) >= len(b) {
		res := strings.Contains(a, b)
		if res {
			return b
		}
	} else {
		res := strings.Contains(b, a)
		if res {
			return a
		}
	}
	return ""
}

func main() {
	fmt.Println(Compare("AKA", "AKASHI"))
	fmt.Println(Compare("KANGOORO", "KANG"))
	fmt.Println(Compare("KI", "KIJANG"))
	fmt.Println(Compare("KUPU-KUPU", "KUPU"))
	fmt.Println(Compare("ILALANG", "ILA"))

}
