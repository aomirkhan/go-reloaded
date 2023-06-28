package Solve

import "strings"

func SingleQuotes(s string) string {
	kavichka := false
	s = strings.ReplaceAll(s, " ' ", "$")
	s = strings.ReplaceAll(s, " '", "$")
	s = strings.ReplaceAll(s, "' ", "$")
	s = strings.ReplaceAll(s, "'", "$")
	for strings.Contains(s, "$") {
		for i := range s {
			if i != len(s)-1 && string(s[i]) == "$" && string(s[i+1]) == "t" && i-1 > 0 && string(s[i-1]) == "n" {
				s = replaceAtIndex(s, "'", i)
			}
			if string(s[i]) == "$" && kavichka == false {

				s = replaceAtIndex(s, " '", i)
				kavichka = true

			} else if string(s[i]) == "$" && kavichka == true {
				if i == 0 {
					kavichka = false
				} else {
					s = replaceAtIndex(s, "' ", i)
					kavichka = false
				}
			}
		}
	}
	return s
}

// Ставит ковычки в правильном порядке.
