package es_encode

const (
	byteDigestLength = 16
)

func ByteDigest(b []byte) map[string]interface{} {
	var d []byte
	if len(b) > byteDigestLength {
		copy(d, b[:byteDigestLength])
	} else {
		copy(d, b)
	}
	return map[string]interface{}{
		"prefix": d,
		"len":    len(b),
	}
}
