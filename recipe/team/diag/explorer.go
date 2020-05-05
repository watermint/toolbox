package diag

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
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
	"strings"
)

type exploreRecipe struct {
	recipe rc_recipe.Recipe
	custom func(r rc_recipe.Recipe)
}

func (z exploreRecipe) exec(c app_control.Control) error {
	spec := rc_spec.New(z.recipe)
	name := strings.ReplaceAll(spec.CliPath(), " ", "-")
	l := c.Log().With(es_log.String("Name", name))
	ui := c.UI()
	ui.Info(spec.Title())
	l.Info("Execute:")

	return app_workspace.WithFork(c.WorkBundle(), name, func(fwb app_workspace.Bundle) error {
		cf := c.WithBundle(fwb)
		if err := rc_exec.Exec(cf, z.recipe, z.custom); err != nil {
			l.Error("Recipe failed", es_log.Error(err))
			return err
		}
		return nil
	})
}

type Explorer struct {
	Info                          dbx_conn.ConnBusinessInfo
	File                          dbx_conn.ConnBusinessFile
	Mgmt                          dbx_conn.ConnBusinessMgmt
	All                           bool
	RecipeInfo                    *team.Info
	RecipeFeature                 *team.Feature
	RecipeGroupList               *group.List
	RecipeGroupMemberList         *groupmember.List
	RecipeMemberList              *member.List
	RecipeMemberQuotaList         *memberquota.List
	RecipeMemberQuotaUsage        *memberquota.Usage
	RecipeTeamFolderList          *teamfolder.List
	RecipeTeamDeviceList          *teamdevice.List
	RecipeTeamFilerequestList     *teamfilerequest.List
	RecipeTeamLinkedappList       *teamlinkedapp.List
	RecipeTeamSharedlinkList      *teamsharedlink.List
	RecipeTeamNamespaceList       *teamnamespace.List
	RecipeTeamNamespaceMemberList *teamnamespacemember.List
	RecipeTeamNamespaceFileList   *namespacefile.List
	RecipeTeamNamespaceFileSize   *namespacefile.Size
}

func (z *Explorer) Preset() {
}

func (z *Explorer) Exec(c app_control.Control) error {
	ers := []exploreRecipe{
		{
			recipe: z.RecipeInfo,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*team.Info)
				rc.Peer = z.Info
			},
		},
		{
			recipe: z.RecipeFeature,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*team.Feature)
				rc.Peer = z.Info
			},
		},
		{
			recipe: z.RecipeGroupList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*group.List)
				rc.Peer = z.Info
			},
		},
		{
			recipe: z.RecipeGroupMemberList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*groupmember.List)
				rc.Peer = z.Info
			},
		},
		{
			recipe: z.RecipeMemberList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*member.List)
				rc.Peer = z.Info
			},
		},
		{
			recipe: z.RecipeMemberQuotaList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*memberquota.List)
				rc.Peer = z.Mgmt
			},
		},
		{
			recipe: z.RecipeMemberQuotaUsage,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*memberquota.Usage)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamDeviceList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamdevice.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamFilerequestList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamfilerequest.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamLinkedappList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamlinkedapp.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamFolderList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamfolder.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamNamespaceList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamnamespace.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamNamespaceMemberList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamnamespacemember.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamSharedlinkList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*teamsharedlink.List)
				rc.Peer = z.File
			},
		},
		{
			recipe: z.RecipeTeamNamespaceFileList,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*namespacefile.List)
				rc.Peer = z.File
				rc.IncludeMemberFolder = true
				rc.IncludeDeleted = true
				rc.IncludeSharedFolder = true
				rc.IncludeMediaInfo = true
			},
		},
		{
			recipe: z.RecipeTeamNamespaceFileSize,
			custom: func(r rc_recipe.Recipe) {
				rc := r.(*namespacefile.Size)
				rc.Peer = z.File
				rc.IncludeMemberFolder = true
				rc.IncludeSharedFolder = true
				rc.IncludeTeamFolder = true
				rc.IncludeAppFolder = true
			},
		},
	}

	for _, er := range ers {
		if err := er.exec(c); err != nil {
			return err
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
