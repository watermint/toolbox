package stage

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_folder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"strings"
	"time"
)

type Teamfolder struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkExperimental
	Peer dbx_conn.ConnScopedTeam
}

func (z *Teamfolder) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
		dbx_auth.ScopeGroupsWrite,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *Teamfolder) Exec(c app_control.Control) error {
	runId := fmt.Sprintf("%08x", time.Now().Unix())
	teamFolderName := "Tokyo Branch " + runId
	nestedFolderPlainName := "Organization"
	nestedFolderSharedName := "Sales"
	nestedFolderRestrictedName := "Report"
	restedFolderRestrictedNoSyncName := "Finance"
	teamFolderMarketing := "Marketing " + runId
	nestedFolderMarketingPlan := "Plan"
	adminGroupName := "toolbox-admin"
	sampleGroupName := "toolbox-sample"

	// [Tokyo Branch] (Team folder, [editor=toolbox-admin])
	//  |
	//  +-- [Organization] (plain folder, not_synced)
	//  |
	//  +-- [Sales] (nested folder, not_synced)
	//  |    |
	//  |    +-- [Report] (nested folder, do not inherit, no external sharing, [editor=toolbox-sample])
	//  |
	//  +-- [Finance] (nested folder, not_synced, do not inherit)
	//
	// [Marketing] (Team folder, acl_policy=owner only)
	//       |
	//       +-- [Plan] (nested folder, no inherit)

	l := c.Log()

	// find admin
	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		return err
	}

	// create team folder
	tf, err := sv_teamfolder.New(z.Peer.Client()).Create(teamFolderName)
	de := dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("teamfolder", tf))
		break

	case de.IsFolderNameAlreadyUsed():
		l.Info("The folder already created")
		teamfolders, err := sv_teamfolder.New(z.Peer.Client()).List()
		if err != nil {
			l.Warn("Unable to retrieve team folder list", esl.Error(err))
			return err
		}

		for _, teamfolder := range teamfolders {
			if strings.ToLower(teamfolder.Name) == strings.ToLower(teamFolderName) {
				tf = teamfolder
				break
			}
		}
		if tf == nil {
			l.Warn("Team folder not found")
			return errors.New("team folder not found")
		}

		break

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	tfCtx := z.Peer.Client().AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Namespace(tf.TeamFolderId))

	// create sub folder : Organization
	folderOrganization, err := sv_file_folder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderPlainName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderOrganization))
		break

	case de.Path().IsConflict():
		l.Info("The folder already created")
		folderOrganization, err = sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderPlainName))
		if err != nil {
			l.Warn("Unable to identify sub folder", esl.Error(err))
			return err
		}
		break

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// create nested folder : Sales
	folderSales, err := sv_sharedfolder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderSales))
		break

	case de.BadPath().IsAlreadyShared():
		l.Info("The folder is already shared")
		folderSalesMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}

		folderSales, err = sv_sharedfolder.New(tfCtx).Resolve(folderSalesMeta.Concrete().SharedFolderId)
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}
		l.Info("Nested folder resolved", esl.Any("folder", folderSales))

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// create nested folder : Sales
	folderSalesReport, err := sv_sharedfolder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderSharedName + "/" + nestedFolderRestrictedName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderSalesReport))
		break

	case de.BadPath().IsAlreadyShared():
		l.Info("The folder is already shared")
		folderSalesReportMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName + "/" + nestedFolderRestrictedName))
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}

		folderSalesReport, err = sv_sharedfolder.New(tfCtx).Resolve(folderSalesReportMeta.Concrete().SharedFolderId)
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}
		l.Info("Nested folder resolved", esl.Any("folder", folderSales))

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// Change sync setting
	folderOrganizationMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderPlainName))
	if err != nil {
		l.Warn("Unable to find meta", esl.Error(err))
		return err
	}
	folderSalesMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
	if err != nil {
		l.Warn("Unable to find meta", esl.Error(err))
		return err
	}

	updated, err := sv_teamfolder.New(z.Peer.Client()).UpdateSyncSetting(tf,
		sv_teamfolder.AddNestedSetting(folderOrganizationMeta, sv_teamfolder.SyncSettingNotSynced),
		sv_teamfolder.AddNestedSetting(folderSalesMeta, sv_teamfolder.SyncSettingNotSynced),
	)
	if err != nil {
		l.Warn("Unable to change : sync setting", esl.Error(err))
		return err
	}
	l.Info("Sync settings updated", esl.Any("updated", updated))

	// Create toolbox admin group
	adminGroup, err := sv_group.New(z.Peer.Client()).Create(
		adminGroupName,
		sv_group.CompanyManaged(),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The admin group created", esl.Any("group", adminGroup))

	case de.IsGroupNameAlreadyUsed():
		l.Info("The admin group already created")
		adminGroup, err = sv_group.New(z.Peer.Client()).ResolveByName(adminGroupName)
		if err != nil {
			l.Warn("Unable to find the admin group", esl.Error(err))
			return err
		}

	default:
		l.Warn("Unable to create the admin group", esl.Error(err))
		return err
	}

	// Add the admin to the admin group
	updatedAdminGroup, err := sv_group_member.NewByGroupId(z.Peer.Client(), adminGroup.GroupId).Add(
		sv_group_member.ByTeamMemberId(admin.TeamMemberId),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The admin successfully added to the admin group", esl.Any("group", updatedAdminGroup))

	case de.IsDuplicateUser():
		l.Info("The admin is already added to the admin group", esl.Any("group", updatedAdminGroup))

	default:
		l.Warn("Unable to add member", esl.Error(err))
		return err
	}

	// Create toolbox sample group
	sampleGroup, err := sv_group.New(z.Peer.Client()).Create(
		sampleGroupName,
		sv_group.UserManaged(),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The sample group created", esl.Any("group", sampleGroup))

	case de.IsGroupNameAlreadyUsed():
		l.Info("The sample group already created")
		sampleGroup, err = sv_group.New(z.Peer.Client()).ResolveByName(sampleGroupName)
		if err != nil {
			l.Warn("Unable to find the sample group", esl.Error(err))
			return err
		}

	default:
		l.Warn("Unable to create the sample group", esl.Error(err))
		return err
	}

	// Add admin group to the team folder
	err = sv_sharedfolder_member.NewByTeamFolder(z.Peer.Client().AsAdminId(admin.TeamMemberId), tf).Add(
		sv_sharedfolder_member.AddByGroup(adminGroup, "editor"),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The admin group added to the team folder as editor")

	default:
		l.Warn("Unable to update members", esl.Error(err))
	}

	// Do not inherit permission from parent : Sales/Report
	for {
		updatedFolderSalesReport, err := sv_sharedfolder.New(z.Peer.Client().AsMemberId(admin.TeamMemberId)).UpdateInheritance(folderSalesReport.SharedFolderId, sv_sharedfolder.AccessInheritanceNoInherit)
		de = dbx_error.NewErrors(err)
		if de == nil {
			l.Info("The sample group added to the team folder as editor", esl.Any("updated", updatedFolderSalesReport))
			break
		}
		switch {
		case de.HasPrefix("access_error/not_a_member"):
			l.Warn("Permission look like not ready", esl.Error(err))
			time.Sleep(3 * time.Second)
			l.Info("Retry")
			continue

		default:
			l.Warn("Unable to update members", esl.Error(err))
			return err
		}
	}

	// Add sample group to the nested folder
	err = sv_sharedfolder_member.NewBySharedFolderId(z.Peer.Client().AsAdminId(admin.TeamMemberId), folderSalesReport.SharedFolderId).Add(
		sv_sharedfolder_member.AddByGroup(sampleGroup, "editor"),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The sample group added to the team folder as editor")

	default:
		l.Warn("Unable to update members", esl.Error(err))
		return err
	}

	// Change folder policy : Sales
	updatedSalesPolicy, err := sv_sharedfolder.New(z.Peer.Client().AsAdminId(admin.TeamMemberId)).UpdatePolicy(
		folderSales.SharedFolderId,
		sv_sharedfolder.MemberPolicy("team"),
		sv_sharedfolder.AclUpdatePolicy("owner"),
		sv_sharedfolder.SharedLinkPolicy("team"),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The sales folder policy successfully updated", esl.Any("updated", updatedSalesPolicy))

	default:
		l.Warn("Unable to update policies", esl.Error(err))
	}

	// Restricted & no sync
	// Apply no sync to Finance: 1. create folder
	folderFinance, err := sv_sharedfolder.New(tfCtx).Create(mo_path.NewDropboxPath("/" + restedFolderRestrictedNoSyncName))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderFinance))
		break

	case de.BadPath().IsAlreadyShared():
		l.Info("The folder is already shared")
		folderFinanceMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + restedFolderRestrictedNoSyncName))
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}

		folderFinance, err = sv_sharedfolder.New(tfCtx).Resolve(folderFinanceMeta.Concrete().SharedFolderId)
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}
		l.Info("Nested folder resolved", esl.Any("folder", folderSales))

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	folderFinanceMeta, err := sv_file.NewFiles(tfCtx).Resolve(mo_path.NewDropboxPath("/" + restedFolderRestrictedNoSyncName))
	if err != nil {
		l.Warn("Unable to find meta", esl.Error(err))
		return err
	}

	// 2. set un-sync
	updatedFinance, err := sv_teamfolder.New(z.Peer.Client()).UpdateSyncSetting(tf,
		sv_teamfolder.AddNestedSetting(folderOrganizationMeta, sv_teamfolder.SyncSettingNotSynced),
		sv_teamfolder.AddNestedSetting(folderFinanceMeta, sv_teamfolder.SyncSettingNotSynced),
	)
	if err != nil {
		l.Warn("Unable to change", esl.Error(err))
		return err
	}
	l.Info("Sync settings updated", esl.Any("updated", updatedFinance))

	// 3. set no_inherit
	updatedFinanceInherit, err := sv_sharedfolder.New(z.Peer.Client().AsMemberId(admin.TeamMemberId)).UpdateInheritance(folderFinance.SharedFolderId, sv_sharedfolder.AccessInheritanceNoInherit)
	if err != nil {
		l.Warn("Unable to change: inherit", esl.Error(err))
		return err
	}
	l.Info("Sync access inheritance updated", esl.Any("updated", updatedFinanceInherit))

	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>
	//       |
	//       +-- [Plan] (nested folder, no inherit) <<nestedFolderMarketingPlan>>

	folderMarketing, err := sv_teamfolder.New(z.Peer.Client()).Create(teamFolderMarketing)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("teamfolder", folderMarketing))
		break

	case de.IsFolderNameAlreadyUsed():
		l.Info("The folder already created")
		teamfolders, err := sv_teamfolder.New(z.Peer.Client()).List()
		if err != nil {
			l.Warn("Unable to retrieve team folder list", esl.Error(err))
			return err
		}

		for _, teamfolder := range teamfolders {
			if strings.ToLower(teamfolder.Name) == strings.ToLower(teamFolderName) {
				folderMarketing = teamfolder
				break
			}
		}
		if folderMarketing == nil {
			l.Warn("Team folder not found")
			return errors.New("team folder not found")
		}

		break

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	marketingCtx := z.Peer.Client().AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Namespace(folderMarketing.TeamFolderId))

	// Set acl_policy
	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>

	updatedMarketingPolicy, err := sv_sharedfolder.New(z.Peer.Client().AsAdminId(admin.TeamMemberId)).UpdatePolicy(
		folderMarketing.TeamFolderId,
		sv_sharedfolder.AclUpdatePolicy("owner"),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The sales folder policy successfully updated", esl.Any("updated", updatedMarketingPolicy))

	default:
		l.Warn("Unable to update policies", esl.Error(err))
	}

	// Set acl_policy
	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>
	// Add admin group to the team folder
	err = sv_sharedfolder_member.NewByTeamFolder(z.Peer.Client().AsAdminId(admin.TeamMemberId), folderMarketing).Add(
		sv_sharedfolder_member.AddByGroup(adminGroup, "editor"),
	)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("The admin group added to the team folder as editor")

	default:
		l.Warn("Unable to update members", esl.Error(err))
	}

	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>
	//       |
	//       +-- [Plan] (nested folder, no inherit) <<nestedFolderMarketingPlan>>      <=== create here

	folderMarketingPlan, err := sv_sharedfolder.New(marketingCtx).Create(mo_path.NewDropboxPath("/" + nestedFolderMarketingPlan))
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Team folder created", esl.Any("folder", folderMarketingPlan))
		break

	case de.BadPath().IsAlreadyShared():
		l.Info("The folder is already shared")
		folderMarketingPlanMeta, err := sv_file.NewFiles(marketingCtx).Resolve(mo_path.NewDropboxPath("/" + nestedFolderSharedName))
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}

		folderMarketingPlan, err = sv_sharedfolder.New(marketingCtx).Resolve(folderMarketingPlanMeta.Concrete().SharedFolderId)
		if err != nil {
			l.Warn("Unable to resolve nested folder", esl.Error(err))
			return err
		}
		l.Info("Nested folder resolved", esl.Any("folder", folderMarketingPlan))

	default:
		l.Warn("Unable to create team folder", esl.Error(err))
		return err
	}

	// Workaround: Turn acl_policy=editor
	// backup  acl_policy
	folderMarketingPlanMeta, err := sv_sharedfolder.New(z.Peer.Client().AsAdminId(admin.TeamMemberId)).Resolve(folderMarketingPlan.SharedFolderId)
	de = dbx_error.NewErrors(err)
	switch {
	case de == nil:
		l.Info("Acquired folder metadata", esl.Any("marketing/plan", folderMarketingPlanMeta),
			esl.String("acl_policy", folderMarketingPlanMeta.PolicyManageAccess))

	default:
		l.Warn("Unable to retrieve the folder metadata", esl.Error(err))
		return err
	}

	// Workaround: Turn acl_policy=editor
	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>
	//       |
	//       +-- [Plan] (nested folder, no inherit) <<nestedFolderMarketingPlan>>      <=== change it to acl_policy=editor
	policyFolderMarketingPlanMetaChanged := false
	if folderMarketingPlanMeta.PolicyManageAccess != "editors" {
		updated, err := sv_sharedfolder.New(z.Peer.Client().AsAdminId(admin.TeamMemberId)).UpdatePolicy(
			folderMarketingPlan.SharedFolderId,
			sv_sharedfolder.AclUpdatePolicy("editors"),
		)
		de = dbx_error.NewErrors(err)
		switch {
		case de == nil:
			l.Info("The admin group added to the team folder as editor", esl.Any("updated", updated))
			policyFolderMarketingPlanMetaChanged = true

		default:
			l.Warn("Unable to update members", esl.Error(err))
			return err
		}
	} else {
		l.Info("Skip policy update", esl.String("policy_manage_access", folderMarketingPlanMeta.PolicyManageAccess))
	}

	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>
	//       |
	//       +-- [Plan] (nested folder, no inherit) <<nestedFolderMarketingPlan>>      <=== set no inherit

	updatedFolderMarketingPlan, err := sv_sharedfolder.New(z.Peer.Client().AsMemberId(admin.TeamMemberId)).UpdateInheritance(folderMarketingPlan.SharedFolderId, sv_sharedfolder.AccessInheritanceNoInherit)
	if err != nil {
		l.Warn("Unable to change: inherit", esl.Error(err))
		return err
	}
	l.Info("Sync access inheritance updated", esl.Any("updated", updatedFolderMarketingPlan))

	// Workaround: restore acl_policy
	//  +-- [Marketing] (nested folder, acl_policy=owner only)  <<teamFolderMarketing>>
	//       |
	//       +-- [Plan] (nested folder, no inherit) <<nestedFolderMarketingPlan>>      <=== restore
	if policyFolderMarketingPlanMetaChanged {
		restored, err := sv_sharedfolder.New(z.Peer.Client().AsAdminId(admin.TeamMemberId)).UpdatePolicy(
			folderMarketingPlan.SharedFolderId,
			sv_sharedfolder.AclUpdatePolicy(folderMarketingPlanMeta.PolicyManageAccess),
		)
		de = dbx_error.NewErrors(err)
		switch {
		case de == nil:
			l.Info("The admin group added to the team folder as editor", esl.Any("restored", restored))

		default:
			l.Warn("Unable to update members", esl.Error(err))
			return err
		}
	}

	return nil
}

func (z *Teamfolder) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Teamfolder{}, rc_recipe.NoCustomValues)
}
