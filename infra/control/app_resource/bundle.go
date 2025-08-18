package app_resource

import "github.com/watermint/toolbox/essentials/es_go/es_resource"

var (
	bundle = es_resource.EmptyBundle()
)

func Bundle() es_resource.Bundle {
	return bundle
}

func SetBundle(b es_resource.Bundle) {
	bundle = b
}
