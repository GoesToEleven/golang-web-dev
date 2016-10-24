package pack

import (
	"strconv"
)

type NumberData struct {
	value      float64
	isNumeric  bool
	isNegative bool
	isInteger  bool
}

func NumberEvaluator(num string) *NumberData {
	var result NumberData

	val, err := strconv.ParseFloat(num, 64)

	if err != nil {
		result.isNumeric = false
	} else {
		result.isNumeric = true
		result.value = val
	}

	if result.isNumeric {
		if val < 0 {
			result.isNegative = true
		}
		if int(val*10)%10 == 0 {
			result.isInteger = true
		}
	}

	return &result
}
