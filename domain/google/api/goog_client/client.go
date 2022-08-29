package goog_client

import "github.com/watermint/toolbox/infra/api/api_client"

type Client interface {
	api_client.Client
	api_client.Get
	api_client.Post
	api_client.Put
	api_client.Delete
	api_client.UI
}
