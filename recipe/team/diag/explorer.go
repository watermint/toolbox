package diag

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/recipe/group"
	groupmember "github.com/watermint/toolbox/recipe/group/member"
	"github.com/watermint/toolbox/recipe/member"
	memberquota "github.com/watermint/toolbox/recipe/member/quota"
	"github.com/watermint/toolbox/recipe/team"
	teamdevice "github.com/watermint/toolbox/recipe/team/device"
	teamfilerequest "github.com/watermint/toolbox/recipe/team/filerequest"
	teamlinkedapp "github.com/watermint/toolbox/recipe/team/linkedapp"
	teamnamespace "github.com/watermint/toolbox/recipe/team/namespace"
	namespacefile "github.com/watermint/toolbox/recipe/team/namespace/file"
	teamnamespacemember "github.com/watermint/toolbox/recipe/team/namespace/member"
	teamsharedlink "github.com/watermint/toolbox/recipe/team/sharedlink"
	"github.com/watermint/toolbox/recipe/teamfolder"
	"go.uber.org/zap"
)

type Explorer struct {
	Info                        dbx_conn.ConnBusinessInfo
	File                        dbx_conn.ConnBusinessFile
	Mgmt                        dbx_conn.ConnBusinessMgmt
	All                         bool
	RecipeInfo                  *team.Info
	RecipeFeature               *team.Feature
	RecipeGroupList             *group.List
	RecipeGroupMemberList       *groupmember.List
	RecipeMemberList            *member.List
	RecipeMemberQuotaList       *memberquota.List
	RecipeMemberQuotaUsage      *memberquota.Usage
	RecipeTeamDeviceList        *teamdevice.List
	RecipeTeamFilerequestList   *teamfilerequest.List
	RecipeTeamLinkedappList     *teamlinkedapp.List
	RecipeTeamSharedlinkList    *teamsharedlink.List
	RecipeTeamNamespaceList     *teamnamespace.List
	RecipeTeamNamespaceFileList *namespacefile.List
	RecipeTeamNamespaceFileSize *namespacefile.Size
}

func (z *Explorer) Preset() {
}

func (z *Explorer) Exec(c app_control.Control) error {
	l := c.Log()
	{
		l.Info("Scanning info")
		fc, err := c.(app_control_launcher.ControlFork).Fork("info")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeInfo, func(r rc_recipe.Recipe) {
			rc := r.(*team.Info)
			rc.Peer = z.Info
		})
		if err != nil {
			l.Error("`team info` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning feature")
		fc, err := c.(app_control_launcher.ControlFork).Fork("feature")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeFeature, func(r rc_recipe.Recipe) {
			rc := r.(*team.Feature)
			rc.Peer = z.Info
		})
		if err != nil {
			l.Error("`team feature` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning group")
		fc, err := c.(app_control_launcher.ControlFork).Fork("group_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeGroupList, func(r rc_recipe.Recipe) {
			rc := r.(*group.List)
			rc.Peer = z.Info
		})
		if err != nil {
			l.Error("`group list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning group members")
		fc, err := c.(app_control_launcher.ControlFork).Fork("group_member_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeGroupMemberList, func(r rc_recipe.Recipe) {
			rc := r.(*groupmember.List)
			rc.Peer = z.Info
		})
		if err != nil {
			l.Error("`group member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning members")
		fc, err := c.(app_control_launcher.ControlFork).Fork("member_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeMemberList, func(r rc_recipe.Recipe) {
			rc := r.(*member.List)
			rc.Peer = z.Info
		})
		if err != nil {
			l.Error("`member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning member quota")
		fc, err := c.(app_control_launcher.ControlFork).Fork("member_quota_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeMemberQuotaList, func(r rc_recipe.Recipe) {
			rc := r.(*memberquota.List)
			rc.Peer = z.Mgmt
		})
		if err != nil {
			l.Error("`member quota list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning member usage")
		fc, err := c.(app_control_launcher.ControlFork).Fork("member_quota_usage")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeMemberQuotaUsage, func(r rc_recipe.Recipe) {
			rc := r.(*memberquota.Usage)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`member quota usage` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning devices")
		fc, err := c.(app_control_launcher.ControlFork).Fork("team_device_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeTeamDeviceList, func(r rc_recipe.Recipe) {
			rc := r.(*teamdevice.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`team device list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning file requests")
		fc, err := c.(app_control_launcher.ControlFork).Fork("team_filerequest_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeTeamFilerequestList, func(r rc_recipe.Recipe) {
			rc := r.(*teamfilerequest.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`team filerequest list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning linked apps")
		fc, err := c.(app_control_launcher.ControlFork).Fork("team_linkedapp_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeTeamLinkedappList, func(r rc_recipe.Recipe) {
			rc := r.(*teamlinkedapp.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`team linkedapp list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning team folders")
		fc, err := c.(app_control_launcher.ControlFork).Fork("teamfolder_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, &teamfolder.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*teamfolder.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`teamfolder list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning namespaces")
		fc, err := c.(app_control_launcher.ControlFork).Fork("team_namespace_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, &teamnamespace.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*teamnamespace.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`team namespace list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning namespace members")
		fc, err := c.(app_control_launcher.ControlFork).Fork("team_namespace_member_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, &teamnamespacemember.List{}, func(r rc_recipe.Recipe) {
			rc := r.(*teamnamespacemember.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`team namespace member list` failed", zap.Error(err))
			return err
		}
	}

	{
		l.Info("Scanning shared links")
		fc, err := c.(app_control_launcher.ControlFork).Fork("team_sharedlink_list")
		if err != nil {
			return err
		}
		err = rc_exec.Exec(fc, z.RecipeTeamSharedlinkList, func(r rc_recipe.Recipe) {
			rc := r.(*teamsharedlink.List)
			rc.Peer = z.File
		})
		if err != nil {
			l.Error("`team sharedlink list` failed", zap.Error(err))
			return err
		}
	}

	if z.All {
		l.Info("Scanning namespace file list")
		{
			fc, err := c.(app_control_launcher.ControlFork).Fork("team_namespace_file_list")
			if err != nil {
				return err
			}
			err = rc_exec.Exec(fc, z.RecipeTeamNamespaceFileList, func(r rc_recipe.Recipe) {
				rc := r.(*namespacefile.List)
				rc.Peer = z.File
				rc.IncludeMemberFolder = true
				rc.IncludeDeleted = true
				rc.IncludeSharedFolder = true
				rc.IncludeMediaInfo = true
			})
			if err != nil {
				l.Error("`team sharedlink list` failed", zap.Error(err))
				return err
			}
		}

		l.Info("Scanning namespace file size")
		{
			fc, err := c.(app_control_launcher.ControlFork).Fork("team_namespace_file_size")
			if err != nil {
				return err
			}
			err = rc_exec.Exec(fc, z.RecipeTeamNamespaceFileSize, func(r rc_recipe.Recipe) {
				rc := r.(*namespacefile.Size)
				rc.Peer = z.File
				rc.IncludeMemberFolder = true
				rc.IncludeSharedFolder = true
				rc.IncludeTeamFolder = true
				rc.IncludeAppFolder = true
			})
			if err != nil {
				l.Error("`team sharedlink list` failed", zap.Error(err))
				return err
			}
		}
	}

	return nil
}

func (z *Explorer) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Explorer{}, func(r rc_recipe.Recipe) {
		rc := r.(*Explorer)
		rc.All = false
	})
}
