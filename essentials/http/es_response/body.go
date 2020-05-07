package es_response

import (
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

const (
	MaximumJsonSize = 16 * 1048576 // 16MiB
)

var (
	ErrorContentIsNotAJSON = errors.New("contents are not json")
	ErrorContentIsTooLarge = errors.New("contents are too large to process")
	ErrorNoContent         = errors.New("no content")
)

type Body interface {
	// Length of the read content in bytes
	ContentLength() int64

	// Body bytes. Returns empty array if the body written in to the file.
	Body() []byte

	// Body in string. Returns empty string if the body written into the file.
	BodyString() string

	// Body file. Return empty string if the body loaded on the memory.
	File() string

	// True when the body written into the file.
	IsFile() bool

	// Retrieve body as file. Returns empty string & error if an error happened during write.
	AsFile() (string, error)

	// Parse body as JSON.
	AsJson() (es_json.Json, error)

	// Parse body & returns non nil json instance.
	Json() es_json.Json
}
