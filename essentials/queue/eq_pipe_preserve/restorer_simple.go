package eq_pipe_preserve

import (
	"archive/zip"
	"io"
	"path/filepath"

	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewRestorer(l esl.Logger, basePath, sessionId string) Restorer {
	return &simpleRestorer{
		l:         l,
		basePath:  basePath,
		sessionId: sessionId,
	}
}

type simpleRestorer struct {
	l         esl.Logger
	basePath  string
	sessionId string
}

func (z *simpleRestorer) Restore(infoLoader func(info []byte) error, loader func(d []byte) error) (err error) {
	l := z.l.With(esl.String("basePath", z.basePath), esl.String("sessionId", z.sessionId))

	filePath := filepath.Join(z.basePath, z.sessionId+sessionFileSuffix)

	l.Debug("Load session file", esl.String("filePath", filePath))
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		l.Debug("Unable to load the session file", esl.Error(err))
		return err
	}

	defer func() {
		l.Debug("Close the session file")
		_ = archive.Close()
	}()

	for i, entry := range archive.File {
		ll := l.With(esl.Int("index", i), esl.Any("entry", entry))

		ll.Debug("Open data file")
		f, err := entry.Open()
		if err != nil {
			ll.Debug("Unable to load the entry", esl.Error(err))
			return err
		}

		ll.Debug("Read data file")
		d, err := io.ReadAll(f)
		_ = f.Close()
		if err != nil {
			ll.Debug("Unable to read the data file", esl.Error(err))
			return err
		}

		if entry.Name == sessionInfoFile {
			ll.Debug("Info file found")
			if err := infoLoader(d); err != nil {
				ll.Debug("The info loader returned an error, abort", esl.Error(err))
				return err
			}
		} else {
			ll.Debug("Passed to the loader")
			if err := loader(d); err != nil {
				ll.Debug("The loader returned an error, abort", esl.Error(err))
				return err
			}
		}
	}

	l.Debug("Session data load finished")

	return
}
