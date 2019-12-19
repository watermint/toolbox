package diag

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	memberquota "github.com/watermint/toolbox/recipe/member/quota"
	"github.com/watermint/toolbox/recipe/team"
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

type ExplorerVO struct {
	Peer rc_conn.ConnBusinessFile
	All  bool
}

type Explorer struct {
}

func (z *Explorer) Console() {
}

func (z *Explorer) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Explorer) Requirement() rc_vo.ValueObject {
	return &ExplorerVO{}
}

func (z *Explorer) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*ExplorerVO)
	l := k.Log()
	pn := vo.Peer.(*rc_conn_impl.ConnBusinessFile).PeerName
	{
		l.Info("Scanning info")
		r := team.Info{}
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &team.InfoVO{
			Peer: &rc_conn_impl.ConnBusinessInfo{
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
		r := team.Feature{}
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &team.FeatureVO{
			Peer: &rc_conn_impl.ConnBusinessInfo{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &group.ListVO{
			Peer: &rc_conn_impl.ConnBusinessInfo{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &groupmember.ListVO{
			Peer: &rc_conn_impl.ConnBusinessInfo{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &member.ListVO{
			Peer: &rc_conn_impl.ConnBusinessInfo{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &memberquota.ListVO{
			Peer: &rc_conn_impl.ConnBusinessMgmt{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`member quota list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning member usage")
		r := memberquota.Usage{}
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &memberquota.UsageVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`member quota usage` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning devices")
		r := teamdevice.List{}
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &teamdevice.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &teamfilerequest.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &teamlinkedapp.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &teamfolder.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &namespace.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &namespacemember.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
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
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &teamsharedlink.ListVO{
			Peer: &rc_conn_impl.ConnBusinessFile{
				PeerName: pn,
			},
		}))
		if err != nil {
			l.Error("`team sharedlink list` failed", zap.Error(err))
			return err
		}
	}

	if vo.All {
		l.Info("Scanning namespace file list")
		{
			r := namespacefile.List{}
			err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &namespacefile.ListVO{
				Peer: &rc_conn_impl.ConnBusinessFile{
					PeerName: pn,
				},
				IncludeMemberFolder: true,
				IncludeDeleted:      true,
				IncludeSharedFolder: true,
				IncludeMediaInfo:    true,
			}))
			if err != nil {
				l.Error("`team sharedlink list` failed", zap.Error(err))
				return err
			}
		}

		l.Info("Scanning namespace file size")
		{
			r := namespacefile.Size{}
			err := r.Exec(rc_kitchen.NewKitchen(k.Control(), &namespacefile.SizeVO{
				Peer: &rc_conn_impl.ConnBusinessFile{
					PeerName: pn,
				},
				IncludeMemberFolder: true,
				IncludeSharedFolder: true,
				IncludeTeamFolder:   true,
				IncludeAppFolder:    true,
			}))
			if err != nil {
				l.Error("`team sharedlink list` failed", zap.Error(err))
				return err
			}
		}
	}

	return nil
}

func (z *Explorer) Test(c app_control.Control) error {
	lvo := &ExplorerVO{
		All: false,
	}
	if !qt_recipe.ApplyTestPeers(c, lvo) {
		return qt_recipe.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return nil
}
