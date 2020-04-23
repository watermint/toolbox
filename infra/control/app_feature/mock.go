package app_feature

import (
	"github.com/watermint/toolbox/infra/control/app_config"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

func NewMock() Feature {
	return &mockFeature{}
}

type mockFeature struct {
}

func (z mockFeature) IsProduction() bool {
	return false
}

func (z mockFeature) IsDebug() bool {
	return false
}

func (z mockFeature) IsTest() bool {
	return true
}

func (z mockFeature) IsQuiet() bool {
	return false
}

func (z mockFeature) IsSecure() bool {
	return false
}

func (z mockFeature) IsLowMemory() bool {
	return false
}

func (z mockFeature) IsAutoOpen() bool {
	return false
}

func (z mockFeature) UIFormat() string {
	return ""
}

func (z mockFeature) Config() app_config.Config {
	panic("the function is not available on the mock")
}

func (z mockFeature) OptInGet(oi OptIn) (f OptIn, found bool) {
	return nil, false
}

func (z mockFeature) OptInUpdate(oi OptIn) error {
	return qt_errors.ErrorMock
}
