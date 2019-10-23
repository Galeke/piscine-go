package main 

import ( 
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	i := 0
	for c := range argument {
		i = c
	}
	for j := i; j >= 1; j-- {
		for _, c := range argument[j] {
			z01.PrintRune(c)
		}
		z01.PrintRune(10)
	}
}
