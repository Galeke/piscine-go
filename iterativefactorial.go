package piscine

func IterativeFactorial(nb int) int {
	result := 1
	if nb >= 0 {
		for i := nb; i >= 1; i-- {
			result = i * result
		}
	} else {
		result = 0
	}
	return result
}
