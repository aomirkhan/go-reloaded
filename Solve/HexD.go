package Solve

import (
	"fmt"
	"strconv"
)

func HexD(s string) string {
	res, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		fmt.Println("There is no hex number.")
	}
	result := strconv.Itoa(int(res))
	return result
}

// Функция HexD вызывается в Солвере того чтобы перевсти шестнадцатеричную версию в десятичную версию
