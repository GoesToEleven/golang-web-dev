package pack

func QuickSort(nums ...float64) *[]float64 {
	partition := func(arr *[]float64, low, high int) int {
		a := *arr
		pivot := a[high]
		i := low
		for j := low; j < high; j++ {
			if a[j] <= pivot {
				a[i], a[j] = a[j], a[i]
				i++
			}
		}
		a[i], a[high] = a[high], a[i]
		return i
	}
	var sort func(*[]float64, int, int)
	sort = func(a *[]float64, low, high int) {
		if low < high {
			p := partition(a, low, high)
			sort(a, low, p-1)
			sort(a, p+1, high)
		}
	}
	a := nums[:]

	sort(&a, 0, len(nums)-1)

	return &a
}
