package ei_exif

import (
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_image"
	"os"
)

type nativeImpl struct {
	logger esl.Logger
}

func (z nativeImpl) Parse(path string) (exifData mo_image.Exif, err error) {
	l := z.logger

	var f *os.File
	f, err = os.Open(path)
	if err != nil {
		return
	}
	defer func() {
		_ = f.Close()
	}()

	exif.RegisterParsers(mknote.All...)
	data, err := exif.Decode(f)
	if err != nil {
		l.Debug("Unable to decode EXIF data", esl.Error(err))
		return
	}

	dataJsonBytes, err := data.MarshalJSON()
	if err != nil {
		l.Debug("Unable to marshal exif data to JSON", esl.Error(err))
		return
	}

	dataJson, err := es_json.Parse(dataJsonBytes)
	if err != nil {
		l.Debug("Unable to convert to JSON", esl.Error(err))
		return
	}

	err = dataJson.Model(&exifData)
	return
}
