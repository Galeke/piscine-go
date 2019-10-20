package psicine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	i := 1
	if n < 0 {
		i = -1
		z01.PrintRune('-')
	}
	if n != 0 {
		j := (n / 10) * i
		if f != 0 {
			z01.PrintRune(j)
		}
		k := (n % 10 * t) + '0'
		z01.PrintRune(rune(k))
	} else {
		z01.PrintRune('0')
	}
}
