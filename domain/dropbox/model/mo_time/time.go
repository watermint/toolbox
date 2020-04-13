package mo_time

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_essential"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"time"
)

type Time interface {
	// Returns Time in ISO8601 (yyyy-MM-ddThh:mm:ssZ) format as UTC.
	// Returns an empty string if an impl. has TimeOptional, and the instance marked as unset.
	Iso8601() string

	// Same as Iso8601
	Value() string

	// Returns time instance
	Time() time.Time

	// True when the time is zero
	IsZero() bool
}

type TimeOptional interface {
	Time
	mo_essential.Optional
}

var (
	ErrorInvalidTimeFormat = errors.New("invalid time format")
)

func Zero() (tm Time) {
	return &TimeImpl{time: time.Time{}, isSet: true}
}

func New(t time.Time) Time {
	return &TimeImpl{time: t, isSet: true}
}

func NewOptional(t time.Time) TimeOptional {
	return &TimeImpl{time: t, isSet: true}
}

type TimeImpl struct {
	time  time.Time
	isSet bool
}

func (z *TimeImpl) Unset() {
	z.isSet = false
}

func (z *TimeImpl) Ok() bool {
	return !z.IsZero() && z.isSet
}

func (z *TimeImpl) IsZero() bool {
	return z.time.IsZero()
}

func (z *TimeImpl) Time() time.Time {
	return z.time
}

func (z *TimeImpl) Iso8601() string {
	if z.Ok() {
		return dbx_util.RebaseAsString(z.time)
	} else {
		return ""
	}
}

func (z *TimeImpl) Value() string {
	if z.Ok() {
		return z.Iso8601()
	} else {
		return ""
	}
}

func (z *TimeImpl) UpdateTime(dateTime string) error {
	ts, valid := ut_time.ParseTimestamp(dateTime)
	if !valid {
		return ErrorInvalidTimeFormat
	}
	z.time = ts
	z.isSet = true
	return nil
}
