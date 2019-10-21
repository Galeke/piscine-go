package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if power > 0 {
		for i := 1; i < power; i++ {
			if i < power {
				result = result * nb
			}
		}
	} else {
		result = 1
	}
	return result
}
