package work_client

import "github.com/watermint/toolbox/infra/api/api_client"

type Client interface {
	api_client.Client
	api_client.Get
	api_client.Post
}
