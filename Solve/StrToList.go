package Solve

func StrToList(s string) []string {
	var x []string
	b := 0
	ss := []rune(s)
	k := false
	for i := 0; i <= len(ss)-1; i++ {
		if ss[i] == '(' {
			k = true
		} else if ss[i] == ')' {
			k = false
		}
		if k == true {
			continue
		}
		if ss[i] == ' ' {
			x = append(x, string(ss[b:i]))
			b = i + 1
		}
		if i == len(ss)-1 && ss[i] != ' ' {
			x = append(x, string(ss[b:len(s)]))
		}
	}
	return x
}

// Из стринга делает [список стрингов] разделенный пробелами
