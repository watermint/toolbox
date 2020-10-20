package ei_exif

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_image"
)

type Parser interface {
	Parse(path string) (exif mo_image.Exif, err error)
}

// Auto select parser implementation.
func Auto(logger esl.Logger) Parser {
	eti := &exifToolImpl{l: logger}
	if eti.IsAvailable() {
		return eti
	}

	return &nativeImpl{
		logger: logger,
	}
}
