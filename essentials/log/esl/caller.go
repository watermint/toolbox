package esl

import (
	"errors"
	"github.com/tidwall/gjson"
	"strings"
	"time"
)

var (
	ErrorCallerTestMessageNotFound = errors.New("test message not found in the output")
	ErrorCallerTestLineIsNotJSON   = errors.New("line is not json format")
)

// Test caller skip
func EnsureCallerSkip(l Logger, msgKey, callerKey string, fetchOut func() string) error {
	msg := "EnsureCaller:" + time.Now().String()
	l.Info(msg)

	if err := l.Sync(); err != nil {
		return err
	}

	out := fetchOut()
	lines := strings.Split(out, "\n")

	for _, line := range lines {
		if found, err := ensureCallerSkipLine(msg, msgKey, callerKey, line); err != nil {
			return err
		} else if found {
			return nil
		}
	}

	return ErrorCallerTestMessageNotFound
}

func ensureCallerSkipLine(testMsg, msgKey, callerKey string, line string) (found bool, err error) {
	expectedCallerPrefix := "es_log/caller.go:"
	if !gjson.Valid(line) {
		return false, ErrorCallerTestLineIsNotJSON
	}
	j := gjson.Parse(line)
	if j.Get(msgKey).String() != testMsg {
		return false, nil
	}
	caller := j.Get(callerKey).String()
	if strings.HasPrefix(caller, expectedCallerPrefix) {
		return true, nil
	}

	return false, errors.New("invalid caller: " + caller)
}
