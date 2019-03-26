package api_rpc_impl

import "time"

const (
	RpcEndpoint             = "api.dropboxapi.com"
	SameErrorRetryCount     = 5
	SameErrorRetryWait      = time.Duration(60) * time.Second
	PrecautionRateLimitWait = time.Duration(2) * time.Second
	ErrorBadInputParam      = 400
	ErrorBadOrExpiredToken  = 401
	ErrorAccessError        = 403
	ErrorEndpointSpecific   = 409
	ErrorNoPermission       = 422
	ErrorRateLimit          = 429
)
