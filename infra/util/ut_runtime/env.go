package ut_runtime

import (
	"os"
	"strings"
)

func EnvMap() map[string]string {
	em := make(map[string]string)
	for _, e := range os.Environ() {
		v := strings.Split(e, "=")
		if len(v) > 1 {
			em[v[0]] = v[1]
		}
	}
	return em
}
