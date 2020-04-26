package dbx_util

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"regexp"
	"strings"
)

type MsgError struct {
	NoError      app_msg.Message
	ErrorGeneral app_msg.Message
}

var (
	MError = app_msg.Apply(&MsgError{}).(*MsgError)
)

var (
	errorSummaryPostfix = regexp.MustCompile(`/\.+$`)
)

// Returns `error_summary` if an error is DropboxError. Otherwise return "".
func ErrorSummary(err error) string {
	switch re := err.(type) {
	case dbx_error.DropboxError:
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

// Returns `user_message` if an error is DropboxError. Otherwise return Error().
func ErrorUserMessage(err error) string {
	switch re := err.(type) {
	case dbx_error.DropboxError:
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
		return MError.NoError
	}
	summary := ErrorSummary(err)
	userMessage := ErrorUserMessage(err)
	switch {
	case summary == "" && userMessage != "":
		return MError.ErrorGeneral.With("Error", userMessage)

	case summary == "":
		return MError.ErrorGeneral.With("Error", userMessage)

	default:
		errMsgKey := "dbx.err." + summary
		return app_msg.CreateMessage(errMsgKey)
	}
}
