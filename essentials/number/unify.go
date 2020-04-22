package number

// Convert all numbers in same precision numbers.
// Invalid numbers will be filtered out.
func Unify(nums ...interface{}) []Number {
	useFloat := false
	nn := make([]Number, 0)
	for _, nr := range nums {
		n := New(nr)
		if !n.IsValid() {
			continue
		}
		nn = append(nn, n)
		if n.IsFloat() {
			useFloat = true
		}
	}
	fl := make([]Number, len(nn))
	for i, n := range nn {
		if useFloat {
			fl[i] = floatImpl{v: n.Float64()}
		} else {
			fl[i] = intImpl{v: n.Int64()}
		}
	}
	return fl
}
