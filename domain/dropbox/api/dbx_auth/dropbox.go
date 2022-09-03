package dbx_auth

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
)

var (
	DropboxIndividual = api_auth2.OAuthAppData{
		AppKeyName:       app.ServiceDropboxIndividual,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth2.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}

	DropboxTeam = api_auth2.OAuthAppData{
		AppKeyName:       app.ServiceDropboxTeam,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth2.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}
)
