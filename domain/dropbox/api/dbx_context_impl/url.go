package dbx_context_impl

import "fmt"

const (
	RpcEndpoint     = "api.dropboxapi.com"
	NotifyEndpoint  = "notify.dropboxapi.com"
	ContentEndpoint = "content.dropboxapi.com"
)

func RpcRequestUrl(base, endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", base, endpoint)
}

func ContentRequestUrl(endpoint string) string {
	return fmt.Sprintf("https://%s/2/%s", ContentEndpoint, endpoint)
}
