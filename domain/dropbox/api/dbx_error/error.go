package dbx_error

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

type ApiErrorRateLimit struct {
	RetryAfter int
}

func (z ApiErrorRateLimit) Error() string {
	return fmt.Sprintf("API Rate limit (retry after %d sec)", z.RetryAfter)
}

func ParseApiError(responseBody string) (ae ApiError) {
	ae.ErrorTag = gjson.Get(responseBody, "error.\\.tag").String()
	ae.ErrorSummary = gjson.Get(responseBody, "error_summary").String()
	ae.UserMessageLocale = gjson.Get(responseBody, "user_message.locale").String()
	ae.UserMessage = gjson.Get(responseBody, "user_message.text").String()
	ae.ErrorBody = json.RawMessage(gjson.Get(responseBody, "error").Raw)
	ae.UserMessageBody = json.RawMessage(gjson.Get(responseBody, "user_message").Raw)

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
	ErrorTag          string          `json:"error,omitempty"`
	ErrorSummary      string          `json:"error_summary,omitempty"`
	ErrorBody         json.RawMessage `json:"error,omitempty"`
	UserMessageLocale string          `json:"user_message_lang,omitempty"`
	UserMessage       string          `json:"user_message,omitempty"`
	UserMessageBody   json.RawMessage `json:"user_message,omitempty"`
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
