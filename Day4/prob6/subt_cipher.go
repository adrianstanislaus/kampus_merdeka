package main

import "fmt"

type Student struct {
	name       string
	nameEncode string
}

type Chiper interface {
	Encode() string
	Decode() string
}

func (s *Student) Encode() string {
	var nameEncode = ""
	for _, v := range s.name {
		angka := 122 - (int(v) - 97)
		nameEncode = nameEncode + string(rune(angka))
	}
	return nameEncode

}

func (s *Student) Decode() string {
	var nameDecode = ""
	for _, v := range s.nameEncode {
		angka := 122 - (int(v) - 97)
		nameDecode = nameDecode + string(rune(angka))
	}
	return nameDecode

}

func main() {

	var menu int

	var a = Student{}

	var c Chiper = &a

	fmt.Print("[1] Encrypt \n[2] Decrypt \nChoose your menu ?")

	fmt.Scan(&menu)

	if menu == 1 {

		fmt.Print("\nInput Student’s Name : ")

		fmt.Scan(&a.name)

		fmt.Print("\nEncode of Student’s Name " + a.name + " is : " + c.Encode())

	} else if menu == 2 {

		fmt.Print("\nInput Student’s Encode Name : ")

		fmt.Scan(&a.nameEncode)

		fmt.Print("\nDecode of Student’s Name " + a.nameEncode + " is : " + c.Decode())

	} else {

		fmt.Println("Wrong input name menu")

	}

}
