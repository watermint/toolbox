package ut_format

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"time"
)

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

type DayRange struct {
	Start string
	End   string
}

func Daily(start, end string) ([]*DayRange, error) {
	dr := make([]*DayRange, 0)
	startTime, ok := ParseTimestamp(start)
	if !ok || startTime.IsZero() {
		return nil, errors.New("start date is required")
	}
	endTime, ok := ParseTimestamp(end)
	if !ok || endTime.IsZero() {
		endTime = time.Now()
	}

	if endTime.Before(startTime) {
		return nil, errors.New("end date is before start date")
	}

	p := time.Date(startTime.Year(), startTime.Month(), startTime.Day(), 0, 0, 0, 0, startTime.Location())
	p = p.Add(24 * time.Hour)
	q := startTime

	for endTime.After(p) {
		dr = append(dr, &DayRange{
			Start: dbx_util.ToApiTimeString(q),
			End:   dbx_util.ToApiTimeString(p),
		})
		q = p
		p = p.Add(24 * time.Hour)
	}
	dr = append(dr, &DayRange{
		Start: dbx_util.ToApiTimeString(q),
		End:   dbx_util.ToApiTimeString(endTime),
	})

	return dr, nil
}
