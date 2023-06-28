package Solve

func RemoveCaps(arr []string) []string {
	for i := 0; len(arr) > i; i++ {
		r := arr[i]
		if r[0] == '(' && r[len(r)-1] == ')' {
			arr = Re(arr, i)
			i = 0
		}
	}
	return arr
}

// RemoveCaps убирает все элементы слайса, где есть скобки
