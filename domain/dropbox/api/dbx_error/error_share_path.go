package dbx_error

func NewSharePath(prefix string, de ErrorInfo) ErrorSharePath {
	return &errorSharePath{
		prefix: prefix + "/",
		de:     de,
	}
}

type errorSharePath struct {
	prefix string
	de     ErrorInfo
}

func (z errorSharePath) IsAlreadyShared() bool {
	return z.de.HasPrefix(z.prefix + "already_shared")
}
