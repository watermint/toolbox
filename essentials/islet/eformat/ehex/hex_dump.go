package ehex

const (
	hexStringLower = "0123456789abcdef"
	hexStringUpper = "0123456789ABCDEF"
)

// ToHexString returns lower case hex string
func ToHexString(b []byte) string {
	size := len(b)
	s := make([]byte, size*2)
	for i := 0; i < size; i++ {
		j := i * 2
		s[j] = hexStringLower[int(b[i]>>4)]
		s[j+1] = hexStringLower[int(b[i]&0x0f)]
	}
	return string(s)
}

// ToUpperHexString returns upper case hex string
func ToUpperHexString(b []byte) string {
	size := len(b)
	s := make([]byte, size*2)
	for i := 0; i < size; i++ {
		j := i * 2
		s[j] = hexStringUpper[int(b[i]>>4)]
		s[j+1] = hexStringUpper[int(b[i]&0x0f)]
	}
	return string(s)
}
