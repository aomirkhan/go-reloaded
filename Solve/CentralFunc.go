package Solve

import (
	"strings"
)

func CentralFunc(s string) string {
	// не проверил
	s = EditSpaces(s)
	if string(s[0]) == " " {
		s = s[1:]
	}
	s = strings.ReplaceAll(s, "\"", "\" ")

	s = strings.ReplaceAll(s, "(", " (")
	s = strings.ReplaceAll(s, ")", ") ")
	s = strings.ReplaceAll(s, ".", ". ")
	s = strings.ReplaceAll(s, ",", ", ")
	s = strings.ReplaceAll(s, "!", "! ")
	s = strings.ReplaceAll(s, "?", "? ")
	s = strings.ReplaceAll(s, ":", ": ")
	s = strings.ReplaceAll(s, ";", "; ")
	s = strings.ReplaceAll(s, "  ", " ")

	s = EditSpaces(s)

	a := StrToList(s)

	a = RemoveCaps(Solver(a))

	s = strings.Join(a, " ")

	s = strings.ReplaceAll(s, " .", ".")
	s = strings.ReplaceAll(s, " ,", ",")
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " ?", "?")
	s = strings.ReplaceAll(s, " :", ":")
	s = strings.ReplaceAll(s, " ;", ";")
	s = SingleQuotes(s)
	if len(s) > 2 {
		if string(s[len(s)-1]) == "$" {
			s = replaceAtIndex(s, "' ", len(s)-1)
		}

		s = DoubleQ(s)
		if string(s[len(s)-1]) == "@" {
			s = replaceAtIndex(s, "\" ", len(s)-1)
		}

		s = EditSpaces(s)
		if string(s[0]) == " " {
			s = s[1:]
		}
		// s = Concat(s)

	}
	s = strings.ReplaceAll(s, " !", "!")
	s = strings.ReplaceAll(s, " ?", "?")
	s = strings.ReplaceAll(s, " .", ".")
	s = strings.ReplaceAll(s, " ,", ",")
	return s
}

// Эта функция вызывает все основные функции(солвер, СтрТоЛист, РемувКапс, СинглКуотс) для того чтобы отформатировать текст
