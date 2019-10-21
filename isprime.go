package piscine

func IsPrime(nb int) bool {
	if nb == 0 {
		return true
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	for i := -2; i < -nb; i-- {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
