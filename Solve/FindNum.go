package Solve

func findNum(s string) int {
	result := 0
	r := []rune(s)
	for i := range r {
		if r[i] >= 48 && r[i] <= 57 {
			result = result*10 + int(r[i]) - 48
		}
	}
	return result
}

// Эта функция вызывается в Солвере для того чтобы определить цифру в скобках и перевсти ее с стринга в инт
