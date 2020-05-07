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

// Deprecated:
// Returns `error_summary` if an error is DropboxError. Otherwise return "".
func ErrorSummary(err error) string {
	ers := dbx_error.NewErrors(err)
	es := errorSummaryPostfix.ReplaceAllString(ers.Summary(), "")
	es = strings.Trim(es, "/")
	return es
}

// Deprecated:
func ErrorSummaryPrefix(err error, prefix string) bool {
	ers := dbx_error.NewErrors(err)
	return strings.HasPrefix(ers.Summary(), prefix)
}
