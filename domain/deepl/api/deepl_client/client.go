package deepl_client

import "github.com/watermint/toolbox/essentials/api/api_client"

type Client interface {
	api_client.Client
	api_client.Post
}
