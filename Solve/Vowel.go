package Solve

func Vowel(s string) bool {
	x := []string{"a", "e", "i", "o", "u", "h"}
	for i := range x {
		if x[i] == string(s[0]) {
			return true
		}
	}
	return false
}

// Функция Воуэл вызывается в солвере и решает проблему сопределением a/an
