package piscine

func AlphNum(t rune) bool {
	if t >= 'A' && t <= 'Z' || t >= 'a' && t <= 'z' || t >= '0' && t <= '9' {
		return true
	}
	return false
}

func Capitalize(s string) string {
	Arr1 := []rune(s)
	for index, letter := range Arr1 {
		if AlphNum(letter) {
			if index == 0 || AlphNum(Arr1[index-1]) == false {
				if letter >= 'a' && letter <= 'z' {
					Arr1[index] = letter - 32
				}
			} else {
				if letter >= 'A' && letter <= 'Z' {
					Arr1[index] = letter + 32
				}
			}
		}
	}
	return string(Arr1)
}
