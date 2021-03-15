package delete

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_team_sharedlink"
	"github.com/watermint/toolbox/essentials/kvs/kv_kvs"
	"github.com/watermint/toolbox/essentials/kvs/kv_storage"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type TargetLinks struct {
	Url string `json:"url"`
}

const (
	linkStatusEnqueue = "e"
	linkStatusDeleted = "d"
	linkStatusFailure = "f"
)

type Links struct {
	rc_recipe.RemarkIrreversible
	Peer         dbx_conn.ConnScopedTeam
	File         fd_file.RowFeed
	OperationLog rp_model.TransactionReport
	Target       kv_storage.Storage
	LinkNotFound app_msg.Message
}

func (z *Links) Preset() {
	z.File.SetModel(&TargetLinks{})
	z.OperationLog.SetModel(
		&TargetLinks{},
		&mo_sharedlink.SharedLinkMember{},
		rp_model.HiddenColumns(
			"result.shared_link_id",
			"result.account_id",
			"result.team_member_id",
			"result.status",
		),
	)
}

func (z *Links) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}
	loadErr := z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*TargetLinks)
		return z.Target.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutString(r.Url, linkStatusEnqueue)
		})
	})
	if loadErr != nil {
		return loadErr
	}

	members, dErr := sv_member.New(z.Peer.Context()).List()
	if dErr != nil {
		return dErr
	}

	var onDeleteSuccess uc_team_sharedlink.DeleteOnSuccess = func(t *uc_team_sharedlink.DeleteTarget) {
		l := c.Log()

		// Mark the link as deleted
		kvErr := z.Target.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutString(t.Entry.Url, linkStatusDeleted)
		})
		if kvErr != nil {
			l.Warn("Unable to record status link. Report might not accurate", esl.Error(kvErr))
		}

		// Report
		z.OperationLog.Success(&TargetLinks{Url: t.Entry.Url}, t.Entry)
	}
	var onDeleteFailure uc_team_sharedlink.DeleteOnFailure = func(t *uc_team_sharedlink.DeleteTarget, cause error) {
		l := c.Log()

		// Mark the link as deleted
		kvErr := z.Target.Update(func(kvs kv_kvs.Kvs) error {
			return kvs.PutString(t.Entry.Url, linkStatusFailure)
		})
		if kvErr != nil {
			l.Warn("Unable to record status link. Report might not accurate", esl.Error(kvErr))
		}

		// Report
		z.OperationLog.Failure(cause, &TargetLinks{Url: t.Entry.Url})
	}

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("delete_link", uc_team_sharedlink.DeleteMemberLink, c, z.Peer.Context(), onDeleteSuccess, onDeleteFailure)
		var onSharedLink uc_team_sharedlink.OnSharedLinkMember = func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember) {
			l := c.Log().With(esl.Any("member", member), esl.Any("entry", entry))
			shouldDelete := false
			kvErr := z.Target.View(func(kvs kv_kvs.Kvs) error {
				v, kvErr := kvs.GetString(entry.Url)
				if kvErr == kv_kvs.ErrorNotFound {
					return nil
				}
				if v == linkStatusEnqueue {
					shouldDelete = true
				}
				return kvErr
			})
			if kvErr != nil {
				l.Debug("Abort delete because of KVS error", esl.Error(kvErr))
				return
			}
			if shouldDelete {
				qml := s.Get("delete_link")
				qml.Enqueue(&uc_team_sharedlink.DeleteTarget{
					Member: member,
					Entry:  entry,
				})
			}
		}
		s.Define("scan_member", uc_team_sharedlink.RetrieveMemberLinks, c, z.Peer.Context(), onSharedLink)
		qsm := s.Get("scan_member")
		for _, member := range members {
			qsm.Enqueue(member)
		}
	})

	return z.Target.View(func(kvs kv_kvs.Kvs) error {
		return kvs.ForEach(func(key string, value []byte) error {
			s := string(value)
			if s == linkStatusEnqueue {
				z.OperationLog.Skip(z.LinkNotFound, &TargetLinks{Url: key})
			}
			return nil
		})
	})
}

func (z *Links) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("links", "https://www.dropbox.com/scl/fo/fir9vjelf\nhttps://www.dropbox.com/scl/fo/fir9vjelg")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.ExecMock(c, &Links{}, func(r rc_recipe.Recipe) {
		m := r.(*Links)
		m.File.SetFilePath(f)
	})
}
