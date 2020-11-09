// +build windows

package ei_exif

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_image"
)

type exifToolImpl struct {
	l esl.Logger
}

func (z exifToolImpl) IsAvailable() bool {
	return false
}

func (z exifToolImpl) Parse(path string) (exif mo_image.Exif, err error) {
	err = ErrorExiftoolNotAvailable
	return
}
