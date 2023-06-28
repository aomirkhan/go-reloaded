package Solve

import "strings"

func replaceAtIndex(input string, replacement string, index int) string {
	return strings.Join([]string{input[:index], replacement, input[index+1:]}, "")
}

// По данному в аргументах индексу меняет нужный мне символ в string
