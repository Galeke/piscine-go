package piscine

func StrRev(s string) string {
	c := []rune(s)
	d := 0
	for i := range c {
		d = i
	}
	for i, j := 0, d; i < j; i, j = i+1, j-1 {
		c[i], c[j] = c[j], c[i]

	}
	return string(c)
}
