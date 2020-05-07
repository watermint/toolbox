package dbx_error

func NewErrorEndpoint(de DropboxError) ErrorEndpoint {
	return &errorEndpoint{
		de: de,
	}
}

type errorEndpoint struct {
	de DropboxError
}

func (z errorEndpoint) IsRateLimit() bool {
	return z.de.HasPrefix("rate_limit")
}
