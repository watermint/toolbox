package rc_value

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/essentials/api/api_conn"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_control"
	"testing"
)

type ValueConnBaseRecipe struct {
	Peer dbx_conn.ConnScopedIndividual
}

func (z *ValueConnBaseRecipe) Preset() {
	z.Peer.SetPeerName("value_test")
}

func (z *ValueConnBaseRecipe) Exec(c app_control.Control) error {
	return nil
}

func (z *ValueConnBaseRecipe) Test(c app_control.Control) error {
	return nil
}

func TestValueConnBaseRecipe(t *testing.T) {
	err := qt_control.WithControl(func(c app_control.Control) error {
		rcp0 := &ValueConnBaseRecipe{}
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
		mod1 := rcp1.(*ValueConnBaseRecipe)
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
		mod2 := rcp2.(*ValueConnBaseRecipe)
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

func TestValueConnBase_Restore(t *testing.T) {
	err := qt_control.WithControl(func(ctl app_control.Control) error {
		v := NewValueConn((*dbx_conn.ConnScopedIndividual)(nil), func(peerName string) api_conn.Connection {
			return dbx_conn_impl.NewConnScopedIndividual(peerName)
		})
		v.ApplyPreset(dbx_conn_impl.NewConnScopedIndividual("123"))

		vc, err := v.Capture(ctl)
		if err != nil {
			t.Error(err)
		}

		capData, err := json.Marshal(vc)
		if err != nil {
			t.Error(err)
		}

		capJson, err := es_json.Parse(capData)
		if err != nil {
			t.Error(err)
		}

		v2 := NewValueConn((*dbx_conn.ConnScopedIndividual)(nil), func(peerName string) api_conn.Connection {
			return dbx_conn_impl.NewConnScopedIndividual(peerName)
		})

		err = v2.Restore(capJson, ctl)
		if err != nil {
			t.Error(err)
		}

		v2b := v2.Bind().(*string)
		if *v2b != "123" {
			t.Error(v2b)
		}
		return err
	})
	if err != nil {
		t.Error(err)
	}
}
