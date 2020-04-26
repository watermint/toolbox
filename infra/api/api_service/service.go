package api_service

import (
	"github.com/watermint/toolbox/essentials/http/response"
	"github.com/watermint/toolbox/infra/api/api_auth"
)

type Service interface {
	// OAuth app information
	App() api_auth.App

	// Returns custom retry & rate limit error.
	TransportHandler(res response.Response) error
}
