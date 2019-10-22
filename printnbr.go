package piscine

import "github.com/01-edu/z01"

func WhatNbr(x int) {
	w := '0'
	if x == 0 {
		z01.PrintRune(w)
		return
	}
	for j := 1; j <= x%10; j++ {
		w++
	}
	for j := -1; j >= x%10; j-- {
		w++
	}
	if x/10 != 0 {
		WhatNbr(x / 10)
	}
	z01.PrintRune(w)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	WhatNbr(n)

}
