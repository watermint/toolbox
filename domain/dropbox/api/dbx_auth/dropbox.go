package dbx_auth

import "github.com/watermint/toolbox/infra/api/api_auth"

var (
	DropboxIndividual = api_auth.OAuthAppData{
		AppKeyName:       api_auth.DropboxIndividual,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}

	DropboxTeam = api_auth.OAuthAppData{
		AppKeyName:       api_auth.DropboxTeam,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}
)
