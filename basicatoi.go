package piscine

func BasicAtoi(s string) int {
	a := []rune(s)
	n := len(s)
	b := 0
	for i := 0; i < n; i++ {
		if a[i] < '0' || a[i] > '9' {

			return b
		} else {
			b *= 10
			b += int(a[i]) - '0'

		}
	}
	return b
}