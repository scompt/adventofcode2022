package util

func SumIntSlice(ints []int) (m int) {
	sum := 0
	for _, v := range ints {
		sum += v
	}
	return sum
}
