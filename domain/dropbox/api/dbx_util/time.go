package dbx_util

import (
	"errors"
	"time"
)

const (
	DateTimeFormat = "2006-01-02T15:04:05Z"
)

func RebaseTime(t time.Time) time.Time {
	return t.UTC().Round(1000 * time.Millisecond)
}

func ToApiTimeString(t time.Time) string {
	return RebaseTime(t).Format(DateTimeFormat)
}

func Parse(iso8601 string) (t time.Time, err error) {
	if iso8601 == "" {
		return time.Unix(0, 0), errors.New("empty")
	}
	t, err = time.Parse(DateTimeFormat, iso8601)
	return
}
