package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_namespace"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_namespace"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder_mount"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharing"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"gorm.io/gorm"
	"reflect"
)

type IndividualScanner interface {
	// ScanCurrentUser scans current user files and sharing information
	ScanCurrentUser() (err error)
}

type TeamScanner interface {
	// ScanTeam scans all team information
	ScanTeam() (err error)

	RetryErrors() (err error)
}

func newDatabase(ctl app_control.Control, path string) (*gorm.DB, error) {
	l := ctl.Log().With(esl.String("path", path))
	db, err := ctl.NewOrm(path)
	if err != nil {
		return nil, err
	}

	tables := []interface{}{
		&GroupMember{},
		&Group{},
		&Member{},
		&Mount{},
		&NamespaceDetail{},
		&NamespaceEntryError{},
		&NamespaceEntry{},
		&NamespaceMember{},
		&Namespace{},
		&ReceivedFile{},
		&SharedLink{},
		&TeamFolder{},
	}

	for _, t := range tables {
		tableName := reflect.ValueOf(t).Elem().Type().Name()
		l.Debug("Migrating", esl.String("table", tableName))
		if err = db.AutoMigrate(t); err != nil {
			l.Debug("Unable to migrate", esl.Error(err), esl.String("table", tableName))
			return nil, err
		}
	}

	return db, nil
}
func NewTeamScanner(ctl app_control.Control, client dbx_client.Client, path string) (TeamScanner, error) {
	db, err := newDatabase(ctl, path)
	if err != nil {
		return nil, err
	}
	return &tsImpl{
		ctl:    ctl,
		client: client,
		db:     db,
	}, nil
}

const (
	teamScanQueueGroup           = "scan_group"
	teamScanQueueGroupMember     = "scan_group_member"
	teamScanQueueMember          = "scan_member"
	teamScanQueueMount           = "scan_mount"
	teamScanQueueNamespace       = "scan_team_namespace"
	teamScanQueueNamespaceDetail = "scan_namespace"
	teamScanQueueNamespaceEntry  = "scan_folder"
	teamScanQueueNamespaceMember = "scan_namespace_member"
	teamScanQueueReceivedFile    = "scan_received_file"
	teamScanQueueSharedLink      = "scan_shared_link"
	teamScanQueueTeamFolder      = "scan_team_folder"
)

type NamespaceEntryParam struct {
	NamespaceId string `json:"namespaceId" path:"namespace_id"`
	Path        string `json:"path" path:"path"`
	IsRetry     bool   `json:"isRetry" path:"is_retry"`
}

type tsImpl struct {
	ctl    app_control.Control
	client dbx_client.Client
	db     *gorm.DB
}

func (z tsImpl) dispatchMember(member *mo_member.Member, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	if member.Status == "removed" {
		return nil
	}
	qMount := stage.Get(teamScanQueueMount)
	qMount.Enqueue(member.TeamMemberId)
	qReceivedFile := stage.Get(teamScanQueueReceivedFile)
	qReceivedFile.Enqueue(member.TeamMemberId)
	qSharedLink := stage.Get(teamScanQueueSharedLink)
	qSharedLink.Enqueue(member.TeamMemberId)

	return nil
}

func (z tsImpl) scanMembers(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	var lastErr error
	opErr := sv_member.New(z.client).ListEach(func(member *mo_member.Member) bool {
		m, err := NewMemberFromJson(es_json.MustParse(member.Raw))
		if err != nil {
			lastErr = err
			return false
		}
		z.db.Create(m)
		if z.db.Error != nil {
			lastErr = z.db.Error
			return false
		}
		if err = z.dispatchMember(member, stage, admin); err != nil {
			lastErr = err
			return false
		}

		return true
	}, sv_member.IncludeDeleted(true))

	return es_lang.NewMultiErrorOrNull(opErr, lastErr)
}

func (z tsImpl) scanGroup(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	gmq := stage.Get(teamScanQueueGroupMember)

	groups, err := sv_group.New(z.client).List()
	if err != nil {
		return err
	}
	for _, group := range groups {
		g, err := NewGroupFromJson(es_json.MustParse(group.Raw))
		if err != nil {
			return err
		}
		z.db.Create(g)
		if z.db.Error != nil {
			return z.db.Error
		}
		gmq.Enqueue(g.GroupId)
	}
	return nil
}

func (z tsImpl) scanGroupMember(groupId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	members, err := sv_group_member.NewByGroupId(z.client, groupId).List()
	if err != nil {
		return err
	}
	for _, member := range members {
		m, err := NewGroupMemberFromJson(groupId, es_json.MustParse(member.Raw))
		if err != nil {
			return err
		}
		z.db.Create(m)
		if z.db.Error != nil {
			return z.db.Error
		}
	}
	return nil
}

func (z tsImpl) scanNamespaces(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	qne := stage.Get(teamScanQueueNamespaceEntry)
	qnd := stage.Get(teamScanQueueNamespaceDetail)
	qnm := stage.Get(teamScanQueueNamespaceMember)

	var lastErr error
	opErr := sv_namespace.New(z.client).ListEach(func(namespace *mo_namespace.Namespace) bool {
		ns, err := NewNamespaceFromJson(es_json.MustParse(namespace.Raw))
		if err != nil {
			lastErr = err
			return false
		}
		z.db.Create(ns)
		if z.db.Error != nil {
			lastErr = z.db.Error
			return false
		}

		qne.Enqueue(&NamespaceEntryParam{
			NamespaceId: ns.NamespaceId,
			Path:        "",
		})
		if ns.NamespaceType != "team_member_folder" && ns.NamespaceType != "app_folder" {
			qnd.Enqueue(ns.NamespaceId)
			qnm.Enqueue(ns.NamespaceId)
		}

		return true
	})

	return es_lang.NewMultiErrorOrNull(opErr, lastErr)
}

func (z tsImpl) scanNamespaceDetail(namespaceId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	ns, err := sv_sharedfolder.New(z.client.AsAdminId(admin.TeamMemberId)).Resolve(namespaceId)
	if err != nil {
		return err
	}
	n, err := NewNamespaceDetail(es_json.MustParse(ns.Raw))
	if err != nil {
		return err
	}
	z.db.Create(n)
	if z.db.Error != nil {
		return z.db.Error
	}
	return nil
}

func (z tsImpl) scanNamespaceMember(namespaceId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	members, err := sv_sharedfolder_member.NewBySharedFolderId(z.client.AsAdminId(admin.TeamMemberId), namespaceId).List()
	if err != nil {
		return err
	}
	for _, member := range members {
		m := NewNamespaceMember(namespaceId, member)
		z.db.Create(m)
		if z.db.Error != nil {
			return z.db.Error
		}
	}
	return nil
}

func (z tsImpl) scanNamespaceEntry(param *NamespaceEntryParam, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	qne := stage.Get(teamScanQueueNamespaceEntry)

	client := z.client.AsAdminId(admin.TeamMemberId).WithPath(dbx_client.Namespace(param.NamespaceId))
	err = sv_file.NewFiles(client).ListEach(mo_path.NewDropboxPath(param.Path),
		func(entry mo_file.Entry) {
			ce := entry.Concrete()
			f, err := NewNamespaceEntry(param.NamespaceId, param.Path, es_json.MustParse(ce.Raw))
			if err != nil {
				return
			}
			z.db.Create(f)

			if ce.IsFolder() {
				qne.Enqueue(&NamespaceEntryParam{
					NamespaceId: param.NamespaceId,
					Path:        f.PathLower,
				})
			}
		},
		sv_file.Recursive(false),
		sv_file.IncludeDeleted(true),
		sv_file.IncludeHasExplicitSharedMembers(true),
	)

	switch err {
	case nil:
		if param.IsRetry {
			z.db.Delete(&NamespaceEntryError{
				NamespaceId: param.NamespaceId,
				Path:        param.Path,
			})
		}

	default:
		z.db.Create(&NamespaceEntryError{
			NamespaceId: param.NamespaceId,
			Path:        param.Path,
			Error:       err.Error(),
		})
	}
	return err
}

func (z tsImpl) scanReceivedFile(teamMemberId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	client := z.client.AsMemberId(teamMemberId)
	received, err := sv_sharing.NewReceived(client).List()
	if err != nil {
		return err
	}
	for _, rf := range received {
		r, err := NewReceivedFileFromJsonWithTeamMemberId(teamMemberId, es_json.MustParse(rf.Raw))
		if err != nil {
			return err
		}
		z.db.Create(r)
	}
	return nil
}

func (z tsImpl) scanMount(teamMemberId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	client := z.client.AsMemberId(teamMemberId)

	mountables, err := sv_sharedfolder_mount.New(client).Mountables()
	if err != nil {
		return err
	}

	for _, mount := range mountables {
		m, err := NewMountFromJsonWithTeamMemberId(teamMemberId, es_json.MustParse(mount.Raw))
		if err != nil {
			return err
		}
		z.db.Create(m)
	}

	return nil
}

func (z tsImpl) scanSharedLink(teamMemberId string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	client := z.client.AsMemberId(teamMemberId)

	links, err := sv_sharedlink.New(client).List()
	if err != nil {
		return err
	}

	for _, link := range links {
		l, err := NewSharedLinkWithTeamMemberId(teamMemberId, es_json.MustParse(link.Metadata().Raw))
		if err != nil {
			return err
		}
		z.db.Create(l)
	}
	return nil
}

func (z tsImpl) scanTeamFolder(dummy string, stage eq_sequence.Stage, admin *mo_profile.Profile) (err error) {
	folders, err := sv_teamfolder.New(z.client).List()
	if err != nil {
		return err
	}

	for _, folder := range folders {
		f, err := NewTeamFolder(es_json.MustParse(folder.Raw))
		if err != nil {
			return err
		}
		z.db.Create(f)
	}
	return nil
}

func (z tsImpl) defineQueues(s eq_sequence.Stage, admin *mo_profile.Profile) {
	s.Define(teamScanQueueGroup, z.scanGroup, s, admin)
	s.Define(teamScanQueueGroupMember, z.scanGroupMember, s, admin)
	s.Define(teamScanQueueMember, z.scanMembers, s, admin)
	s.Define(teamScanQueueMount, z.scanMount, s, admin)
	s.Define(teamScanQueueNamespace, z.scanNamespaces, s, admin)
	s.Define(teamScanQueueNamespaceDetail, z.scanNamespaceDetail, s, admin)
	s.Define(teamScanQueueNamespaceMember, z.scanNamespaceMember, s, admin)
	s.Define(teamScanQueueNamespaceEntry, z.scanNamespaceEntry, s, admin)
	s.Define(teamScanQueueReceivedFile, z.scanReceivedFile, s, admin)
	s.Define(teamScanQueueSharedLink, z.scanSharedLink, s, admin)
	s.Define(teamScanQueueTeamFolder, z.scanTeamFolder, s, admin)
}

func (z tsImpl) ScanTeam() (err error) {
	admin, err := sv_profile.NewTeam(z.client).Admin()
	if err != nil {
		return err
	}

	var lastErr error
	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineQueues(s, admin)

		qMember := s.Get(teamScanQueueMember)
		qMember.Enqueue("")
		qNamespace := s.Get(teamScanQueueNamespace)
		qNamespace.Enqueue("")
		qGroup := s.Get(teamScanQueueGroup)
		qGroup.Enqueue("")
		qTeamFolder := s.Get(teamScanQueueTeamFolder)
		qTeamFolder.Enqueue("")

	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	db, err := z.db.DB()
	if err != nil {
		return err
	}
	_ = db.Close()

	return lastErr
}

func (z tsImpl) RetryErrors() error {
	admin, err := sv_profile.NewTeam(z.client).Admin()
	if err != nil {
		return err
	}

	queueSize := z.ctl.Feature().Concurrency() * 2
	offset := 0

	var lastErr error
	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		z.defineQueues(s, admin)

		for {
			namespaceEntryErrors := make([]NamespaceEntryError, 0)
			tx := z.db.Limit(queueSize).Offset(offset).First(&namespaceEntryErrors)
			offset += int(tx.RowsAffected)

		}

	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	db, err := z.db.DB()
	if err != nil {
		return err
	}
	_ = db.Close()

	return lastErr
}
