package team

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_conn_impl"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	memberquota "github.com/watermint/toolbox/recipe/member/quota"
	teamdevice "github.com/watermint/toolbox/recipe/team/device"
	teamfilerequest "github.com/watermint/toolbox/recipe/team/filerequest"
	teamlinkedapp "github.com/watermint/toolbox/recipe/team/linkedapp"
	"github.com/watermint/toolbox/recipe/team/namespace"
	namespacefile "github.com/watermint/toolbox/recipe/team/namespace/file"
	namespacemember "github.com/watermint/toolbox/recipe/team/namespace/member"
	teamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	"github.com/watermint/toolbox/recipe/teamfolder"
	"go.uber.org/zap"
)

type DiagnosisVO struct {
	Peer            app_conn.ConnBusinessFile
	IncludeFileList bool
}

type Diagnosis struct {
}

func (z *Diagnosis) Hidden() {
}

func (z *Diagnosis) Requirement() app_vo.ValueObject {
	return &DiagnosisVO{}
}

func (z *Diagnosis) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*DiagnosisVO)
	l := k.Log()
	pn := vo.Peer.(*app_conn_impl.ConnBusinessFile).PeerName
	{
		l.Info("Scanning info")
		r := Info{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &InfoVO{
			Peer: &app_conn_impl.ConnBusinessInfo{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team info` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning feature")
		r := Feature{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &FeatureVO{
			Peer: &app_conn_impl.ConnBusinessInfo{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team feature` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning group")
		r := group.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &group.ListVO{
			Peer: &app_conn_impl.ConnBusinessInfo{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`group list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning group members")
		r := groupmember.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &groupmember.ListVO{
			Peer: &app_conn_impl.ConnBusinessInfo{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`group member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning members")
		r := member.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &member.ListVO{
			Peer: &app_conn_impl.ConnBusinessInfo{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning member quota")
		r := memberquota.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &memberquota.ListVO{
			Peer: &app_conn_impl.ConnBusinessMgmt{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`member quota list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning devices")
		r := teamdevice.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &teamdevice.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team device list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning file requests")
		r := teamfilerequest.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &teamfilerequest.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team filerequest list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning linked apps")
		r := teamlinkedapp.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &teamlinkedapp.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team linkedapp list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning team folders")
		r := teamfolder.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &teamfolder.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`teamfolder list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning namespaces")
		r := namespace.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &namespace.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team namespace list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning namespace members")
		r := namespacemember.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &namespacemember.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team namespace member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning shared links")
		r := teamsharedlink.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &teamsharedlink.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team sharedlink list` failed", zap.Error(err))
			return err
		}
	}

	if vo.IncludeFileList {
		l.Info("Scanning namespace file list")
		r := namespacefile.List{}
		err := r.Exec(app_kitchen.NewKitchen(k.Control(), &namespacefile.ListVO{
			Peer: &app_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team sharedlink list` failed", zap.Error(err))
			return err
		}
	}

	return nil
}

func (z *Diagnosis) Test(c app_control.Control) error {
	lvo := &DiagnosisVO{
		IncludeFileList: false,
	}
	if !app_test.ApplyTestPeers(c, lvo) {
		return nil
	}
	if err := z.Exec(app_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return nil
}
