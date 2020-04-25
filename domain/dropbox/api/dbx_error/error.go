package dbx_error

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/essentials/http/response"
	"strings"
)

var (
	ErrorPathNotFound = errors.New("path not found")
)

type ApiErrorRateLimit struct {
	RetryAfter int
}

func (z ApiErrorRateLimit) Error() string {
	return fmt.Sprintf("API Rate limit (retry after %d sec)", z.RetryAfter)
}

func IsApiError(res response.Response) error {
	if res.CodeCategory() == response.Code2xxSuccess {
		return nil
	}
	ae := &ApiError{}
	if _, err := res.Success().Json().Model(ae); err != nil {
		return nil
	}
	switch {
	case strings.HasPrefix(ae.ErrorSummary, "path/not_found"):
		return ErrorPathNotFound
	}
	return ae
}

// Deprecated: use IsApiError
func ParseApiError(responseBody string) (ae ApiError) {
	ae.ErrorTag = gjson.Get(responseBody, "error.\\.tag").String()
	ae.ErrorSummary = gjson.Get(responseBody, "error_summary").String()
	ae.UserMessageLocale = gjson.Get(responseBody, "user_message.locale").String()
	ae.UserMessage = gjson.Get(responseBody, "user_message.text").String()

	return
}

func ParseAccessError(responseBody string) (ae AccessError) {
	ae.PaperAccessDenied = gjson.Get(responseBody, "invalid_account_type.\\.tag").String()
	ae.InvalidAccountType = gjson.Get(responseBody, "paper_access_denied.\\.tag").String()
	ae.ErrorBody = json.RawMessage(responseBody)

	return
}

type ServerError struct {
	StatusCode int
}

func (z ServerError) Error() string {
	return fmt.Sprintf("An error occurred on the Dropbox servers (%d). Check status.dropbox.com for announcements about Dropbox service issues.", z.StatusCode)
}

type ApiError struct {
	ErrorTag          string `path:"error.\\.tag" json:"error,omitempty"`
	ErrorSummary      string `path:"error_summary" json:"error_summary,omitempty"`
	UserMessageLocale string `json:"user_message_lang,omitempty"`
	UserMessage       string `json:"user_message,omitempty"`
}

func (z ApiError) Error() string {
	return fmt.Sprintf("Endpoint specific error[%s] %s", z.ErrorTag, z.ErrorSummary)
}

type AccessError struct {
	InvalidAccountType string          `json:"invalid_account_type,omitempty"`
	PaperAccessDenied  string          `json:"paper_access_denied,omitempty"`
	ErrorBody          json.RawMessage `json:"error,omitempty"`
}

func (z AccessError) Error() string {
	if z.InvalidAccountType != "" {
		return z.InvalidAccountType
	}
	if z.PaperAccessDenied != "" {
		return z.PaperAccessDenied
	}
	return "The user or team account doesn't have access to the endpoint or feature"
}
