package pack

func Add(nums ...int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
}
