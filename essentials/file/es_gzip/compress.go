package es_gzip

import (
	"compress/gzip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
)

const (
	SuffixCompress = ".gz"
)

func Compress(path string) (err error) {
	cp := path + SuffixCompress
	l := esl.Default().With(esl.String("path", cp))
	l.Debug("Compress")
	s, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open source", esl.Error(err))
		return err
	}
	c, err := os.Create(cp)
	if err != nil {
		l.Debug("Unable to create compressed", esl.Error(err))
		_ = s.Close()
		return err
	}

	zc := gzip.NewWriter(c)

	if _, err := io.Copy(zc, s); err != nil {
		l.Debug("Unable to copy", esl.Error(err))
		_ = c.Close()
		if err = os.Remove(cp); err != nil {
			l.Debug("unable to remove", esl.Error(err))
		}
		return err
	}
	if err := zc.Flush(); err != nil {
		l.Debug("Unable to flush", esl.Error(err))
		return err
	}

	// gzip writer's Close flushes & writes footer. Should catch an error.
	if err := zc.Close(); err != nil {
		l.Debug("Unable to close", esl.Error(err))
		return err
	}
	_ = s.Close()
	_ = c.Close()
	if err = os.Remove(path); err != nil {
		l.Debug("unable to remove", esl.Error(err))
	}
	return nil
}
