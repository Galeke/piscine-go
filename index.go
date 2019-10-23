package piscine

func Index(s string, toFind string) int {
	stoRune := []rune(s)
	toFindtoRune := []rune(toFind)
	kS := 0
	kF := 0
	for index := range toFindtoRune {
		index = index
		kF++
	}
	if kF == 0 {
		return 0
	}
	for index := range stoRune {
		index = index
		kS++
	}
	for index, letter := range stoRune {
		if letter == toFindtoRune[0] && kS >= kF+index-1 {
			m := 1
			for i := 1; i < kF; i++ {
				if toFindtoRune[i] == stoRune[index+i] {
					m++
				}
			}
			if m == kF {
				return index
			}
		}
	}
	return -1
}
