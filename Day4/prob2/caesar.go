package main

import "fmt"

func Caesar(offset int, input string) string {
	var hasil string
	for _, v := range input {
		angka := int(v) + (offset % 26)
		if angka > 122 {
			angka = angka - 122 + 96
		}
		hasil = hasil + string(rune(angka))
	}
	return hasil
}

func main() {
	fmt.Println(Caesar(3, "abc"))
	fmt.Println(Caesar(2, "alta"))
	fmt.Println(Caesar(10, "alterraacademy"))
	fmt.Println(Caesar(1, "abcdefghijklmnopqrstuvwxyz"))
	fmt.Println(Caesar(1000, "abcdefghijklmnopqrstuvwxyz"))

}
