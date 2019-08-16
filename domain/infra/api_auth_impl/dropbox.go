package api_auth_impl

import (
	"errors"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
)

func DropboxOAuthEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}

// Returns description of the account
func VerifyToken(tokenType string, ctx api_context.Context) (desc string, err error) {
	switch tokenType {
	case api_auth.DropboxTokenFull, api_auth.DropboxTokenApp:
		p, err := ctx.Request("users/get_current_account").Call()
		if err != nil {
			ctx.Log().Debug("Unable to verify token", zap.Error(err))
			return "", err
		}

		j, err := p.Json()
		if err != nil {
			ctx.Log().Debug("Unable to retrieve JSON response", zap.Error(err))
			return "", errors.New("unable to retrieve json response")
		}
		desc := j.Get("name.display_name").String()
		ctx.Log().Debug("Token Verified", zap.String("desc", desc))

		return desc, nil

	case api_auth.DropboxTokenBusinessInfo,
		api_auth.DropboxTokenBusinessManagement,
		api_auth.DropboxTokenBusinessFile,
		api_auth.DropboxTokenBusinessAudit:
		p, err := ctx.Request("team/get_info").Call()
		if err != nil {
			ctx.Log().Debug("Unable to verify token", zap.Error(err))
			return "", err
		}
		j, err := p.Json()
		if err != nil {
			ctx.Log().Debug("Unable to retrieve JSON response", zap.Error(err))
			return "", errors.New("unable to retrieve json response")
		}

		desc := j.Get("name").String()
		ctx.Log().Debug("Token Verified", zap.String("desc", desc))

		return desc, nil

	default:
		return "", nil
	}
}
