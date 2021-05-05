package ut_format

import (
	"testing"
	"time"
)

func TestParseTimestamp(t *testing.T) {
	tm, ok := ParseTimestamp("2020-06-30T06:00:05Z")
	if !ok {
		t.Error(ok)
	}
	if !tm.Equal(time.Unix(1593496805, 0)) {
		t.Error(tm)
	}
}

func TestDaily(t *testing.T) {
	{
		dr, err := Daily("2019-10-01T10:11:12Z", "2019-10-05T23:38:29Z")
		if err != nil {
			t.Error("invalid error", err)
		}
		if len(dr) != 5 {
			t.Error("Invalid range")
		}
		if dr[0].Start != "2019-10-01T10:11:12Z" {
			t.Error("Invalid start")
		}
		if dr[0].End != "2019-10-02T00:00:00Z" {
			t.Error("Invalid start")
		}
		if dr[4].Start != "2019-10-05T00:00:00Z" {
			t.Error("Invalid end")
		}
		if dr[4].End != "2019-10-05T23:38:29Z" {
			t.Error("Invalid end")
		}
	}

	{
		_, err := Daily("", "2019-10-05T23:28:29Z")
		if err == nil {
			t.Error("should fail")
		}
	}

	{
		_, err := Daily("2020-10-05T23:28:29Z", "2019-10-05T23:28:29Z")
		if err == nil {
			t.Error("should fail")
		}
	}

}

func TestParseDuration(t *testing.T) {
	jitter := 1 * time.Second
	inJitter := func(tm, expect time.Time) {
		if expect.Add(-jitter).After(tm) {
			t.Error(tm, expect, "invalid time")
		}
		if expect.Add(jitter).Before(tm) {
			t.Error(tm, expect, "invalid time")
		}
	}

	if tm, valid := ParseTimestamp("-24h"); valid {
		inJitter(tm, time.Now().Add(-1*time.Hour*24))
	}
	if tm, valid := ParseTimestamp("24h"); valid {
		inJitter(tm, time.Now().Add(1*time.Hour*24))
	}
}
