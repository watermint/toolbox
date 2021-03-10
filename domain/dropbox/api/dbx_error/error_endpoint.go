package dbx_error

func NewErrorEndpoint(de ErrorInfo) ErrorEndpoint {
	return &errorEndpoint{
		de: de,
	}
}

type errorEndpoint struct {
	de ErrorInfo
}

func (z errorEndpoint) IsRateLimit() bool {
	return z.de.HasPrefix("rate_limit")
}
