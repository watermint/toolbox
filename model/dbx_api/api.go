package dbx_api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/app"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var (
	ReqHeaderSelectUser     = "Dropbox-API-Select-User"
	ReqHeaderSelectAdmin    = "Dropbox-API-Select-Admin"
	ReqHeaderPathRoot       = "Dropbox-API-Path-Root"
	ResHeaderRetryAfter     = "Retry-After"
	ResJsonDotTag           = "\\.tag"
	DefaultClientTimeout    = time.Duration(60) * time.Second
	DateTimeFormat          = "2006-01-02T15:04:05Z"
	ErrorBadInputParam      = 400
	ErrorBadOrExpiredToken  = 401
	ErrorAccessError        = 403
	ErrorEndpointSpecific   = 409
	ErrorNoPermission       = 422
	ErrorRateLimit          = 429
	ErrorSuccess            = 0
	ErrorTransport          = 1000
	ErrorUnknown            = 1001
	ErrorUnexpectedDataType = 1002
	ErrorOperationFailed    = 1003
	ErrorServerError        = 1500
)

// Data structure for `Dropbox-API-Path-Root: {".tag": "root", "root": "<namespace_id>"}`
type PathRootRoot struct {
	Tag         string `json:".tag"`
	NamespaceId string `json:"root"`
}

func NewPathRootRoot(nsId string) PathRootRoot {
	return PathRootRoot{
		Tag:         "root",
		NamespaceId: nsId,
	}
}

// Data structure for `Dropbox-API-Path-Root: {".tag": "home"}`
type PathRootHome struct {
	Tag string `json:".tag"`
}

func NewPathRootHome() PathRootHome {
	return PathRootHome{
		Tag: "home",
	}
}

// Data structure for `Dropbox-API-Path-Root: {".tag": "namespace_id", "namespace_id": "<namespace_id>"}`
type PathRootNamespace struct {
	Tag         string `json:".tag"`
	NamespaceId string `json:"namespace_id"`
}

func NewPathRootNamespace(nsId string) PathRootNamespace {
	return PathRootNamespace{
		Tag:         "namespace_id",
		NamespaceId: nsId,
	}
}

func ParseApiError(responseBody string) (ae ApiError) {
	ae.ErrorTag = gjson.Get(responseBody, "error."+ResJsonDotTag).String()
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

func (e ServerError) Error() string {
	return fmt.Sprintf("An error occurred on the Dropbox servers (%d). Check status.dropbox.com for announcements about Dropbox service issues.", e.StatusCode)
}

type ApiError struct {
	ErrorTag          string          `json:"error,omitempty"`
	ErrorSummary      string          `json:"error_summary,omitempty"`
	ErrorBody         json.RawMessage `json:"error,omitempty"`
	UserMessageLocale string          `json:"user_message_lang,omitempty"`
	UserMessage       string          `json:"user_message,omitempty"`
	UserMessageBody   json.RawMessage `json:"user_message,omitempty"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("Endpoint specific error[%s] %s", e.ErrorTag, e.ErrorSummary)
}

type AccessError struct {
	InvalidAccountType string          `json:"invalid_account_type,omitempty"`
	PaperAccessDenied  string          `json:"paper_access_denied,omitempty"`
	ErrorBody          json.RawMessage `json:"error,omitempty"`
}

func (a AccessError) Error() string {
	if a.InvalidAccountType != "" {
		return a.InvalidAccountType
	}
	if a.PaperAccessDenied != "" {
		return a.PaperAccessDenied
	}
	return "The user or team account doesn't have access to the endpoint or feature"
}

type ApiErrorRateLimit struct {
	RetryAfter int
}

func (e ApiErrorRateLimit) Error() string {
	return fmt.Sprintf("API Rate limit (retry after %d sec)", e.RetryAfter)
}

func RebaseTimeForAPI(t time.Time) time.Time {
	return t.UTC().Round(time.Second)
}

type Context struct {
	Token      string
	TokenType  string
	Client     *http.Client
	RetryAfter time.Time
	LastErrors []error
	ec         *app.ExecContext
}

func (z *Context) Log() *zap.Logger {
	return z.ec.Log().With(zap.String("token", z.TokenType))
}

func (z *Context) ParseModel(v interface{}, j gjson.Result) error {
	return parseModel(z.Log(), v, j)
}

func (z *Context) ParseModelJson(v interface{}, raw json.RawMessage) error {
	return parseModelJson(z.Log(), v, raw)
}

func NewContext(ec *app.ExecContext, tokenType, token string) *Context {
	return &Context{
		Token:     token,
		TokenType: tokenType,
		Client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		LastErrors: make([]error, 0),
		ec:         ec,
	}
}

func ParserError(msg string, body string, logger *zap.Logger, onError func(err error) bool) bool {
	logger.Debug(msg, zap.String("body", body))
	err := errors.New(msg)
	return onError(err)
}
