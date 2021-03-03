package dbx_error

func NewErrorWrite(prefix string, de ErrorInfo) ErrorWrite {
	return &errorWriteImpl{
		prefix: prefix,
		de:     de,
	}
}

type errorWriteImpl struct {
	// prefix without suffix /
	prefix string
	de     ErrorInfo
}

func (z errorWriteImpl) IsConflict() bool {
	return z.de.HasPrefix(z.prefix + "/" + "conflict")
}
