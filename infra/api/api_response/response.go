package api_response

import "github.com/tidwall/gjson"

const (
	ResHeaderRetryAfter    = "Retry-After"
	ResHeaderApiResult     = "Dropbox-API-Result"
	ErrorBadInputParam     = 400
	ErrorBadOrExpiredToken = 401
	ErrorAccessError       = 403
	ErrorEndpointSpecific  = 409
	ErrorNoPermission      = 422
	ErrorRateLimit         = 429
)

type Response interface {
	// Response code. Returns -1 if a response does not contain status code.
	StatusCode() int

	// Returns body string. Returns empty & error if a response does not contain body.
	Result() (body string, err error)

	// Returns body string.
	ResultString() string

	// Content length.
	ContentLength() int64

	// True if the content stored in the file
	IsContentDownloaded() bool

	// Path to the content saved, if the content downloaded. Otherwise empty string.
	ContentFilePath() string

	// Header for key. Returns empty if no header for the key.
	Header(key string) string

	// Key value pair of headers
	Headers() map[string]string

	// Returns JSON result. Returns empty & error if a response is not a JSON document.
	Json() (res gjson.Result, err error)

	// Returns first element of the array.
	// Returns empty & error if a response is not an array of JSON
	JsonArrayFirst() (res gjson.Result, err error)

	// Parse model.
	Model(v interface{}) error

	// Parse model with given JSON path.
	ModelWithPath(v interface{}, path string) error

	// Parse model for a first element of the array of JSON.
	ModelArrayFirst(v interface{}) error
}
