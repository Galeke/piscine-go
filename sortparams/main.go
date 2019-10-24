package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args 
	n := 0 
	for index := range argument {
		n = index
	}
	argument = argument [1 : n+1]

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if argument[i] > argument[j] {
				argument[i], argument[j] = argument[j], argument[i]
			}
		}
	}
	for _, argument := range argument {
		for _, letter := range argument {
			z01.PrintRune(letter)
		}
		z01.PrintRune(10)
	}
}