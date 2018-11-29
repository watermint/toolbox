package oper_api

import "github.com/watermint/toolbox/model/dbx_api"

type DropboxApiToken interface {
	// identifier of the token. Token identifier includes type label of token,
	// and might include account_id or team_id in the future implementation.
	Tag() string

	// Api context.
	Api() *dbx_api.Context

	// Api app key.
	ApiKey() string

	// Api app secret.
	ApiSecret() string

	// New instance with Api context from authenticator
	WithApi(api *dbx_api.Context) DropboxApiToken

	// Access type label for UI message.
	AppTypeLabel() string

	// Type of Access label for UI message.
	AppAccessLabel() string
}
