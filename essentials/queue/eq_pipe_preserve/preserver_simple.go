package eq_pipe_preserve

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/strings/es_uuid"
	"os"
	"path/filepath"
	"time"
)

var (
	ErrorSessionIsNotAvailable = errors.New("session is not available")
)

const (
	sessionFileSuffix = ".session"
	sessionInfoFile   = "info.dat"
)

func NewPreserver(l esl.Logger, basePath string) Preserver {
	l = l.With(esl.String("basePath", basePath))
	l.Debug("Create preserver")

	sessionId := es_uuid.NewV4().String()
	l = l.With(esl.String("sessionId", sessionId))
	l.Debug("SessionID generated")

	return &simplePreserver{
		l:         l,
		sessionId: sessionId,
		basePath:  basePath,
	}
}

// This implementation is not thread safe.
type simplePreserver struct {
	l           esl.Logger
	sessionId   string
	basePath    string
	filePath    string
	count       int64
	archive     *zip.Writer
	archiveFile *os.File
}

func (z *simplePreserver) logger() esl.Logger {
	return z.l.With(esl.Int64("count", z.count))
}

func (z *simplePreserver) Start() (err error) {
	l := z.logger()
	l.Debug("Start")
	z.filePath = filepath.Join(z.basePath, z.sessionId+sessionFileSuffix)
	z.archiveFile, err = os.Create(z.filePath)
	if err != nil {
		l.Debug("Unable to create a file", esl.Error(err))
		return
	}

	l.Debug("Create zip archiver")
	z.archive = zip.NewWriter(z.archiveFile)

	l.Debug("Set counter")
	z.count = 0

	return nil
}

func (z *simplePreserver) cleanUpOnError() {
	l := z.logger()
	l.Debug("Clean up")
	if z.archive == nil {
		l.Debug("Archive is not yet created, skip operation")
		return
	}

	l.Debug("Close archive")
	err := z.archive.Close()
	l.Debug("Close archive: Done", esl.Error(err))

	l.Debug("Remove archive")
	err = os.Remove(z.filePath)
	l.Debug("Remove archive: Done", esl.Error(err))

	z.archive = nil
	z.filePath = ""
	z.count = 0
}

func (z *simplePreserver) Add(d []byte) (err error) {
	l := z.logger()
	l.Debug("Add", esl.Int("size", len(d)))

	if z.archive == nil {
		l.Warn("The session file is not available")
		return ErrorSessionIsNotAvailable
	}

	index := z.count
	z.count++

	name := fmt.Sprintf("%016x", index)

	l.Debug("Create archive", esl.String("name", name))
	w, err := z.archive.CreateHeader(&zip.FileHeader{
		Name:     name,
		Modified: time.Now(),
	})
	if err != nil {
		l.Debug("Unable to create a data file in the archive", esl.Error(err))
		z.cleanUpOnError()
		return err
	}

	_, err = w.Write(d)
	if err != nil {
		l.Debug("Unable to write a data", esl.Error(err))
		z.cleanUpOnError()
		return err
	}

	l.Debug("Data successfully added")
	return nil
}

func (z *simplePreserver) Commit(info []byte) (sessionId string, err error) {
	l := z.logger()
	l.Debug("Commit")

	if z.archive == nil {
		l.Warn("The session file is not available")
		return "", ErrorSessionIsNotAvailable
	}

	l.Debug("Write info")
	f, err := z.archive.CreateHeader(&zip.FileHeader{
		Name:     sessionInfoFile,
		Modified: time.Now(),
	})
	if err != nil {
		l.Debug("Unable to create info file", esl.Error(err))
		return "", err
	}
	if _, err := f.Write(info); err != nil {
		l.Debug("Unable to write into the info file", esl.Error(err))
		return "", err
	}

	l.Debug("Flush")
	if err = z.archive.Flush(); err != nil {
		l.Debug("Unable to flush", esl.Error(err))
		z.cleanUpOnError()
		return "", err
	}

	if err = z.archive.Close(); err != nil {
		l.Debug("Unable to close", esl.Error(err))
		z.cleanUpOnError()
		return "", err
	}

	l.Debug("Commit done")
	return z.sessionId, nil
}
