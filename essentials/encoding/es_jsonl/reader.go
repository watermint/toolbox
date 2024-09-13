package es_jsonl

import (
	"bufio"
	"encoding/json"
	"io"
	"strings"
)

func ReadEachJson(r io.Reader, f func(line []byte) error) error {
	s := bufio.NewScanner(r)
	lines := make([]string, 0)
	isEachLine := false
	isFirstLine := true
	for s.Scan() {
		line := s.Text()
		if isFirstLine {
			var obj interface{}
			err := json.Unmarshal([]byte(line), &obj)
			if err == nil {
				isEachLine = true
				if fErr := f([]byte(line)); fErr != nil {
					return fErr
				}
			} else if _, ok := err.(*json.SyntaxError); ok {
				isEachLine = false
				lines = append(lines, line)
			} else {
				return err
			}
		} else {
			if isEachLine {
				if fErr := f([]byte(line)); fErr != nil {
					return fErr
				}
			} else {
				lines = append(lines, line)
			}
		}
		isFirstLine = false
	}
	if !isEachLine && len(lines) > 0 {
		entire := strings.Join(lines, "\n")
		if fErr := f([]byte(entire)); fErr != nil {
			return fErr
		}
	}

	return nil
}
