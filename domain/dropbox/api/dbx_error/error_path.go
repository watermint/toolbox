package dbx_error

func NewErrorPath(de DropboxError) ErrorEndpointPath {
	return &errorPathImpl{
		de: de,
	}
}

type errorPathImpl struct {
	de DropboxError
}

func (z errorPathImpl) IsNotFound() bool {
	return z.de.HasPrefix("path/not_found")
}

func (z errorPathImpl) IsMalformedPath() bool {
	return z.de.HasPrefix("path/malformed_path")
}
