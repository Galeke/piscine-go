package piscine 

func AppendRange(min, max int) []int {
	var Arr1 []int 
	for i := min; i < max; i++ {
		Arr1 = append(Arr1, i)

	}
	return Arr1
}
