package Solve

import (
	"math"
)

func BinaryToDecimal(num string) int {
	num1 := findNum(num)
	decimal := 0
	counter := 0.0
	remainder := 0

	for num1 != 0 {
		remainder = num1 % 10
		decimal += remainder * int(math.Pow(2.0, counter))
		num1 = num1 / 10
		counter++
	}

	return decimal
}

// Эта функция вызывается в Солвере и помогает определить бинарную цифру
