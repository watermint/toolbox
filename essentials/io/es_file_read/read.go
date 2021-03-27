package es_file_read

import (
	"bufio"
	"compress/gzip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
	"strings"
)

func ReadFileLines(path string, h func(line []byte) error) error {
	l := esl.Default().With(esl.String("path", path))
	f, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	return ReadLines(f, h)
}

func ReadLines(r io.Reader, h func(line []byte) error) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		if err := h(scanner.Bytes()); err != nil {
			return err
		}
	}
	return scanner.Err()
}

func ReadFileOrArchived(path string, handler func(r io.Reader) error) error {
	l := esl.Default().With(esl.String("path", path))
	l.Debug("Open data")
	f, err := os.Open(path)
	if err != nil {
		l.Debug("Unable to open the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	var r io.Reader
	if strings.HasSuffix(path, ".gz") {
		r, err = gzip.NewReader(f)
		if err != nil {
			l.Debug("Unable to read gzipped file", esl.Error(err))
			return err
		}
	} else {
		r = f
	}
	return handler(r)
}
