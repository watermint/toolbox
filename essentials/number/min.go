package number

func Min(nums ...interface{}) Number {
	unified := Unify(nums...)
	switch len(unified) {
	case 0:
		return Zero()
	case 1:
		return unified[0]
	default:
		if unified[0].IsFloat() {
			return minFloat(unified)
		} else {
			return minInt(unified)
		}
	}
}

func minInt(nums []Number) Number {
	var min = nums[0].Int64()
	for i := 1; i < len(nums); i++ {
		n := nums[i].Int64()
		if n < min {
			min = n
		}
	}
	return New(min)
}

func minFloat(nums []Number) Number {
	var min = nums[0].Float64()
	for i := 1; i < len(nums); i++ {
		n := nums[i].Float64()
		if n < min {
			min = n
		}
	}
	return New(min)
}
