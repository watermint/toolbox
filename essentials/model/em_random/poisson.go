package em_random

import "math/rand"

// Poisson distribution
// https://blog.monochromegane.com/blog/2019/10/11/random_number_gen_using_go/
func Poisson(r *rand.Rand, lambda float64) float64 {
	p := 0.0
	for i := 0; ; i++ {
		p += r.ExpFloat64() / lambda
		if p > 1.0 {
			return float64(i)
		}
	}
}

func PoissonWithRange(r *rand.Rand, lambda, min, max float64) float64 {
	m := max - min
	p := Poisson(r, lambda)
	if m < p {
		return m
	}
	return p + min
}
