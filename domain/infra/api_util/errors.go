package api_util

import (
	"encoding/json"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_ui"
	"github.com/watermint/toolbox/domain/infra/api_rpc"
	"github.com/watermint/toolbox/experimental/app_msg"
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

func MsgFromError(err error) app_msg.Message {
	if err == nil {
		return app_msg.M("api.error.no_error")
	}
	summary := ErrorSummary(err)
	userMessage := ErrorUserMessage(err)
	switch {
	case summary == "" && userMessage != "":
		return app_msg.M(
			"api.error.general_error",
			app_msg.P("Error", userMessage),
		)
	case summary == "":
		return app_msg.M(
			"api.error.general_error",
			app_msg.P("Error", err.Error()),
		)

	default:
		errMsgKey := "dbx.err." + summary
		return app_msg.M(errMsgKey)
	}
}

func UIMsgFromError(err error) app_ui.UIMessage {
	ec := app.Root()
	if err == nil {
		return ec.Msg(app.MsgNoError)
	}
	summary := ErrorSummary(err)
	if summary == "" {
		return ec.Msg("app.common.api.err.general_error").WithData(struct {
			Error string
		}{
			Error: err.Error(),
		})
	} else {
		errMsgKey := "dbx.err." + summary
		userMessage := ErrorUserMessage(err)

		if ec.MessageContainer().MsgExists(errMsgKey) {
			errDesc := app.Root().Msg(errMsgKey).T()
			return ec.Msg("app.common.api.err.api_error").WithData(struct {
				Tag   string
				Error string
			}{
				Tag:   summary,
				Error: errDesc,
			})
		}

		return ec.Msg("app.common.api.err.api_error").WithData(struct {
			Tag   string
			Error string
		}{
			Tag:   summary,
			Error: userMessage,
		})
	}
}
