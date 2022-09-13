package dbx_auth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
)

var (
	DropboxIndividual = api_auth.OAuthAppData{
		AppKeyName:       app.ServiceDropboxIndividual,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}

	DropboxTeam = api_auth.OAuthAppData{
		AppKeyName:       app.ServiceDropboxTeam,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}
)
