package api_util

import (
	"github.com/watermint/toolbox/model/dbx_api"
	"regexp"
)

var (
	errorSummaryPostfix = regexp.MustCompile(`/\.+$`)
)

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorSummary(err error) string {
	switch re := err.(type) {
	case dbx_api.ApiError:
		es := errorSummaryPostfix.ReplaceAllString(re.ErrorSummary, "")
		return es

	default:
		return ""
	}
}

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorTag(err error) string {
	switch re := err.(type) {
	case dbx_api.ApiError:
		return re.ErrorTag

	default:
		return ""
	}
}

// Returns `user_message` if an error is ApiError. Otherwise return Error().
func ErrorUserMessage(err error) string {
	switch re := err.(type) {
	case dbx_api.ApiError:
		if re.UserMessage == "" {
			return re.Error()
		}
		return re.UserMessage

	default:
		return re.Error()
	}
}
