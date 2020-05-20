package app_msg

const (
	KeyRaw             = "raw"
	KeyParamRaw        = "Raw"
	KeyComplex         = "complex"
	KeyComplexMessages = "Messages"
)

func IsSpecialKey(key string) bool {
	switch key {
	case KeyRaw, KeyComplex:
		return true
	}
	return false
}
