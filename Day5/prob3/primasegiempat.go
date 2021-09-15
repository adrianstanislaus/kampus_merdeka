package main

import "fmt"

func primaSegiEmpat(high, wide, start int) {
	total := 0
	for i := 1; i <= wide; i++ {
		for j := 1; j <= high; {
			prima := true
			start += 1
			for x := 2; x <= (start); {
				if (start)%x == 0 {
					prima = false
					break
				}
			}
			if prima {
				j++
				total += start
				fmt.Print(start)
			}
		}
		fmt.Println("")
	}
	fmt.Println(total)
}

func main() {
	primaSegiEmpat(2, 3, 13)
}
