package piscine

func StrRev(s string) string {
	var c string
	for i := len(s) - 1; i >= 0; i-- {
		c += string(s[i])

	}
	return c
}
