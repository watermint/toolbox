package dbx_filesystem

type RootInfo struct {
	RootNamespaceId string `json:"root_namespace_id" path:"root_info.root_namespace_id"`
	HomeNamespaceId string `json:"home_namespace_id" path:"root_info.home_namespace_id"`
}
