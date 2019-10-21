
package piscine

func RecursivePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power > 0 {
		return RecursivePower(nb, power-1)*nb
	}
	return 0
}