package es_file_read

import (
	"bufio"
	"bytes"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
	"os"
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
	l := esl.Default()
	br := bufio.NewReader(r)

	prefix := &bytes.Buffer{}
	for {
		line, isPrefix, err := br.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			l.Debug("Error on read", esl.Error(err))
			return err
		}
		if isPrefix {
			_, err := prefix.Write(line)
			if err != nil {
				l.Debug("Unable to append prefix", esl.Error(err))

				// reset prefix and continue
				prefix.Reset()
				continue
			}
			continue
		}

		if prefix.Len() < 1 {
			if err := h(line); err != nil {
				l.Debug("Failed process line", esl.Error(err))
			}
		} else {
			_, err := prefix.Write(line)
			if err != nil {
				l.Debug("Unable to append prefix", esl.Error(err))

				// reset prefix and continue
				prefix.Reset()
				continue
			}
			if err := h(prefix.Bytes()); err != nil {
				l.Debug("Failed process line", esl.Error(err))
			}
			prefix.Reset()
		}
	}
}
