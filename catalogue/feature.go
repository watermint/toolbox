package catalogue

// Code generated by dev catalogue command DO NOT EDIT

import (
	essentialsapiapi_auth_oauth "github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_feature"
	infrareportrp_artifact_feature "github.com/watermint/toolbox/infra/report/rp_artifact_feature"
	ingredientig_bootstrap "github.com/watermint/toolbox/ingredient/ig_bootstrap"
)

func AutoDetectedFeatures() []app_feature.OptIn {
	return []app_feature.OptIn{
		&essentialsapiapi_auth_oauth.OptInFeatureRedirect{},
		&infrareportrp_artifact_feature.OptInFeatureAutoOpen{},
		&ingredientig_bootstrap.OptInFeatureAutodelete{},
	}
}
