package Solve

import "strings"

func DoubleQ(s string) string {
	s = strings.ReplaceAll(s, " \" ", "@")
	s = strings.ReplaceAll(s, " \"", "@")
	s = strings.ReplaceAll(s, "\" ", "@")
	s = strings.ReplaceAll(s, "\"", "@")

	k := false
	for strings.Contains(s, "@") {
		for i := range s {
			if string(s[i]) == "@" && k == false {

				s = replaceAtIndex(s, " \"", i)
				k = true

			} else if string(s[i]) == "@" && k == true {
				if i == 0 {
					k = false
				} else {
					s = replaceAtIndex(s, "\" ", i)
					k = false
				}
			}
		}
	}
	return s
}

// Ставит двойные ковычки в правильном порядке.
