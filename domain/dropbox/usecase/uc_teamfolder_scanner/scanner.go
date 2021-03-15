package uc_teamfolder_scanner

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"math"
	"path/filepath"
	"time"
)

var (
	scanShortTimeout = 3 * time.Minute
	scanLongTimeout  = 3 * time.Hour
)

const (
	ScanTimeoutShort   ScanTimeoutMode = "short"
	ScanTimeoutLong    ScanTimeoutMode = "long"
	ScanTimeoutAltPath                 = "/:ERROR-SCAN-TIMEOUT:/"
)

type ScanTimeoutMode string

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
	queueIdExtractTeamFolder     = "extract_teamfolder"
)

type Scanner interface {
	Scan(filter mo_filter.Filter) (teamFolders []*TeamFolder, err error)
}

func New(ctl app_control.Control, ctx dbx_context.Context, scanTimeout ScanTimeoutMode) Scanner {
	return &scanImpl{
		ctl:         ctl,
		ctx:         ctx,
		scanTimeout: scanTimeout,
	}
}

type scanImpl struct {
	ctl         app_control.Control
	ctx         dbx_context.Context
	scanTimeout ScanTimeoutMode
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

	traverse := make(map[string]bool)
	deleted := make(map[string]bool)
	for _, d := range teamfolder.Descendants {
		traverse[d] = false
		deleted[d] = false
	}
	completed := func() bool {
		for _, t := range traverse {
			if !t {
				return false
			}
		}
		return true
	}

	integrityTest := func() {
		// Integrity test
		for _, descendant := range teamfolder.Descendants {
			nested := &TeamFolderNested{}
			err0 := storageNested.View(func(kvs kv_kvs.Kvs) error {
				return kvs.GetJsonModel(descendant, nested)
			})
			if err0 != nil {
				if deleted[descendant] {
					l.Debug("Deleted descendant", esl.String("descendant", descendant))
				} else {
					l.Debug("Not found", esl.String("teamfolderId", teamfolder.NamespaceId),
						esl.String("namespaceId", descendant))
				}
			}
		}
	}

	defer integrityTest()

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
		l.Debug("Unable to store team folder data")
		return err
	}

	l.Debug("Looking for descendant folder status")
	for _, d := range teamfolder.Descendants {
		ll := l.With(esl.String("descendant", d))
		ll.Debug("Retrieve file info")
		info, err := sv_file.NewFiles(z.ctx.AsAdminId(admin.TeamMemberId)).Resolve(
			mo_path.NewDropboxPath("ns:"+d),
			sv_file.ResolveIncludeDeleted(true),
		)
		if err != nil {
			dbxErr := dbx_error.NewErrors(err)
			if dbxErr.Path().IsNotFound() {
				ll.Debug("Assuming the folder is deleted", esl.String("response", dbxErr.Summary()))
				traverse[d] = true
				deleted[d] = true
			} else {
				ll.Debug("Unable to retrieve folder info", esl.Error(err))
			}
			continue
		}
		if _, ok := info.Deleted(); ok {
			ll.Debug("The folder is deleted, mark as traversed", esl.Any("deleted", info))
			traverse[d] = true
			deleted[d] = true
		}
		if f, ok := info.Folder(); ok {
			ll.Debug("Folder info found", esl.Any("folderInfo", f))
		}
	}

	// looking for team member id who have access to the team folder.
	scanBySearch := func() {
		if len(teamfolder.Descendants) < 1 {
			return
		}

		l.Debug("Looking for root folder members")
		rootMemberTeamMemberId := ""
		rootMembers, err := sv_sharedfolder_member.NewBySharedFolderId(z.ctx.AsAdminId(admin.TeamMemberId), teamfolder.NamespaceId).List()
		if err != nil {
			l.Debug("Unable to retrieve root folder members", esl.Error(err))
			return
		}

		l.Debug("Scan folder members")
		for _, member := range rootMembers {
			if u, ok := member.User(); ok {
				if u.TeamMemberId != "" && u.IsSameTeam {
					l.Debug("Team member found", esl.String("teamMemberId", u.TeamMemberId))
					rootMemberTeamMemberId = u.TeamMemberId
					break
				}
			}
			if g, ok := member.Group(); ok {
				groupMembers, err := sv_group_member.NewByGroupId(z.ctx, g.GroupId).List()
				if err != nil {
					l.Debug("Unable to retrieve group members", esl.Error(err))
					continue
				}
				for _, gm := range groupMembers {
					if gm.TeamMemberId != "" {
						l.Debug("Team member found", esl.String("teamMemberId", gm.TeamMemberId))
						rootMemberTeamMemberId = gm.TeamMemberId
						break
					}
				}
			}
		}

		if rootMemberTeamMemberId == "" {
			l.Debug("no root member found, skip")
			return
		}

		ctx := z.ctx.WithPath(dbx_context.Namespace(teamfolder.NamespaceId)).AsMemberId(rootMemberTeamMemberId).NoRetry()

		for _, descendantNamespaceId := range teamfolder.Descendants {
			ll := l.With(esl.String("descendant", descendantNamespaceId))
			ll.Debug("Retrieve metadata for search")
			descendant := &mo_sharedfolder.SharedFolder{}
			err = storageMeta.View(func(kvs kv_kvs.Kvs) error {
				return kvs.GetJsonModel(descendantNamespaceId, descendant)
			})
			if err != nil {
				ll.Debug("Unable to unmarshal", esl.Error(err))
				continue
			}

			ll.Debug("Search")
			matches, err := sv_file.NewFiles(ctx).Search(descendant.Name, sv_file.SearchFileNameOnly(), sv_file.SearchCategories("folder"))
			if err != nil {
				ll.Debug("Unable to search", esl.Error(err))
				continue
			}

			for _, match := range matches {
				entry := match.Concrete()
				if entry.SharedFolderId == descendantNamespaceId {
					ll.Debug("Descendant found")
					traverse[descendant.SharedFolderId] = true
					err := storageNested.Update(func(kvs kv_kvs.Kvs) error {
						tfn := &TeamFolderNested{
							NamespaceId:   descendantNamespaceId,
							NamespaceName: entry.Name,
							RelativePath:  filepath.Join(teamFolderName, entry.PathDisplay),
						}
						return kvs.PutJsonModel(descendantNamespaceId, tfn)
					})
					if err != nil {
						ll.Debug("Unable to store search result", esl.Error(err))
					}
					break
				}
			}
		}
	}

	ErrorScanCompleted := errors.New("scan completed")
	ErrorScanTimeout := errors.New("scan timeout")

	timeout := time.Now()
	switch z.scanTimeout {
	case ScanTimeoutShort:
		timeout = timeout.Add(scanShortTimeout)
	case ScanTimeoutLong:
		timeout = timeout.Add(scanLongTimeout)
	}

	scanDeferred := make([]mo_path.DropboxPath, 0)

	var scan func(path mo_path.DropboxPath, depth, maxDepth int) error
	scan = func(path mo_path.DropboxPath, depth, maxDepth int) error {
		ll := l.With(esl.String("path", path.Path()), esl.Int("depth", depth), esl.Int("maxDepth", maxDepth))
		if maxDepth < depth {
			ll.Debug("Defer scan", esl.String("path", path.Path()))
			scanDeferred = append(scanDeferred, path)
			return nil
		}

		ll.Debug("Scan path")
		ctx := z.ctx.WithPath(dbx_context.Namespace(teamfolder.NamespaceId)).AsAdminId(admin.TeamMemberId)
		entries, err := sv_file.NewFiles(ctx).List(path)
		if err != nil {
			l.Debug("Unable to retrieve entries", esl.Error(err), esl.String("path", path.Path()))
			return err
		}
		ll.Debug("Entries", esl.Any("entries", entries))

		// Mark nested folders
		for _, entry := range entries {
			if f, ok := entry.Folder(); ok {
				if f.EntrySharedFolderId != "" {
					lll := ll.With(esl.Any("folder", f))
					lll.Debug("Descendant found")
					traverse[f.EntrySharedFolderId] = true
					rp := path.ChildPath(f.EntryName)
					err := storageNested.Update(func(kvs kv_kvs.Kvs) error {
						tfn := &TeamFolderNested{
							NamespaceId:   f.EntrySharedFolderId,
							NamespaceName: f.EntryName,
							RelativePath:  teamFolderName + rp.Path(),
						}
						return kvs.PutJsonModel(f.EntrySharedFolderId, tfn)
					})
					if err != nil {
						lll.Debug("Unable to store", esl.Error(err))
						return err
					}
				}
			}
		}

		// Return if the scan completed
		if completed() {
			ll.Debug("Scan completed")
			return ErrorScanCompleted
		}

		// Return if the time exceed timeout
		if time.Now().After(timeout) {
			ll.Debug("Scan timeout")
			return ErrorScanTimeout
		}

		// Dive into descendants
		for _, entry := range entries {
			if f, ok := entry.Folder(); ok {
				ll.Debug("Dive into descendant", esl.Any("folder", f))
				if err := scan(path.ChildPath(f.Name()), depth+1, maxDepth); err != nil {
					ll.Debug("Got an error", esl.Error(err))
					return err
				}
			}
		}
		return nil
	}

	handleScanResult := func(errScan error) error {
		switch errScan {
		case nil, ErrorScanCompleted:
			l.Debug("Scan finished")
			return nil

		case ErrorScanTimeout:
			l.Debug("Scan timeout")
			for nsid, ok := range traverse {
				ll := l.With(esl.String("descendant", nsid))
				ll.Debug("Ensure descendants")
				if ok {
					ll.Debug("Skip traversed folder")
					continue
				}

				ll.Debug("Retrieve shared folder metadata")
				sf, err := sv_sharedfolder.New(z.ctx).Resolve(nsid)
				if err != nil {
					ll.Debug("Unable to retrieve folder metadata", esl.Error(err))
					return err
				}

				ll.Debug("Store shared folder data", esl.Any("sf", sf))
				err = storageNested.Update(func(kvs kv_kvs.Kvs) error {
					return kvs.PutJsonModel(sf.SharedFolderId, &TeamFolderNested{
						NamespaceId:   sf.SharedFolderId,
						NamespaceName: sf.Name,
						RelativePath:  teamFolderName + ScanTimeoutAltPath + sf.Name,
					})
				})
				if err != nil {
					ll.Debug("Unable to store data", esl.Error(err))
					return err
				}
			}
			return nil

		default:
			l.Debug("The error occurred on scanning team folder", esl.Error(err))
			return err
		}
	}

	// Finish if there is no descendants
	if completed() {
		l.Debug("No descendants found")
		return nil
	}

	l.Debug("search nested folders", esl.Strings("descendants", teamfolder.Descendants), esl.Time("timeout", timeout))

	// Scan first two levels
	if errScan := handleScanResult(scan(mo_path.NewDropboxPath(""), 1, 2)); errScan != nil {
		l.Debug("Unable to scan", esl.Error(errScan))
		return errScan
	}

	// If not found in first two levels, search
	if !completed() {
		l.Debug("scan by search")
		scanBySearch()
	}

	// If still not found, search the entire tree
	if !completed() {
		l.Debug("scan deferred", esl.Int("deferred", len(scanDeferred)))
		for _, dp := range scanDeferred {
			l.Debug("Scanning deferred path", esl.String("path", dp.Path()))
			if errScan := handleScanResult(scan(dp, 1, math.MaxInt32)); errScan != nil {
				l.Debug("The error occurred on scanning team folder", esl.Error(err))
				return err
			}
		}
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
func (z scanImpl) extractTeamFolder(entry *TeamFolderEntry, storageMeta, storageNested, storageTeamFolder kv_storage.Storage) error {
	l := z.ctl.Log().With(esl.String("nsid", entry.NamespaceId), esl.Strings("descendants", entry.Descendants))
	l.Debug("Compose TeamFolder object")
	teamFolderMeta := &mo_sharedfolder.SharedFolder{}
	err := storageMeta.View(func(kvs kv_kvs.Kvs) error {
		return kvs.GetJsonModel(entry.NamespaceId, teamFolderMeta)
	})
	if err != nil {
		l.Debug("Unable to retrieve team folder data", esl.Error(err))
		return err
	}

	descendantMetadata := make(map[string]*mo_sharedfolder.SharedFolder)
	for _, descendant := range entry.Descendants {
		ll := l.With(esl.String("descendant", descendant))
		ll.Debug("Compose Descendant")
		meta := &mo_sharedfolder.SharedFolder{}
		err = storageMeta.View(func(kvs kv_kvs.Kvs) error {
			return kvs.GetJsonModel(descendant, meta)
		})
		if err != nil {
			ll.Debug("Unable to retrieve descendant data, skip", esl.Error(err))
			continue
		}

		n := &TeamFolderNested{}
		err = storageNested.View(func(kvs kv_kvs.Kvs) error {
			return kvs.GetJsonModel(descendant, n)
		})
		if err != nil {
			ll.Debug("Unable to retrieve descendant relative path, assuming the folder is deleted", esl.Error(err))
			continue
		}

		descendantMetadata[n.RelativePath] = meta
	}

	teamFolder := &TeamFolder{
		TeamFolder:    teamFolderMeta,
		NestedFolders: descendantMetadata,
	}

	return storageTeamFolder.Update(func(kvs kv_kvs.Kvs) error {
		return kvs.PutJsonModel(entry.NamespaceId, teamFolder)
	})
}

func (z scanImpl) Scan(filter mo_filter.Filter) (teamFolders []*TeamFolder, err error) {
	l := z.ctl.Log()
	scanSessionId := sc_random.MustGetSecureRandomString(8)

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

	storageMeta, err := z.ctl.NewKvs("meta_" + scanSessionId)
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

	storageTeamFolder, err := z.ctl.NewKvs("teamfolder_" + scanSessionId)
	if err != nil {
		l.Debug("Unable to create temporary storage", esl.Error(err))
		return nil, err
	}
	defer storageTeamFolder.Close()

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

	thirdStage := secondStage.DoThen(func(s eq_sequence.Stage) {
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

	thirdStage.Do(func(s eq_sequence.Stage) {
		s.Define(queueIdExtractTeamFolder, z.extractTeamFolder, storageMeta, storageNested, storageTeamFolder)
		q := s.Get(queueIdExtractTeamFolder)
		for nsid, descendants := range nested {
			q.Enqueue(&TeamFolderEntry{
				NamespaceId: nsid,
				Descendants: descendants,
			})
		}
	})

	teamFolders = make([]*TeamFolder, 0)
	err = storageTeamFolder.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&TeamFolder{}, func(key string, m interface{}) error {
			f := m.(*TeamFolder)
			teamFolders = append(teamFolders, f)
			return nil
		})
	})
	return
}
