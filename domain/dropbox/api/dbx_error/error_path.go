package dbx_error

func NewErrorPath(de ErrorInfo) ErrorEndpointPath {
	return &errorPathImpl{
		de: de,
	}
}

type errorPathImpl struct {
	de ErrorInfo
}

func (z errorPathImpl) IsTooManyWriteOperations() bool {
	return z.de.HasPrefix("path/too_many_write_operations")
}

func (z errorPathImpl) IsConflict() bool {
	return z.de.HasPrefix("path/conflict")
}

func (z errorPathImpl) IsNotFound() bool {
	return z.de.HasPrefix("path/not_found")
}

func (z errorPathImpl) IsMalformedPath() bool {
	return z.de.HasPrefix("path/malformed_path")
}
