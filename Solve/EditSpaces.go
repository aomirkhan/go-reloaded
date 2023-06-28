package Solve

import "strings"

func EditSpaces(s string) string {
	for strings.Contains(s, "  ") == true {
		s = strings.ReplaceAll(s, "  ", " ")
	}
	return s
}
