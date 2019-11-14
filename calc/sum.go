package calc

func Sum(nums ...int) (s int) {
	for _, num := range nums {
		s += num
	}
	return s
}

func Mul(a int, b int) int {
	return a * b
}
