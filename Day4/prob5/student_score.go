package main

import "fmt"

type Student struct {
	name  []string
	score []int
}

func (s Student) Average() float64 {
	var total int
	for _, v := range s.score {
		total += v
	}
	return float64(total / len(s.score))
}

func (s Student) Min() (min int, name string) {
	var namanilai = map[int]string{}
	var nilaimin int = s.score[0]
	for i := 0; i < len(s.name); i++ {
		namanilai[s.score[i]] = s.name[i]
		if s.score[i] < nilaimin {
			nilaimin = s.score[i]
		}
	}
	return nilaimin, namanilai[nilaimin]
}

func (s Student) Max() (min int, name string) {
	var namanilai = map[int]string{}
	var nilaimax int = s.score[0]
	for i := 0; i < len(s.name); i++ {
		namanilai[s.score[i]] = s.name[i]
		if s.score[i] > nilaimax {
			nilaimax = s.score[i]
		}
	}
	return nilaimax, namanilai[nilaimax]
}

func main() {
	var a = Student{}

	for i := 1; i < 6; i++ {

		var name string

		fmt.Print("Input " + string(rune(i)) + "Student's name : ")

		fmt.Scan(&name)

		a.name = append(a.name, name)

		var score int

		fmt.Print("Input " + name + " Score : ")

		fmt.Scan(&score)

		a.score = append(a.score, score)

	}

	fmt.Println("\n\nAvarage Score Students is", a.Average())

	scoreMax, nameMax := a.Max()

	fmt.Println("Max Score Students is :"+nameMax+"(", scoreMax, ")")

	scoreMin, nameMin := a.Min()

	fmt.Println("Min Score Students is :"+nameMin+"(", scoreMin, ")")

}
