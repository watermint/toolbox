package api_auth_impl

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_context"
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
func VerifyToken(tokenType string, ctx api_context.Context) (desc, suppl string, err error) {
	switch tokenType {
	case api_auth.DropboxTokenFull, api_auth.DropboxTokenApp:
		p, err := ctx.Request("users/get_current_account").Call()
		if err != nil {
			ctx.Log().Debug("Unable to verify token", zap.Error(err))
			return "", "", err
		}

		j, err := p.Json()
		if err != nil {
			ctx.Log().Debug("Unable to retrieve JSON response", zap.Error(err))
			return "", "", errors.New("unable to retrieve json response")
		}
		desc := j.Get("name.display_name").String()
		suppl := j.Get("email").String()
		ctx.Log().Debug("Token Verified", zap.String("desc", desc))

		return desc, suppl, nil

	case api_auth.DropboxTokenBusinessInfo,
		api_auth.DropboxTokenBusinessManagement,
		api_auth.DropboxTokenBusinessFile,
		api_auth.DropboxTokenBusinessAudit:
		p, err := ctx.Request("team/get_info").Call()
		if err != nil {
			ctx.Log().Debug("Unable to verify token", zap.Error(err))
			return "", "", err
		}
		j, err := p.Json()
		if err != nil {
			ctx.Log().Debug("Unable to retrieve JSON response", zap.Error(err))
			return "", "", errors.New("unable to retrieve json response")
		}

		desc := j.Get("name").String()
		supplLic := j.Get("num_licensed_users").Int()
		suppl := fmt.Sprintf("%d License(s)", supplLic)
		ctx.Log().Debug("Token Verified", zap.String("desc", desc))

		return desc, suppl, nil

	default:
		return "", "", nil
	}
}
