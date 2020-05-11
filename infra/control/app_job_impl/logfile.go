package app_job_impl

import (
	"compress/gzip"
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_job"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	ErrorFileIsNotALog = errors.New("the file is not a log")
)

func newLogFile(path string) (app_job.LogFile, error) {
	basename := filepath.Base(path)
	lft, found := logFileTypeFromBasename(basename)
	if !found {
		return nil, ErrorFileIsNotALog
	}

	lbn := strings.ToLower(basename)
	switch {
	case strings.HasSuffix(lbn, ".log"):
		return &logFileImpl{
			path:       path,
			fileType:   lft,
			basename:   basename,
			compressed: false,
		}, nil
	case strings.HasSuffix(lbn, ".log.gz"):
		return &logFileImpl{
			path:       path,
			fileType:   lft,
			basename:   basename,
			compressed: true,
		}, nil
	}

	return nil, ErrorFileIsNotALog
}

func logFileTypeFromBasename(basename string) (lft app_job.LogFileType, found bool) {
	name := strings.ToLower(basename)
	switch {
	case strings.HasPrefix(name, string(app_job.LogFileTypeToolbox)):
		return app_job.LogFileTypeToolbox, true
	case strings.HasPrefix(name, string(app_job.LogFileTypeCapture)):
		return app_job.LogFileTypeCapture, true
	case strings.HasPrefix(name, string(app_job.LogFileTypeSummary)):
		return app_job.LogFileTypeSummary, true
	}
	return "", false
}

type logFileImpl struct {
	path       string
	fileType   app_job.LogFileType
	basename   string
	compressed bool
}

func (z logFileImpl) Type() app_job.LogFileType {
	return z.fileType
}

func (z logFileImpl) Name() string {
	return z.basename
}

func (z logFileImpl) Path() string {
	return z.path
}

func (z logFileImpl) IsCompressed() bool {
	return z.compressed
}

func (z logFileImpl) copyToCompressed(writer io.Writer) error {
	l := esl.Default().With(esl.String("path", z.path))
	f, err := os.Open(z.path)
	if err != nil {
		l.Debug("unable to open the log file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	g, err := gzip.NewReader(f)
	if err != nil {
		l.Debug("unable to create gzip reader", esl.Error(err))
		return err
	}
	defer func() {
		_ = g.Close()
	}()

	if written, err := io.Copy(writer, g); err != nil {
		l.Debug("unable to copy", esl.Error(err))
		return err
	} else {
		l.Debug("entire log file copied", esl.Int64("written", written))
		return nil
	}
}

func (z logFileImpl) copyToUncompressed(writer io.Writer) error {
	l := esl.Default().With(esl.String("path", z.path))
	f, err := os.Open(z.path)
	if err != nil {
		l.Debug("unable to open the log file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	if written, err := io.Copy(writer, f); err != nil {
		l.Debug("unable to copy", esl.Error(err))
		return err
	} else {
		l.Debug("entire log file copied", esl.Int64("written", written))
		return nil
	}
}

func (z logFileImpl) CopyTo(writer io.Writer) error {
	if z.IsCompressed() {
		return z.copyToCompressed(writer)
	} else {
		return z.copyToUncompressed(writer)
	}
}
