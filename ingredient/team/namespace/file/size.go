package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_size"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_file_size"
	"github.com/watermint/toolbox/essentials/file/es_size"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage_impl"
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
	NamespaceSize       rp_model.RowReport
}

func (z *Size) Preset() {
	z.NamespaceSize.SetModel(
		&mo_file_size.NamespaceSize{},
	)
	z.Folder.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_filter.NewNameSuffixFilter(),
	)
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

	namespaceSizes := sync.Map{}
	for _, namespace := range namespaces {
		namespaceSizes.Store(namespace.NamespaceId, uc_file_size.NewSum(z.Depth.Value()))
	}

	cta := z.Peer.Context().AsAdminId(admin.TeamMemberId)

	scanFolderQueueId := "scan_folder"
	scanSessionQueueId := "scan_session"

	factory := kv_storage_impl.NewFactory(c)
	defer func() {
		factory.Close()
	}()

	sizeCtx := es_size.New(c.Log(), scanFolderQueueId, factory)

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(scanFolderQueueId, es_size.ScanFolder, sizeCtx)
		s.Define(scanSessionQueueId, sizeCtx.StartSession, s)

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

			dbxCtx := cta.WithPath(dbx_context.Namespace(namespace.NamespaceId))
			dbxFs := filesystem.NewFileSystem(dbxCtx)
			sessionId := namespace.NamespaceId

			sizeCtx.New(
				sessionId,
				filesystem.NewPath(namespace.NamespaceId, mo_path.NewDropboxPath("/")),
				s,
				dbxFs,
				namespace,
			)

			s.Get(scanSessionQueueId).Batch(namespace.NamespaceId).Enqueue(sessionId)
		}
	})

	return sizeCtx.ListEach(z.Depth.Value(), func(sessionId string, meta interface{}, size es_size.FolderSize) {
		ns, ok := meta.(*mo_namespace.Namespace)
		if !ok {
			l.Debug("Unable to cast to namespace")
			return
		}

		z.NamespaceSize.Row(&mo_file_size.NamespaceSize{
			NamespaceName:     ns.Name,
			NamespaceId:       ns.NamespaceId,
			NamespaceType:     ns.NamespaceType,
			OwnerTeamMemberId: ns.TeamMemberId,
			Path:              size.Path,
			CountFile:         size.NumFile,
			CountFolder:       size.NumFolder,
			CountDescendant:   size.NumFile + size.NumFolder,
			Size:              size.Size,
			Depth:             size.Depth,
			ModTimeEarliest:   size.ModTimeEarliest,
			ModTimeLatest:     size.ModTimeLatest,
			ApiComplexity:     size.OperationalComplexity,
		})
	})
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
