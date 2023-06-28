package Solve

import (
	"fmt"
	"strconv"
	"strings"
)

func Solver(s []string) []string {
	var run []rune
	punc := ",.!@#$%^&{}][?!"

	for i := range s {
		run = []rune(s[i])
		if len(run) == 0 {
			break
		}
		if run[0] == '(' && run[len(run)-1] == ')' && (len(run) == 5 || len(run) == 4) {
			if string(run[1:4]) == "cap" {
				s[i] = "(cap, 1)"
			} else if string(run[1:4]) == "low" {
				s[i] = "(low, 1)"
			} else if string(run[1:3]) == "up" {
				s[i] = "(up, 1)"
			} else if string(run[1:4]) == "bin" {
				s[i] = "(bin, 1)"
			} else if string(run[1:4]) == "hex" {
				s[i] = "(hex, 1)"
			}
		}

		run = []rune(s[i])

		// fmt.Println(string(run))
		if run[0] == '(' && run[len(run)-1] == ')' {
			if string(run[1:4]) == "cap" {

				num := findNum(s[i])
				for j := 1; num >= j; j++ {
					if i-j < 0 {
						fmt.Println("There aren't enough arguments.")
						s = []string{}
						return s
					}
					for k := range punc {
						if s[i-j] == string(punc[k]) {
							fmt.Println("There is unavailabe character to turn.")
							s = []string{}
							return s

						}
					}
					rr := []rune(s[i-j])
					if rr[0] == '(' && rr[len(rr)-1] == ')' {
						num = num + 1
					} else {
						s[i-j] = strings.Title(s[i-j])
					}
				}
			} else if string(run[1:4]) == "low" {
				num := findNum(s[i])

				for j := 1; num >= j; j++ {
					if i-j < 0 {
						fmt.Println("There aren't enough arguments.")
						s = []string{}
						return s
					}
					for k := range punc {
						if s[i-j] == string(punc[k]) {
							fmt.Println("There is unavailabe character to turn.")
							s = []string{}
							return s

						}
					}

					rr := []rune(s[i-j])

					if rr[0] == '(' && rr[len(rr)-1] == ')' {
						num = num + 1
					} else {
						s[i-j] = strings.ToLower(s[i-j])
					}
				}
			} else if string(run[1:3]) == "up" {
				num := findNum(s[i])

				for j := 1; num >= j; j++ {
					if i-j < 0 {
						fmt.Println("There aren't enough arguments.")
						s = []string{}
						return s
					}
					for k := range punc {
						if s[i-j] == string(punc[k]) {
							fmt.Println("There is unavailabe character to turn.")
							s = []string{}
							return s

						}
					}
					rr := []rune(s[i-j])
					if rr[0] == '(' && rr[len(rr)-1] == ')' {
						num = num + 1
					} else {
						s[i-j] = strings.ToUpper(s[i-j])
					}
				}
			} else if string(run[1:4]) == "bin" {
				num := findNum(s[i])
				for j := 1; num >= j; j++ {
					if i-j < 0 {
						fmt.Println("There aren't enough arguments.")
						s = []string{}
						return s
					}
					for k := range punc {
						if s[i-j] == string(punc[k]) {
							fmt.Println("There is unavailabe character to turn.")
							s = []string{}
							return s

						}
					}
					rr := []rune(s[i-j])
					if rr[0] == '(' && rr[len(rr)-1] == ')' {
						num = num + 1
					} else {
						s[i-j] = strconv.Itoa(BinaryToDecimal(s[i-j]))
						if s[i-j] == "0" {
							fmt.Println("Its not a binary number.")
						}
					}
				}
			} else if string(run[1:4]) == "hex" {
				num := findNum(s[i])
				for j := 1; num >= j; j++ {
					if i-j < 0 {
						fmt.Println("There aren't enough arguments.")
						s = []string{}
						return s
					}
					for k := range punc {
						if s[i-j] == string(punc[k]) {
							fmt.Println("There is unavailabe character to turn.")
							s = []string{}
							return s

						}
					}
					rr := []rune(s[i-j])
					if rr[0] == '(' && rr[len(rr)-1] == ')' {
						num = num + 1
					} else {
						s[i-j] = HexD(s[i-j])
					}
				}
			}
		} else if s[i] == "a" && i != len(s)-1 {
			if Vowel(string(s[i+1][0])) == true {
				s[i] = "an"
			}
		} else if s[i] == "A" && i != len(s)-1 {
			if Vowel(string(s[i+1][0])) == true {
				s[i] = "An"
			}
		}
	}

	return s
}

// Фнукция Солвер нужна для форматирования текста по заданным скобкам. Внутри Солвера мы вызываем функции BinaryToDecimal и HexD
// Кроме этого Солвер решает задачу с a/an
