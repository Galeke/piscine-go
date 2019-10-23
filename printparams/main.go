package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for index, argument := range argument {
		if index != 0 {
			for _, c := range argument {
				z01.PrintRune(c)
			}
			z01.PrintRune(10)

		}
	}
}
