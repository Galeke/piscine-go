package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var t []int
	if n == 0 {
		t = append(t, 0)
	}
	for i := 0; n > 0; i-- {
		t = append(t, n%10)
		n /= 10
	}

	for i := range t {
		z01.PrintRune(rune(t[i]) + '0')
	}
}
