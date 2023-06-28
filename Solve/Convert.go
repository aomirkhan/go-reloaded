package Solve

import (
	"fmt"
	"os"
)

func Convert(f string) string {
	b, err := os.ReadFile(f)
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

// конвертирует текст файла в 1 стринг
