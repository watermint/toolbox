package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueGhConnGithubPublicRecipe struct {
	Peer gh_conn.ConnGithubPublic
}

func (z *ValueGhConnGithubPublicRecipe) Preset() {
	z.Peer.SetPeerName("value_test")
}

func (z *ValueGhConnGithubPublicRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueGhConnGithubPublicRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueGhConnGithubPublic(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueGhConnGithubPublicRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())

		// Spin up
		ct := c.WithFeature(c.Feature().AsTest(true))
		rcp2, err := repo.SpinUp(ct)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueGhConnGithubPublicRecipe)
		if mod2.Peer.Client().ClientHash() == "" {
			t.Error(mod2)
		}

		if err := repo.SpinDown(ct); err != nil {
			t.Error(err)
			return err
		}

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
