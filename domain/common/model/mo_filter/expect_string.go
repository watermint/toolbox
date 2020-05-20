package mo_filter

import "fmt"

func ExpectString(v interface{}, f func(s string) bool) bool {
	switch v0 := v.(type) {
	case string:
		return f(v0)
	case fmt.Stringer:
		return f(v0.String())
	default:
		return f(fmt.Sprintf("%v", v0))
	}
}
