package piscine 

func Index(s string, toFind string) int {
	sToRune := []rune(s)
	toFindtoRune := []rune(toFind)
	length := 0
	for range toFindtoRune {
		length++
	}
	for index, str := range sToRune {
		if length > 0 && str == toFindtoRune[0] {
			return index
		}
	}
	return -1
}
