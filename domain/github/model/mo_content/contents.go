package mo_content

import (
	"errors"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
)

var (
	ErrorUnexpectedFormat = errors.New("unexpected format")
)

type Contents interface {
	File() (c Content, found bool)
	Dir() (c []Content, found bool)
	Symlink() (c Content, found bool)
	Submodule() (c Content, found bool)
}

func NewContents(j es_json.Json) (c Contents, err error) {
	l := esl.Default()

	if entries, ok := j.Array(); ok {
		cts := make([]Content, len(entries))
		for i, entry := range entries {
			if err := entry.Model(&cts[i]); err != nil {
				return nil, err
			}
		}
		return ctsDir{entries: cts}, nil
	}

	cts := Content{}
	if err := j.Model(&cts); err != nil {
		return nil, err
	}
	switch cts.Type {
	case "file":
		return ctsFile{c: cts}, nil
	case "symlink":
		return ctsSymlink{c: cts}, nil
	case "submodule":
		return ctsSubmodule{c: cts}, nil
	default:
		l.Debug("Unknown or unexpected tag", esl.String("type", cts.Type), esl.String("json", j.RawString()))
		return nil, ErrorUnexpectedFormat
	}
}
