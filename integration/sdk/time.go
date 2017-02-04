package sdk

import "time"

func RebaseTimeForAPI(t time.Time) time.Time {
	return t.Round(time.Second).UTC()
}
