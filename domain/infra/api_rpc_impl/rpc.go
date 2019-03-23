package api_rpc_impl

import "time"

const (
	RpcEndpoint            = "api.dropboxapi.com"
	SameErrorRetryCount    = 5
	SameErrorRetryWait     = time.Duration(60) * time.Second
	ReqHeaderSelectUser    = "Dropbox-API-Select-User"
	ReqHeaderSelectAdmin   = "Dropbox-API-Select-Admin"
	ReqHeaderPathRoot      = "Dropbox-API-Path-Root"
	ResHeaderRetryAfter    = "Retry-After"
	ErrorBadInputParam     = 400
	ErrorBadOrExpiredToken = 401
	ErrorAccessError       = 403
	ErrorEndpointSpecific  = 409
	ErrorNoPermission      = 422
	ErrorRateLimit         = 429
	ErrorSuccess           = 0
	ErrorTransport         = 1000
	ErrorUnknown           = 1001
)
