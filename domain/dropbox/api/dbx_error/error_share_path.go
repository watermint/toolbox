package dbx_error

func NewSharePath(prefix string, de DropboxError) ErrorSharePath {
	return &errorSharePath{
		prefix: prefix + "/",
		de:     de,
	}
}

type errorSharePath struct {
	prefix string
	de     DropboxError
}

func (z errorSharePath) IsAlreadyShared() bool {
	return z.de.HasPrefix(z.prefix + "already_shared")
}
