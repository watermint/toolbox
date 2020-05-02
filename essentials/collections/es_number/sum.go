package es_number

func Sum(nums ...interface{}) Number {
	unified := Unify(nums...)
	if len(unified) < 1 {
		return Zero()
	}
	if unified[0].IsFloat() {
		return sumFloat(unified)
	} else {
		return sumInt(unified)
	}
}

func sumFloat(nums []Number) Number {
	var sum float64
	for _, n := range nums {
		sum += n.Float64()
	}
	return New(sum)
}

func sumInt(nums []Number) Number {
	var sum int64
	for _, n := range nums {
		sum += n.Int64()
	}
	return New(sum)
}
