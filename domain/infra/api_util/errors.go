package api_util

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"regexp"
	"strings"
)

var (
	errorSummaryPostfix = regexp.MustCompile(`/\.+$`)
)

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorSummary(err error) string {
	switch re := err.(type) {
	case api_rpc.ApiError:
		es := errorSummaryPostfix.ReplaceAllString(re.ErrorSummary, "")
		return es

	default:
		return ""
	}
}

func ErrorSummaryPrefix(err error, prefix string) bool {
	return strings.HasPrefix(ErrorSummary(err), prefix)
}

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorBody(err error) json.RawMessage {
	switch re := err.(type) {
	case api_rpc.ApiError:
		return re.ErrorBody

	default:
		return nil
	}
}

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorTag(err error) string {
	switch re := err.(type) {
	case api_rpc.ApiError:
		return re.ErrorTag

	default:
		return ""
	}
}

// Returns `user_message` if an error is ApiError. Otherwise return Error().
func ErrorUserMessage(err error) string {
	switch re := err.(type) {
	case api_rpc.ApiError:
		if re.UserMessage == "" {
			return re.Error()
		}
		return re.UserMessage

	default:
		return re.Error()
	}
}
