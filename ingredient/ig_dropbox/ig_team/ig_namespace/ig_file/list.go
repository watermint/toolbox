package ig_file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_traverse"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer                dbx_conn.ConnScopedTeam
	IncludeDeleted      bool
	IncludeMemberFolder bool
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	NamespaceFile       rp_model.RowReport
	Folder              mo_filter.Filter
	Errors              rp_model.TransactionReport
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeTeamDataMember,
		dbx_auth.ScopeTeamDataTeamSpace,
		dbx_auth.ScopeTeamInfoRead,
	)
	z.IncludeTeamFolder = true
	z.IncludeSharedFolder = true
	z.IncludeMemberFolder = false
	z.NamespaceFile.SetModel(
		&mo_namespace.NamespaceEntry{},
		rp_model.HiddenColumns(
			"namespace_id",
			"file_id",
			"revision",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
	z.Errors.SetModel(&uc_file_traverse.TraverseEntry{}, nil)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
}

func (z *List) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.NamespaceFile.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Client()).Admin()
	if err != nil {
		return err
	}
	l.Debug("Run as admin", esl.Any("admin", admin))

	members, err := sv_member.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}
	idToMember := mo_member.MapByTeamMemberId(members)

	namespaces, err := sv_namespace.New(z.Peer.Client()).List()
	if err != nil {
		return err
	}

	if err := z.Errors.Open(); err != nil {
		return err
	}

	cta := z.Peer.Client().AsAdminId(admin.TeamMemberId)

	handlerEntries := func(te uc_file_traverse.TraverseEntry, entries []mo_file.Entry) {
		for _, entry := range entries {
			ne := mo_namespace.NewNamespaceEntry(te.Namespace, entry.Concrete())
			if m, e := idToMember[te.Namespace.TeamMemberId]; e {
				ne.NamespaceMemberEmail = m.Email
			}
			z.NamespaceFile.Row(ne)
		}
	}
	handlerError := func(te uc_file_traverse.TraverseEntry, err error) {
		z.Errors.Failure(err, &te)
	}

	traverseQueueId := "namespace"
	traverse := uc_file_traverse.NewTraverse(
		cta,
		c,
		traverseQueueId,
		handlerEntries,
		handlerError,
		sv_file.IncludeDeleted(z.IncludeDeleted),
		sv_file.IncludeHasExplicitSharedMembers(true),
	)

	c.Sequence().Do(func(stage eq_sequence.Stage) {
		stage.Define(traverseQueueId, traverse.Traverse, stage)
		q := stage.Get(traverseQueueId)
		for _, namespace := range namespaces {
			process := false
			switch {
			case z.IncludeTeamFolder && namespace.NamespaceType == "team_folder":
				process = true
			case z.IncludeSharedFolder && namespace.NamespaceType == "shared_folder":
				process = true
			case z.IncludeMemberFolder && namespace.NamespaceType == "team_member_folder":
				process = true
			case z.IncludeMemberFolder && namespace.NamespaceType == "team_member_root":
				process = true
			case z.IncludeMemberFolder && namespace.NamespaceType == "app_folder":
				process = true
			}
			if !process {
				l.Debug("Skip", esl.Any("namespace", namespace))
				continue
			}
			if !z.Folder.Accept(namespace.Name) {
				l.Debug("Skip", esl.Any("namespace", namespace))
				continue
			}

			qb := q.Batch(namespace.NamespaceId)
			qb.Enqueue(uc_file_traverse.TraverseEntry{
				Namespace: namespace,
				Path:      "/",
			})
		}
	})

	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		rc := r.(*List)
		rc.Folder.SetOptions(
			mo_filter.NewTestNameFilter(qtr_endtoend.TestTeamFolderName),
		)
	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "namespace_file", func(cols map[string]string) error {
		if _, ok := cols["namespace_name"]; !ok {
			return errors.New("`namespace_name` is not found")
		}
		if _, ok := cols["path_display"]; !ok {
			return errors.New("`path_display` is not found")
		}
		return nil
	})
}
