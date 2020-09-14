package dbx_error

func NewErrorWrite(prefix string, de DropboxError) ErrorWrite {
	return &errorWriteImpl{
		prefix: prefix,
		de:     de,
	}
}

type errorWriteImpl struct {
	// prefix without suffix /
	prefix string
	de     DropboxError
}

func (z errorWriteImpl) IsConflict() bool {
	return z.de.HasPrefix(z.prefix + "/" + "conflict")
}
