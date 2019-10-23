package piscine

func IsPrintable(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, print := range s {
		if print >= 32 && print <= 127 {
			continue
		} else {
			return false
		}
	}
	return true
}
