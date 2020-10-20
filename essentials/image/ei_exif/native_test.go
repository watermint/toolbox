package ei_exif

import (
	"github.com/watermint/toolbox/essentials/go/es_project"
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
	"testing"
)

func TestNativeImpl_Parse(t *testing.T) {
	eni := &nativeImpl{logger: esl.Default()}

	rr, err := es_project.DetectRepositoryRoot()
	if err != nil {
		t.Error(err)
		return
	}

	exif, err := eni.Parse(filepath.Join(rr, "test/data/exif_test001.jpg"))
	if err != nil {
		t.Error(err)
	}

	if x := exif.DateTimeOriginal; x != "2012:07:15 15:27:05" {
		t.Error(x)
	}
	if x := exif.Model; x != "NIKON D800" {
		t.Error(x)
	}
}
