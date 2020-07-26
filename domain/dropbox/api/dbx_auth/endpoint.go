package dbx_auth

import "golang.org/x/oauth2"

func DropboxOAuthEndpoint() oauth2.Endpoint {
	return oauth2.Endpoint{
		AuthURL:  "https://www.dropbox.com/oauth2/authorize",
		TokenURL: "https://api.dropboxapi.com/oauth2/token",
	}
}
