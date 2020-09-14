package uc_member_folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_random"
)

type MemberNamespace struct {
	Namespace       *mo_sharedfolder.SharedFolder `json:"namespace"`
	TeamMemberId    string                        `json:"team_member_id"`
	TeamMemberEmail string                        `json:"team_member_email"`
}

type Scanner interface {
	Scan(filter mo_filter.Filter) (namespaces []*MemberNamespace, err error)
}

func New(ctl app_control.Control, ctx dbx_context.Context) Scanner {
	return &scanImpl{
		ctl: ctl,
		ctx: ctx,
	}
}

const (
	queueIdScanMember        = "scan_members"
	queueIdScanMemberFolders = "scan_member_folders"
)

type scanImpl struct {
	ctl app_control.Control
	ctx dbx_context.Context
}

func (z scanImpl) scanMember(sessionId string, stage eq_sequence.Stage) error {
	l := z.ctl.Log().With(esl.String("sessionId", sessionId))
	l.Debug("Scan members")

	return sv_member.New(z.ctx).ListEach(func(member *mo_member.Member) bool {
		q := stage.Get(queueIdScanMemberFolders)
		q.Enqueue(member)
		return true
	})
}

func (z scanImpl) scanNamespace(member *mo_member.Member, storageNamespace kv_storage.Storage) error {
	l := z.ctl.Log().With(esl.String("teamMemberId", member.TeamMemberId), esl.String("teamMemberEmail", member.Email))
	l.Debug("Scan member folders")

	folders, err := sv_sharedfolder.New(z.ctx.AsMemberId(member.TeamMemberId)).List()
	if err != nil {
		l.Debug("Unable to scan team member folders", esl.Error(err))
		return err
	}

	for _, folder := range folders {
		err = storageNamespace.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutJsonModel(folder.SharedFolderId, &MemberNamespace{
				TeamMemberId:    member.TeamMemberId,
				TeamMemberEmail: member.Email,
				Namespace:       folder,
			})
		})
	}

	l.Debug("Finished")
	return nil
}

func (z scanImpl) Scan(filter mo_filter.Filter) (namespaces []*MemberNamespace, err error) {
	l := z.ctl.Log()
	scanSessionId := sc_random.MustGenerateRandomString(8)
	namespaces = make([]*MemberNamespace, 0)

	storageNamespace, err := z.ctl.NewKvs("namespace_" + scanSessionId)
	if err != nil {
		l.Debug("Unable to create temporary storage", esl.Error(err))
		return nil, err
	}
	defer storageNamespace.Close()

	z.ctl.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define(queueIdScanMember, z.scanMember, s)
		s.Define(queueIdScanMemberFolders, z.scanNamespace, storageNamespace)
		q := s.Get(queueIdScanMember)
		q.Enqueue(scanSessionId)
	})

	err = storageNamespace.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEachModel(&MemberNamespace{}, func(key string, m interface{}) error {
			namespace := m.(*MemberNamespace)
			if filter.Accept(namespace.Namespace.Name) {
				namespaces = append(namespaces, namespace)
			}
			return nil
		})
	})
	return
}
