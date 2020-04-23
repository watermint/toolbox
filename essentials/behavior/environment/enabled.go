package environment

import (
	"os"
	"strconv"
	"strings"
)

// Returns true, if a feature enabled by env variable.
func IsEnabled(envName string) bool {
	for _, nv := range os.Environ() {
		pair := strings.Split(nv, "=")
		if len(pair) < 2 {
			continue
		}
		name := pair[0]
		valu := pair[1]
		if name != envName {
			continue
		}
		if b, err := strconv.ParseBool(valu); err != nil {
			return false
		} else {
			return b
		}
	}
	return false
}
