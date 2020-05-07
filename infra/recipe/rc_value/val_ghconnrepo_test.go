package rc_value

import (
	"flag"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueGhConnGithubPrivateRecipe struct {
	Peer dbx_conn.ConnBusinessMgmt
}

func (z *ValueGhConnGithubPrivateRecipe) Preset() {
	z.Peer.SetPeerName("value_test")
}

func (z *ValueGhConnGithubPrivateRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueGhConnGithubPrivateRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueGhConnGithubPrivate(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueGhConnGithubPrivateRecipe{}
		repo := NewRepository(rcp0)

		// Parse flags
		flg := flag.NewFlagSet("value", flag.ContinueOnError)
		repo.ApplyFlags(flg, c.UI())
		if err := flg.Parse([]string{"-peer", "by_argument"}); err != nil {
			t.Error(err)
			return err
		}

		// Apply parsed values
		rcp1 := repo.Apply()
		mod1 := rcp1.(*ValueGhConnGithubPrivateRecipe)
		if mod1.Peer.PeerName() != "by_argument" {
			t.Error(mod1)
		}

		// Spin up
		ct := c.WithFeature(c.Feature().AsTest(true))
		rcp2, err := repo.SpinUp(ct)
		if err != nil {
			t.Error(err)
			return err
		}
		mod2 := rcp2.(*ValueGhConnGithubPrivateRecipe)
		if mod1.Peer.PeerName() != "by_argument" {
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
