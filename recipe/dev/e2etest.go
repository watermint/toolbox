package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_conn_impl"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type EndToEndTestVO struct {
	Full  app_conn.ConnUserFile
	Info  app_conn.ConnBusinessInfo
	File  app_conn.ConnBusinessFile
	Audit app_conn.ConnBusinessAudit
	Mgmt  app_conn.ConnBusinessMgmt
}

type EndToEnd struct {
}

func (z *EndToEnd) Hidden() {
}

func (z *EndToEnd) Console() {
}

func (z *EndToEnd) Requirement() app_vo.ValueObject {
	return &EndToEndTestVO{
		Full:  &app_conn_impl.ConnUserFile{PeerName: app_test.EndToEndPeer},
		Info:  &app_conn_impl.ConnBusinessInfo{PeerName: app_test.EndToEndPeer},
		File:  &app_conn_impl.ConnBusinessFile{PeerName: app_test.EndToEndPeer},
		Audit: &app_conn_impl.ConnBusinessAudit{PeerName: app_test.EndToEndPeer},
		Mgmt:  &app_conn_impl.ConnBusinessMgmt{PeerName: app_test.EndToEndPeer},
	}
}

func (z *EndToEnd) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	evo := vo.(*EndToEndTestVO)
	if _, err := evo.Full.Connect(k.Control()); err != nil {
		return err
	}
	if _, err := evo.Info.Connect(k.Control()); err != nil {
		return err
	}
	if _, err := evo.File.Connect(k.Control()); err != nil {
		return err
	}
	if _, err := evo.Mgmt.Connect(k.Control()); err != nil {
		return err
	}
	if _, err := evo.Audit.Connect(k.Control()); err != nil {
		return err
	}
	return nil
}

func (z *EndToEnd) Test(c app_control.Control) error {
	return nil
}
