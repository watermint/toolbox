package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	infraapiapi_auth_oauth "github.com/watermint/toolbox/infra/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_feature"
	infrareportrp_artifact_feature "github.com/watermint/toolbox/infra/report/rp_artifact_feature"
	ingredientbootstrap "github.com/watermint/toolbox/ingredient/bootstrap"
)

func AutoDetectedFeatures() []app_feature.OptIn {
	return []app_feature.OptIn{
		&infraapiapi_auth_oauth.OptInFeatureRedirect{},
		&infrareportrp_artifact_feature.OptInFeatureAutoOpen{},
		&ingredientbootstrap.OptInFeatureAutodelete{},
	}
}
