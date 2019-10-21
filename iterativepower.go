package piscine

func IterativePower(nb int, power int) int {
	result := nb
	if power > 0 {
		for i := 1; i < power; i++ {
			result = nb * result
		}
	} else {
		result = 0
	}
	return result
}
