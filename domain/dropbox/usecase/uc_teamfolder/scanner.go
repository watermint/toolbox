package uc_teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
)

type TeamFolder struct {
	// team folder metadata
	TeamFolder *mo_sharedfolder.SharedFolder `json:"team_folder"`

	// relative path to nested folder
	NestedFolders map[string]*mo_sharedfolder.SharedFolder `json:"nested_folders"`
}

type TeamFolderNested struct {
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	RelativePath  string `json:"relative_path"`
}

type TeamFolderEntry struct {
	NamespaceId string   `json:"namespace_id"`
	Descendants []string `json:"descendants"`
}

const (
	queueIdScanTeamNamespace     = "scan_team"
	queueIdScanNamespaceMetadata = "scan_namespace"
	queueIdScanTeamFolder        = "scan_teamfolder"
)

type Scanner interface {
	Scan(filter mo_filter.Filter) (teamFolders []*TeamFolder, err error)
}

func New(ctl app_control.Control, ctx dbx_context.Context) Scanner {
	return &scanImpl{
		ctl: ctl,
		ctx: ctx,
	}
}

type scanImpl struct {
	ctl app_control.Control
	ctx dbx_context.Context
}

func (z scanImpl) scanNamespace(sessionId string, stg eq_sequence.Stage, storageNamespace kv_storage.Storage) (err error) {
	l := z.ctl.Log().With(esl.String("sessionId", sessionId))
	l.Debug("Scan namespace")
	q := stg.Get(queueIdScanNamespaceMetadata)
	var lastErr error
	err = sv_namespace.New(z.ctx).ListEach(func(namespace *mo_namespace.Namespace) bool {
		l.Debug("Namespace", esl.Any("namespace", namespace))
		switch namespace.NamespaceType {
		case "app_folder", "team_member_folder":
			l.Debug("Ignore member folders")
		default:
			q.Enqueue(namespace)
			err0 := storageNamespace.Update(func(kvs kv_kvs.Kvs) error {
				return kvs.PutJsonModel(namespace.NamespaceId, namespace)
			})
			if err0 != nil {
				lastErr = err0
			}
		}
		return true
	})
	if err == nil && lastErr != nil {
		err = lastErr
	}
	l.Debug("Scan namespace: done", esl.Error(err))

	return err
}

func (z scanImpl) scanNamespaceMetadata(namespace *mo_namespace.Namespace, storageMeta kv_storage.Storage, admin *mo_profile.Profile) error {
	l := z.ctl.Log().With(esl.Any("namespace", namespace))
	l.Debug("Scan namespace")

	meta, err := sv_sharedfolder.New(z.ctx.AsAdminId(admin.TeamMemberId)).Resolve(namespace.NamespaceId)
	if err != nil {
		l.Debug("Unable to retrieve metadata", esl.Error(err))
		return err
	}

	return storageMeta.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(namespace.NamespaceId, meta)
	})
}

func (z scanImpl) scanTeamFolder(teamfolder *TeamFolderEntry, storageMeta, storageNested kv_storage.Storage, admin *mo_profile.Profile) (err error) {
	l := z.ctl.Log().With(esl.Any("teamfolder", teamfolder))
	l.Debug("Scan team folder")
	teamFolderName := ""
	l.Debug("lookup team folder name")

	err = storageMeta.View(func(kvs kv_kvs.Kvs) error {
		tf := &mo_sharedfolder.SharedFolder{}
		if err := kvs.GetJsonModel(teamfolder.NamespaceId, tf); err != nil {
			return err
		}
		teamFolderName = tf.Name
		return nil
	})
	if err != nil {
		return err
	}
	l = l.With(esl.String("teamFolderName", teamFolderName))

	err = storageNested.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(teamfolder.NamespaceId, &TeamFolderNested{
			NamespaceId:   teamfolder.NamespaceId,
			NamespaceName: teamFolderName,
			RelativePath:  mo_path.NewDropboxPath(teamFolderName).Path(),
		})
	})
	if err != nil {
		return err
	}

	l.Debug("search nested folders", esl.Strings("descendants", teamfolder.Descendants))
	traverse := make(map[string]bool)
	for _, d := range teamfolder.Descendants {
		traverse[d] = false
	}
	completed := func() bool {
		for _, t := range traverse {
			if !t {
				return false
			}
		}
		return true
	}
	ErrorScanCompleted := errors.New("scan completed")

	ctx := z.ctx.WithPath(dbx_context.Namespace(teamfolder.NamespaceId)).AsAdminId(admin.TeamMemberId)
	var scan func(path mo_path.DropboxPath) error
	scan = func(path mo_path.DropboxPath) error {
		entries, err := sv_file.NewFiles(ctx).List(path)
		if err != nil {
			l.Debug("Unable to retrieve entries", esl.Error(err), esl.String("path", path.Path()))
			return err
		}

		// Mark nested folders
		for _, entry := range entries {
			if f, ok := entry.Folder(); ok {
				if f.EntrySharedFolderId != "" {
					traverse[f.EntrySharedFolderId] = true
					rp := path.ChildPath(f.EntryName)
					err := storageNested.Update(func(kvs kv_kvs.Kvs) error {
						return kvs.PutJsonModel(f.EntrySharedFolderId, &TeamFolderNested{
							NamespaceId:   f.EntrySharedFolderId,
							NamespaceName: f.EntryName,
							RelativePath:  teamFolderName + rp.Path(),
						})
					})
					if err != nil {
						return err
					}
				}
			}
		}

		// Return if the scan completed
		if completed() {
			return ErrorScanCompleted
		}

		// Dive into descendants
		for _, entry := range entries {
			if f, ok := entry.Folder(); ok {
				if err := scan(path.ChildPath(f.Name())); err != nil {
					return err
				}
			}
		}
		return nil
	}

	if errScan := scan(mo_path.NewDropboxPath("")); errScan != nil && errScan != ErrorScanCompleted {
		l.Debug("The error occurred on scanning team folder", esl.Error(err))
		return err
	}

	l.Debug("Scan completed")

	return
}

func (z scanImpl) parentChildRelationship(storageMeta kv_storage.Storage) (relation map[string]string, err error) {
	l := z.ctl.Log()
	l.Debug("Making mapping of parent-child relationship")

	// namespace_id -> parent namespace_id
	relation = make(map[string]string)

	err = storageMeta.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&mo_sharedfolder.SharedFolder{}, func(key string, m interface{}) error {
			ns := m.(*mo_sharedfolder.SharedFolder)
			relation[ns.SharedFolderId] = ns.ParentSharedFolderId
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	l.Debug("Relation", esl.Any("relation", relation))

	return relation, nil
}

func (z scanImpl) namespaceToTopNamespaceId(storageMeta kv_storage.Storage) (top map[string]string, err error) {
	l := z.ctl.Log()
	// namespace_id -> top level namespace_id
	top = make(map[string]string)

	relation, err := z.parentChildRelationship(storageMeta)
	if err != nil {
		return nil, err
	}

	l.Debug("Making child - top level folder namespace mapping")
	for nsid := range relation {
		ll := l.With(esl.String("namespace_id", nsid))
		top[nsid] = ""
		chain := make([]string, 0)
		parent := relation[nsid]
		current := parent
		for {
			chain = append(chain, parent)

			var ok bool
			current, ok = relation[parent]
			if current == "" || !ok {
				break
			}
			parent = current
		}
		top[nsid] = parent
		ll.Debug("Top folder", esl.String("top", parent))
	}

	return top, nil
}

func (z scanImpl) nestedFolderNamespaceIds(storageMeta kv_storage.Storage) (nested map[string][]string, others []string, err error) {
	l := z.ctl.Log()

	// team_folder's namespace_id -> array of nested team folder namespace_id
	nested = make(map[string][]string)

	// other un-nested namespaces
	others = make([]string, 0)

	top, err := z.namespaceToTopNamespaceId(storageMeta)
	if err != nil {
		return nil, nil, err
	}

	l.Debug("Aggregate nested folders")
	for nsid, t := range top {
		if t == "" {
			others = append(others, nsid)
			continue
		}
		if _, ok := nested[t]; !ok {
			nested[t] = make([]string, 0)
		}
		nested[t] = append(nested[t], nsid)
	}

	l.Debug("Team folders and nested folders", esl.Any("nested", nested))
	l.Debug("Others", esl.Strings("others", others))

	return nested, others, nil
}

func (z scanImpl) Scan(filter mo_filter.Filter) (teamFolders []*TeamFolder, err error) {
	l := z.ctl.Log()
	scanSessionId := sc_random.MustGenerateRandomString(8)
	teamFolders = make([]*TeamFolder, 0)

	admin, err := sv_profile.NewTeam(z.ctx).Admin()
	if err != nil {
		l.Debug("Unable to identify admin")
		return nil, err
	}
	storageNamespace, err := z.ctl.NewKvs("namespace_" + scanSessionId)
	if err != nil {
		l.Debug("Unable to create temporary storage", esl.Error(err))
		return nil, err
	}
	defer storageNamespace.Close()

	storageMeta, err := z.ctl.NewKvs("teamfolder_" + scanSessionId)
	if err != nil {
		l.Debug("Unable to create temporary storage", esl.Error(err))
		return nil, err
	}
	defer storageMeta.Close()

	storageNested, err := z.ctl.NewKvs("nested_" + scanSessionId)
	if err != nil {
		l.Debug("Unable to create temporary storage", esl.Error(err))
		return nil, err
	}
	defer storageNested.Close()

	secondStage := z.ctl.Sequence().DoThen(func(s eq_sequence.Stage) {
		s.Define(queueIdScanTeamNamespace, z.scanNamespace, s, storageNamespace)
		s.Define(queueIdScanNamespaceMetadata, z.scanNamespaceMetadata, storageMeta, admin)
		q := s.Get(queueIdScanTeamNamespace)
		q.Enqueue(scanSessionId)
	})

	var nested map[string][]string
	nested, _, err = z.nestedFolderNamespaceIds(storageMeta)
	if err != nil {
		return
	}

	secondStage.Do(func(s eq_sequence.Stage) {
		s.Define(queueIdScanTeamFolder, z.scanTeamFolder, storageMeta, storageNested, admin)
		q := s.Get(queueIdScanTeamFolder)

		for nsid, descendants := range nested {
			meta := &mo_sharedfolder.SharedFolder{}
			err = storageNamespace.View(func(kvs kv_kvs.Kvs) error {
				return kvs.GetJsonModel(nsid, meta)
			})
			if filter.Accept(meta.Name) {
				q.Enqueue(&TeamFolderEntry{
					NamespaceId: nsid,
					Descendants: descendants,
				})
			}
		}
	})

	for nsid, descendants := range nested {
		teamFolderMeta := &mo_sharedfolder.SharedFolder{}
		err = storageMeta.View(func(kvs kv_kvs.Kvs) error {
			return kvs.GetJsonModel(nsid, teamFolderMeta)
		})
		if err != nil {
			l.Debug("Unable to retrieve team folder data", esl.Error(err))
			return nil, err
		}

		descendantMetadata := make(map[string]*mo_sharedfolder.SharedFolder)
		for _, descendant := range descendants {
			meta := &mo_sharedfolder.SharedFolder{}
			err = storageMeta.View(func(kvs kv_kvs.Kvs) error {
				return kvs.GetJsonModel(descendant, meta)
			})
			if err != nil {
				l.Debug("Unable to retrieve descendant data", esl.Error(err))
				return nil, err
			}

			n := &TeamFolderNested{}
			err = storageNested.View(func(kvs kv_kvs.Kvs) error {
				return kvs.GetJsonModel(descendant, n)
			})
			if err != nil {
				l.Debug("Unable to retrieve descendant relative path", esl.Error(err))
				return nil, err
			}

			descendantMetadata[n.RelativePath] = meta
		}

		teamFolder := &TeamFolder{
			TeamFolder:    teamFolderMeta,
			NestedFolders: descendantMetadata,
		}

		teamFolders = append(teamFolders, teamFolder)
	}

	return
}
