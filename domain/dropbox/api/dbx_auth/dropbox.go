package dbx_auth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_definitions"
)

var (
	DropboxIndividual = api_auth.OAuthAppData{
		AppKeyName:       app_definitions.AppKeyDropboxIndividual,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}

	DropboxTeam = api_auth.OAuthAppData{
		AppKeyName:       app_definitions.AppKeyDropboxTeam,
		EndpointAuthUrl:  "https://www.dropbox.com/oauth2/authorize",
		EndpointTokenUrl: "https://api.dropboxapi.com/oauth2/token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          true,
		RedirectUrl:      "",
	}
)

func IsTeamAppKey(appKey string) bool {
	return appKey == app_definitions.AppKeyDropboxTeam
}

func IsIndividualAppKey(appKey string) bool {
	return appKey == app_definitions.AppKeyDropboxIndividual
}
