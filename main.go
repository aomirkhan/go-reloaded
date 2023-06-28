package main

import (
	"fmt"
	"log"
	"os"
	"piscine/Solve"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Wrong Number of Arguments!")
	} else if string(os.Args[1][len(os.Args[1])-4:]) != ".txt" || string(os.Args[2][len(os.Args[2])-4:]) != ".txt" {
		fmt.Println("Wrong Type of File!")
	} else if len(os.Args) == 3 {
		args := os.Args
		s := Solve.Convert(args[1])

		x := Solve.CentralFunc(s)

		f, err := os.Create(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()

		_, err2 := f.WriteString(x)

		if err2 != nil {
			log.Fatal(err2)
		}
	}
}

// Выбираем 2 аргумента и вызываем центральную функцию CentralFunc
// В конечном результате мы получаем файл, куда записывается обработанный текст
