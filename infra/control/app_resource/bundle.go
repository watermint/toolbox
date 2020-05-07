package app_resource

import "github.com/watermint/toolbox/essentials/go/es_resource"

var (
	bundle es_resource.Bundle = es_resource.EmptyBundle()
)

func Bundle() es_resource.Bundle {
	return bundle
}

func SetBundle(b es_resource.Bundle) {
	bundle = b
}
