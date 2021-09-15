//Problem 3
package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scanf("%d\r", &bilangan)
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Println(i)
		}
	}

}
