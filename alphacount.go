package piscine

func AlphaCount(str string) int {
	a := 0
	for _, c := range str {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			a++

		}
	}
	return a
}
