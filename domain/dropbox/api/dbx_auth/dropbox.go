package dbx_auth

import (
	"errors"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"golang.org/x/oauth2"
)

var (
	ErrorNoAuthDefined = errors.New("no auth defined")
	ErrorUserCancelled = errors.New("user cancelled")
)

type MsgCcAuth struct {
	FailedOrCancelled app_msg.Message
	GeneratedToken1   app_msg.Message
	GeneratedToken2   app_msg.Message
	OauthSeq1         app_msg.Message
	OauthSeq2         app_msg.Message
}

var (
	MCcAuth = app_msg.Apply(&MsgCcAuth{}).(*MsgCcAuth)
)

func DropboxOAuthEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}
