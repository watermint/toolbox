package policy

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_content"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer     dbx_conn.ConnBusinessFile
	Metadata kv_storage.Storage
	Tree     kv_storage.Storage
	Policy   rp_model.RowReport
	Folder   mo_filter.Filter
}

type FolderPolicy struct {
	NamespaceId        string `json:"namespace_id"`
	NamespaceName      string `json:"namespace_name"`
	Path               string `json:"path"`
	IsTeamFolder       bool   `json:"is_team_folder"`
	OwnerTeamId        string `json:"owner_team_id"`
	OwnerTeamName      string `json:"owner_team_name"`
	PolicyManageAccess string `json:"policy_manage_access"`
	PolicySharedLink   string `json:"policy_shared_link"`
	PolicyMember       string `json:"policy_member"`
	PolicyViewerInfo   string `json:"policy_viewer_info"`
}

func (z *List) Preset() {
	z.Policy.SetModel(
		&FolderPolicy{},
		rp_model.HiddenColumns(
			"owner_team_id",
			"namespace_id",
			"namespace_name",
		),
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()

	q := c.NewQueue()
	s := &uc_team_content.TeamScanner{
		Ctx:    z.Peer.Context(),
		Ctl:    c,
		Queue:  q,
		Filter: z.Folder,
		Scanner: &uc_team_content.ScanNamespaceMetadata{
			Metadata: z.Metadata,
			Queue:    q,
		},
	}
	if err := s.Scan(); err != nil {
		return err
	}
	q.Wait()

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}

	st := &uc_team_content.TeamFolderScanner{
		Ctl:      c,
		Ctx:      z.Peer.Context().AsAdminId(admin.TeamMemberId),
		Metadata: z.Metadata,
		Tree:     z.Tree,
	}
	if err := st.Scan(); err != nil {
		return err
	}
	if err := z.Policy.Open(); err != nil {
		return err
	}
	return z.Tree.View(func(treeKvs kv_kvs.Kvs) error {
		return treeKvs.ForEachModel(&uc_team_content.Tree{}, func(key string, m interface{}) error {
			t := m.(*uc_team_content.Tree)
			ll := l.With(esl.String("nsid", t.NamespaceId))
			ll.Debug("Preparing for report")
			meta := &mo_sharedfolder.SharedFolder{}
			err := z.Metadata.View(func(metaKvs kv_kvs.Kvs) error {
				return metaKvs.GetJsonModel(t.NamespaceId, meta)
			})
			if err != nil {
				ll.Debug("Unable to get metadata for the namespace", esl.Error(err))
				return err
			}
			z.Policy.Row(&FolderPolicy{
				NamespaceId:        t.NamespaceId,
				NamespaceName:      t.NamespaceName,
				Path:               t.RelativePath,
				IsTeamFolder:       meta.IsTeamFolder || meta.IsInsideTeamFolder,
				OwnerTeamId:        meta.OwnerTeamId,
				OwnerTeamName:      meta.OwnerTeamName,
				PolicyManageAccess: meta.PolicyManageAccess,
				PolicySharedLink:   meta.PolicySharedLink,
				PolicyMember:       meta.PolicyMember,
				PolicyViewerInfo:   meta.PolicyViewerInfo,
			})
			return nil
		})
	})
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
