package main

import (
	"fmt"
	"time"
)

func main() {
	input := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"
	var abcdefg = map[string]int{"a": 0, "b": 0, "c": 0, "d": 0, "e": 0, "f": 0, "g": 0}
	var hijklmn = map[string]int{"h": 0, "i": 0, "j": 0, "k": 0, "l": 0, "m": 0, "n": 0}
	var opqrstu = map[string]int{"o": 0, "p": 0, "q": 0, "r": 0, "s": 0, "t": 0, "u": 0}
	var vwxyz = map[string]int{"v": 0, "w": 0, "x": 0, "y": 0, "z": 0}
	go func() {
		for _, v := range input {
			x := string(v)
			var _, isExist = abcdefg[x]
			if isExist {
				abcdefg[x] += 1
				fmt.Printf("%s : %d\n", x, abcdefg[x])
			}

		}
	}()
	go func() {
		for _, v := range input {
			x := string(v)
			var _, isExist = hijklmn[x]
			if isExist {
				hijklmn[x] += 1
				fmt.Printf("%s : %d\n", x, hijklmn[x])
			}
		}

	}()
	go func() {
		for _, v := range input {
			x := string(v)
			var _, isExist = opqrstu[x]
			if isExist {
				opqrstu[x] += 1
				fmt.Printf("%s : %d\n", x, opqrstu[x])
			}
		}

	}()
	go func() {
		for _, v := range input {
			x := string(v)
			var _, isExist = vwxyz[x]
			if isExist {
				vwxyz[x] += 1
				fmt.Printf("%s : %d\n", x, vwxyz[x])
			}
		}

	}()
	time.Sleep(1 * time.Second)
}
