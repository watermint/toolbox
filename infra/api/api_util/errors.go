package api_util

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/api/api_error"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"regexp"
	"strings"
)

var (
	errorSummaryPostfix = regexp.MustCompile(`/\.+$`)
)

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorSummary(err error) string {
	switch re := err.(type) {
	case api_error.ApiError:
		es := errorSummaryPostfix.ReplaceAllString(re.ErrorSummary, "")
		es = strings.Trim(es, "/")
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
	case api_error.ApiError:
		return re.ErrorBody

	default:
		return nil
	}
}

// Returns `error_summary` if an error is ApiError. Otherwise return "".
func ErrorTag(err error) string {
	switch re := err.(type) {
	case api_error.ApiError:
		return re.ErrorTag

	default:
		return ""
	}
}

// Returns `user_message` if an error is ApiError. Otherwise return ErrorK().
func ErrorUserMessage(err error) string {
	switch re := err.(type) {
	case api_error.ApiError:
		if re.UserMessage == "" {
			return re.Error()
		}
		return re.UserMessage

	default:
		return re.Error()
	}
}

func MsgFromError(err error) app_msg.Message {
	if err == nil {
		return app_msg.M("api.error.no_error")
	}
	summary := ErrorSummary(err)
	userMessage := ErrorUserMessage(err)
	switch {
	case summary == "" && userMessage != "":
		return app_msg.M(
			"dbx.err.general_error",
			app_msg.P{"ErrorK": userMessage},
		)
	case summary == "":
		return app_msg.M(
			"dbx.err.general_error",
			app_msg.P{"ErrorK": err.Error()},
		)

	default:
		errMsgKey := "dbx.err." + summary
		return app_msg.M(errMsgKey)
	}
}
