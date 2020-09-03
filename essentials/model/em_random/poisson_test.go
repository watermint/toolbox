package em_random

import (
	"math/rand"
	"testing"
	"time"
)

func TestPoisson(t *testing.T) {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)
	for i := 0; i < 10; i++ {
		if p := Poisson(r, 0); 0 < p {
			t.Error(seed, i, p)
		}
		if p := Poisson(r, 100); p < 0 {
			t.Error(seed, i, p)
		}
		if p := PoissonWithRange(r, 100, 50, 150); p < 50 || 150 < p {
			t.Error(seed, i, p)
		}
	}
}
