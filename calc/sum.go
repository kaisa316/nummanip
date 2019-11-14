package calc

func Sum(nums ...int) (s int) {
	for _, num := range nums {
		s += num
	}
	return s
}
