package number

func Max(nums ...interface{}) Number {
	unified := Unify(nums...)
	switch len(unified) {
	case 0:
		return Zero()
	case 1:
		return unified[0]
	default:
		if unified[0].IsFloat() {
			return maxFloat(unified)
		} else {
			return maxInt(unified)
		}
	}
}

func maxInt(nums []Number) Number {
	var max = nums[0].Int64()
	for i := 1; i < len(nums); i++ {
		n := nums[i].Int64()
		if max < n {
			max = n
		}
	}
	return New(max)
}

func maxFloat(nums []Number) Number {
	var max = nums[0].Float64()
	for i := 1; i < len(nums); i++ {
		n := nums[i].Float64()
		if max < n {
			max = n
		}
	}
	return New(max)
}
