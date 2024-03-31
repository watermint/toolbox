package gh_client

import (
	"github.com/watermint/toolbox/essentials/api/api_client"
)

type Client interface {
	api_client.Client
	api_client.Post
	api_client.Get
	api_client.Put
	api_client.Upload
	api_client.Patch
	api_client.Delete
}
