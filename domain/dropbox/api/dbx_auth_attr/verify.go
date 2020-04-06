package dbx_auth_attr

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"go.uber.org/zap"
)

var (
	ErrorUnableToRetrieveResponse = errors.New("unable to retrieve json response")
	ErrorNoVerification           = errors.New("no verification")
)

// Returns description of the account
func VerifyToken(ctx api_auth.Context, ctl app_control.Control) (actx api_auth.Context, err error) {
	l := ctl.Log().With(zap.String("peerName", ctx.PeerName()), zap.String("scope", ctx.Scope()))
	ui := ctl.UI()

	switch ctx.Scope() {
	case api_auth.DropboxTokenFull, api_auth.DropboxTokenApp:
		apiCtx := dbx_context_impl.New(ctl, ctx)
		p, err := apiCtx.Rpc("users/get_current_account").Call()
		if err != nil {
			l.Debug("Unable to verify token", zap.Error(err))
			return nil, err
		}

		j, err := p.Json()
		if err != nil {
			l.Debug("Unable to retrieve JSON response", zap.Error(err))
			return nil, ErrorUnableToRetrieveResponse
		}
		desc := j.Get("name.display_name").String()
		suppl := j.Get("email").String()
		l.Debug("Token Verified", zap.String("desc", desc))

		return dbx_auth.NewContextWithAttr(ctx, desc, suppl), nil

	case api_auth.DropboxTokenBusinessInfo,
		api_auth.DropboxTokenBusinessManagement,
		api_auth.DropboxTokenBusinessFile,
		api_auth.DropboxTokenBusinessAudit:
		apiCtx := dbx_context_impl.New(ctl, ctx)
		p, err := apiCtx.Rpc("team/get_info").Call()
		if err != nil {
			l.Debug("Unable to verify token", zap.Error(err))
			return nil, err
		}
		j, err := p.Json()
		if err != nil {
			l.Debug("Unable to retrieve JSON response", zap.Error(err))
			return nil, ErrorUnableToRetrieveResponse
		}

		desc := j.Get("name").String()
		supplLic := j.Get("num_licensed_users").Int()
		suppl := ui.Text(MAttr.AttrTeamLicenses.With("Licenses", supplLic))
		l.Debug("Token Verified", zap.String("desc", desc), zap.String("suppl", suppl))

		return dbx_auth.NewContextWithAttr(ctx, desc, suppl), nil

	default:
		return nil, ErrorNoVerification
	}
}
