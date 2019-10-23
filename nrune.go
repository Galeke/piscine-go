package piscine

func NRune(s string, n int) rune {
	Arr1 := []rune(s)
	for index, char := range Arr1 {
		if index == n-1 {
			return char
		}
	}

	return 0
}
