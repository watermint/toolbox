package resources

import (
	"embed"
	"github.com/watermint/toolbox/essentials/go/es_resource"
)

//go:embed templates/*
var resTemplates embed.FS

//go:embed messages/*
var resMessages embed.FS

//go:embed web/*
var resWeb embed.FS

//go:embed keys/*
var resKeys embed.FS

//go:embed images/*
var resImages embed.FS

//go:embed data/*
var resData embed.FS

func NewBundle() es_resource.Bundle {
	return es_resource.New(
		es_resource.NewResource("templates", resTemplates),
		es_resource.NewSecureResource("messages", resMessages),
		es_resource.NewResource("web", resWeb),
		es_resource.NewSecureResource("keys", resKeys),
		es_resource.NewResource("images", resImages),
		es_resource.NewSecureResource("data", resData),
	)
}
