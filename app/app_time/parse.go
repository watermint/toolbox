package app_time

import "time"

// Parse timestamp from command line input. This function supports multiple time format
func ParseTimestamp(ts string) (p time.Time, valid bool) {
	formats := []string{
		time.RFC822,
		time.RFC822Z,
		time.RFC3339,
		time.RFC1123,
		time.RFC1123Z,
		time.UnixDate,
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"2016.01.02 15:04:05",
		"2006-01-02",
		"2006/01/02",
		"2006.01.02",
	}

	loc := time.Now().Location()
	for _, f := range formats {
		t, err := time.ParseInLocation(f, ts, loc)
		if err == nil {
			return t, true
		}
	}

	return time.Unix(0, 0), false
}
