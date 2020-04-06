package dbx_context_impl

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

func ClientHash(seeds []string) string {
	return fmt.Sprintf("%x", sha256.Sum224([]byte(strings.Join(seeds, ","))))
}
