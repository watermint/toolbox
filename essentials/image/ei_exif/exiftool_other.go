// +build !windows

package ei_exif

import (
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_image"
	"os/exec"
)

const (
	exiftoolExecutable = "exiftool"
)

type exifToolImpl struct {
	l esl.Logger
}

func (z exifToolImpl) IsAvailable() bool {
	l := z.l
	c := exec.Command(exiftoolExecutable, "-ver")
	if out, err := c.Output(); err != nil {
		l.Debug("Unable to retrieve output", esl.Error(err))
		return false
	} else {
		l.Debug("Tool found", esl.ByteString("out", out))
		return true
	}
}

func (z exifToolImpl) Parse(path string) (exif mo_image.Exif, err error) {
	l := z.l.With(esl.String("path", path))
	c := exec.Command(exiftoolExecutable, "-json", path)

	var exifBytes []byte
	exifBytes, err = c.Output()
	if err != nil {
		l.Debug("Unable to retrieve exiftool output", esl.Error(err))
		return exif, ErrorExiftoolNotAvailable
	}

	var exifJson es_json.Json
	exifJson, err = es_json.Parse(exifBytes)
	if err != nil {
		l.Debug("Unable to parse JSON data", esl.Error(err))
		return
	}

	if exifJsonArray, found := exifJson.Array(); !found || len(exifJsonArray) < 1 {
		l.Debug("Data not found", esl.ByteString("json", exifJson.Raw()))
		err = errors.New("data not found")
	} else if err = exifJsonArray[0].Model(&exif); err != nil {
		l.Debug("Unable to retrieve data", esl.Error(err))
	}
	return
}
