package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_size"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_traverse"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"sync"
)

type Size struct {
	Peer                dbx_conn.ConnBusinessFile
	IncludeSharedFolder bool
	IncludeTeamFolder   bool
	IncludeMemberFolder bool
	IncludeAppFolder    bool
	Folder              mo_filter.Filter
	Depth               mo_int.RangeInt
	NamespaceSize       rp_model.TransactionReport
	Errors              rp_model.TransactionReport
}

func (z *Size) Preset() {
	z.NamespaceSize.SetModel(
		&mo_namespace.Namespace{},
		&mo_file_size.NamespaceSize{},
		rp_model.HiddenColumns(
			"result.namespace_name",
			"result.namespace_id",
			"result.namespace_type",
			"result.owner_team_member_id",
			"input.team_member_id",
			"input.namespace_id",
		),
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
	z.Errors.SetModel(&uc_file_traverse.TraverseEntry{}, nil)
	z.IncludeSharedFolder = true
	z.IncludeTeamFolder = true
	z.Depth.SetRange(1, 300, 1)
}

func (z *Size) Exec(c app_control.Control) error {
	l := c.Log()

	if err := z.NamespaceSize.Open(); err != nil {
		return err
	}

	admin, err := sv_profile.NewTeam(z.Peer.Context()).Admin()
	if err != nil {
		return err
	}
	l.Debug("Run as admin", esl.Any("admin", admin))

	namespaces, err := sv_namespace.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	namespaceDict := make(map[string]*mo_namespace.Namespace)
	for _, ns := range namespaces {
		namespaceDict[ns.NamespaceId] = ns
	}

	if err := z.Errors.Open(); err != nil {
		return err
	}

	namespaceSizes := sync.Map{}
	for _, namespace := range namespaces {
		namespaceSizes.Store(namespace.NamespaceId, uc_file_size.NewSum(z.Depth.Value()))
	}

	cta := z.Peer.Context().AsAdminId(admin.TeamMemberId)

	handlerEntries := func(te uc_file_traverse.TraverseEntry, entries []mo_file.Entry) {
		if size, ok := namespaceSizes.Load(te.Namespace.NamespaceId); ok {
			s := size.(uc_file_size.Sum)
			s.Eval(te.Path, entries)
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
	)

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(traverseQueueId, traverse.Traverse, s)
		for _, namespace := range namespaces {
			process := false
			switch {
			case z.IncludeTeamFolder && namespace.NamespaceType == "team_folder":
				process = true
			case z.IncludeSharedFolder && namespace.NamespaceType == "shared_folder":
				process = true
			case z.IncludeMemberFolder && namespace.NamespaceType == "team_member_folder":
				process = true
			case z.IncludeAppFolder && namespace.NamespaceType == "app_folder":
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

			q := s.Get(traverseQueueId).Batch(namespace.NamespaceId)
			q.Enqueue(uc_file_traverse.TraverseEntry{
				Namespace: namespace,
				Path:      "/",
			})
		}
	})

	namespaceSizes.Range(func(key, value interface{}) bool {
		size := value.(uc_file_size.Sum)
		namespaceId := key.(string)
		namespace := namespaceDict[namespaceId]

		size.Each(func(path mo_path.DropboxPath, size mo_file_size.Size) {
			z.NamespaceSize.Success(namespace, size)
		})
		return true
	})

	return nil
}

func (z *Size) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &Size{}, func(r rc_recipe.Recipe) {
		rc := r.(*Size)
		rc.Folder.SetOptions(mo_filter.NewTestNameFilter(qtr_endtoend.TestTeamFolderName))
		rc.IncludeTeamFolder = false
	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "namespace_size", func(cols map[string]string) error {
		if _, ok := cols["input.namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		if _, ok := cols["result.size"]; !ok {
			return errors.New("`size` is not found")
		}
		return nil
	})
}
