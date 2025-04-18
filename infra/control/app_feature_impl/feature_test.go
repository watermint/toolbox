package app_feature_impl

import (
	"os"
	"testing"

	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
)

type SampleFeatureOptIn struct {
	app_feature.OptInStatus
}

func TestFeatureImpl_OptInGetSet(t *testing.T) {
	p, err := os.MkdirTemp("", "feature")
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
	sfo.OptInCommit(true)
	if err := fe.OptInUpdate(sfo); err != nil {
		t.Error(err)
	}
	sfo2 := &SampleFeatureOptIn{}
	if sfo3, found := fe.OptInGet(sfo2); !found || !sfo3.OptInIsEnabled() {
		t.Error(sfo3, found)
	}

	sfo.OptInCommit(false)
	if err := fe.OptInUpdate(sfo); err != nil {
		t.Error(err)
	}
	if sfo3, found := fe.OptInGet(sfo); !found || sfo3.OptInIsEnabled() {
		t.Error(sfo3, found)
	}
}

func TestFeatureImpl_Experiment(t *testing.T) {
	p, err := os.MkdirTemp("", "feature")
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
	com.Experiment = "experiment1,experiment2"
	fe := NewFeature(com, ws, true)

	if e := fe.Experiment("experiment1"); !e {
		t.Error(e)
	}
	if e := fe.Experiment("experiment2"); !e {
		t.Error(e)
	}
	if e := fe.Experiment("experiment3"); e {
		t.Error(e)
	}
}
