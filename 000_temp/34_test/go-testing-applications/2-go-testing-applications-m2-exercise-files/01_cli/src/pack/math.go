package pack

func Add(nums ...int) int {
	var result int
	for _, i := range nums {
		result += i
	}

	return result
}

func Subtract(initial int, nums ...int) int {
	for _, i := range nums {
		initial -= i
	}

	return initial
}
