package piscine

func IsNumeric(str string) bool {
	if str == "" {
		return false
	}
	s := []rune(str)
	for _, index := range s {
		if index >= '0' && index <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
