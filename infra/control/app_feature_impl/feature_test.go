package app_feature_impl

import (
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"io/ioutil"
	"os"
	"testing"
)

type SampleFeatureOptIn struct {
	app_feature.OptInStatus
}

func TestFeatureImpl_OptInGetSet(t *testing.T) {
	p, err := ioutil.TempDir("", "feature")
	if err != nil {
		t.Error(err)
		return
	}
	ws, err := app_workspace.NewWorkspace(p, false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()
	com := app_opt.Default()
	fe := NewFeature(com, ws, true)

	sfo := &SampleFeatureOptIn{}
	if sfo1, found := fe.OptInGet(sfo); found {
		t.Error(sfo1)
	}
	sfo.OptInCommit(true)
	if err := fe.OptInUpdate(sfo); err != nil {
		t.Error(err)
	}
	sfo2 := &SampleFeatureOptIn{}
	if sfo3, found := fe.OptInGet(sfo2); !found || !sfo3.OptInIsEnabled() {
		t.Error(sfo3, found)
	}
}
