package test

import (
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type AuthVO struct {
	Full  rc_conn.ConnUserFile
	Info  rc_conn.ConnBusinessInfo
	File  rc_conn.ConnBusinessFile
	Audit rc_conn.ConnBusinessAudit
	Mgmt  rc_conn.ConnBusinessMgmt
}

type Auth struct {
}

func (z *Auth) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Auth) Hidden() {
}

func (z *Auth) Console() {
}

func (z *Auth) Requirement() rc_vo.ValueObject {
	return &AuthVO{
		Full:  &rc_conn_impl.ConnUserFile{PeerName: qt_recipe.EndToEndPeer},
		Info:  &rc_conn_impl.ConnBusinessInfo{PeerName: qt_recipe.EndToEndPeer},
		File:  &rc_conn_impl.ConnBusinessFile{PeerName: qt_recipe.EndToEndPeer},
		Audit: &rc_conn_impl.ConnBusinessAudit{PeerName: qt_recipe.EndToEndPeer},
		Mgmt:  &rc_conn_impl.ConnBusinessMgmt{PeerName: qt_recipe.EndToEndPeer},
	}
}

func (z *Auth) Exec(k rc_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	evo := vo.(*AuthVO)
	l := k.Log()
	l.Info("Please proceed auth for `User Full`")
	if _, err := evo.Full.Connect(k.Control()); err != nil {
		return err
	}
	l.Info("Please proceed auth for `Business Info`")
	if _, err := evo.Info.Connect(k.Control()); err != nil {
		return err
	}
	l.Info("Please proceed auth for `Business File`")
	if _, err := evo.File.Connect(k.Control()); err != nil {
		return err
	}
	l.Info("Please proceed auth for `Business Mgmt`")
	if _, err := evo.Mgmt.Connect(k.Control()); err != nil {
		return err
	}
	l.Info("Please proceed auth for `Business Audit`")
	if _, err := evo.Audit.Connect(k.Control()); err != nil {
		return err
	}
	if err := api_auth_impl.CreateCompatible(k.Control(), qt_recipe.EndToEndPeer); err != nil {
		return err
	}
	return nil
}

func (z *Auth) Test(c app_control.Control) error {
	return qt_recipe.NoTestRequired()
}
