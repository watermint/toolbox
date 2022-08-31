package api_service

import (
	"github.com/watermint/toolbox/essentials/http/es_response"
	"github.com/watermint/toolbox/infra/api/api_auth"
)

type Service interface {
	// OAuth app information
	App() api_auth.OAuthAppLegacy

	// Returns custom retry & rate limit error.
	TransportHandler(res es_response.Response) error
}
