package diag

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_conn_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
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
	Peer rc_conn.OldConnBusinessFile
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
		err := rc_exec.Exec(k.Control(), &team.Info{}, func(r rc_recipe.Recipe) {
			rc := r.(*team.Info)
			rc.Peer.SetPeerName(pn)
		})
		if err != nil {
			l.Error("`team info` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning feature")
		err := rc_exec.Exec(k.Control(), &team.Feature{}, func(r rc_recipe.Recipe) {
			rc := r.(*team.Feature)
			rc.Peer.SetPeerName(pn)
		})
		if err != nil {
			l.Error("`team feature` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning group")
		err := rc_exec.Exec(k.Control(), &group.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*group.List)
			rc.Peer.SetPeerName(pn)
		})
		if err != nil {
			l.Error("`group list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning group members")
		err := rc_exec.Exec(k.Control(), &groupmember.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*groupmember.List)
			rc.Peer.SetPeerName(pn)
		})
		if err != nil {
			l.Error("`group member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning members")
		err := rc_exec.Exec(k.Control(), &member.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*member.List)
			rc.Peer.SetPeerName(pn)
		})
		if err != nil {
			l.Error("`member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning member quota")
		err := rc_exec.Exec(k.Control(), &memberquota.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*memberquota.List)
			rc.Peer.SetPeerName(pn)
		})
		if err != nil {
			l.Error("`member quota list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning member usage")
		err := rc_exec.Exec(k.Control(), &memberquota.Usage{}, func(r rc_recipe.Recipe) {
			rc := r.(*memberquota.Usage)
			rc.Peer.SetPeerName(pn)
		})
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
		err := rc_exec.Exec(k.Control(), &teamfolder.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*teamfolder.List)
			rc.Peer.SetPeerName(pn)
		})
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
		return qt_endtoend.NotEnoughResource()
	}
	if err := z.Exec(rc_kitchen.NewKitchen(c, lvo)); err != nil {
		return err
	}
	return nil
}
