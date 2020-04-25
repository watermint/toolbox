package response

const (
	Code1xxInformational = 1
	Code2xxSuccess       = 2
	Code3xxRedirection   = 3
	Code4xxClientErrors  = 4
	Code5xxServerErrors  = 5
)

type CodeCategory int

type Response interface {
	// Status code.
	Code() int

	// Status code category.
	CodeCategory() CodeCategory

	// Response headers.
	Headers() map[string]string

	// Get header value. Ignore cases.
	// Returns empty string, if no header found in the response.
	Header(header string) string

	// True on the response recognized as success.
	IsSuccess() bool

	// Response body on success.
	// Returns empty body when the response is not recognized as success.
	Success() Body

	// Alternative response body. Returns empty body on success.
	Alt() Body

	// Error on IO. Returns nil if no errors during the process.
	Error() error
}
