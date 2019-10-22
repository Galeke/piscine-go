package piscine

import "github.com/01-edu/z01"

func intToDigits(n int) (digits []int) {
	for n > 0 {
		if n == 0 {
			digits = append(digits, 0)
		} else {
			digits = append(digits, n%10)
		}
		n /= 10
	}
	return
}
func sort(arr []int) []int {
	count := 0
	for range arr {
		count = count + 1
	}
	for i := count; i > 0; i-- {
		for j := 1; j < i; j++ {
			if arr[j-1] > arr[j] {
				value := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = value
			}
		}
	}
	return arr
}

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for _, c := range sort(intToDigits(n)) {
		z01.PrintRune(rune(c) + '0')
	}
}
