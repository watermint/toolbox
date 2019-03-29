package api_rpc_impl

import "time"

const (
	RpcEndpoint = "api.dropboxapi.com"

	// Threshold of abort retry on error.
	// Number of same errors happened in last ten errors.
	// SameErrorRetryCount must less then 10.
	SameErrorRetryCount     = 8
	SameErrorRetryWait      = time.Duration(30) * time.Second
	PrecautionRateLimitWait = time.Duration(2) * time.Second
	ErrorBadInputParam      = 400
	ErrorBadOrExpiredToken  = 401
	ErrorAccessError        = 403
	ErrorEndpointSpecific   = 409
	ErrorNoPermission       = 422
	ErrorRateLimit          = 429
)
