package piscine

func ToUpper(s string) string {
	Arr1 := []rune(s)
	for index, letter := range Arr1 {
		if letter >= 'a' && letter <= 'z' {
			Arr1[index] = Arr1[index] - 32 //shift down in ASCII table is so stoopid jesus
		}
	}
	caps:= string(Arr1)
	return caps
}