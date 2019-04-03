package api_util

import "time"

const (
	DateTimeFormat = "2006-01-02T15:04:05Z"
)

func RebaseTime(t time.Time) time.Time {
	return t.UTC().Round(time.Second)
}

func RebaseAsString(t time.Time) string {
	return RebaseTime(t).Format(DateTimeFormat)
}
